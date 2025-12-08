package types

import "encoding/binary"

const (
	ModuleName        = "stablecoin"
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	ModuleAccountName = ModuleName
	StablecoinDenom   = "ssusd"

	EventTypeVaultCreated        = "vault_created"
	EventTypeCollateralDeposited = "collateral_deposited"
	EventTypeCollateralWithdrawn = "collateral_withdrawn"
	EventTypeStablecoinMinted    = "stablecoin_minted"
	EventTypeStablecoinRepaid    = "stablecoin_repaid"
	EventTypeVaultLiquidated     = "vault_liquidated"

	AttributeKeyOwner      = "owner"
	AttributeKeyVaultID    = "vault_id"
	AttributeKeyAmount     = "amount"
	AttributeKeyCollateral = "collateral"
	AttributeKeyLiquidator = "liquidator"
)

var (
	VaultKeyPrefix = []byte{0x01}
)

func VaultStoreKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, VaultKeyPrefix...), bz...)
}
