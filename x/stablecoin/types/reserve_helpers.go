package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ReserveAssetType defines the type of reserve asset.
type ReserveAssetType = string

const (
	ReserveAssetCash              ReserveAssetType = "cash"
	ReserveAssetTBill             ReserveAssetType = "t_bill"
	ReserveAssetTNote             ReserveAssetType = "t_note"
	ReserveAssetTBond             ReserveAssetType = "t_bond"
	ReserveAssetTokenizedTreasury ReserveAssetType = "tokenized_treasury"
	ReserveAssetRepoAgreement     ReserveAssetType = "repo"
	ReserveAssetMMF               ReserveAssetType = "mmf"
)

// DepositStatus represents the status of a reserve deposit.
type DepositStatus = string

const (
	DepositStatusActive    DepositStatus = "active"
	DepositStatusRedeeming DepositStatus = "redeeming"
	DepositStatusRedeemed  DepositStatus = "redeemed"
)

// RedeemStatus represents the status of a redemption request.
type RedeemStatus = string

const (
	RedeemStatusPending   RedeemStatus = "pending"
	RedeemStatusExecuted  RedeemStatus = "executed"
	RedeemStatusCancelled RedeemStatus = "cancelled"
)

// ApprovedReserveAssets returns all approved reserve asset types.
func ApprovedReserveAssets() []ReserveAssetType {
	return []ReserveAssetType{
		ReserveAssetCash,
		ReserveAssetTBill,
		ReserveAssetTNote,
		ReserveAssetTBond,
		ReserveAssetTokenizedTreasury,
		ReserveAssetRepoAgreement,
		ReserveAssetMMF,
	}
}

// IsApprovedReserveAsset checks if an asset type is approved.
func IsApprovedReserveAsset(assetType ReserveAssetType) bool {
	for _, approved := range ApprovedReserveAssets() {
		if approved == assetType {
			return true
		}
	}
	return false
}

func (c TokenizedTreasuryConfig) ValidateBasic() error {
	if c.Denom == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "denom cannot be empty")
	}
	if c.Issuer == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "issuer cannot be empty")
	}
	if c.HaircutBps > 5000 {
		return errorsmod.Wrap(ErrInvalidReserve, "haircut cannot exceed 50%")
	}
	if c.MaxAllocationBps > 10000 {
		return errorsmod.Wrap(ErrInvalidReserve, "max allocation cannot exceed 100%")
	}
	return nil
}

func DefaultReserveParams() ReserveParams {
	return ReserveParams{
		MinReserveRatioBps:    10000,
		TargetReserveRatioBps: 10200,
		MintFeeBps:            10,
		RedeemFeeBps:          10,
		MinMintAmount:         sdkmath.NewInt(100_000_000),
		MinRedeemAmount:       sdkmath.NewInt(100_000_000),
		RedemptionDelay:       0,
		MaxDailyMint:          sdkmath.NewInt(100_000_000_000_000),
		MaxDailyRedeem:        sdkmath.NewInt(100_000_000_000_000),
		TokenizedTreasuries:   DefaultTokenizedTreasuries(),
		RequireKyc:            true,
		MintPaused:            false,
		RedeemPaused:          false,
	}
}

func DefaultTokenizedTreasuries() []TokenizedTreasuryConfig {
	return []TokenizedTreasuryConfig{
		{
			Denom:            "usdy",
			Issuer:           "ondo",
			UnderlyingType:   ReserveAssetTBill,
			Active:           true,
			HaircutBps:       50,
			MaxAllocationBps: 5000,
			OracleDenom:      "usdy",
		},
		{
			Denom:            "stbt",
			Issuer:           "matrixdock",
			UnderlyingType:   ReserveAssetTBill,
			Active:           true,
			HaircutBps:       50,
			MaxAllocationBps: 3000,
			OracleDenom:      "stbt",
		},
		{
			Denom:            "ousg",
			Issuer:           "ondo",
			UnderlyingType:   ReserveAssetTBond,
			Active:           true,
			HaircutBps:       100,
			MaxAllocationBps: 3000,
			OracleDenom:      "ousg",
		},
		{
			Denom:            "tbill",
			Issuer:           "openeden",
			UnderlyingType:   ReserveAssetTBill,
			Active:           true,
			HaircutBps:       50,
			MaxAllocationBps: 4000,
			OracleDenom:      "tbill",
		},
		{
			Denom:            "usdc",
			Issuer:           "circle",
			UnderlyingType:   ReserveAssetCash,
			Active:           true,
			HaircutBps:       0,
			MaxAllocationBps: 2000,
			OracleDenom:      "usdc",
		},
	}
}

func (p ReserveParams) ValidateBasic() error {
	if p.MinReserveRatioBps < 10000 {
		return errorsmod.Wrap(ErrInvalidReserve, "minimum reserve ratio must be at least 100%")
	}
	if p.TargetReserveRatioBps < p.MinReserveRatioBps {
		return errorsmod.Wrap(ErrInvalidReserve, "target ratio must be >= minimum ratio")
	}
	if p.MintFeeBps > 1000 {
		return errorsmod.Wrap(ErrInvalidReserve, "mint fee cannot exceed 10%")
	}
	if p.RedeemFeeBps > 1000 {
		return errorsmod.Wrap(ErrInvalidReserve, "redeem fee cannot exceed 10%")
	}
	for _, tt := range p.TokenizedTreasuries {
		if err := tt.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}

func (p ReserveParams) GetTokenizedTreasury(denom string) (TokenizedTreasuryConfig, bool) {
	for _, tt := range p.TokenizedTreasuries {
		if tt.Denom == denom {
			return tt, true
		}
	}
	return TokenizedTreasuryConfig{}, false
}

func (r Reserve) GetReserveRatio() uint32 {
	if r.TotalMinted.IsZero() {
		return 10000
	}

	ratio := r.TotalValue.Mul(sdkmath.NewInt(10000)).Quo(r.TotalMinted)
	if ratio.GT(sdkmath.NewInt(100000)) {
		return 100000
	}
	return uint32(ratio.Int64())
}

func (r Reserve) IsHealthy(minRatioBps uint32) bool {
	return r.GetReserveRatio() >= minRatioBps
}

func (a OffChainReserveAttestation) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(a.Attester); err != nil {
		return errorsmod.Wrap(ErrInvalidReserve, "invalid attester address")
	}
	if a.TotalValue.IsNegative() {
		return errorsmod.Wrap(ErrInvalidReserve, "total value cannot be negative")
	}
	if a.CustodianName == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "custodian name required")
	}
	return nil
}

// TotalReserves aggregates on-chain and off-chain reserves.
func (t TotalReserves) CalculateReserveRatio() uint32 {
	if t.TotalSupply.IsZero() {
		return 10000
	}
	ratio := t.TotalValue.Mul(sdkmath.NewInt(10000)).Quo(t.TotalSupply)
	if ratio.GT(sdkmath.NewInt(100000)) {
		return 100000
	}
	return uint32(ratio.Int64())
}

func zeroTime() time.Time { return time.Time{} }
