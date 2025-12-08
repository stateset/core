package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ReserveSnapshot captures the off-chain reserve attestations aligned with on-chain supply.
type ReserveSnapshot struct {
	Id            uint64    `json:"id" yaml:"id"`
	Reporter      string    `json:"reporter" yaml:"reporter"`
	TotalSupply   sdk.Coin  `json:"total_supply" yaml:"total_supply"`
	FiatReserves  sdk.Coin  `json:"fiat_reserves" yaml:"fiat_reserves"`
	OtherReserves sdk.Coins `json:"other_reserves" yaml:"other_reserves"`
	Timestamp     time.Time `json:"timestamp" yaml:"timestamp"`
	Metadata      string    `json:"metadata" yaml:"metadata"`
}

func (r ReserveSnapshot) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(r.Reporter); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, err.Error())
	}
	if !r.TotalSupply.IsValid() {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid total supply")
	}
	if !r.FiatReserves.IsValid() {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid fiat reserves")
	}
	if err := r.OtherReserves.Validate(); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, err.Error())
	}
	return nil
}
