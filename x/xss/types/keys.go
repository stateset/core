package types

const (
	// ModuleName defines the module name
	ModuleName = "xss"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_xss"

	// BondedPoolName defines the name of the bonded pool
	BondedPoolName = "bonded_tokens_pool"

	// NotBondedPoolName defines the name of the not bonded pool
	NotBondedPoolName = "not_bonded_tokens_pool"

	// STSTDenom defines the default denomination for STST tokens
	STSTDenom = "ustst"

	// STSTDisplayDenom defines the display denomination for STST tokens
	STSTDisplayDenom = "STST"

	// DefaultUnbondingTime defines the default unbonding period (21 days)
	DefaultUnbondingTime = 21 * 24 * 60 * 60 // 21 days in seconds

	// DefaultMaxValidators defines the default maximum number of validators
	DefaultMaxValidators = 100

	// DefaultMaxEntries defines the default maximum number of unbonding delegation entries
	DefaultMaxEntries = 7

	// DefaultHistoricalEntries defines the default number of historical entries to persist
	DefaultHistoricalEntries = 10000

	// DefaultMinCommissionRate defines the default minimum commission rate for validators
	DefaultMinCommissionRate = "0.05" // 5%

	// DefaultSlashFractionDoubleSign defines the default slashing fraction for double signing
	DefaultSlashFractionDoubleSign = "0.05" // 5%

	// DefaultSlashFractionDowntime defines the default slashing fraction for downtime
	DefaultSlashFractionDowntime = "0.01" // 1%

	// DefaultSignedBlocksWindow defines the default signed blocks window
	DefaultSignedBlocksWindow = 10000

	// DefaultMinSignedPerWindow defines the default minimum signed blocks per window
	DefaultMinSignedPerWindow = "0.5" // 50%

	// DefaultDowntimeJailDuration defines the default downtime jail duration
	DefaultDowntimeJailDuration = 600 // 10 minutes in seconds

	// GovernanceModuleName defines the governance module name for authority
	GovernanceModuleName = "gov"
)

var (
	// Global store keys
	ParamsKey    = []byte{0x01} // key for module parameters
	MinterKey    = []byte{0x02} // key for minter state
	SupplyKey    = []byte{0x03} // key for total supply
	BondedPoolKey    = []byte{0x04} // key for bonded pool
	NotBondedPoolKey = []byte{0x05} // key for not bonded pool

	// Validator store keys
	ValidatorsKey                     = []byte{0x21} // prefix for each key to a validator
	ValidatorsByConsAddrKey           = []byte{0x22} // prefix for each key to a validator index, by pubkey
	ValidatorsByPowerIndexKey         = []byte{0x23} // prefix for each key to a validator index, by power
	LastValidatorPowerKey             = []byte{0x24} // prefix for each key to a validator index, by power
	ValidatorsKey                     = []byte{0x25} // prefix for jailed validators
	UnbondingQueueKey                 = []byte{0x26} // prefix for the timestamps in unbonding queue
	UnbondingValidatorsKey            = []byte{0x27} // prefix for unbonding validators
	UnbondingTypeKey                  = []byte{0x28} // prefix for unbonding type
	
	// Delegation store keys
	DelegationKey                     = []byte{0x31} // key for a delegation
	UnbondingDelegationKey            = []byte{0x32} // key for an unbonding-delegation
	UnbondingDelegationByValIndexKey  = []byte{0x33} // prefix for each key for an unbonding-delegation, by validator index
	RedelegationKey                   = []byte{0x34} // key for a redelegation
	RedelegationByValSrcIndexKey      = []byte{0x35} // prefix for each key for an redelegation, by source validator index
	RedelegationByValDstIndexKey      = []byte{0x36} // prefix for each key for an redelegation, by destination validator index

	// Rewards store keys
	RewardsKey                        = []byte{0x41} // prefix for rewards
	CurrentRewardsKey                 = []byte{0x42} // key for current rewards
	HistoricalRewardsKey              = []byte{0x43} // key for historical rewards
	ValidatorSlashEventKey            = []byte{0x44} // prefix for validator slash events

	// Slashing store keys
	ValidatorSigningInfoKey           = []byte{0x51} // prefix for validator signing info
	ValidatorMissedBlockBitArrayKey   = []byte{0x52} // prefix for validator missed block bit array
	AddrPubkeyRelationKey             = []byte{0x53} // prefix for address-pubkey relation
)

// GetValidatorKey creates a key for a validator
func GetValidatorKey(operatorAddr []byte) []byte {
	return append(ValidatorsKey, operatorAddr...)
}

