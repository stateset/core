package types

import (
	"context"
	"time"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgServer is the server API for oracle module messages
type MsgServer interface {
	SubmitPriceFeed(context.Context, *MsgSubmitPriceFeed) (*MsgSubmitPriceFeedResponse, error)
	RegisterOracle(context.Context, *MsgRegisterOracle) (*MsgRegisterOracleResponse, error)
	UpdateOracle(context.Context, *MsgUpdateOracle) (*MsgUpdateOracleResponse, error)
	RemoveOracle(context.Context, *MsgRemoveOracle) (*MsgRemoveOracleResponse, error)
	RequestPrice(context.Context, *MsgRequestPrice) (*MsgRequestPriceResponse, error)
}

// Response types for oracle messages

type MsgSubmitPriceFeedResponse struct {
	FeedId  string `json:"feed_id"`
	Success bool   `json:"success"`
}

type MsgRegisterOracleResponse struct {
	Success bool `json:"success"`
}

type MsgUpdateOracleResponse struct {
	Success bool `json:"success"`
}

type MsgRemoveOracleResponse struct {
	Success bool `json:"success"`
}

type MsgRequestPriceResponse struct {
	Asset        string    `json:"asset"`
	Price        sdk.Dec   `json:"price"`
	LastUpdate   time.Time `json:"last_update"`
	NextUpdate   time.Time `json:"next_update"`
	Confidence   sdk.Dec   `json:"confidence"`
	NumProviders uint32    `json:"num_providers"`
}

// Placeholder for service descriptor (would be generated from proto in production)
var _Msg_serviceDesc = struct{}{}

// GetSignBytes implementations for messages
func (msg *MsgSubmitPriceFeed) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRegisterOracle) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgUpdateOracle) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRemoveOracle) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRequestPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}