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
)

var (
	// Keys for store prefixes
	ReserveParamsKey           = []byte{0x01}
	ReserveKey                 = []byte{0x02}
	NextDepositIDKey           = []byte{0x03}
	ReserveDepositKeyPrefix    = []byte{0x04}
	NextRedemptionIDKey        = []byte{0x05}
	RedemptionRequestKeyPrefix = []byte{0x06}
	DailyMintStatsKeyPrefix    = []byte{0x07}
	NextAttestationIDKey       = []byte{0x08}
	OffChainAttestationKeyPrefix = []byte{0x09}
	ApprovedAttesterKeyPrefix  = []byte{0x0A}
	LockedReservesKey          = []byte{0x0B}

	// Vault keys
	VaultKeyPrefix             = []byte{0x10}
	VaultCountKey              = []byte{0x11}
)

func DailyMintStatsKey(date string) []byte {
	return append(DailyMintStatsKeyPrefix, []byte(date)...)
}

func ApprovedAttesterKey(addr string) []byte {
	return append(ApprovedAttesterKeyPrefix, []byte(addr)...)
}

// Event Types
const (
	EventTypeVaultCreated        = "vault_created"
	EventTypeCollateralDeposited = "collateral_deposited"
	EventTypeCollateralWithdrawn = "collateral_withdrawn"
	EventTypeStablecoinMinted    = "stablecoin_minted"
	EventTypeStablecoinRepaid    = "stablecoin_repaid"
	EventTypeVaultLiquidated     = "vault_liquidated"
	EventTypeReserveDeposit      = "reserve_deposit"
	EventTypeRedemptionRequested = "redemption_requested"
	EventTypeRedemptionExecuted  = "redemption_executed"
	EventTypeRedemptionCancelled = "redemption_cancelled"
	EventTypeReserveParamsUpdated = "reserve_params_updated"
	EventTypeReserveAttestation  = "reserve_attestation"
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
)