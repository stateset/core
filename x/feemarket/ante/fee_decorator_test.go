package ante_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	protov2 "google.golang.org/protobuf/proto"

	"github.com/stateset/core/x/feemarket/ante"
	feemarketkeeper "github.com/stateset/core/x/feemarket/keeper"
	feemarkettypes "github.com/stateset/core/x/feemarket/types"
)

type testFeeTx struct {
	fee sdk.Coins
	gas uint64
}

func (t testFeeTx) GetMsgs() []sdk.Msg { return nil }
func (t testFeeTx) GetMsgsV2() ([]protov2.Message, error) {
	return nil, nil
}
func (t testFeeTx) GetGas() uint64     { return t.gas }
func (t testFeeTx) GetFee() sdk.Coins  { return t.fee }
func (t testFeeTx) FeePayer() []byte   { return nil }
func (t testFeeTx) FeeGranter() []byte { return nil }

func TestFeeMarketCheckTxFeeWithMinGasPrices_ZeroFeeDoesNotPanic(t *testing.T) {
	key := storetypes.NewKVStoreKey(feemarkettypes.StoreKey)
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockTime(time.Now().UTC())

	ctx = ctx.WithMinGasPrices(
		sdk.NewDecCoins(
			sdk.NewDecCoinFromDec("ustate", sdkmath.LegacyMustNewDecFromStr("0.1")),
		),
	)

	k := feemarketkeeper.NewKeeper(nil, key, "authority")
	checker := ante.FeeMarketCheckTxFeeWithMinGasPrices(k)

	tx := testFeeTx{fee: nil, gas: 1000}
	_, _, err := checker(ctx, tx)
	require.Error(t, err)
	require.ErrorIs(t, err, feemarkettypes.ErrInsufficientFee)
}