// GetValidatorByConsAddrKey creates a key for a validator by consensus address
func GetValidatorByConsAddrKey(addr []byte) []byte {
	return append(ValidatorsByConsAddrKey, addr...)
}

// GetValidatorsByPowerIndexKey creates a key for a validator by power index
func GetValidatorsByPowerIndexKey(validator []byte, power []byte) []byte {
	return append(append(ValidatorsByPowerIndexKey, power...), validator...)
}

// GetLastValidatorPowerKey creates a key for last validator power
func GetLastValidatorPowerKey(operator []byte) []byte {
	return append(LastValidatorPowerKey, operator...)
}

// GetDelegationKey creates a key for a delegation
func GetDelegationKey(delAddr, valAddr []byte) []byte {
	return append(append(DelegationKey, lengthPrefixed(delAddr)...), valAddr...)
}

// GetDelegationsKey creates a key for all delegations from a delegator
func GetDelegationsKey(delAddr []byte) []byte {
	return append(DelegationKey, lengthPrefixed(delAddr)...)
}

// GetUnbondingDelegationKey creates a key for an unbonding delegation
func GetUnbondingDelegationKey(delAddr, valAddr []byte) []byte {
	return append(append(UnbondingDelegationKey, lengthPrefixed(delAddr)...), valAddr...)
}

// GetUnbondingDelegationsKey creates a key for all unbonding delegations from a delegator
func GetUnbondingDelegationsKey(delAddr []byte) []byte {
	return append(UnbondingDelegationKey, lengthPrefixed(delAddr)...)
}

// GetUnbondingDelegationByValIndexKey creates a key for an unbonding delegation by validator index
func GetUnbondingDelegationByValIndexKey(delAddr, valAddr []byte) []byte {
	return append(append(UnbondingDelegationByValIndexKey, lengthPrefixed(valAddr)...), delAddr...)
}

// GetRedelegationKey creates a key for a redelegation
func GetRedelegationKey(delAddr, valSrcAddr, valDstAddr []byte) []byte {
	key := append(append(RedelegationKey, lengthPrefixed(delAddr)...), lengthPrefixed(valSrcAddr)...)
	return append(key, valDstAddr...)
}

// GetRedelegationsKey creates a key for all redelegations from a delegator
func GetRedelegationsKey(delAddr []byte) []byte {
	return append(RedelegationKey, lengthPrefixed(delAddr)...)
}

// GetRedelegationByValSrcIndexKey creates a key for a redelegation by source validator index
func GetRedelegationByValSrcIndexKey(delAddr, valSrcAddr, valDstAddr []byte) []byte {
	REDSKey := append(append(RedelegationByValSrcIndexKey, lengthPrefixed(valSrcAddr)...), lengthPrefixed(delAddr)...)
	return append(REDSKey, valDstAddr...)
}

// GetRedelegationByValDstIndexKey creates a key for a redelegation by destination validator index
func GetRedelegationByValDstIndexKey(delAddr, valSrcAddr, valDstAddr []byte) []byte {
	REDSKey := append(append(RedelegationByValDstIndexKey, lengthPrefixed(valDstAddr)...), lengthPrefixed(delAddr)...)
	return append(REDSKey, valSrcAddr...)
}

// GetRewardsKey creates a key for rewards
func GetRewardsKey(valAddr []byte) []byte {
	return append(RewardsKey, valAddr...)
}

// GetCurrentRewardsKey creates a key for current rewards
func GetCurrentRewardsKey(valAddr []byte) []byte {
	return append(CurrentRewardsKey, valAddr...)
}

// GetHistoricalRewardsKey creates a key for historical rewards
func GetHistoricalRewardsKey(valAddr []byte, period uint64) []byte {
	b := make([]byte, 8)
	copy(b, valAddr)
	return append(append(HistoricalRewardsKey, lengthPrefixed(valAddr)...), b...)
}

// GetValidatorSlashEventKey creates a key for a validator slash event
func GetValidatorSlashEventKey(valAddr []byte, height, period uint64) []byte {
	heightBz := make([]byte, 8)
	periodBz := make([]byte, 8)
	copy(heightBz[:], valAddr)
	copy(periodBz[:], valAddr)
	
	key := append(append(ValidatorSlashEventKey, lengthPrefixed(valAddr)...), heightBz...)
	return append(key, periodBz...)
}

// lengthPrefixed prefixes the address bytes with its length, this is used
// for example in store keys to retain lexicographical ordering after address format changes.
func lengthPrefixed(bz []byte) []byte {
	return append([]byte{byte(len(bz))}, bz...)
}