package types

// Event types for the XSS module
const (
	EventTypeStakeTokens          = "stake_tokens"
	EventTypeUnstakeTokens        = "unstake_tokens"
	EventTypeWithdrawRewards      = "withdraw_rewards"
	EventTypeCreateValidator      = "create_validator"
	EventTypeEditValidator        = "edit_validator"
	EventTypeUnjailValidator      = "unjail_validator"
	EventTypeSlashValidator       = "slash_validator"
	EventTypeCompleteUnbonding    = "complete_unbonding"
	EventTypeCompleteRedelegation = "complete_redelegation"
	EventTypeBond                 = "bond"
	EventTypeUnbond               = "unbond"
	EventTypeRedelegate           = "redelegate"
	EventTypeSetWithdrawAddress   = "set_withdraw_address"
	EventTypeRewards              = "rewards"
	EventTypeCommission           = "commission"
	EventTypeSlash                = "slash"
	EventTypeLiveness             = "liveness"
	EventTypeMint                 = "mint"
	EventTypeBurn                 = "burn"
)

// Event attribute keys
const (
	AttributeKeyValidator         = "validator"
	AttributeKeyDelegator         = "delegator"
	AttributeKeyAmount            = "amount"
	AttributeKeyCompletionTime    = "completion_time"
	AttributeKeyCreationHeight    = "creation_height"
	AttributeKeyDestinationValidator = "destination_validator"
	AttributeKeySourceValidator   = "source_validator"
	AttributeKeyCommissionRate    = "commission_rate"
	AttributeKeyMinSelfDelegation = "min_self_delegation"
	AttributeKeySlashAmount       = "slash_amount"
	AttributeKeySlashFraction     = "slash_fraction"
	AttributeKeyInfractionHeight  = "infraction_height"
	AttributeKeyInfractionTime    = "infraction_time"
	AttributeKeyJailed            = "jailed"
	AttributeKeyReason            = "reason"
	AttributeKeyPower             = "power"
	AttributeKeyMissedBlocks      = "missed_blocks"
	AttributeKeyBondedRatio       = "bonded_ratio"
	AttributeKeyInflation         = "inflation"
	AttributeKeyProvisions        = "provisions"
	AttributeKeyTotalSupply       = "total_supply"
	AttributeKeyBondedTokens      = "bonded_tokens"
	AttributeKeyNotBondedTokens   = "not_bonded_tokens"
	
	// Common attribute values
	AttributeValueCategory        = ModuleName
	AttributeValueAction         = "action"
	AttributeValueDoubleSign     = "double_sign"
	AttributeValueMissingSignature = "missing_signature"
)