package types

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Params represents module parameters
type Params struct {
	MaxStablecoins   uint64 `json:"max_stablecoins" yaml:"max_stablecoins"`
	MinInitialSupply string `json:"min_initial_supply" yaml:"min_initial_supply"`
	MaxInitialSupply string `json:"max_initial_supply" yaml:"max_initial_supply"`
	CreationFee      string `json:"creation_fee" yaml:"creation_fee"`
	MinReserveRatio  string `json:"min_reserve_ratio" yaml:"min_reserve_ratio"`
	MaxFeePercentage string `json:"max_fee_percentage" yaml:"max_fee_percentage"`
}

// Message type URLs
const (
	TypeMsgCreateStablecoin      = "create_stablecoin"
	TypeMsgUpdateStablecoin      = "update_stablecoin"
	TypeMsgMintStablecoin        = "mint_stablecoin"
	TypeMsgBurnStablecoin        = "burn_stablecoin"
	TypeMsgPauseStablecoin       = "pause_stablecoin"
	TypeMsgUnpauseStablecoin     = "unpause_stablecoin"
	TypeMsgUpdatePriceData       = "update_price_data"
	TypeMsgUpdateReserves        = "update_reserves"
	TypeMsgWhitelistAddress      = "whitelist_address"
	TypeMsgBlacklistAddress      = "blacklist_address"
	TypeMsgRemoveFromWhitelist   = "remove_from_whitelist"
	TypeMsgRemoveFromBlacklist   = "remove_from_blacklist"
)

// Message type implementations for Cosmos SDK messages

// MsgCreateStablecoin
var _ sdk.Msg = &MsgCreateStablecoin{}

func NewMsgCreateStablecoin(
	creator string,
	denom string,
	name string,
	symbol string,
	decimals uint32,
	description string,
	maxSupply sdk.Int,
	pegInfo *PegInfo,
	reserveInfo *ReserveInfo,
	stabilityMechanism string,
	feeInfo *FeeInfo,
	accessControl *AccessControlInfo,
	metadata string,
) *MsgCreateStablecoin {
	return &MsgCreateStablecoin{
		Creator:            creator,
		Denom:             denom,
		Name:              name,
		Symbol:            symbol,
		Decimals:          decimals,
		Description:       description,
		MaxSupply:         maxSupply,
		PegInfo:           pegInfo,
		ReserveInfo:       reserveInfo,
		StabilityMechanism: stabilityMechanism,
		FeeInfo:           feeInfo,
		AccessControl:     accessControl,
		Metadata:          metadata,
	}
}

func (msg *MsgCreateStablecoin) Route() string {
	return RouterKey
}

func (msg *MsgCreateStablecoin) Type() string {
	return TypeMsgCreateStablecoin
}

func (msg *MsgCreateStablecoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStablecoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if strings.TrimSpace(msg.Denom) == "" {
		return sdkerrors.Wrap(ErrInvalidStablecoinDenom, "denom cannot be empty")
	}

	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrInvalidName, "name cannot be empty")
	}

	if strings.TrimSpace(msg.Symbol) == "" {
		return sdkerrors.Wrap(ErrInvalidSymbol, "symbol cannot be empty")
	}

	if msg.Decimals > 18 {
		return sdkerrors.Wrap(ErrInvalidDecimalPlaces, "decimals cannot exceed 18")
	}

	if msg.MaxSupply.IsNil() || !msg.MaxSupply.IsPositive() {
		return sdkerrors.Wrap(ErrInvalidAmount, "max supply must be positive")
	}

	return nil
}

// MsgUpdateStablecoin
var _ sdk.Msg = &MsgUpdateStablecoin{}

func NewMsgUpdateStablecoin(
	creator string,
	denom string,
	name string,
	description string,
	pegInfo *PegInfo,
	feeInfo *FeeInfo,
	accessControl *AccessControlInfo,
	metadata string,
) *MsgUpdateStablecoin {
	return &MsgUpdateStablecoin{
		Creator:       creator,
		Denom:        denom,
		Name:         name,
		Description:  description,
		PegInfo:      pegInfo,
		FeeInfo:      feeInfo,
		AccessControl: accessControl,
		Metadata:     metadata,
	}
}

func (msg *MsgUpdateStablecoin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStablecoin) Type() string {
	return TypeMsgUpdateStablecoin
}

func (msg *MsgUpdateStablecoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStablecoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if strings.TrimSpace(msg.Denom) == "" {
		return sdkerrors.Wrap(ErrInvalidStablecoinDenom, "denom cannot be empty")
	}

	return nil
}

