package types

import (
	"encoding/json"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// JSON marshaling helpers for extended types
func MarshalJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MustMarshalJSON(v interface{}) []byte {
	bz, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bz
}

func UnmarshalJSON(bz []byte, v interface{}) error {
	return json.Unmarshal(bz, v)
}

func MustUnmarshalJSON(bz []byte, v interface{}) {
	if err := json.Unmarshal(bz, v); err != nil {
		panic(err)
	}
}

// ============================================================================
// PSM (Peg Stability Module) Types
// ============================================================================

// PSMConfig defines configuration for a PSM asset (e.g., USDC, USDT).
type PSMConfig struct {
	// Denom is the token denomination (e.g., "ibc/USDC_HASH").
	Denom string `json:"denom"`
	// Active indicates whether the PSM accepts this asset.
	Active bool `json:"active"`
	// MintFeeBps is the fee charged when minting ssUSD from this asset (in basis points).
	MintFeeBps uint32 `json:"mint_fee_bps"`
	// RedeemFeeBps is the fee charged when redeeming ssUSD for this asset (in basis points).
	RedeemFeeBps uint32 `json:"redeem_fee_bps"`
	// DebtCeiling is the maximum ssUSD that can be minted against this asset.
	DebtCeiling sdkmath.Int `json:"debt_ceiling"`
	// OracleDenom is the oracle price feed identifier for this asset.
	OracleDenom string `json:"oracle_denom"`
}

// PSMState tracks the current state of a PSM asset.
type PSMState struct {
	// Denom is the token denomination.
	Denom string `json:"denom"`
	// TotalDeposited is the total amount of this asset in the PSM.
	TotalDeposited sdkmath.Int `json:"total_deposited"`
	// TotalMinted is the total ssUSD minted against this asset.
	TotalMinted sdkmath.Int `json:"total_minted"`
}

// ============================================================================
// Savings Rate Types
// ============================================================================

// SavingsParams defines parameters for the ssUSD savings rate.
type SavingsParams struct {
	// Enabled indicates whether the savings rate is active.
	Enabled bool `json:"enabled"`
	// SavingsRateBps is the annual percentage yield in basis points (e.g., 500 = 5% APY).
	SavingsRateBps uint32 `json:"savings_rate_bps"`
	// MinDeposit is the minimum ssUSD deposit amount.
	MinDeposit sdkmath.Int `json:"min_deposit"`
	// AccrualIntervalSeconds is how often interest accrues (in seconds).
	AccrualIntervalSeconds int64 `json:"accrual_interval_seconds"`
}

// SavingsDeposit represents a user's savings deposit.
type SavingsDeposit struct {
	// Depositor is the address of the depositor.
	Depositor string `json:"depositor"`
	// Principal is the original deposited amount.
	Principal sdkmath.Int `json:"principal"`
	// AccruedInterest is the accumulated interest.
	AccruedInterest sdkmath.Int `json:"accrued_interest"`
	// LastAccrualTime is the timestamp of the last interest accrual.
	LastAccrualTime time.Time `json:"last_accrual_time"`
	// DepositedAt is when the deposit was created.
	DepositedAt time.Time `json:"deposited_at"`
}

// SavingsStats tracks global savings statistics.
type SavingsStats struct {
	// TotalDeposits is the total ssUSD in savings.
	TotalDeposits sdkmath.Int `json:"total_deposits"`
	// TotalInterestPaid is the cumulative interest paid out.
	TotalInterestPaid sdkmath.Int `json:"total_interest_paid"`
	// DepositorCount is the number of unique depositors.
	DepositorCount uint64 `json:"depositor_count"`
}

// ============================================================================
// Dutch Auction Types
// ============================================================================

// AuctionStatus represents the status of a Dutch auction.
type AuctionStatus int32

const (
	AuctionStatus_AUCTION_STATUS_UNSPECIFIED AuctionStatus = 0
	AuctionStatus_AUCTION_STATUS_ACTIVE      AuctionStatus = 1
	AuctionStatus_AUCTION_STATUS_COMPLETED   AuctionStatus = 2
	AuctionStatus_AUCTION_STATUS_EXPIRED     AuctionStatus = 3
	AuctionStatus_AUCTION_STATUS_CANCELLED   AuctionStatus = 4
)

// DutchAuction represents a Dutch auction for liquidated collateral.
type DutchAuction struct {
	// Id is the unique auction identifier.
	Id uint64 `json:"id"`
	// VaultId is the ID of the liquidated vault.
	VaultId uint64 `json:"vault_id"`
	// Owner is the original vault owner.
	Owner string `json:"owner"`
	// Collateral is the collateral being auctioned.
	Collateral sdk.Coin `json:"collateral"`
	// DebtToCover is the ssUSD debt that must be covered.
	DebtToCover sdkmath.Int `json:"debt_to_cover"`
	// StartPrice is the initial price (in ssUSD per unit of collateral).
	StartPrice sdkmath.LegacyDec `json:"start_price"`
	// EndPrice is the minimum price (floor).
	EndPrice sdkmath.LegacyDec `json:"end_price"`
	// StartedAt is when the auction started.
	StartedAt time.Time `json:"started_at"`
	// Duration is the auction duration.
	Duration time.Duration `json:"duration"`
	// Status is the current auction status.
	Status AuctionStatus `json:"status"`
	// CollateralSold is the amount of collateral already sold.
	CollateralSold sdkmath.Int `json:"collateral_sold"`
	// DebtRaised is the ssUSD raised so far.
	DebtRaised sdkmath.Int `json:"debt_raised"`
}

// AuctionParams defines parameters for Dutch auctions.
type AuctionParams struct {
	// Enabled indicates whether Dutch auctions are active.
	Enabled bool `json:"enabled"`
	// Duration is the default auction duration.
	Duration time.Duration `json:"duration"`
	// StartPriceMultiplierBps is the multiplier for starting price (e.g., 13000 = 130% of oracle price).
	StartPriceMultiplierBps uint32 `json:"start_price_multiplier_bps"`
	// EndPriceMultiplierBps is the multiplier for ending price (e.g., 8000 = 80% of oracle price).
	EndPriceMultiplierBps uint32 `json:"end_price_multiplier_bps"`
	// LiquidationPenaltyBps is the penalty charged on liquidation (e.g., 1300 = 13%).
	LiquidationPenaltyBps uint32 `json:"liquidation_penalty_bps"`
}

// ============================================================================
// Flash Mint Types
// ============================================================================

// FlashMintParams defines parameters for flash minting.
type FlashMintParams struct {
	// Enabled indicates whether flash minting is active.
	Enabled bool `json:"enabled"`
	// FeeBps is the flash mint fee in basis points (e.g., 9 = 0.09%).
	FeeBps uint32 `json:"fee_bps"`
	// MaxFlashMint is the maximum amount that can be flash minted per tx.
	MaxFlashMint sdkmath.Int `json:"max_flash_mint"`
}

// FlashMintStats tracks flash mint statistics.
type FlashMintStats struct {
	// TotalFlashMinted is the cumulative flash mint volume.
	TotalFlashMinted sdkmath.Int `json:"total_flash_minted"`
	// TotalFeesCollected is the cumulative fees collected.
	TotalFeesCollected sdkmath.Int `json:"total_fees_collected"`
}

// ============================================================================
// Message Types (Request/Response)
// ============================================================================

// PSM Messages
type MsgPSMSwapIn struct {
	Sender string   `json:"sender"`
	Amount sdk.Coin `json:"amount"`
}

func (msg *MsgPSMSwapIn) ValidateBasic() error {
	if msg.Sender == "" {
		return ErrInvalidReserve
	}
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return ErrInvalidAmount
	}
	return nil
}

