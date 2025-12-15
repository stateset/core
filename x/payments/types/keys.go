package types

import "encoding/binary"

const (
	ModuleName         = "payments"
	StoreKey           = ModuleName
	RouterKey          = ModuleName
	ModuleAccountName  = ModuleName
	EventTypeCreated   = "payment_created"
	EventTypeSettled   = "payment_settled"
	EventTypeCancelled = "payment_cancelled"
	AttributeKeyPayer  = "payer"
	AttributeKeyPayee  = "payee"
	AttributeKeyID     = "payment_id"
)

var (
	PaymentKeyPrefix      = []byte{0x01}
	PaymentRouteKeyPrefix = []byte{0x03}
)

func PaymentStoreKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, PaymentKeyPrefix...), bz...)
}
