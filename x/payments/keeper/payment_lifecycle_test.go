package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/payments/types"
)

// Alias for existing test helpers
var (
	setupPaymentKeeper = setupPaymentsKeeper
	newPaymentAddress  = newPaymentsAddress
)

// Helper to set sanctioned status in mock compliance keeper
func (m *mockComplianceKeeper) SetSanctioned(addr string, sanctioned bool) {
	if sanctioned {
		m.blocked[addr] = true
	} else {
		delete(m.blocked, addr)
	}
}

// Helper to set balance using string address
func (m *mockBankKeeper) SetBalance(addr string, coins sdk.Coins) {
	m.balances[addr] = coins.Sort()
}

// TestPaymentLifecycle_CreatePayment tests creating a payment intent
func TestPaymentLifecycle_CreatePayment(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund payer
	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	intent := types.PaymentIntent{
		Payer:    payer.String(),
		Payee:    payee.String(),
		Amount:   amount,
		Metadata: "test payment",
	}

	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)
	require.Equal(t, uint64(1), id)

	// Verify payment was stored
	payment, found := k.GetPayment(ctx, id)
	require.True(t, found)
	require.Equal(t, payer.String(), payment.Payer)
	require.Equal(t, payee.String(), payment.Payee)
	require.Equal(t, amount, payment.Amount)
	require.Equal(t, types.PaymentStatusPending, payment.Status)
}

func TestPaymentLifecycle_CreatePayment_InvalidPayer(t *testing.T) {
	k, ctx, _, _ := setupPaymentKeeper(t)

	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	intent := types.PaymentIntent{
		Payer:  "invalid",
		Payee:  payee.String(),
		Amount: amount,
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidAddress)
}

func TestPaymentLifecycle_CreatePayment_InvalidPayee(t *testing.T) {
	k, ctx, _, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  "invalid",
		Amount: amount,
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidAddress)
}

func TestPaymentLifecycle_CreatePayment_SamePayerPayee(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	addr := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(addr.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	intent := types.PaymentIntent{
		Payer:  addr.String(),
		Payee:  addr.String(), // Same as payer
		Amount: amount,
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidAddress)
}

func TestPaymentLifecycle_CreatePayment_ZeroAmount(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidAmount)
}

func TestPaymentLifecycle_CreatePayment_NegativeAmount(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: sdk.Coin{Denom: "ssusd", Amount: sdkmath.NewInt(-1000)},
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
}

func TestPaymentLifecycle_CreatePayment_InsufficientBalance(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund payer with less than needed
	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(500000))))

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInsufficientBalance)
}

func TestPaymentLifecycle_CreatePayment_SanctionedPayer(t *testing.T) {
	k, ctx, bankKeeper, compKeeper := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))
	compKeeper.SetSanctioned(payer.String(), true)

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
}

func TestPaymentLifecycle_CreatePayment_SanctionedPayee(t *testing.T) {
	k, ctx, bankKeeper, compKeeper := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))
	compKeeper.SetSanctioned(payee.String(), true)

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}

	_, err := k.CreatePayment(ctx, intent)
	require.Error(t, err)
}

// TestPaymentLifecycle_SettlePayment tests settling a payment
func TestPaymentLifecycle_SettlePayment(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Settle payment
	err = k.SettlePayment(ctx, id, payee)
	require.NoError(t, err)

	// Verify payment status
	payment, found := k.GetPayment(ctx, id)
	require.True(t, found)
	require.Equal(t, types.PaymentStatusSettled, payment.Status)
	require.False(t, payment.SettledTime.IsZero())
}

func TestPaymentLifecycle_SettlePayment_NotFound(t *testing.T) {
	k, ctx, _, _ := setupPaymentKeeper(t)

	payee := newPaymentAddress()

	err := k.SettlePayment(ctx, 9999, payee)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPaymentNotFound)
}

func TestPaymentLifecycle_SettlePayment_WrongPayee(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	wrongPayee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Try to settle with wrong payee
	err = k.SettlePayment(ctx, id, wrongPayee)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPaymentLifecycle_SettlePayment_AlreadySettled(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Settle payment
	err = k.SettlePayment(ctx, id, payee)
	require.NoError(t, err)

	// Try to settle again
	err = k.SettlePayment(ctx, id, payee)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPaymentCompleted)
}

func TestPaymentLifecycle_SettlePayment_Cancelled(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Cancel payment
	err = k.CancelPayment(ctx, id, payer)
	require.NoError(t, err)

	// Try to settle cancelled payment
	err = k.SettlePayment(ctx, id, payee)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPaymentCancelled)
}

// TestPaymentLifecycle_CancelPayment tests cancelling a payment
func TestPaymentLifecycle_CancelPayment(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Cancel payment
	err = k.CancelPayment(ctx, id, payer)
	require.NoError(t, err)

	// Verify payment status
	payment, found := k.GetPayment(ctx, id)
	require.True(t, found)
	require.Equal(t, types.PaymentStatusCancelled, payment.Status)
}

func TestPaymentLifecycle_CancelPayment_NotFound(t *testing.T) {
	k, ctx, _, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()

	err := k.CancelPayment(ctx, 9999, payer)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPaymentNotFound)
}