type MsgPSMSwapInResponse struct {
	SsusdMinted string `json:"ssusd_minted"`
	FeeCharged  string `json:"fee_charged"`
}

type MsgPSMSwapOut struct {
	Sender      string `json:"sender"`
	SsusdAmount string `json:"ssusd_amount"`
	OutputDenom string `json:"output_denom"`
}

func (msg *MsgPSMSwapOut) ValidateBasic() error {
	if msg.Sender == "" {
		return ErrInvalidReserve
	}
	if msg.SsusdAmount == "" {
		return ErrInvalidAmount
	}
	if msg.OutputDenom == "" {
		return ErrInvalidReserve
	}
	return nil
}

type MsgPSMSwapOutResponse struct {
	OutputAmount string `json:"output_amount"`
	FeeCharged   string `json:"fee_charged"`
}

type MsgUpdatePSMConfig struct {
	Authority string      `json:"authority"`
	Configs   []PSMConfig `json:"configs"`
}

func (msg *MsgUpdatePSMConfig) ValidateBasic() error {
	if msg.Authority == "" {
		return ErrUnauthorized
	}
	return nil
}

type MsgUpdatePSMConfigResponse struct{}

// Savings Messages
type MsgDepositSavings struct {
	Depositor string `json:"depositor"`
	Amount    string `json:"amount"`
}

