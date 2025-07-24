package types

// CCTP module event types
const (
	// Message events
	EventTypeMessageSent     = "message_sent"
	EventTypeMessageReceived = "message_received"

	// Deposit for burn events
	EventTypeDepositForBurn = "deposit_for_burn"

	// Mint and withdraw events
	EventTypeMintAndWithdraw = "mint_and_withdraw"

	// Ownership events
	EventTypeOwnershipTransferStarted = "ownership_transfer_started"
	EventTypeOwnerUpdated            = "owner_updated"

	// Attester events
	EventTypeAttesterEnabled         = "attester_enabled"
	EventTypeAttesterDisabled        = "attester_disabled"
	EventTypeAttesterManagerUpdated  = "attester_manager_updated"
	EventTypeSignatureThresholdUpdated = "signature_threshold_updated"

	// Token events
	EventTypeTokenPairLinked      = "token_pair_linked"
	EventTypeTokenPairUnlinked    = "token_pair_unlinked"
	EventTypeTokenControllerUpdated = "token_controller_updated"
	EventTypeSetBurnLimitPerMessage = "set_burn_limit_per_message"

	// Remote token messenger events
	EventTypeRemoteTokenMessengerAdded   = "remote_token_messenger_added"
	EventTypeRemoteTokenMessengerRemoved = "remote_token_messenger_removed"

	// Pause/unpause events
	EventTypeBurningAndMintingPaused        = "burning_and_minting_paused"
	EventTypeBurningAndMintingUnpaused      = "burning_and_minting_unpaused"
	EventTypeSendingAndReceivingPaused      = "sending_and_receiving_messages_paused"
	EventTypeSendingAndReceivingUnpaused    = "sending_and_receiving_messages_unpaused"

	// Configuration events
	EventTypePauserUpdated             = "pauser_updated"
	EventTypeMaxMessageBodySizeUpdated = "max_message_body_size_updated"
)

// CCTP module event attribute keys
const (
	// Common attributes
	AttributeKeyModule = ModuleName

	// Message attributes
	AttributeKeyMessage            = "message"
	AttributeKeyMessageHash        = "message_hash"
	AttributeKeyNonce             = "nonce"
	AttributeKeySourceDomain      = "source_domain"
	AttributeKeyDestinationDomain = "destination_domain"
	AttributeKeySender            = "sender"
	AttributeKeyRecipient         = "recipient"
	AttributeKeyCaller            = "caller"
	AttributeKeyDestinationCaller = "destination_caller"
	AttributeKeyMessageBody       = "message_body"

	// Deposit for burn attributes
	AttributeKeyAmount                    = "amount"
	AttributeKeyDepositor                = "depositor"
	AttributeKeyMintRecipient            = "mint_recipient"
	AttributeKeyBurnToken                = "burn_token"
	AttributeKeyMintToken                = "mint_token"
	AttributeKeyDestinationTokenMessenger = "destination_token_messenger"

	// Ownership attributes
	AttributeKeyPreviousOwner = "previous_owner"
	AttributeKeyNewOwner      = "new_owner"

	// Attester attributes
	AttributeKeyAttester                 = "attester"
	AttributeKeyPreviousAttesterManager  = "previous_attester_manager"
	AttributeKeyNewAttesterManager       = "new_attester_manager"
	AttributeKeyPreviousSignatureThreshold = "previous_signature_threshold"
	AttributeKeyNewSignatureThreshold   = "new_signature_threshold"

	// Token attributes
	AttributeKeyLocalToken              = "local_token"
	AttributeKeyRemoteToken             = "remote_token"
	AttributeKeyRemoteDomain            = "remote_domain"
	AttributeKeyPreviousTokenController = "previous_token_controller"
	AttributeKeyNewTokenController      = "new_token_controller"
	AttributeKeyBurnLimit              = "burn_limit"

	// Remote token messenger attributes
	AttributeKeyDomainId = "domain_id"
	AttributeKeyAddress  = "address"

	// Configuration attributes
	AttributeKeyPreviousPauser       = "previous_pauser"
	AttributeKeyNewPauser           = "new_pauser"
	AttributeKeyPreviousMaxMessageBodySize = "previous_max_message_body_size"
	AttributeKeyNewMaxMessageBodySize = "new_max_message_body_size"

	// Status attributes
	AttributeKeyPausedStatus = "paused_status"
)