package types

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys for ReserveParams
var (
	KeyMinReserveRatioBps    = []byte("MinReserveRatioBps")
	KeyTargetReserveRatioBps = []byte("TargetReserveRatioBps")
	KeyMintFeeBps            = []byte("MintFeeBps")
	KeyRedeemFeeBps          = []byte("RedeemFeeBps")
	KeyMinMintAmount         = []byte("MinMintAmount")
	KeyMinRedeemAmount       = []byte("MinRedeemAmount")
	KeyRedemptionDelay       = []byte("RedemptionDelay")
	KeyMaxDailyMint          = []byte("MaxDailyMint")
	KeyMaxDailyRedeem        = []byte("MaxDailyRedeem")
	KeyTokenizedTreasuries   = []byte("TokenizedTreasuries")
	KeyRequireKyc            = []byte("RequireKyc")
	KeyMintPaused            = []byte("MintPaused")	// No limit
	KeyRedeemPaused          = []byte("RedeemPaused")	// No limit
)

// ReserveParamKeyTable for stablecoin module's reserve parameters
func ReserveParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&ReserveParams{})
}

// ParamKeyTable is an alias for ReserveParamKeyTable for backwards compatibility
func ParamKeyTable() paramtypes.KeyTable {
	return ReserveParamKeyTable()
}

// DefaultReserveParams returns a default set of parameters for the reserve-backed stablecoin.
func DefaultReserveParams() ReserveParams {
	return ReserveParams{
		MinReserveRatioBps:    10000, // 100% (must be at least 100% per validation)
		TargetReserveRatioBps: 10000, // 100% (cannot exceed 100% per validation)
		MintFeeBps:            0,    // 0%
		RedeemFeeBps:          0,    // 0%
		MinMintAmount:         sdkmath.NewInt(1_000_000), // 1 ssUSD
		MinRedeemAmount:       sdkmath.NewInt(1_000_000), // 1 ssUSD
		RedemptionDelay:       0 * time.Minute,
		MaxDailyMint:          sdkmath.NewInt(0), // No limit
		MaxDailyRedeem:        sdkmath.NewInt(0), // No limit
		TokenizedTreasuries: []TokenizedTreasuryConfig{
			{
				Denom:            "ustn", // Default tokenized US Treasury Note
				Issuer:           "Stateset Treasury",
				UnderlyingType:   ReserveAssetTNote,
				Active:           true,
				HaircutBps:       500, // 5% haircut
				MaxAllocationBps: 10000, // 100% of reserves
				OracleDenom:      "USTN",
			},
			{
				Denom:            "ibc/OPENEDEN_TBILL_HASH", // Assumed IBC denom for OpenEden TBill
				Issuer:           "OpenEden",
				UnderlyingType:   ReserveAssetTBill,
				Active:           true,
				HaircutBps:       100, // 1% haircut (example)
				MaxAllocationBps: 5000, // 50% max allocation (example)
				OracleDenom:      "USDTBILL", // Oracle denom for OpenEden TBill
			},
		},
		RequireKyc:   true,
		MintPaused:   false,
		RedeemPaused: false,
	}
}