func (msg *MsgDepositSavings) ValidateBasic() error {
	if msg.Depositor == "" {
		return ErrInvalidReserve
	}
	if msg.Amount == "" {
		return ErrInvalidAmount
	}
	return nil
}

type MsgDepositSavingsResponse struct {
	TotalDeposit string `json:"total_deposit"`
}

type MsgWithdrawSavings struct {
	Depositor string `json:"depositor"`
	Amount    string `json:"amount"`
}

func (msg *MsgWithdrawSavings) ValidateBasic() error {
	if msg.Depositor == "" {
		return ErrInvalidReserve
	}
	if msg.Amount == "" {
		return ErrInvalidAmount
	}
	return nil
}

type MsgWithdrawSavingsResponse struct {
	AmountWithdrawn string `json:"amount_withdrawn"`
	InterestEarned  string `json:"interest_earned"`
}

type MsgClaimSavingsInterest struct {
	Depositor string `json:"depositor"`
}

func (msg *MsgClaimSavingsInterest) ValidateBasic() error {
	if msg.Depositor == "" {
		return ErrInvalidReserve
	}
	return nil
}

type MsgClaimSavingsInterestResponse struct {
	InterestClaimed string `json:"interest_claimed"`
}

type MsgUpdateSavingsParams struct {
	Authority string        `json:"authority"`
	Params    SavingsParams `json:"params"`
}

func (msg *MsgUpdateSavingsParams) ValidateBasic() error {
	if msg.Authority == "" {
		return ErrUnauthorized
	}
	return nil
}

type MsgUpdateSavingsParamsResponse struct{}

// Auction Messages
type MsgBidAuction struct {
	Bidder              string `json:"bidder"`
	AuctionId           uint64 `json:"auction_id"`
	MaxCollateralAmount string `json:"max_collateral_amount"`
	MaxSsusdToSpend     string `json:"max_ssusd_to_spend"`
}

func (msg *MsgBidAuction) ValidateBasic() error {
	if msg.Bidder == "" {
		return ErrInvalidReserve
	}
	if msg.AuctionId == 0 {
		return ErrInvalidVault
	}
	return nil
}

type MsgBidAuctionResponse struct {
	CollateralPurchased string `json:"collateral_purchased"`
	SsusdSpent          string `json:"ssusd_spent"`
	PricePerUnit        string `json:"price_per_unit"`
}

type MsgUpdateAuctionParams struct {
	Authority string        `json:"authority"`
	Params    AuctionParams `json:"params"`
}

func (msg *MsgUpdateAuctionParams) ValidateBasic() error {
	if msg.Authority == "" {
		return ErrUnauthorized
	}
	return nil
}

type MsgUpdateAuctionParamsResponse struct{}

// Flash Mint Messages
type MsgFlashMint struct {
	Sender       string `json:"sender"`
	Amount       string `json:"amount"`
	CallbackData []byte `json:"callback_data"`
}

func (msg *MsgFlashMint) ValidateBasic() error {
	if msg.Sender == "" {
		return ErrInvalidReserve
	}
	if msg.Amount == "" {
		return ErrInvalidAmount
	}
	return nil
}

type MsgFlashMintResponse struct {
	AmountMinted string `json:"amount_minted"`
	FeePaid      string `json:"fee_paid"`
}

type MsgFlashMintCallback struct {
	Sender         string `json:"sender"`
	AmountToReturn string `json:"amount_to_return"`
}

func (msg *MsgFlashMintCallback) ValidateBasic() error {
	if msg.Sender == "" {
		return ErrInvalidReserve
	}
	if msg.AmountToReturn == "" {
		return ErrInvalidAmount
	}
	return nil
}

type MsgFlashMintCallbackResponse struct{}

type MsgUpdateFlashMintParams struct {
	Authority string          `json:"authority"`
	Params    FlashMintParams `json:"params"`
}

func (msg *MsgUpdateFlashMintParams) ValidateBasic() error {
	if msg.Authority == "" {
		return ErrUnauthorized
	}
	return nil
}

type MsgUpdateFlashMintParamsResponse struct{}
