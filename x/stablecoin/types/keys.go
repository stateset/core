package types

const (
	// ModuleName defines the module name
	ModuleName = "stablecoin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stablecoin"

	// ModuleAccountName is the name of the module account
	ModuleAccountName = ModuleName

	// VaultParamspace is the paramstore key for vault parameters
	VaultParamspace = "VaultParams"

	// ReserveParamspace is the paramstore key for reserve parameters
	ReserveParamspace = "ReserveParams"

	// StablecoinDenom is the denomination of the stablecoin (ssUSD)
	StablecoinDenom = "ussUSD"
)

var (
	// Keys for store prefixes
	ReserveParamsKey             = []byte{0x01}
	ReserveKey                   = []byte{0x02}
	NextDepositIDKey             = []byte{0x03}
	ReserveDepositKeyPrefix      = []byte{0x04}
	NextRedemptionIDKey          = []byte{0x05}
	RedemptionRequestKeyPrefix   = []byte{0x06}
	DailyMintStatsKeyPrefix      = []byte{0x07}
	NextAttestationIDKey         = []byte{0x08}
	OffChainAttestationKeyPrefix = []byte{0x09}
	ApprovedAttesterKeyPrefix    = []byte{0x0A}
	LockedReservesKey            = []byte{0x0B}

	// Vault keys
	VaultKeyPrefix = []byte{0x10}
	VaultCountKey  = []byte{0x11}

	// PSM (Peg Stability Module) keys
	PSMConfigKeyPrefix = []byte{0x20}
	PSMStateKeyPrefix  = []byte{0x21}

	// Savings Rate keys
	SavingsParamsKey         = []byte{0x30}
	SavingsDepositKeyPrefix  = []byte{0x31}
	SavingsStatsKey          = []byte{0x32}

	// Dutch Auction keys
	AuctionParamsKey       = []byte{0x40}
	NextAuctionIDKey       = []byte{0x41}
	AuctionKeyPrefix       = []byte{0x42}
	ActiveAuctionKeyPrefix = []byte{0x43}

	// Flash Mint keys
	FlashMintParamsKey       = []byte{0x50}
	FlashMintStatsKey        = []byte{0x51}
	FlashMintSessionKey      = []byte{0x52} // Tracks active flash mint sessions within a tx
)

func DailyMintStatsKey(date string) []byte {
	return append(DailyMintStatsKeyPrefix, []byte(date)...)
}

func ApprovedAttesterKey(addr string) []byte {
	return append(ApprovedAttesterKeyPrefix, []byte(addr)...)
}

// PSM key helpers
func PSMConfigKey(denom string) []byte {
	return append(PSMConfigKeyPrefix, []byte(denom)...)
}

func PSMStateKey(denom string) []byte {
	return append(PSMStateKeyPrefix, []byte(denom)...)
}

// Savings key helpers
func SavingsDepositKey(depositor string) []byte {
	return append(SavingsDepositKeyPrefix, []byte(depositor)...)
}

// Auction key helpers
func AuctionKey(id uint64) []byte {
	bz := make([]byte, 8)
	bz[0] = byte(id >> 56)
	bz[1] = byte(id >> 48)
	bz[2] = byte(id >> 40)
	bz[3] = byte(id >> 32)
	bz[4] = byte(id >> 24)
	bz[5] = byte(id >> 16)
	bz[6] = byte(id >> 8)
	bz[7] = byte(id)
	return append(AuctionKeyPrefix, bz...)
}

func ActiveAuctionKey(id uint64) []byte {
	bz := make([]byte, 8)
	bz[0] = byte(id >> 56)
	bz[1] = byte(id >> 48)
	bz[2] = byte(id >> 40)
	bz[3] = byte(id >> 32)
	bz[4] = byte(id >> 24)
	bz[5] = byte(id >> 16)
	bz[6] = byte(id >> 8)
	bz[7] = byte(id)
	return append(ActiveAuctionKeyPrefix, bz...)
}

