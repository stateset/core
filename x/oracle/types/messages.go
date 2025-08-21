package types

import (
	"time"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSubmitPriceFeed = "submit_price_feed"
	TypeMsgRegisterOracle = "register_oracle"
	TypeMsgUpdateOracle = "update_oracle"
	TypeMsgRemoveOracle = "remove_oracle"
	TypeMsgRequestPrice = "request_price"
)

var (
	_ sdk.Msg = &MsgSubmitPriceFeed{}
	_ sdk.Msg = &MsgRegisterOracle{}
	_ sdk.Msg = &MsgUpdateOracle{}
	_ sdk.Msg = &MsgRemoveOracle{}
	_ sdk.Msg = &MsgRequestPrice{}
)

// MsgSubmitPriceFeed submits a price feed from an oracle provider
type MsgSubmitPriceFeed struct {
	Provider   string    `json:"provider"`
	FeedID     string    `json:"feed_id"`
	Asset      string    `json:"asset"`
	Price      sdk.Dec   `json:"price"`
	Confidence sdk.Dec   `json:"confidence"`
	Volume     sdk.Int   `json:"volume"`
	Expiry     time.Time `json:"expiry"`
}

func NewMsgSubmitPriceFeed(provider, feedID, asset string, price, confidence sdk.Dec, volume sdk.Int, expiry time.Time) *MsgSubmitPriceFeed {
	return &MsgSubmitPriceFeed{
		Provider:   provider,
		FeedID:     feedID,
		Asset:      asset,
		Price:      price,
		Confidence: confidence,
		Volume:     volume,
		Expiry:     expiry,
	}
}

func (msg *MsgSubmitPriceFeed) GetSigners() []sdk.AccAddress {
	provider, err := sdk.AccAddressFromBech32(msg.Provider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{provider}
}

func (msg *MsgSubmitPriceFeed) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Provider); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	if msg.FeedID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "feed ID cannot be empty")
	}
	if msg.Asset == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "asset cannot be empty")
	}
	if msg.Price.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "price cannot be negative")
	}
	if msg.Confidence.LT(sdk.ZeroDec()) || msg.Confidence.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "confidence must be between 0 and 1")
	}
	if msg.Volume.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "volume cannot be negative")
	}
	return nil
}

// MsgRegisterOracle registers a new oracle provider
type MsgRegisterOracle struct {
	Authority     string  `json:"authority"`
	Name          string  `json:"name"`
	ProviderAddr  string  `json:"provider_addr"`
	Priority      uint32  `json:"priority"`
	MinSubmissions uint32 `json:"min_submissions"`
	MaxDeviation  sdk.Dec `json:"max_deviation"`
}

func NewMsgRegisterOracle(authority, name, providerAddr string, priority, minSubmissions uint32, maxDeviation sdk.Dec) *MsgRegisterOracle {
	return &MsgRegisterOracle{
		Authority:      authority,
		Name:          name,
		ProviderAddr:  providerAddr,
		Priority:      priority,
		MinSubmissions: minSubmissions,
		MaxDeviation:  maxDeviation,
	}
}

func (msg *MsgRegisterOracle) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgRegisterOracle) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ProviderAddr); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}
	if msg.MinSubmissions == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "minimum submissions must be greater than 0")
	}
	if msg.MaxDeviation.IsNegative() || msg.MaxDeviation.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max deviation must be between 0 and 1")
	}
	return nil
}

// MsgUpdateOracle updates an existing oracle provider
type MsgUpdateOracle struct {
	Authority     string  `json:"authority"`
	ProviderAddr  string  `json:"provider_addr"`
	Active        bool    `json:"active"`
	Priority      uint32  `json:"priority"`
	MinSubmissions uint32 `json:"min_submissions"`
	MaxDeviation  sdk.Dec `json:"max_deviation"`
}

func NewMsgUpdateOracle(authority, providerAddr string, active bool, priority, minSubmissions uint32, maxDeviation sdk.Dec) *MsgUpdateOracle {
	return &MsgUpdateOracle{
		Authority:      authority,
		ProviderAddr:  providerAddr,
		Active:        active,
		Priority:      priority,
		MinSubmissions: minSubmissions,
		MaxDeviation:  maxDeviation,
	}
}

func (msg *MsgUpdateOracle) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgUpdateOracle) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ProviderAddr); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	if msg.MinSubmissions == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "minimum submissions must be greater than 0")
	}
	if msg.MaxDeviation.IsNegative() || msg.MaxDeviation.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max deviation must be between 0 and 1")
	}
	return nil
}

// MsgRemoveOracle removes an oracle provider
type MsgRemoveOracle struct {
	Authority    string `json:"authority"`
	ProviderAddr string `json:"provider_addr"`
	Reason       string `json:"reason"`
}

func NewMsgRemoveOracle(authority, providerAddr, reason string) *MsgRemoveOracle {
	return &MsgRemoveOracle{
		Authority:    authority,
		ProviderAddr: providerAddr,
		Reason:       reason,
	}
}

func (msg *MsgRemoveOracle) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgRemoveOracle) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ProviderAddr); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	return nil
}

// MsgRequestPrice requests a price update for an asset
type MsgRequestPrice struct {
	Requester string `json:"requester"`
	Asset     string `json:"asset"`
	Urgent    bool   `json:"urgent"`
}

func NewMsgRequestPrice(requester, asset string, urgent bool) *MsgRequestPrice {
	return &MsgRequestPrice{
		Requester: requester,
		Asset:     asset,
		Urgent:    urgent,
	}
}

func (msg *MsgRequestPrice) GetSigners() []sdk.AccAddress {
	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{requester}
}

func (msg *MsgRequestPrice) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Requester); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid requester address (%s)", err)
	}
	if msg.Asset == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "asset cannot be empty")
	}
	return nil
}