func TestPaymentLifecycle_CancelPayment_WrongPayer(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	wrongPayer := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Try to cancel with wrong payer
	err = k.CancelPayment(ctx, id, wrongPayer)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestPaymentLifecycle_CancelPayment_AlreadySettled(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Settle payment
	err = k.SettlePayment(ctx, id, payee)
	require.NoError(t, err)

	// Try to cancel settled payment
	err = k.CancelPayment(ctx, id, payer)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPaymentCompleted)
}

func TestPaymentLifecycle_CancelPayment_AlreadyCancelled(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create payment
	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: amount,
	}
	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	// Cancel payment
	err = k.CancelPayment(ctx, id, payer)
	require.NoError(t, err)

	// Try to cancel again
	err = k.CancelPayment(ctx, id, payer)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPaymentCancelled)
}

// TestPaymentLifecycle_MultiplePayments tests creating multiple payments
func TestPaymentLifecycle_MultiplePayments(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee1 := newPaymentAddress()
	payee2 := newPaymentAddress()
	payee3 := newPaymentAddress()

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create multiple payments
	payees := []sdk.AccAddress{payee1, payee2, payee3}
	for i, payee := range payees {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(int64((i+1)*100000)))
		intent := types.PaymentIntent{
			Payer:  payer.String(),
			Payee:  payee.String(),
			Amount: amount,
		}
		id, err := k.CreatePayment(ctx, intent)
		require.NoError(t, err)
		require.Equal(t, uint64(i+1), id)
	}

	// Verify all payments exist
	for i := uint64(1); i <= 3; i++ {
		payment, found := k.GetPayment(ctx, i)
		require.True(t, found)
		require.Equal(t, types.PaymentStatusPending, payment.Status)
	}
}

// TestPaymentLifecycle_IteratePayments tests payment iteration
func TestPaymentLifecycle_IteratePayments(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create multiple payments
	for i := 0; i < 5; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		intent := types.PaymentIntent{
			Payer:  payer.String(),
			Payee:  payee.String(),
			Amount: amount,
		}
		_, err := k.CreatePayment(ctx, intent)
		require.NoError(t, err)
	}

	// Iterate and count
	count := 0
	k.IteratePayments(ctx, func(p types.PaymentIntent) bool {
		count++
		return false
	})
	require.Equal(t, 5, count)
}

func TestPaymentLifecycle_IteratePayments_EarlyStop(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create multiple payments
	for i := 0; i < 10; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		intent := types.PaymentIntent{
			Payer:  payer.String(),
			Payee:  payee.String(),
			Amount: amount,
		}
		_, err := k.CreatePayment(ctx, intent)
		require.NoError(t, err)
	}

	// Iterate and stop after 3
	count := 0
	k.IteratePayments(ctx, func(p types.PaymentIntent) bool {
		count++
		return count >= 3
	})
	require.Equal(t, 3, count)
}

// TestPaymentLifecycle_GenesisExportImport tests genesis state
func TestPaymentLifecycle_GenesisExportImport(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create payments
	for i := 0; i < 3; i++ {
		amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		intent := types.PaymentIntent{
			Payer:  payer.String(),
			Payee:  payee.String(),
			Amount: amount,
		}
		_, err := k.CreatePayment(ctx, intent)
		require.NoError(t, err)
	}

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Len(t, genesis.Payments, 3)
	require.Equal(t, uint64(4), genesis.NextPaymentId)

	// Create new keeper and import
	k2, ctx2, _, _ := setupPaymentKeeper(t)
	k2.InitGenesis(ctx2, genesis)

	// Verify imported payments
	count := 0
	k2.IteratePayments(ctx2, func(p types.PaymentIntent) bool {
		count++
		return false
	})
	require.Equal(t, 3, count)

	// Verify next ID
	nextID := k2.getNextID(ctx2)
	require.Equal(t, uint64(4), nextID)
}

// TestPaymentLifecycle_LargeAmount tests handling of large amounts
func TestPaymentLifecycle_LargeAmount(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()
	// 1 billion units
	largeAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000000))

	bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000000000))))

	intent := types.PaymentIntent{
		Payer:  payer.String(),
		Payee:  payee.String(),
		Amount: largeAmount,
	}

	id, err := k.CreatePayment(ctx, intent)
	require.NoError(t, err)

	payment, found := k.GetPayment(ctx, id)
	require.True(t, found)
	require.Equal(t, largeAmount, payment.Amount)
}

// TestPaymentLifecycle_DifferentDenominations tests different coin denominations
func TestPaymentLifecycle_DifferentDenominations(t *testing.T) {
	k, ctx, bankKeeper, _ := setupPaymentKeeper(t)

	payer := newPaymentAddress()
	payee := newPaymentAddress()

	// Setup different denominations
	denoms := []string{"ssusd", "atom", "osmo"}
	for _, denom := range denoms {
		bankKeeper.SetBalance(payer.String(), sdk.NewCoins(sdk.NewCoin(denom, sdkmath.NewInt(2000000))))

		amount := sdk.NewCoin(denom, sdkmath.NewInt(1000000))
		intent := types.PaymentIntent{
			Payer:  payer.String(),
			Payee:  payee.String(),
			Amount: amount,
		}

		id, err := k.CreatePayment(ctx, intent)
		require.NoError(t, err)

		payment, found := k.GetPayment(ctx, id)
		require.True(t, found)
		require.Equal(t, denom, payment.Amount.Denom)
	}
}