// MsgMintStablecoin
var _ sdk.Msg = &MsgMintStablecoin{}

func NewMsgMintStablecoin(
	creator string,
	denom string,
	amount sdk.Int,
	recipient string,
) *MsgMintStablecoin {
	return &MsgMintStablecoin{
		Creator:   creator,
		Denom:    denom,
		Amount:   amount,
		Recipient: recipient,
	}
}

func (msg *MsgMintStablecoin) Route() string {
	return RouterKey
}

func (msg *MsgMintStablecoin) Type() string {
	return TypeMsgMintStablecoin
}

func (msg *MsgMintStablecoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintStablecoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if strings.TrimSpace(msg.Denom) == "" {
		return sdkerrors.Wrap(ErrInvalidStablecoinDenom, "denom cannot be empty")
	}

	if msg.Amount.IsNil() || !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(ErrInvalidAmount, "amount must be positive")
	}

	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	return nil
}

// MsgBurnStablecoin
var _ sdk.Msg = &MsgBurnStablecoin{}

func NewMsgBurnStablecoin(
	creator string,
	denom string,
	amount sdk.Int,
) *MsgBurnStablecoin {
	return &MsgBurnStablecoin{
		Creator: creator,
		Denom:  denom,
		Amount: amount,
	}
}

func (msg *MsgBurnStablecoin) Route() string {
	return RouterKey
}

func (msg *MsgBurnStablecoin) Type() string {
	return TypeMsgBurnStablecoin
}

func (msg *MsgBurnStablecoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnStablecoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if strings.TrimSpace(msg.Denom) == "" {
		return sdkerrors.Wrap(ErrInvalidStablecoinDenom, "denom cannot be empty")
	}

	if msg.Amount.IsNil() || !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(ErrInvalidAmount, "amount must be positive")
	}

	return nil
}

// MsgPauseStablecoin
var _ sdk.Msg = &MsgPauseStablecoin{}

func NewMsgPauseStablecoin(
	creator string,
	denom string,
	operation string,
	reason string,
) *MsgPauseStablecoin {
	return &MsgPauseStablecoin{
		Creator:   creator,
		Denom:    denom,
		Operation: operation,
		Reason:   reason,
	}
}

func (msg *MsgPauseStablecoin) Route() string {
	return RouterKey
}

func (msg *MsgPauseStablecoin) Type() string {
	return TypeMsgPauseStablecoin
}

func (msg *MsgPauseStablecoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPauseStablecoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPauseStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if strings.TrimSpace(msg.Denom) == "" {
		return sdkerrors.Wrap(ErrInvalidStablecoinDenom, "denom cannot be empty")
	}

	validOperations := []string{"mint", "burn", "transfer", "all"}
	valid := false
	for _, op := range validOperations {
		if msg.Operation == op {
			valid = true
			break
		}
	}
	if !valid {
		return sdkerrors.Wrap(ErrOperationNotAllowed, "invalid operation type")
	}

	return nil
}

// Additional message types would follow the same pattern...
// For brevity, I'll implement a few more key ones

// Helper function to create a new stablecoin with default values
func NewStablecoin(
	denom string,
	name string,
	symbol string,
	decimals uint32,
	description string,
	issuer string,
	admin string,
	maxSupply sdk.Int,
	pegInfo *PegInfo,
	reserveInfo *ReserveInfo,
	stabilityMechanism string,
	feeInfo *FeeInfo,
	accessControl *AccessControlInfo,
	metadata string,
) Stablecoin {
	return Stablecoin{
		Denom:                  denom,
		Name:                   name,
		Symbol:                 symbol,
		Decimals:               decimals,
		Description:            description,
		Issuer:                 issuer,
		Admin:                  admin,
		TotalSupply:            sdk.ZeroInt(),
		MaxSupply:              maxSupply,
		PegInfo:                pegInfo,
		ReserveInfo:            reserveInfo,
		Active:                 true,
		MintPaused:             false,
		BurnPaused:             false,
		TransferPaused:         false,
		Metadata:               metadata,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
		CollateralizationRatio: "1.00",
		StabilityMechanism:     stabilityMechanism,
		FeeInfo:                feeInfo,
		AccessControl:          accessControl,
	}
}

// Message type URLs for ssUSD functionality
const (
	TypeMsgInitializeSSUSD = "initialize_ssusd"
	TypeMsgIssueSSUSD      = "issue_ssusd"
	TypeMsgRedeemSSUSD     = "redeem_ssusd"
)

