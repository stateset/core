package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/math"
)

// Stablecoin represents a stablecoin in the system
type Stablecoin struct {
	Denom            string      `json:"denom"`
	Name             string      `json:"name"`
	Symbol           string      `json:"symbol"`
	TotalSupply      sdk.Coins   `json:"total_supply"`
	Paused           bool        `json:"paused"`
	CollateralRatio  math.LegacyDec `json:"collateral_ratio"`
	Creator          string      `json:"creator"`
	MintingEnabled   bool        `json:"minting_enabled"`
	BurningEnabled   bool        `json:"burning_enabled"`
}

// PriceData represents price information for a stablecoin
type PriceData struct {
	Denom     string      `json:"denom"`
	Price     math.LegacyDec `json:"price"`
	Timestamp int64       `json:"timestamp"`
	Source    string      `json:"source"`
}

// MintRequest represents a request to mint stablecoins
type MintRequest struct {
	Id              uint64    `json:"id"`
	Denom           string    `json:"denom"`
	Requester       string    `json:"requester"`
	Amount          sdk.Coins `json:"amount"`
	Collateral      sdk.Coins `json:"collateral"`
	Status          string    `json:"status"`
	ApprovedBy      string    `json:"approved_by"`
	ApprovedAt      int64     `json:"approved_at"`
	CreatedAt       int64     `json:"created_at"`
}

// BurnRequest represents a request to burn stablecoins
type BurnRequest struct {
	Id              uint64    `json:"id"`
	Denom           string    `json:"denom"`
	Requester       string    `json:"requester"`
	Amount          sdk.Coins `json:"amount"`
	CollateralToReturn sdk.Coins `json:"collateral_to_return"`
	Status          string    `json:"status"`
	ApprovedBy      string    `json:"approved_by"`
	ApprovedAt      int64     `json:"approved_at"`
	CreatedAt       int64     `json:"created_at"`
}

// MsgCreateStablecoin message structure
type MsgCreateStablecoin struct {
	Creator            string             `json:"creator"`
	Denom              string             `json:"denom"`
	Name               string             `json:"name"`
	Symbol             string             `json:"symbol"`
	Decimals           uint32             `json:"decimals"`
	Description        string             `json:"description"`
	MaxSupply          math.Int           `json:"max_supply"`
	PegInfo            *PegInfo           `json:"peg_info"`
	ReserveInfo        *ReserveInfo       `json:"reserve_info"`
	StabilityMechanism string             `json:"stability_mechanism"`
	FeeInfo            *FeeInfo           `json:"fee_info"`
	AccessControl      *AccessControlInfo `json:"access_control"`
	Metadata           string             `json:"metadata"`
}

// ProtoMessage implements proto.Message
func (*MsgCreateStablecoin) ProtoMessage() {}

// Reset implements proto.Message
func (m *MsgCreateStablecoin) Reset() {
	*m = MsgCreateStablecoin{}
}

// String implements proto.Message
func (m *MsgCreateStablecoin) String() string {
	bz := ModuleCdc.MustMarshalJSON(m)
	return string(sdk.MustSortJSON(bz))
}

// PegInfo represents peg information for a stablecoin
type PegInfo struct {
	PegType      string   `json:"peg_type"`
	PeggedTo     string   `json:"pegged_to"`
	TargetValue  math.LegacyDec `json:"target_value"`
	AllowedAssets []string `json:"allowed_assets"`
}

// ReserveInfo represents reserve information for a stablecoin
type ReserveInfo struct {
	ReserveType     string    `json:"reserve_type"`
	ReserveAssets   sdk.Coins `json:"reserve_assets"`
	ReserveRatio    math.LegacyDec `json:"reserve_ratio"`
	MinReserveRatio math.LegacyDec `json:"min_reserve_ratio"`
}

// FeeInfo represents fee configuration for a stablecoin
type FeeInfo struct {
	MintFee       math.LegacyDec `json:"mint_fee"`
	BurnFee       math.LegacyDec `json:"burn_fee"`
	TransferFee   math.LegacyDec `json:"transfer_fee"`
	FeeCollector  string     `json:"fee_collector"`
}