// Event Types
const (
	EventTypeVaultCreated         = "vault_created"
	EventTypeCollateralDeposited  = "collateral_deposited"
	EventTypeCollateralWithdrawn  = "collateral_withdrawn"
	EventTypeStablecoinMinted     = "stablecoin_minted"
	EventTypeStablecoinRepaid     = "stablecoin_repaid"
	EventTypeVaultLiquidated      = "vault_liquidated"
	EventTypeReserveDeposit       = "reserve_deposit"
	EventTypeRedemptionRequested  = "redemption_requested"
	EventTypeRedemptionExecuted   = "redemption_executed"
	EventTypeRedemptionCancelled  = "redemption_cancelled"
	EventTypeReserveParamsUpdated = "reserve_params_updated"
	EventTypeReserveAttestation   = "reserve_attestation"
	EventTypeSolvencyEmergency    = "solvency_emergency"

	// PSM Events
	EventTypePSMSwapIn       = "psm_swap_in"
	EventTypePSMSwapOut      = "psm_swap_out"
	EventTypePSMConfigUpdate = "psm_config_update"

	// Savings Events
	EventTypeSavingsDeposit        = "savings_deposit"
	EventTypeSavingsWithdraw       = "savings_withdraw"
	EventTypeSavingsInterestClaim  = "savings_interest_claim"
	EventTypeSavingsParamsUpdate   = "savings_params_update"
	EventTypeSavingsInterestAccrue = "savings_interest_accrue"

	// Auction Events
	EventTypeAuctionCreated   = "auction_created"
	EventTypeAuctionBid       = "auction_bid"
	EventTypeAuctionCompleted = "auction_completed"
	EventTypeAuctionExpired   = "auction_expired"
	EventTypeAuctionCancelled = "auction_cancelled"

	// Flash Mint Events
	EventTypeFlashMint         = "flash_mint"
	EventTypeFlashMintCallback = "flash_mint_callback"
)

// Attribute Keys
const (
	AttributeKeyOwner        = "owner"
	AttributeKeyVaultID      = "vault_id"
	AttributeKeyCollateral   = "collateral"
	AttributeKeyAmount       = "amount"
	AttributeKeyLiquidator   = "liquidator"
	AttributeKeyDepositor    = "depositor"
	AttributeKeyDepositID    = "deposit_id"
	AttributeKeyReserveAsset = "reserve_asset"
	AttributeKeyUsdValue     = "usd_value"
	AttributeKeyPrice        = "price"
	AttributeKeySsusdAmount  = "ssusd_amount"
	AttributeKeyFeeAmount    = "fee_amount"
	AttributeKeyRedemptionID = "redemption_id"
	AttributeKeyReserveRatio = "reserve_ratio"
	AttributeKeyAttester     = "attester"
	AttributeKeyAction       = "action"

	// PSM Attributes
	AttributeKeyInputDenom  = "input_denom"
	AttributeKeyOutputDenom = "output_denom"
	AttributeKeySwapFee     = "swap_fee"

	// Savings Attributes
	AttributeKeyPrincipal       = "principal"
	AttributeKeyInterest        = "interest"
	AttributeKeySavingsRate     = "savings_rate"
	AttributeKeyTotalDeposits   = "total_deposits"

	// Auction Attributes
	AttributeKeyAuctionID         = "auction_id"
	AttributeKeyBidder            = "bidder"
	AttributeKeyStartPrice        = "start_price"
	AttributeKeyEndPrice          = "end_price"
	AttributeKeyCurrentPrice      = "current_price"
	AttributeKeyCollateralSold    = "collateral_sold"
	AttributeKeyDebtRaised        = "debt_raised"
	AttributeKeyDebtToCover       = "debt_to_cover"

	// Flash Mint Attributes
	AttributeKeyFlashMintAmount = "flash_mint_amount"
	AttributeKeyFlashMintFee    = "flash_mint_fee"
	AttributeKeySender          = "sender"
)