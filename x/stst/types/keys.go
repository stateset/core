package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "stst"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stst"

	// DefaultTokenDenom defines the default token denomination for STST
	DefaultTokenDenom = "stst"
)

// KVStore keys
var (
	// ParamsKey defines the key to store the module parameters
	ParamsKey = []byte{0x01}

	// StakingStateKey defines the key to store the staking state
	StakingStateKey = []byte{0x02}

	// FeeBurnStateKey defines the key to store the fee burn state
	FeeBurnStateKey = []byte{0x03}

	// VestingScheduleKeyPrefix defines the prefix for vesting schedule keys
	VestingScheduleKeyPrefix = []byte{0x04}

	// StakerInfoKeyPrefix defines the prefix for staker info keys
	StakerInfoKeyPrefix = []byte{0x05}

	// DelegationKeyPrefix defines the prefix for delegation keys
	DelegationKeyPrefix = []byte{0x06}

	// ProposalKeyPrefix defines the prefix for proposal keys
	ProposalKeyPrefix = []byte{0x07}

	// VoteKeyPrefix defines the prefix for vote keys
	VoteKeyPrefix = []byte{0x08}

	// BurnRateHistoryKeyPrefix defines the prefix for burn rate history keys
	BurnRateHistoryKeyPrefix = []byte{0x09}

	// NextProposalIDKey defines the key to store the next proposal ID
	NextProposalIDKey = []byte{0x0A}

	// UnstakingQueueKeyPrefix defines the prefix for unstaking queue keys
	UnstakingQueueKeyPrefix = []byte{0x0B}
)

// VestingScheduleKey returns the key for a vesting schedule by category
func VestingScheduleKey(category string) []byte {
	return append(VestingScheduleKeyPrefix, []byte(category)...)
}

// StakerInfoKey returns the key for staker info by staker address
func StakerInfoKey(stakerAddr string) []byte {
	return append(StakerInfoKeyPrefix, []byte(stakerAddr)...)
}

// DelegationKey returns the key for a delegation by staker and validator address
func DelegationKey(stakerAddr, validatorAddr string) []byte {
	key := append(DelegationKeyPrefix, []byte(stakerAddr)...)
	key = append(key, []byte("/")...)
	return append(key, []byte(validatorAddr)...)
}

// ProposalKey returns the key for a proposal by ID
func ProposalKey(proposalID uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, proposalID)
	return append(ProposalKeyPrefix, bz...)
}

// VoteKey returns the key for a vote by proposal ID and voter address
func VoteKey(proposalID uint64, voterAddr string) []byte {
	key := ProposalKey(proposalID)
	key = append(key, VoteKeyPrefix...)
	return append(key, []byte(voterAddr)...)
}

// BurnRateHistoryKey returns the key for burn rate history by block height
func BurnRateHistoryKey(blockHeight int64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(blockHeight))
	return append(BurnRateHistoryKeyPrefix, bz...)
}

// UnstakingQueueKey returns the key for unstaking queue by completion time and staker address
func UnstakingQueueKey(completionTime int64, stakerAddr string) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(completionTime))
	key := append(UnstakingQueueKeyPrefix, bz...)
	return append(key, []byte(stakerAddr)...)
}

// ParseDelegationKey parses a delegation key to extract staker and validator addresses
func ParseDelegationKey(key []byte) (stakerAddr, validatorAddr string) {
	// Remove the prefix
	key = key[len(DelegationKeyPrefix):]
	
	// Find the separator
	sepIndex := -1
	for i, b := range key {
		if b == '/' {
			sepIndex = i
			break
		}
	}
	
	if sepIndex == -1 {
		return "", ""
	}
	
	stakerAddr = string(key[:sepIndex])
	validatorAddr = string(key[sepIndex+1:])
	return
}

// GetVestingScheduleIteratorKey returns iterator key for vesting schedules
func GetVestingScheduleIteratorKey() []byte {
	return VestingScheduleKeyPrefix
}

// GetStakerInfoIteratorKey returns iterator key for staker info
func GetStakerInfoIteratorKey() []byte {
	return StakerInfoKeyPrefix
}

// GetProposalIteratorKey returns iterator key for proposals
func GetProposalIteratorKey() []byte {
	return ProposalKeyPrefix
}