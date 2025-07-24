package types

import (
	"encoding/binary"
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "cctp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cctp"

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	// State store keys
	OwnerKey                     = []byte{0x01}
	PendingOwnerKey              = []byte{0x02}
	AttesterManagerKey           = []byte{0x03}
	TokenControllerKey           = []byte{0x04}
	PauserKey                    = []byte{0x05}
	SignatureThresholdKey        = []byte{0x06}
	NextAvailableNonceKey        = []byte{0x07}
	BurningAndMintingPausedKey   = []byte{0x08}
	SendingAndReceivingPausedKey = []byte{0x09}
	MaxMessageBodySizeKey        = []byte{0x10}

	// Collection store key prefixes
	AttesterKeyPrefix             = []byte{0x20}
	RemoteTokenMessengerKeyPrefix = []byte{0x21}
	TokenPairKeyPrefix            = []byte{0x22}
	UsedNonceKeyPrefix            = []byte{0x23}
	PerMessageBurnLimitKeyPrefix  = []byte{0x24}
	SentMessageKeyPrefix          = []byte{0x25}
	ReceivedMessageKeyPrefix      = []byte{0x26}
)

// AttesterKey returns the store key for an attester
func AttesterKey(attester string) []byte {
	return append(AttesterKeyPrefix, []byte(attester)...)
}

// RemoteTokenMessengerKey returns the store key for a remote token messenger
func RemoteTokenMessengerKey(domain uint32) []byte {
	key := make([]byte, len(RemoteTokenMessengerKeyPrefix)+4)
	copy(key, RemoteTokenMessengerKeyPrefix)
	binary.BigEndian.PutUint32(key[len(RemoteTokenMessengerKeyPrefix):], domain)
	return key
}

// TokenPairKey returns the store key for a token pair
func TokenPairKey(remoteDomain uint32, remoteToken []byte) []byte {
	key := make([]byte, len(TokenPairKeyPrefix)+4+len(remoteToken))
	copy(key, TokenPairKeyPrefix)
	binary.BigEndian.PutUint32(key[len(TokenPairKeyPrefix):], remoteDomain)
	copy(key[len(TokenPairKeyPrefix)+4:], remoteToken)
	return key
}

// UsedNonceKey returns the store key for a used nonce
func UsedNonceKey(sourceDomain uint32, nonce uint64) []byte {
	key := make([]byte, len(UsedNonceKeyPrefix)+4+8)
	copy(key, UsedNonceKeyPrefix)
	binary.BigEndian.PutUint32(key[len(UsedNonceKeyPrefix):], sourceDomain)
	binary.BigEndian.PutUint64(key[len(UsedNonceKeyPrefix)+4:], nonce)
	return key
}

// PerMessageBurnLimitKey returns the store key for per message burn limit
func PerMessageBurnLimitKey(denom string) []byte {
	return append(PerMessageBurnLimitKeyPrefix, []byte(denom)...)
}

// SentMessageKey returns the store key for a sent message
func SentMessageKey(domain uint32, nonce uint64) []byte {
	key := make([]byte, len(SentMessageKeyPrefix)+4+8)
	copy(key, SentMessageKeyPrefix)
	binary.BigEndian.PutUint32(key[len(SentMessageKeyPrefix):], domain)
	binary.BigEndian.PutUint64(key[len(SentMessageKeyPrefix)+4:], nonce)
	return key
}

// ReceivedMessageKey returns the store key for a received message
func ReceivedMessageKey(sourceDomain uint32, nonce uint64) []byte {
	key := make([]byte, len(ReceivedMessageKeyPrefix)+4+8)
	copy(key, ReceivedMessageKeyPrefix)
	binary.BigEndian.PutUint32(key[len(ReceivedMessageKeyPrefix):], sourceDomain)
	binary.BigEndian.PutUint64(key[len(ReceivedMessageKeyPrefix)+4:], nonce)
	return key
}

// Helper functions to parse keys back to their components

// ParseRemoteTokenMessengerKey parses a remote token messenger key to extract domain
func ParseRemoteTokenMessengerKey(key []byte) (uint32, error) {
	if len(key) != len(RemoteTokenMessengerKeyPrefix)+4 {
		return 0, fmt.Errorf("invalid remote token messenger key length")
	}
	return binary.BigEndian.Uint32(key[len(RemoteTokenMessengerKeyPrefix):]), nil
}

// ParseTokenPairKey parses a token pair key to extract domain and remote token
func ParseTokenPairKey(key []byte) (uint32, []byte, error) {
	if len(key) < len(TokenPairKeyPrefix)+4 {
		return 0, nil, fmt.Errorf("invalid token pair key length")
	}
	domain := binary.BigEndian.Uint32(key[len(TokenPairKeyPrefix):])
	remoteToken := key[len(TokenPairKeyPrefix)+4:]
	return domain, remoteToken, nil
}

