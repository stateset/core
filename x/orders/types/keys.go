package types

const (
	// ModuleName is the module name constant.
	ModuleName = "orders"

	// StoreKey is the store key string for orders.
	StoreKey = ModuleName

	// RouterKey is the message route for orders.
	RouterKey = ModuleName

	// QuerierRoute is the querier route for orders.
	QuerierRoute = ModuleName

	// ModuleAccountName is the module account name for holding escrow funds.
	ModuleAccountName = "orders_escrow"
)

var (
	// OrderKeyPrefix is the prefix for order storage.
	OrderKeyPrefix = []byte{0x01}

	// OrderByCustomerKeyPrefix indexes orders by customer.
	OrderByCustomerKeyPrefix = []byte{0x02}

	// OrderByMerchantKeyPrefix indexes orders by merchant.
	OrderByMerchantKeyPrefix = []byte{0x03}

	// OrderByStatusKeyPrefix indexes orders by status.
	OrderByStatusKeyPrefix = []byte{0x04}

	// NextOrderIDKey stores the next order ID.
	NextOrderIDKey = []byte{0x05}

	// ParamsKey stores module parameters.
	ParamsKey = []byte{0x06}

	// DisputeKeyPrefix is the prefix for dispute storage.
	DisputeKeyPrefix = []byte{0x07}
)
