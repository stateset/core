package types

const (
	ModuleName   = "compliance"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName

	EventTypeProfileUpserted   = "profile_upserted"
	EventTypeProfileSanctioned = "profile_sanctioned"

	AttributeKeyAddress   = "address"
	AttributeKeyAuthority = "authority"
	AttributeKeySanction  = "sanction"
)

var (
	ProfileKeyPrefix = []byte{0x01}
)
