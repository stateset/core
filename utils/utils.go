package utils

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParseBool parses a string to a boolean value
func ParseBool(s string) (bool, error) {
	switch s {
	case "true", "1", "yes", "on":
		return true, nil
	case "false", "0", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("invalid bool value: %s", s)
	}
}

// StringInSlice checks if a string exists in a slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// MustMarshalJSON marshals an object to JSON or panics
func MustMarshalJSON(o interface{}) []byte {
	bz, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return bz
}

// SortedKeys returns sorted keys from a map
func SortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// ValidateDenom validates a denomination string
func ValidateDenom(denom string) error {
	return sdk.ValidateDenom(denom)
}