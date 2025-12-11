package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ReserveAssetType defines the type of reserve asset
type ReserveAssetType string

const (
	// ReserveAssetCash represents cash reserves (USD held at custodian)
	ReserveAssetCash ReserveAssetType = "cash"
	// ReserveAssetTBill represents US Treasury Bills (< 1 year)
	ReserveAssetTBill ReserveAssetType = "t_bill"
	// ReserveAssetTNote represents US Treasury Notes (1-10 years)
	ReserveAssetTNote ReserveAssetType = "t_note"
	// ReserveAssetTBond represents US Treasury Bonds (> 10 years)
	ReserveAssetTBond ReserveAssetType = "t_bond"
	// ReserveAssetTokenizedTreasury represents on-chain tokenized US Treasuries
	ReserveAssetTokenizedTreasury ReserveAssetType = "tokenized_treasury"
	// ReserveAssetRepoAgreement represents overnight repo agreements
	ReserveAssetRepoAgreement ReserveAssetType = "repo"
	// ReserveAssetMMF represents money market funds (government only)
	ReserveAssetMMF ReserveAssetType = "mmf"
)

// ApprovedReserveAssets returns all approved reserve asset types
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

// IsApprovedReserveAsset checks if an asset type is approved
func IsApprovedReserveAsset(assetType ReserveAssetType) bool {
	for _, approved := range ApprovedReserveAssets() {
		if approved == assetType {
			return true
		}
	}
	return false
}

// TokenizedTreasuryConfig defines approved tokenized treasury tokens
type TokenizedTreasuryConfig struct {
	// Denom is the on-chain token denomination
	Denom string `json:"denom" yaml:"denom"`
	// Issuer is the authorized issuer (e.g., "ondo", "backed", "openeden")
	Issuer string `json:"issuer" yaml:"issuer"`
	// UnderlyingType is the underlying treasury type
	UnderlyingType ReserveAssetType `json:"underlying_type" yaml:"underlying_type"`
	// Active indicates if this token is currently accepted
	Active bool `json:"active" yaml:"active"`
	// Haircut in basis points (e.g., 100 = 1% haircut for safety buffer)
	HaircutBps uint32 `json:"haircut_bps" yaml:"haircut_bps"`
	// MaxAllocation in basis points (max % of reserves in this token)
	MaxAllocationBps uint32 `json:"max_allocation_bps" yaml:"max_allocation_bps"`
	// OracleDenom for price feed (if different from denom)
	OracleDenom string `json:"oracle_denom" yaml:"oracle_denom"`
}

// ValidateBasic performs basic validation
func (c TokenizedTreasuryConfig) ValidateBasic() error {
	if c.Denom == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "denom cannot be empty")
	}
	if c.Issuer == "" {
		return errorsmod.Wrap(ErrInvalidReserve, "issuer cannot be empty")
	}
	if c.HaircutBps > 5000 { // Max 50% haircut
		return errorsmod.Wrap(ErrInvalidReserve, "haircut cannot exceed 50%")
	}
	if c.MaxAllocationBps > 10000 {
		return errorsmod.Wrap(ErrInvalidReserve, "max allocation cannot exceed 100%")
	}
	return nil
}

// ReserveParams defines parameters for reserve management
type ReserveParams struct {
	// MinReserveRatio in basis points (10000 = 100% backing required)
	MinReserveRatioBps uint32 `json:"min_reserve_ratio_bps" yaml:"min_reserve_ratio_bps"`
	// TargetReserveRatio in basis points (target buffer, e.g., 10200 = 102%)
	TargetReserveRatioBps uint32 `json:"target_reserve_ratio_bps" yaml:"target_reserve_ratio_bps"`
	// MintFee in basis points
	MintFeeBps uint32 `json:"mint_fee_bps" yaml:"mint_fee_bps"`
	// RedeemFee in basis points
	RedeemFeeBps uint32 `json:"redeem_fee_bps" yaml:"redeem_fee_bps"`
	// MinMintAmount minimum amount to mint
	MinMintAmount sdkmath.Int `json:"min_mint_amount" yaml:"min_mint_amount"`
	// MinRedeemAmount minimum amount to redeem
	MinRedeemAmount sdkmath.Int `json:"min_redeem_amount" yaml:"min_redeem_amount"`
	// RedemptionDelay time before redemption is processed
	RedemptionDelay time.Duration `json:"redemption_delay" yaml:"redemption_delay"`
	// MaxDailyMint maximum amount that can be minted per day
	MaxDailyMint sdkmath.Int `json:"max_daily_mint" yaml:"max_daily_mint"`
	// MaxDailyRedeem maximum amount that can be redeemed per day
	MaxDailyRedeem sdkmath.Int `json:"max_daily_redeem" yaml:"max_daily_redeem"`
	// TokenizedTreasuries approved tokenized treasury tokens
	TokenizedTreasuries []TokenizedTreasuryConfig `json:"tokenized_treasuries" yaml:"tokenized_treasuries"`
	// RequireKYC whether KYC is required for mint/redeem
	RequireKYC bool `json:"require_kyc" yaml:"require_kyc"`
	// MintPaused global mint pause
	MintPaused bool `json:"mint_paused" yaml:"mint_paused"`
	// RedeemPaused global redeem pause
	RedeemPaused bool `json:"redeem_paused" yaml:"redeem_paused"`
}

