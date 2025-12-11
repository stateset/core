package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/keeper"
)

// setupStablecoinKeeper is used by older stablecoin tests. It delegates to setupKeeper.
func setupStablecoinKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *mockBankKeeper, *mockOracleKeeper, *mockComplianceKeeper) {
	return setupKeeper(t)
}

// newStablecoinAddress is used by older stablecoin tests. It delegates to newAddress.
func newStablecoinAddress() sdk.AccAddress {
	return newAddress()
}
