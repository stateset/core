package types

const (
	// ModuleName defines the module name
	ModuleName = "agreement"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_agreement"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SentAgreementKey      = "SentAgreement-value-"
	SentAgreementCountKey = "SentAgreement-count-"
)

const (
	TimedoutAgreementKey      = "TimedoutAgreement-value-"
	TimedoutAgreementCountKey = "TimedoutAgreement-count-"
)

const (
	AgreementKey      = "Agreement-value-"
	AgreementCountKey = "Agreement-count-"
)