// AccessControlInfo represents access control configuration
type AccessControlInfo struct {
	MintPermission     string   `json:"mint_permission"`
	BurnPermission     string   `json:"burn_permission"`
	TransferPermission string   `json:"transfer_permission"`
	AuthorizedMinters  []string `json:"authorized_minters"`
	AuthorizedBurners  []string `json:"authorized_burners"`
}

// MsgUpdateStablecoin message structure for updating stablecoin parameters
type MsgUpdateStablecoin struct {
	Creator       string             `json:"creator"`
	Denom         string             `json:"denom"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	PegInfo       *PegInfo           `json:"peg_info"`
	FeeInfo       *FeeInfo           `json:"fee_info"`
	AccessControl *AccessControlInfo `json:"access_control"`
	Metadata      string             `json:"metadata"`
}

// ProtoMessage implements proto.Message
func (*MsgUpdateStablecoin) ProtoMessage() {}

// Reset implements proto.Message
func (m *MsgUpdateStablecoin) Reset() {
	*m = MsgUpdateStablecoin{}
}

// String implements proto.Message
func (m *MsgUpdateStablecoin) String() string {
	bz := ModuleCdc.MustMarshalJSON(m)
	return string(sdk.MustSortJSON(bz))
}

// MsgMintStablecoin message structure for minting stablecoins
type MsgMintStablecoin struct {
	Creator   string   `json:"creator"`
	Denom     string   `json:"denom"`
	Amount    math.Int `json:"amount"`
	Recipient string   `json:"recipient"`
}

// ProtoMessage implements proto.Message
func (*MsgMintStablecoin) ProtoMessage() {}

// Reset implements proto.Message
func (m *MsgMintStablecoin) Reset() {
	*m = MsgMintStablecoin{}
}

// String implements proto.Message
func (m *MsgMintStablecoin) String() string {
	bz := ModuleCdc.MustMarshalJSON(m)
	return string(sdk.MustSortJSON(bz))
}

// MsgBurnStablecoin message structure for burning stablecoins
type MsgBurnStablecoin struct {
	Creator string   `json:"creator"`
	Denom   string   `json:"denom"`
	Amount  math.Int `json:"amount"`
}

// ProtoMessage implements proto.Message
func (*MsgBurnStablecoin) ProtoMessage() {}

// Reset implements proto.Message
func (m *MsgBurnStablecoin) Reset() {
	*m = MsgBurnStablecoin{}
}

// String implements proto.Message
func (m *MsgBurnStablecoin) String() string {
	bz := ModuleCdc.MustMarshalJSON(m)
	return string(sdk.MustSortJSON(bz))
}

// MsgPauseStablecoin message structure for pausing stablecoins
type MsgPauseStablecoin struct {
	Creator string `json:"creator"`
	Denom   string `json:"denom"`
	Paused  bool   `json:"paused"`
}

// ProtoMessage implements proto.Message
func (*MsgPauseStablecoin) ProtoMessage() {}

// Reset implements proto.Message
func (m *MsgPauseStablecoin) Reset() {
	*m = MsgPauseStablecoin{}
}

// String implements proto.Message
func (m *MsgPauseStablecoin) String() string {
	bz := ModuleCdc.MustMarshalJSON(m)
	return string(sdk.MustSortJSON(bz))
}

// StablecoinSupply represents the supply information of a stablecoin
type StablecoinSupply struct {
	Denom              string    `json:"denom"`
	TotalSupply        sdk.Coins `json:"total_supply"`
	CirculatingSupply  sdk.Coins `json:"circulating_supply"`
	LockedSupply       sdk.Coins `json:"locked_supply"`
	ReserveBalance     sdk.Coins `json:"reserve_balance"`
	CollateralizationRatio math.LegacyDec `json:"collateralization_ratio"`
}

