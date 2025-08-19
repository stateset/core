package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		StablecoinList:   []Stablecoin{},
		PriceDataList:    []PriceData{},
		MintRequestList:  []MintRequest{},
		BurnRequestList:  []BurnRequest{},
		WhitelistEntries: []WhitelistEntry{},
		BlacklistEntries: []BlacklistEntry{},
		StablecoinCount:  0,
		Params:           DefaultParams(),
	}
}

// GenesisState defines the genesis state for the stablecoins module
type GenesisState struct {
	StablecoinList   []Stablecoin      `json:"stablecoin_list"`
	PriceDataList    []PriceData       `json:"price_data_list"`
	MintRequestList  []MintRequest     `json:"mint_request_list"`
	BurnRequestList  []BurnRequest     `json:"burn_request_list"`
	WhitelistEntries []WhitelistEntry  `json:"whitelist_entries"`
	BlacklistEntries []BlacklistEntry  `json:"blacklist_entries"`
	StablecoinCount  uint64            `json:"stablecoin_count"`
	Params           Params            `json:"params"`
}

// WhitelistEntry represents a whitelist entry
type WhitelistEntry struct {
	Denom   string `json:"denom"`
	Address string `json:"address"`
}

// BlacklistEntry represents a blacklist entry
type BlacklistEntry struct {
	Denom   string `json:"denom"`
	Address string `json:"address"`
	Reason  string `json:"reason"`
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in stablecoin
	stablecoinIndexMap := make(map[string]struct{})

	for _, elem := range gs.StablecoinList {
		index := string(StablecoinKey(elem.Denom))
		if _, ok := stablecoinIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for stablecoin")
		}
		stablecoinIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in priceData
	priceDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.PriceDataList {
		index := string(PriceDataKey(elem.Denom))
		if _, ok := priceDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for priceData")
		}
		priceDataIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in mintRequest
	mintRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.MintRequestList {
		index := string(MintRequestKey(fmt.Sprintf("%d", elem.Id)))
		if _, ok := mintRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mintRequest")
		}
		mintRequestIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in burnRequest
	burnRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.BurnRequestList {
		index := string(BurnRequestKey(fmt.Sprintf("%d", elem.Id)))
		if _, ok := burnRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for burnRequest")
		}
		burnRequestIndexMap[index] = struct{}{}
	}

	// Check for duplicated whitelist entries
	whitelistIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhitelistEntries {
		index := string(WhitelistKey(elem.Denom, elem.Address))
		if _, ok := whitelistIndexMap[index]; ok {
			return fmt.Errorf("duplicated whitelist entry")
		}
		whitelistIndexMap[index] = struct{}{}
	}

	// Check for duplicated blacklist entries
	blacklistIndexMap := make(map[string]struct{})

	for _, elem := range gs.BlacklistEntries {
		index := string(BlacklistKey(elem.Denom, elem.Address))
		if _, ok := blacklistIndexMap[index]; ok {
			return fmt.Errorf("duplicated blacklist entry")
		}
		blacklistIndexMap[index] = struct{}{}
	}

	// Validate parameters
	return gs.Params.Validate()
}