// MsgInitializeSSUSD initializes the ssUSD stablecoin
type MsgInitializeSSUSD struct {
	Creator string `json:"creator"`
}

var _ sdk.Msg = &MsgInitializeSSUSD{}

func NewMsgInitializeSSUSD(creator string) *MsgInitializeSSUSD {
	return &MsgInitializeSSUSD{
		Creator: creator,
	}
}

func (msg *MsgInitializeSSUSD) Route() string {
	return RouterKey
}

func (msg *MsgInitializeSSUSD) Type() string {
	return TypeMsgInitializeSSUSD
}

func (msg *MsgInitializeSSUSD) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitializeSSUSD) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitializeSSUSD) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// MsgIssueSSUSD issues new ssUSD tokens backed by reserves
type MsgIssueSSUSD struct {
	Creator        string    `json:"creator"`
	Amount         sdk.Int   `json:"amount"`
	ReservePayment sdk.Coins `json:"reserve_payment"`
}

var _ sdk.Msg = &MsgIssueSSUSD{}

func NewMsgIssueSSUSD(creator string, amount sdk.Int, reservePayment sdk.Coins) *MsgIssueSSUSD {
	return &MsgIssueSSUSD{
		Creator:        creator,
		Amount:         amount,
		ReservePayment: reservePayment,
	}
}

func (msg *MsgIssueSSUSD) Route() string {
	return RouterKey
}

func (msg *MsgIssueSSUSD) Type() string {
	return TypeMsgIssueSSUSD
}

func (msg *MsgIssueSSUSD) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssueSSUSD) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueSSUSD) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Amount.IsNil() || !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be positive")
	}

	if !msg.ReservePayment.IsValid() || msg.ReservePayment.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "reserve payment must be valid and non-zero")
	}

	// Validate reserve asset types
	validAssets := map[string]bool{
		"us_cash_token":       true,
		"treasury_bill_token": true,
		"mmf_token":          true,
		"repo_token":         true,
	}

	for _, coin := range msg.ReservePayment {
		if !validAssets[coin.Denom] {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, 
				"invalid reserve asset type: %s. Valid types: us_cash_token, treasury_bill_token, mmf_token, repo_token", 
				coin.Denom)
		}
	}

	return nil
}

// MsgRedeemSSUSD redeems ssUSD tokens for underlying reserves
type MsgRedeemSSUSD struct {
	Creator        string  `json:"creator"`
	SSUSDAmount    sdk.Int `json:"ssusd_amount"`
	PreferredAsset string  `json:"preferred_asset"`
}

var _ sdk.Msg = &MsgRedeemSSUSD{}

func NewMsgRedeemSSUSD(creator string, ssusdAmount sdk.Int, preferredAsset string) *MsgRedeemSSUSD {
	return &MsgRedeemSSUSD{
		Creator:        creator,
		SSUSDAmount:    ssusdAmount,
		PreferredAsset: preferredAsset,
	}
}

func (msg *MsgRedeemSSUSD) Route() string {
	return RouterKey
}

func (msg *MsgRedeemSSUSD) Type() string {
	return TypeMsgRedeemSSUSD
}

func (msg *MsgRedeemSSUSD) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRedeemSSUSD) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedeemSSUSD) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.SSUSDAmount.IsNil() || !msg.SSUSDAmount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "ssUSD amount must be positive")
	}

	// Validate preferred asset if specified
	if msg.PreferredAsset != "" {
		validAssets := map[string]bool{
			"us_cash_token":       true,
			"treasury_bill_token": true,
			"mmf_token":          true,
			"repo_token":         true,
		}

		if !validAssets[msg.PreferredAsset] {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, 
				"invalid preferred asset: %s. Valid types: us_cash_token, treasury_bill_token, mmf_token, repo_token", 
				msg.PreferredAsset)
		}
	}

	return nil
}

// Response types for ssUSD messages

// MsgInitializeSSUSDResponse response type for MsgInitializeSSUSD
type MsgInitializeSSUSDResponse struct {
	Success bool `json:"success"`
}

// MsgIssueSSUSDResponse response type for MsgIssueSSUSD
type MsgIssueSSUSDResponse struct {
	AmountIssued sdk.Int `json:"amount_issued"`
}

// MsgRedeemSSUSDResponse response type for MsgRedeemSSUSD
type MsgRedeemSSUSDResponse struct {
	AmountRedeemed sdk.Int `json:"amount_redeemed"`
}