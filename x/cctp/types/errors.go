package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CCTP module sentinel errors
var (
	ErrInvalidSigner                  = sdkerrors.Register(ModuleName, 1, "expected gov account as only signer for proposal message")
	ErrInvalidMessage                 = sdkerrors.Register(ModuleName, 2, "invalid message")
	ErrInvalidAttestation            = sdkerrors.Register(ModuleName, 3, "invalid attestation")
	ErrInvalidMessageVersion         = sdkerrors.Register(ModuleName, 4, "invalid message version")
	ErrInvalidMessageBodyVersion     = sdkerrors.Register(ModuleName, 5, "invalid message body version")
	ErrInvalidMessageLength          = sdkerrors.Register(ModuleName, 6, "invalid message length")
	ErrInvalidAddressLength          = sdkerrors.Register(ModuleName, 7, "invalid address length")
	ErrInvalidAddress                = sdkerrors.Register(ModuleName, 8, "invalid address")
	ErrInvalidToken                  = sdkerrors.Register(ModuleName, 9, "invalid token")
	ErrInvalidAmount                 = sdkerrors.Register(ModuleName, 10, "invalid amount")
	ErrInvalidDestinationDomain      = sdkerrors.Register(ModuleName, 11, "invalid destination domain")
	ErrInvalidSourceDomain           = sdkerrors.Register(ModuleName, 12, "invalid source domain")
	ErrInvalidNonce                  = sdkerrors.Register(ModuleName, 13, "invalid nonce")
	ErrInvalidSignature              = sdkerrors.Register(ModuleName, 14, "invalid signature")
	ErrInvalidSignatureRecovery      = sdkerrors.Register(ModuleName, 15, "invalid signature recovery")
	ErrInvalidSignatureThreshold     = sdkerrors.Register(ModuleName, 16, "invalid signature threshold")
	ErrInvalidMessageBody            = sdkerrors.Register(ModuleName, 17, "invalid message body")
	ErrInvalidBurnMessage            = sdkerrors.Register(ModuleName, 18, "invalid burn message")
	ErrInvalidTokenPair              = sdkerrors.Register(ModuleName, 19, "invalid token pair")
	ErrInvalidBurnLimit              = sdkerrors.Register(ModuleName, 20, "invalid burn limit")
	ErrInvalidMaxMessageBodySize     = sdkerrors.Register(ModuleName, 21, "invalid max message body size")
	
	// State related errors
	ErrOwnerNotSet                   = sdkerrors.Register(ModuleName, 30, "owner not set")
	ErrPendingOwnerNotSet           = sdkerrors.Register(ModuleName, 31, "pending owner not set")
	ErrAttesterManagerNotSet        = sdkerrors.Register(ModuleName, 32, "attester manager not set")
	ErrTokenControllerNotSet        = sdkerrors.Register(ModuleName, 33, "token controller not set")
	ErrPauserNotSet                 = sdkerrors.Register(ModuleName, 34, "pauser not set")
	ErrAttesterNotFound             = sdkerrors.Register(ModuleName, 35, "attester not found")
	ErrAttesterAlreadyFound         = sdkerrors.Register(ModuleName, 36, "attester already found")
	ErrTokenPairNotFound            = sdkerrors.Register(ModuleName, 37, "token pair not found")
	ErrTokenPairAlreadyFound        = sdkerrors.Register(ModuleName, 38, "token pair already found")
	ErrRemoteTokenMessengerNotFound = sdkerrors.Register(ModuleName, 39, "remote token messenger not found")
	ErrRemoteTokenMessengerAlreadyFound = sdkerrors.Register(ModuleName, 40, "remote token messenger already found")
	ErrUsedNonceAlreadyFound        = sdkerrors.Register(ModuleName, 41, "used nonce already found")
	ErrBurnLimitNotFound            = sdkerrors.Register(ModuleName, 42, "burn limit not found")
	
	// Permission related errors
	ErrUnauthorized                 = sdkerrors.Register(ModuleName, 50, "unauthorized")
	ErrNotOwner                     = sdkerrors.Register(ModuleName, 51, "not owner")
	ErrNotPendingOwner              = sdkerrors.Register(ModuleName, 52, "not pending owner")
	ErrNotAttesterManager           = sdkerrors.Register(ModuleName, 53, "not attester manager")
	ErrNotTokenController           = sdkerrors.Register(ModuleName, 54, "not token controller")
	ErrNotPauser                    = sdkerrors.Register(ModuleName, 55, "not pauser")
	ErrNotDestinationCaller         = sdkerrors.Register(ModuleName, 56, "not destination caller")
	ErrNotOriginalSender            = sdkerrors.Register(ModuleName, 57, "not original sender")
	
	// Operational errors
	ErrBurningAndMintingPaused      = sdkerrors.Register(ModuleName, 70, "burning and minting paused")
	ErrSendingAndReceivingPaused    = sdkerrors.Register(ModuleName, 71, "sending and receiving messages paused")
	ErrExceedsBurnLimit             = sdkerrors.Register(ModuleName, 72, "exceeds burn limit")
	ErrExceedsMaxMessageBodySize    = sdkerrors.Register(ModuleName, 73, "exceeds max message body size")
	ErrNonceAlreadyUsed             = sdkerrors.Register(ModuleName, 74, "nonce already used")
	ErrInsufficientAttestersToDisable = sdkerrors.Register(ModuleName, 75, "insufficient attesters to disable")
	ErrCannotRemoveLastAttester     = sdkerrors.Register(ModuleName, 76, "cannot remove last attester")
	ErrInsufficientSignatures       = sdkerrors.Register(ModuleName, 77, "insufficient signatures")
	ErrDuplicateSignatures          = sdkerrors.Register(ModuleName, 78, "duplicate signatures")
	ErrSignaturesOutOfOrder         = sdkerrors.Register(ModuleName, 79, "signatures out of order")
	ErrAttesterNotEnabled           = sdkerrors.Register(ModuleName, 80, "attester not enabled")
	ErrInvalidSignatureLength       = sdkerrors.Register(ModuleName, 81, "invalid signature length")
	
	// Token and minting errors
	ErrMintingFailed                = sdkerrors.Register(ModuleName, 90, "minting failed")
	ErrBurningFailed                = sdkerrors.Register(ModuleName, 91, "burning failed")
	ErrInvalidMintRecipient         = sdkerrors.Register(ModuleName, 92, "invalid mint recipient")
	ErrInvalidBurnToken             = sdkerrors.Register(ModuleName, 93, "invalid burn token")
	ErrTokenNotSupported            = sdkerrors.Register(ModuleName, 94, "token not supported")
	ErrInsufficientBalance          = sdkerrors.Register(ModuleName, 95, "insufficient balance")
	
	// Message processing errors
	ErrMessageAlreadyReceived       = sdkerrors.Register(ModuleName, 100, "message already received")
	ErrMessageNotFound              = sdkerrors.Register(ModuleName, 101, "message not found")
	ErrInvalidMessageSender         = sdkerrors.Register(ModuleName, 102, "invalid message sender")
	ErrInvalidMessageRecipient      = sdkerrors.Register(ModuleName, 103, "invalid message recipient")
	ErrMessageExpired               = sdkerrors.Register(ModuleName, 104, "message expired")
	ErrMessageProcessingFailed      = sdkerrors.Register(ModuleName, 105, "message processing failed")
)