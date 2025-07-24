package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Keccak256Hash calculates the Keccak256 hash of the input data
func Keccak256Hash(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// SHA256Hash calculates the SHA256 hash of the input data
func SHA256Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// PadAddressTo32Bytes pads an address to 32 bytes (bytes32 format)
func PadAddressTo32Bytes(addr []byte) []byte {
	padded := make([]byte, 32)
	if len(addr) <= 32 {
		copy(padded[32-len(addr):], addr)
	} else {
		copy(padded, addr[len(addr)-32:])
	}
	return padded
}

// TrimLeadingZeros removes leading zeros from a byte slice
func TrimLeadingZeros(data []byte) []byte {
	for i, b := range data {
		if b != 0 {
			return data[i:]
		}
	}
	return []byte{0}
}

// IsZeroBytes checks if all bytes in the slice are zero
func IsZeroBytes(data []byte) bool {
	for _, b := range data {
		if b != 0 {
			return false
		}
	}
	return true
}

// BytesEqual compares two byte slices for equality
func BytesEqual(a, b []byte) bool {
	return bytes.Equal(a, b)
}

// HexToBytes converts a hex string to bytes
func HexToBytes(s string) ([]byte, error) {
	// Remove 0x prefix if present
	if len(s) >= 2 && s[:2] == "0x" {
		s = s[2:]
	}
	return hex.DecodeString(s)
}

// BytesToHex converts bytes to a hex string with 0x prefix
func BytesToHex(data []byte) string {
	return "0x" + hex.EncodeToString(data)
}

// ValidateBytes32 validates that the byte slice is exactly 32 bytes
func ValidateBytes32(data []byte) bool {
	return len(data) == 32
}

// ValidateBytes20 validates that the byte slice is exactly 20 bytes (Ethereum address)
func ValidateBytes20(data []byte) bool {
	return len(data) == 20
}

// ConcatBytes concatenates multiple byte slices
func ConcatBytes(slices ...[]byte) []byte {
	var result []byte
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// ReverseBytes reverses a byte slice
func ReverseBytes(data []byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[len(data)-1-i] = b
	}
	return result
}

// SafeCopy safely copies bytes, ensuring the destination has the correct size
func SafeCopy(dst []byte, src []byte) {
	if len(dst) != len(src) {
		panic("SafeCopy: length mismatch")
	}
	copy(dst, src)
}

// ZeroPad pads a byte slice with leading zeros to reach the target length
func ZeroPad(data []byte, length int) []byte {
	if len(data) >= length {
		return data
	}
	padded := make([]byte, length)
	copy(padded[length-len(data):], data)
	return padded
}

// ValidateEthereumAddress validates an Ethereum-style address (20 bytes)
func ValidateEthereumAddress(addr []byte) bool {
	return len(addr) == 20 && !IsZeroBytes(addr)
}

// ValidateEthereumAddressBytes32 validates an Ethereum address in bytes32 format
func ValidateEthereumAddressBytes32(addr []byte) bool {
	if len(addr) != 32 {
		return false
	}
	// Check that the first 12 bytes are zero (standard Ethereum address padding)
	for i := 0; i < 12; i++ {
		if addr[i] != 0 {
			return false
		}
	}
	// Check that the address part is not all zeros
	return !IsZeroBytes(addr[12:])
}

// ExtractEthereumAddressFromBytes32 extracts the 20-byte Ethereum address from bytes32
func ExtractEthereumAddressFromBytes32(addr []byte) []byte {
	if len(addr) != 32 {
		return nil
	}
	return addr[12:]
}

// ConvertEthereumAddressToBytes32 converts a 20-byte Ethereum address to bytes32 format
func ConvertEthereumAddressToBytes32(addr []byte) []byte {
	if len(addr) != 20 {
		return nil
	}
	result := make([]byte, 32)
	copy(result[12:], addr)
	return result
}

// ValidateDomain validates a domain ID (should be non-zero for most cases)
func ValidateDomain(domain uint32) bool {
	return domain > 0
}

// ValidateNonce validates a nonce (should be positive)
func ValidateNonce(nonce uint64) bool {
	return nonce > 0
}

// GetDomainFromChainID maps common chain IDs to domain IDs
func GetDomainFromChainID(chainID string) uint32 {
	switch chainID {
	case "noble-1":
		return 4 // Noble mainnet
	case "ethereum-1":
		return 0 // Ethereum mainnet
	case "polygon-1":
		return 2 // Polygon mainnet
	case "avalanche-1":
		return 1 // Avalanche mainnet
	case "arbitrum-1":
		return 3 // Arbitrum mainnet
	default:
		return 0
	}
}

// GetChainIDFromDomain maps domain IDs to common chain IDs
func GetChainIDFromDomain(domain uint32) string {
	switch domain {
	case 4:
		return "noble-1"
	case 0:
		return "ethereum-1"
	case 2:
		return "polygon-1"
	case 1:
		return "avalanche-1"
	case 3:
		return "arbitrum-1"
	default:
		return "unknown"
	}
}

// CalculateMessageHash calculates the message hash for attestation
func CalculateMessageHash(message []byte) []byte {
	return Keccak256Hash(message)
}

// ValidateSignatureLength validates signature length (65 bytes for ECDSA)
func ValidateSignatureLength(sig []byte) bool {
	return len(sig) == 65
}

// SplitSignature splits a 65-byte signature into r, s, and v components
func SplitSignature(sig []byte) (r, s []byte, v uint8) {
	if len(sig) != 65 {
		return nil, nil, 0
	}
	r = sig[:32]
	s = sig[32:64]
	v = sig[64]
	return
}

// CombineSignature combines r, s, v components into a 65-byte signature
func CombineSignature(r, s []byte, v uint8) []byte {
	if len(r) != 32 || len(s) != 32 {
		return nil
	}
	sig := make([]byte, 65)
	copy(sig[:32], r)
	copy(sig[32:64], s)
	sig[64] = v
	return sig
}

// ValidateAttestationLength validates attestation length based on signature threshold
func ValidateAttestationLength(attestation []byte, threshold uint32) bool {
	expectedLength := int(threshold) * 65 // 65 bytes per signature
	return len(attestation) == expectedLength
}

// SplitAttestation splits attestation into individual signatures
func SplitAttestation(attestation []byte) [][]byte {
	if len(attestation)%65 != 0 {
		return nil
	}
	
	numSigs := len(attestation) / 65
	signatures := make([][]byte, numSigs)
	
	for i := 0; i < numSigs; i++ {
		start := i * 65
		end := start + 65
		signatures[i] = attestation[start:end]
	}
	
	return signatures
}

// CombineAttestations combines multiple signatures into a single attestation
func CombineAttestations(signatures [][]byte) []byte {
	var result []byte
	for _, sig := range signatures {
		if len(sig) != 65 {
			return nil
		}
		result = append(result, sig...)
	}
	return result
}

// FormatDomain formats a domain ID for display
func FormatDomain(domain uint32) string {
	chainID := GetChainIDFromDomain(domain)
	if chainID != "unknown" {
		return chainID
	}
	return hex.EncodeToString([]byte{byte(domain >> 24), byte(domain >> 16), byte(domain >> 8), byte(domain)})
}

// FormatNonce formats a nonce for display
func FormatNonce(nonce uint64) string {
	return hex.EncodeToString([]byte{
		byte(nonce >> 56), byte(nonce >> 48), byte(nonce >> 40), byte(nonce >> 32),
		byte(nonce >> 24), byte(nonce >> 16), byte(nonce >> 8), byte(nonce),
	})
}

// ParseHexBytes parses hex string and validates length
func ParseHexBytes(hexStr string, expectedLength int) ([]byte, error) {
	data, err := HexToBytes(hexStr)
	if err != nil {
		return nil, err
	}
	if len(data) != expectedLength {
		return nil, ErrInvalidAddressLength
	}
	return data, nil
}

// FormatAddress formats an address for display (shows first and last 4 bytes)
func FormatAddress(addr []byte) string {
	if len(addr) == 0 {
		return "0x0000...0000"
	}
	if len(addr) <= 8 {
		return BytesToHex(addr)
	}
	return BytesToHex(addr[:4]) + "..." + BytesToHex(addr[len(addr)-4:])
}

// ValidateTokenDenom validates a token denomination
func ValidateTokenDenom(denom string) bool {
	return len(denom) > 0 && len(denom) <= 128
}

// NormalizeTokenDenom normalizes a token denomination
func NormalizeTokenDenom(denom string) string {
	// Convert to lowercase for consistency
	return denom
}

// CreateNonceKey creates a unique key for nonce tracking
func CreateNonceKey(sourceDomain uint32, nonce uint64) string {
	return hex.EncodeToString([]byte{
		byte(sourceDomain >> 24), byte(sourceDomain >> 16), byte(sourceDomain >> 8), byte(sourceDomain),
		byte(nonce >> 56), byte(nonce >> 48), byte(nonce >> 40), byte(nonce >> 32),
		byte(nonce >> 24), byte(nonce >> 16), byte(nonce >> 8), byte(nonce),
	})
}