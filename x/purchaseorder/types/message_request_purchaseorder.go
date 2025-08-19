package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestPurchaseorder{}

func NewMsgRequestPurchaseorder(creator string, did string, uri string, amount string, state string, seller string) *MsgRequestPurchaseorder {
	return &MsgRequestPurchaseorder{
		Creator:   creator,
		Did:       did,
		Uri:       uri,
		Amount:    amount,
		State:     state,
		Purchaser: creator,
		Seller:    seller,
	}
}

func (msg *MsgRequestPurchaseorder) Route() string {
	return RouterKey
}

func (msg *MsgRequestPurchaseorder) Type() string {
	return "RequestPurchaseorder"
}

func (msg *MsgRequestPurchaseorder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestPurchaseorder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestPurchaseorder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
