package types

const (
	// ModuleName defines the module name
	ModuleName = "invoice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_invoice"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	InvoiceKey      = "Invoice-value-"
	InvoiceCountKey = "Invoice-count-"
)

const (
	SentInvoiceKey      = "SentInvoice-value-"
	SentInvoiceCountKey = "SentInvoice-count-"
)

const (
	TimedoutInvoiceKey      = "TimedoutInvoice-value-"
	TimedoutInvoiceCountKey = "TimedoutInvoice-count-"
)