// DefaultReserveParams returns default reserve parameters
func DefaultReserveParams() ReserveParams {
	return ReserveParams{
		MinReserveRatioBps:    10000,                               // 100% backed
		TargetReserveRatioBps: 10200,                               // 102% target (2% buffer)
		MintFeeBps:            10,                                  // 0.1% mint fee
		RedeemFeeBps:          10,                                  // 0.1% redeem fee
		MinMintAmount:         sdkmath.NewInt(100_000_000),         // 100 ssUSD minimum
		MinRedeemAmount:       sdkmath.NewInt(100_000_000),         // 100 ssUSD minimum
		RedemptionDelay:       0,                                   // Instant for tokenized
		MaxDailyMint:          sdkmath.NewInt(100_000_000_000_000), // 100M ssUSD/day
		MaxDailyRedeem:        sdkmath.NewInt(100_000_000_000_000), // 100M ssUSD/day
		TokenizedTreasuries:   DefaultTokenizedTreasuries(),
		RequireKYC:            true,
		MintPaused:            false,
		RedeemPaused:          false,
	}
}

// DefaultTokenizedTreasuries returns default approved tokenized treasury tokens
func DefaultTokenizedTreasuries() []TokenizedTreasuryConfig {
	return []TokenizedTreasuryConfig{
		{
			Denom:            "usdy", // Ondo USDY
			Issuer:           "ondo",
			UnderlyingType:   ReserveAssetTBill,
			Active:           true,
			HaircutBps:       50,   // 0.5% haircut
			MaxAllocationBps: 5000, // Max 50% of reserves
			OracleDenom:      "usdy",
		},
		{
			Denom:            "stbt", // Matrixdock STBT
			Issuer:           "matrixdock",
			UnderlyingType:   ReserveAssetTBill,
			Active:           true,
			HaircutBps:       50,
			MaxAllocationBps: 3000, // Max 30%
			OracleDenom:      "stbt",
		},
		{
			Denom:            "ousg", // Ondo US Government Bond
			Issuer:           "ondo",
			UnderlyingType:   ReserveAssetTBond,
			Active:           true,
			HaircutBps:       100, // 1% haircut for duration risk
			MaxAllocationBps: 3000,
			OracleDenom:      "ousg",
		},
		{
			Denom:            "tbill", // OpenEden T-Bill
			Issuer:           "openeden",
			UnderlyingType:   ReserveAssetTBill,
			Active:           true,
			HaircutBps:       50,
			MaxAllocationBps: 4000,
			OracleDenom:      "tbill",
		},
		{
			Denom:            "usdc", // USDC as cash equivalent
			Issuer:           "circle",
			UnderlyingType:   ReserveAssetCash,
			Active:           true,
			HaircutBps:       0,    // No haircut for cash
			MaxAllocationBps: 2000, // Max 20% in cash
			OracleDenom:      "usdc",
		},
	}
}

