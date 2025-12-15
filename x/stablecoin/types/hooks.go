package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StablecoinHooks event hooks for stablecoin state changes
type StablecoinHooks interface {
	AfterMintStablecoin(ctx sdk.Context, depositor sdk.AccAddress, amount sdk.Coin)
	AfterRedeemStablecoin(ctx sdk.Context, redeemer sdk.AccAddress, amount sdk.Coin)
}

// MultiStablecoinHooks combines multiple stablecoin hooks
type MultiStablecoinHooks []StablecoinHooks

func NewMultiStablecoinHooks(hooks ...StablecoinHooks) MultiStablecoinHooks {
	return hooks
}

func (h MultiStablecoinHooks) AfterMintStablecoin(ctx sdk.Context, depositor sdk.AccAddress, amount sdk.Coin) {
	for i := range h {
		h[i].AfterMintStablecoin(ctx, depositor, amount)
	}
}

func (h MultiStablecoinHooks) AfterRedeemStablecoin(ctx sdk.Context, redeemer sdk.AccAddress, amount sdk.Coin) {
	for i := range h {
		h[i].AfterRedeemStablecoin(ctx, redeemer, amount)
	}
}