// Validate validates the ReserveParams
func (p ReserveParams) Validate() error {
	if p.MinReserveRatioBps > 10000 {
		return fmt.Errorf("min reserve ratio cannot exceed 10000 (100%%)")
	}
	if p.TargetReserveRatioBps > 10000 {
		return fmt.Errorf("target reserve ratio cannot exceed 10000 (100%%)")
	}
	if p.MinReserveRatioBps > p.TargetReserveRatioBps {
		return fmt.Errorf("min reserve ratio cannot be greater than target reserve ratio")
	}
	if p.MintFeeBps > 10000 {
		return fmt.Errorf("mint fee cannot exceed 10000 (100%%)")
	}
	if p.RedeemFeeBps > 10000 {
		return fmt.Errorf("redeem fee cannot exceed 10000 (100%%)")
	}
	if p.MinMintAmount.IsNegative() {
		return fmt.Errorf("min mint amount cannot be negative")
	}
	if p.MinRedeemAmount.IsNegative() {
		return fmt.Errorf("min redeem amount cannot be negative")
	}
	if p.RedemptionDelay < 0 {
		return fmt.Errorf("redemption delay cannot be negative")
	}
	if p.MaxDailyMint.IsNegative() {
		return fmt.Errorf("max daily mint cannot be negative")
	}
	if p.MaxDailyRedeem.IsNegative() {
		return fmt.Errorf("max daily redeem cannot be negative")
	}

	for _, tt := range p.TokenizedTreasuries {
		if err := tt.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// TokenizedTreasuryConfig
func (ttc TokenizedTreasuryConfig) Validate() error {
	if ttc.Denom == "" {
		return fmt.Errorf("tokenized treasury denom cannot be empty")
	}
	if ttc.OracleDenom == "" {
		return fmt.Errorf("tokenized treasury oracle denom cannot be empty")
	}
	if ttc.HaircutBps > 10000 {
		return fmt.Errorf("tokenized treasury haircut cannot exceed 10000 (100%%)")
	}
	if ttc.MaxAllocationBps > 10000 {
		return fmt.Errorf("tokenized treasury max allocation cannot exceed 10000 (100%%)")
	}
	// Add other validation rules as needed
	return nil
}


// Parameter store keys for VaultParams
var (
	KeyCollateralParams    = []byte("CollateralParams")
	KeyVaultMintingEnabled = []byte("VaultMintingEnabled")
)

// VaultParamKeyTable for stablecoin module's vault parameters
func VaultParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters for the stablecoin module.
// These params are for the vault-based CDP system.
func DefaultParams() Params {
	return Params{
		CollateralParams: []CollateralParam{
			{
				Denom:            "uatom",
				LiquidationRatio: sdkmath.LegacyMustNewDecFromStr("1.5"), // 150% collateralization
				StabilityFee:     sdkmath.LegacyMustNewDecFromStr("0.01"), // 1% annual stability fee
				DebtLimit:        sdkmath.NewInt(100_000_000_000_000), // 100 billion ssUSD
				Active:           true,
			},
		},
		VaultMintingEnabled: true,
	}
}

// Validate validates the Params
func (p Params) Validate() error {
	for _, cp := range p.CollateralParams {
		if err := cp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetCollateralParam returns the CollateralParam for a given denom
func (p Params) GetCollateralParam(denom string) (CollateralParam, bool) {
	for _, cp := range p.CollateralParams {
		if cp.Denom == denom {
			return cp, true
		}
	}
	return CollateralParam{}, false
}

// Validate validates the CollateralParam
func (cp CollateralParam) Validate() error {
	if cp.Denom == "" {
		return fmt.Errorf("collateral denom cannot be empty")
	}
	if cp.LiquidationRatio.IsNil() || cp.LiquidationRatio.IsNegative() {
		return fmt.Errorf("liquidation ratio cannot be negative")
	}
	if cp.StabilityFee.IsNil() || cp.StabilityFee.IsNegative() {
		return fmt.Errorf("stability fee cannot be negative")
	}
	if cp.DebtLimit.IsNegative() {
		return fmt.Errorf("debt limit cannot be negative")
	}
	return nil
}

// ParamSetPairs implements the paramtypes.ParamSet interface for Params.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCollateralParams, &p.CollateralParams, validateCollateralParams),
		paramtypes.NewParamSetPair(KeyVaultMintingEnabled, &p.VaultMintingEnabled, validateBool),
	}
}

// ParamSetPairs implements the paramtypes.ParamSet interface for ReserveParams.
func (p *ReserveParams) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinReserveRatioBps, &p.MinReserveRatioBps, validateUint32),
		paramtypes.NewParamSetPair(KeyTargetReserveRatioBps, &p.TargetReserveRatioBps, validateUint32),
		paramtypes.NewParamSetPair(KeyMintFeeBps, &p.MintFeeBps, validateUint32),
		paramtypes.NewParamSetPair(KeyRedeemFeeBps, &p.RedeemFeeBps, validateUint32),
		paramtypes.NewParamSetPair(KeyMinMintAmount, &p.MinMintAmount, validateInt),
		paramtypes.NewParamSetPair(KeyMinRedeemAmount, &p.MinRedeemAmount, validateInt),
		paramtypes.NewParamSetPair(KeyRedemptionDelay, &p.RedemptionDelay, validateDuration),
		paramtypes.NewParamSetPair(KeyMaxDailyMint, &p.MaxDailyMint, validateInt),
		paramtypes.NewParamSetPair(KeyMaxDailyRedeem, &p.MaxDailyRedeem, validateInt),
		paramtypes.NewParamSetPair(KeyTokenizedTreasuries, &p.TokenizedTreasuries, validateTokenizedTreasuries),
		paramtypes.NewParamSetPair(KeyRequireKyc, &p.RequireKyc, validateBool),
		paramtypes.NewParamSetPair(KeyMintPaused, &p.MintPaused, validateBool),
		paramtypes.NewParamSetPair(KeyRedeemPaused, &p.RedeemPaused, validateBool),
	}
}

// Validation functions for param set pairs
func validateCollateralParams(i interface{}) error {
	v, ok := i.([]CollateralParam)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, cp := range v {
		if err := cp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateUint32(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateInt(i interface{}) error {
	v, ok := i.(sdkmath.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("value cannot be negative")
	}
	return nil
}

func validateDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v < 0 {
		return fmt.Errorf("duration cannot be negative")
	}
	return nil
}

func validateTokenizedTreasuries(i interface{}) error {
	v, ok := i.([]TokenizedTreasuryConfig)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, tt := range v {
		if err := tt.Validate(); err != nil {
			return err
		}
	}
	return nil
}