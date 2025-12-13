package types

import "encoding/binary"

const (
	ModuleName        = "stablecoin"
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	ModuleAccountName = ModuleName
	StablecoinDenom   = "ssusd"

	// Legacy vault events (for backwards compatibility)
	EventTypeVaultCreated        = "vault_created"
	EventTypeCollateralDeposited = "collateral_deposited"
	EventTypeCollateralWithdrawn = "collateral_withdrawn"
	EventTypeStablecoinMinted    = "stablecoin_minted"
	EventTypeStablecoinRepaid    = "stablecoin_repaid"
	EventTypeVaultLiquidated     = "vault_liquidated"

	// Reserve-backed stablecoin events
	EventTypeReserveDeposit       = "reserve_deposit"
	EventTypeReserveMint          = "reserve_mint"
	EventTypeRedemptionRequested  = "redemption_requested"
	EventTypeRedemptionExecuted   = "redemption_executed"
	EventTypeRedemptionCancelled  = "redemption_cancelled"
	EventTypeReserveAttestation   = "reserve_attestation"
	EventTypeReserveParamsUpdated = "reserve_params_updated"

	// Attribute keys
	AttributeKeyOwner        = "owner"
	AttributeKeyVaultID      = "vault_id"
	AttributeKeyAmount       = "amount"
	AttributeKeyCollateral   = "collateral"
	AttributeKeyLiquidator   = "liquidator"
	AttributeKeyDepositor    = "depositor"
	AttributeKeyDepositID    = "deposit_id"
	AttributeKeyRedemptionID = "redemption_id"
	AttributeKeyReserveAsset = "reserve_asset"
	AttributeKeyUsdValue     = "usd_value"
	AttributeKeySsusdAmount  = "ssusd_amount"
	AttributeKeyReserveRatio = "reserve_ratio"
	AttributeKeyAttester     = "attester"
	AttributeKeyFeeAmount    = "fee_amount"
)

var (
	// Legacy vault storage
	VaultKeyPrefix = []byte{0x01}

	// Reserve-backed storage prefixes
	ReserveParamsKey             = []byte{0x10}
	ReserveKey                   = []byte{0x11}
	ReserveDepositKeyPrefix      = []byte{0x12}
	RedemptionRequestKeyPrefix   = []byte{0x13}
	DailyMintStatsKeyPrefix      = []byte{0x14}
	OffChainAttestationKeyPrefix = []byte{0x15}
	NextDepositIDKey             = []byte{0x16}
	NextRedemptionIDKey          = []byte{0x17}
	NextAttestationIDKey         = []byte{0x18}
	ApprovedAttesterKeyPrefix    = []byte{0x19}
	LockedReservesKey            = []byte{0x1A}
)

func VaultStoreKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, VaultKeyPrefix...), bz...)
}

func ReserveDepositKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, ReserveDepositKeyPrefix...), bz...)
}

func RedemptionRequestKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, RedemptionRequestKeyPrefix...), bz...)
}

func OffChainAttestationKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, OffChainAttestationKeyPrefix...), bz...)
}

func DailyMintStatsKey(date string) []byte {
	return append(append([]byte{}, DailyMintStatsKeyPrefix...), []byte(date)...)
}

func ApprovedAttesterKey(addr string) []byte {
	return append(append([]byte{}, ApprovedAttesterKeyPrefix...), []byte(addr)...)
}
