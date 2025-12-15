package keeper

import (
	"encoding/binary"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/payments/types"
)

var (
	nextIDKey = []byte{0x02}
)

// Keeper handles payment intent lifecycle.
type Keeper struct {
	storeKey   storetypes.StoreKey
	bankKeeper types.BankKeeper
	compKeeper types.ComplianceKeeper

	moduleName string
}

func NewKeeper(_ codec.BinaryCodec, key storetypes.StoreKey, bank types.BankKeeper, compliance types.ComplianceKeeper, moduleName string) Keeper {
	return Keeper{
		storeKey:   key,
		bankKeeper: bank,
		compKeeper: compliance,
		moduleName: moduleName,
	}
}

func (k Keeper) getNextID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(nextIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(nextIDKey, bz)
}

func (k Keeper) storePayment(ctx sdk.Context, payment types.PaymentIntent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PaymentKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&payment)
	store.Set(mustWriteUint64(payment.Id), bz)
}

func (k Keeper) GetPayment(ctx sdk.Context, id uint64) (types.PaymentIntent, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PaymentKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.PaymentIntent{}, false
	}
	var payment types.PaymentIntent
	types.ModuleCdc.MustUnmarshalJSON(bz, &payment)
	return payment, true
}

func (k Keeper) CreatePayment(ctx sdk.Context, intent types.PaymentIntent) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	// Validate payer address
	payerAddr, err := sdk.AccAddressFromBech32(intent.Payer)
	if err != nil {
		return 0, errorsmod.Wrapf(types.ErrInvalidAddress, "invalid payer address: %s", err)
	}

	// Validate payee address
	payeeAddr, err := sdk.AccAddressFromBech32(intent.Payee)
	if err != nil {
		return 0, errorsmod.Wrapf(types.ErrInvalidAddress, "invalid payee address: %s", err)
	}

	// Ensure payer and payee are different
	if payerAddr.Equals(payeeAddr) {
		return 0, errorsmod.Wrap(types.ErrInvalidAddress, "payer and payee cannot be the same")
	}

	// Validate amount is positive
	if !intent.Amount.IsPositive() {
		return 0, errorsmod.Wrap(types.ErrInvalidAmount, "amount must be positive")
	}

	if err := k.compKeeper.AssertCompliant(wrappedCtx, payerAddr); err != nil {
		return 0, err
	}
	if err := k.compKeeper.AssertCompliant(wrappedCtx, payeeAddr); err != nil {
		return 0, err
	}

	balance := k.bankKeeper.GetBalance(wrappedCtx, payerAddr, intent.Amount.Denom)
	if balance.IsLT(intent.Amount) {
		return 0, types.ErrInsufficientBalance
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, payerAddr, k.moduleName, sdk.NewCoins(intent.Amount)); err != nil {
		return 0, err
	}

	nextID := k.getNextID(ctx)
	intent.Id = nextID
	intent.Status = types.PaymentStatusPending
	intent.CreatedHeight = ctx.BlockHeight()
	intent.CreatedTime = ctx.BlockTime()

	k.storePayment(ctx, intent)

	// Calculate and store optimal route
	route := k.OptimizeRoute(ctx, payerAddr, payeeAddr, intent.Amount)
	k.SetPaymentRoute(ctx, nextID, route)

	k.setNextID(ctx, nextID+1)
	return intent.Id, nil
}

func (k Keeper) SettlePayment(ctx sdk.Context, id uint64, payee sdk.AccAddress) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	payment, found := k.GetPayment(ctx, id)
	if !found {
		return types.ErrPaymentNotFound
	}
	if payment.Status == types.PaymentStatusSettled {
		return types.ErrPaymentCompleted
	}
	if payment.Status == types.PaymentStatusCancelled {
		return types.ErrPaymentCancelled
	}

	expectedPayee, err := sdk.AccAddressFromBech32(payment.Payee)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidAddress, "invalid stored payee address: %s", err)
	}
	if !expectedPayee.Equals(payee) {
		return types.ErrNotAuthorized
	}

	if err := k.compKeeper.AssertCompliant(wrappedCtx, payee); err != nil {
		return err
	}

	coin := payment.Amount
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, k.moduleName, payee, sdk.NewCoins(coin)); err != nil {
		return err
	}

	payment.Status = types.PaymentStatusSettled
	payment.SettledHeight = ctx.BlockHeight()
	payment.SettledTime = ctx.BlockTime()
	k.storePayment(ctx, payment)
	return nil
}

func (k Keeper) CancelPayment(ctx sdk.Context, id uint64, payer sdk.AccAddress) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	payment, found := k.GetPayment(ctx, id)
	if !found {
		return types.ErrPaymentNotFound
	}
	if payment.Status == types.PaymentStatusSettled {
		return types.ErrPaymentCompleted
	}
	if payment.Status == types.PaymentStatusCancelled {
		return types.ErrPaymentCancelled
	}

	payerAddr, err := sdk.AccAddressFromBech32(payment.Payer)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidAddress, "invalid stored payer address: %s", err)
	}
	if !payerAddr.Equals(payer) {
		return types.ErrNotAuthorized
	}

	if err := k.compKeeper.AssertCompliant(wrappedCtx, payer); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, k.moduleName, payer, sdk.NewCoins(payment.Amount)); err != nil {
		return err
	}

	payment.Status = types.PaymentStatusCancelled
	k.storePayment(ctx, payment)
	return nil
}

func (k Keeper) IteratePayments(ctx sdk.Context, cb func(types.PaymentIntent) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PaymentKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var payment types.PaymentIntent
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &payment)
		if cb(payment) {
			break
		}
	}
}

func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.setNextID(ctx, state.NextPaymentId)
	for _, payment := range state.Payments {
		k.storePayment(ctx, payment)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.NextPaymentId = k.getNextID(ctx)
	k.IteratePayments(ctx, func(payment types.PaymentIntent) bool {
		state.Payments = append(state.Payments, payment)
		return false
	})
	return state
}

func mustWriteUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