// ParseUsedNonceKey parses a used nonce key to extract source domain and nonce
func ParseUsedNonceKey(key []byte) (uint32, uint64, error) {
	if len(key) != len(UsedNonceKeyPrefix)+4+8 {
		return 0, 0, fmt.Errorf("invalid used nonce key length")
	}
	sourceDomain := binary.BigEndian.Uint32(key[len(UsedNonceKeyPrefix):])
	nonce := binary.BigEndian.Uint64(key[len(UsedNonceKeyPrefix)+4:])
	return sourceDomain, nonce, nil
}

// ParseSentMessageKey parses a sent message key to extract domain and nonce
func ParseSentMessageKey(key []byte) (uint32, uint64, error) {
	if len(key) != len(SentMessageKeyPrefix)+4+8 {
		return 0, 0, fmt.Errorf("invalid sent message key length")
	}
	domain := binary.BigEndian.Uint32(key[len(SentMessageKeyPrefix):])
	nonce := binary.BigEndian.Uint64(key[len(SentMessageKeyPrefix)+4:])
	return domain, nonce, nil
}

// ParseReceivedMessageKey parses a received message key to extract source domain and nonce
func ParseReceivedMessageKey(key []byte) (uint32, uint64, error) {
	if len(key) != len(ReceivedMessageKeyPrefix)+4+8 {
		return 0, 0, fmt.Errorf("invalid received message key length")
	}
	sourceDomain := binary.BigEndian.Uint32(key[len(ReceivedMessageKeyPrefix):])
	nonce := binary.BigEndian.Uint64(key[len(ReceivedMessageKeyPrefix)+4:])
	return sourceDomain, nonce, nil
}

// Key iteration helpers

// GetAttesterKeyPrefix returns the prefix for attester keys
func GetAttesterKeyPrefix() []byte {
	return AttesterKeyPrefix
}

// GetRemoteTokenMessengerKeyPrefix returns the prefix for remote token messenger keys
func GetRemoteTokenMessengerKeyPrefix() []byte {
	return RemoteTokenMessengerKeyPrefix
}

// GetTokenPairKeyPrefix returns the prefix for token pair keys
func GetTokenPairKeyPrefix() []byte {
	return TokenPairKeyPrefix
}

// GetUsedNonceKeyPrefix returns the prefix for used nonce keys
func GetUsedNonceKeyPrefix() []byte {
	return UsedNonceKeyPrefix
}

// GetPerMessageBurnLimitKeyPrefix returns the prefix for per message burn limit keys
func GetPerMessageBurnLimitKeyPrefix() []byte {
	return PerMessageBurnLimitKeyPrefix
}

// GetSentMessageKeyPrefix returns the prefix for sent message keys
func GetSentMessageKeyPrefix() []byte {
	return SentMessageKeyPrefix
}

// GetReceivedMessageKeyPrefix returns the prefix for received message keys
func GetReceivedMessageKeyPrefix() []byte {
	return ReceivedMessageKeyPrefix
}

// GetUsedNonceKeyPrefixForDomain returns the prefix for used nonce keys for a specific domain
func GetUsedNonceKeyPrefixForDomain(sourceDomain uint32) []byte {
	prefix := make([]byte, len(UsedNonceKeyPrefix)+4)
	copy(prefix, UsedNonceKeyPrefix)
	binary.BigEndian.PutUint32(prefix[len(UsedNonceKeyPrefix):], sourceDomain)
	return prefix
}

// GetTokenPairKeyPrefixForDomain returns the prefix for token pair keys for a specific domain
func GetTokenPairKeyPrefixForDomain(remoteDomain uint32) []byte {
	prefix := make([]byte, len(TokenPairKeyPrefix)+4)
	copy(prefix, TokenPairKeyPrefix)
	binary.BigEndian.PutUint32(prefix[len(TokenPairKeyPrefix):], remoteDomain)
	return prefix
}

// GetSentMessageKeyPrefixForDomain returns the prefix for sent message keys for a specific domain
func GetSentMessageKeyPrefixForDomain(domain uint32) []byte {
	prefix := make([]byte, len(SentMessageKeyPrefix)+4)
	copy(prefix, SentMessageKeyPrefix)
	binary.BigEndian.PutUint32(prefix[len(SentMessageKeyPrefix):], domain)
	return prefix
}

// GetReceivedMessageKeyPrefixForDomain returns the prefix for received message keys for a specific domain
func GetReceivedMessageKeyPrefixForDomain(sourceDomain uint32) []byte {
	prefix := make([]byte, len(ReceivedMessageKeyPrefix)+4)
	copy(prefix, ReceivedMessageKeyPrefix)
	binary.BigEndian.PutUint32(prefix[len(ReceivedMessageKeyPrefix):], sourceDomain)
	return prefix
}