// ValidateBasic validates reserve params
func (p ReserveParams) ValidateBasic() error {
	if p.MinReserveRatioBps < 10000 {
		return errorsmod.Wrap(ErrInvalidReserve, "minimum reserve ratio must be at least 100%")
	}
	if p.TargetReserveRatioBps < p.MinReserveRatioBps {
		return errorsmod.Wrap(ErrInvalidReserve, "target ratio must be >= minimum ratio")
	}
	if p.MintFeeBps > 1000 { // Max 10%
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

// GetTokenizedTreasury returns config for a tokenized treasury denom
func (p ReserveParams) GetTokenizedTreasury(denom string) (TokenizedTreasuryConfig, bool) {
	for _, tt := range p.TokenizedTreasuries {
		if tt.Denom == denom {
			return tt, true
		}
	}
	return TokenizedTreasuryConfig{}, false
}

// Reserve represents the on-chain reserve backing
type Reserve struct {
	// TotalDeposited total tokenized treasuries deposited
	TotalDeposited sdk.Coins `json:"total_deposited" yaml:"total_deposited"`
	// TotalValue USD value of reserves (from oracle)
	TotalValue sdkmath.Int `json:"total_value" yaml:"total_value"`
	// TotalMinted total ssUSD minted against reserves
	TotalMinted sdkmath.Int `json:"total_minted" yaml:"total_minted"`
	// LastUpdatedHeight block height of last update (kept as `last_updated` for backward compatibility)
	LastUpdatedHeight int64 `json:"last_updated" yaml:"last_updated"`
	// LastUpdatedTime wall-clock time of last update
	LastUpdatedTime time.Time `json:"last_updated_time" yaml:"last_updated_time"`
}

// GetReserveRatio returns the current reserve ratio in basis points
func (r Reserve) GetReserveRatio() uint32 {
	if r.TotalMinted.IsZero() {
		return 10000 // 100% if nothing minted
	}
	// (TotalValue * 10000) / TotalMinted
	ratio := r.TotalValue.Mul(sdkmath.NewInt(10000)).Quo(r.TotalMinted)
	if ratio.GT(sdkmath.NewInt(100000)) { // Cap at 1000%
		return 100000
	}
	return uint32(ratio.Int64())
}

// IsHealthy checks if reserve ratio is above minimum
func (r Reserve) IsHealthy(minRatioBps uint32) bool {
	return r.GetReserveRatio() >= minRatioBps
}

// ReserveDeposit represents a deposit of tokenized treasuries
type ReserveDeposit struct {
	Id          uint64        `json:"id" yaml:"id"`
	Depositor   string        `json:"depositor" yaml:"depositor"`
	Amount      sdk.Coin      `json:"amount" yaml:"amount"`
	UsdValue    sdkmath.Int   `json:"usd_value" yaml:"usd_value"`
	SsusdMinted sdkmath.Int   `json:"ssusd_minted" yaml:"ssusd_minted"`
	DepositedAt time.Time     `json:"deposited_at" yaml:"deposited_at"`
	Status      DepositStatus `json:"status" yaml:"status"`
}

// DepositStatus represents the status of a reserve deposit
type DepositStatus string

const (
	DepositStatusActive    DepositStatus = "active"
	DepositStatusRedeeming DepositStatus = "redeeming"
	DepositStatusRedeemed  DepositStatus = "redeemed"
)

// RedemptionRequest represents a request to redeem ssUSD for reserves
type RedemptionRequest struct {
	Id              uint64       `json:"id" yaml:"id"`
	Requester       string       `json:"requester" yaml:"requester"`
	SsusdAmount     sdkmath.Int  `json:"ssusd_amount" yaml:"ssusd_amount"`
	OutputDenom     string       `json:"output_denom" yaml:"output_denom"`
	RequestedAt     time.Time    `json:"requested_at" yaml:"requested_at"`
	ExecutableAfter time.Time    `json:"executable_after" yaml:"executable_after"`
	Status          RedeemStatus `json:"status" yaml:"status"`
	ExecutedAt      time.Time    `json:"executed_at,omitempty" yaml:"executed_at,omitempty"`
	OutputAmount    sdk.Coin     `json:"output_amount,omitempty" yaml:"output_amount,omitempty"`
}

// RedeemStatus represents the status of a redemption request
type RedeemStatus string

const (
	RedeemStatusPending   RedeemStatus = "pending"
	RedeemStatusExecuted  RedeemStatus = "executed"
	RedeemStatusCancelled RedeemStatus = "cancelled"
)

// DailyMintStats tracks daily minting/redemption
type DailyMintStats struct {
	Date          string      `json:"date" yaml:"date"` // YYYY-MM-DD
	TotalMinted   sdkmath.Int `json:"total_minted" yaml:"total_minted"`
	TotalRedeemed sdkmath.Int `json:"total_redeemed" yaml:"total_redeemed"`
}

// OffChainReserveAttestation represents off-chain reserve attestation
type OffChainReserveAttestation struct {
	Id              uint64      `json:"id" yaml:"id"`
	Attester        string      `json:"attester" yaml:"attester"`
	TotalCash       sdkmath.Int `json:"total_cash" yaml:"total_cash"`
	TotalTBills     sdkmath.Int `json:"total_tbills" yaml:"total_tbills"`
	TotalTNotes     sdkmath.Int `json:"total_tnotes" yaml:"total_tnotes"`
	TotalTBonds     sdkmath.Int `json:"total_tbonds" yaml:"total_tbonds"`
	TotalRepos      sdkmath.Int `json:"total_repos" yaml:"total_repos"`
	TotalMMF        sdkmath.Int `json:"total_mmf" yaml:"total_mmf"`
	TotalValue      sdkmath.Int `json:"total_value" yaml:"total_value"`
	CustodianName   string      `json:"custodian_name" yaml:"custodian_name"`
	AuditFirm       string      `json:"audit_firm" yaml:"audit_firm"`
	ReportDate      time.Time   `json:"report_date" yaml:"report_date"`
	AttestationHash string      `json:"attestation_hash" yaml:"attestation_hash"`
	Timestamp       time.Time   `json:"timestamp" yaml:"timestamp"`
}

// ValidateBasic validates the attestation
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

// TotalReserves calculates total reserve value (on-chain + off-chain)
type TotalReserves struct {
	OnChainValue       sdkmath.Int `json:"on_chain_value" yaml:"on_chain_value"`
	OffChainValue      sdkmath.Int `json:"off_chain_value" yaml:"off_chain_value"`
	TotalValue         sdkmath.Int `json:"total_value" yaml:"total_value"`
	TotalSupply        sdkmath.Int `json:"total_supply" yaml:"total_supply"`
	ReserveRatioBps    uint32      `json:"reserve_ratio_bps" yaml:"reserve_ratio_bps"`
	LastOnChainUpdate  time.Time   `json:"last_on_chain_update" yaml:"last_on_chain_update"`
	LastOffChainUpdate time.Time   `json:"last_off_chain_update" yaml:"last_off_chain_update"`
}

// CalculateReserveRatio calculates the reserve ratio
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
