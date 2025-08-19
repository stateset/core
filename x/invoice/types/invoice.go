package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates the invoice
func (i Invoice) Validate() error {
	if i.Id == 0 {
		return fmt.Errorf("invoice ID cannot be 0")
	}
	if i.Did == "" {
		return fmt.Errorf("invoice DID cannot be empty")
	}
	if i.Amount == "" {
		return fmt.Errorf("invoice amount cannot be empty")
	}
	if i.Seller == "" {
		return fmt.Errorf("seller cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(i.Seller); err != nil {
		return fmt.Errorf("invalid seller address: %w", err)
	}
	if i.Purchaser != "" {
		if _, err := sdk.AccAddressFromBech32(i.Purchaser); err != nil {
			return fmt.Errorf("invalid purchaser address: %w", err)
		}
	}
	if i.Factor != "" {
		if _, err := sdk.AccAddressFromBech32(i.Factor); err != nil {
			return fmt.Errorf("invalid factor address: %w", err)
		}
	}
	return nil
}