package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
)

// Message types
const (
	TypeMsgDepositForBurn                       = "deposit_for_burn"
	TypeMsgDepositForBurnWithCaller            = "deposit_for_burn_with_caller"
	TypeMsgReceiveMessage                      = "receive_message"
	TypeMsgSendMessage                         = "send_message"
	TypeMsgSendMessageWithCaller               = "send_message_with_caller"
	TypeMsgReplaceDepositForBurn               = "replace_deposit_for_burn"
	TypeMsgReplaceMessage                      = "replace_message"
	TypeMsgAcceptOwner                         = "accept_owner"
	TypeMsgAddRemoteTokenMessenger             = "add_remote_token_messenger"
	TypeMsgDisableAttester                     = "disable_attester"
	TypeMsgEnableAttester                      = "enable_attester"
	TypeMsgLinkTokenPair                       = "link_token_pair"
	TypeMsgPauseBurningAndMinting              = "pause_burning_and_minting"
	TypeMsgPauseSendingAndReceivingMessages    = "pause_sending_and_receiving_messages"
	TypeMsgRemoveRemoteTokenMessenger          = "remove_remote_token_messenger"
	TypeMsgSetMaxBurnAmountPerMessage          = "set_max_burn_amount_per_message"
	TypeMsgUnlinkTokenPair                     = "unlink_token_pair"
	TypeMsgUnpauseBurningAndMinting            = "unpause_burning_and_minting"
	TypeMsgUnpauseSendingAndReceivingMessages  = "unpause_sending_and_receiving_messages"
	TypeMsgUpdateAttesterManager               = "update_attester_manager"
	TypeMsgUpdateMaxMessageBodySize            = "update_max_message_body_size"
	TypeMsgUpdateOwner                         = "update_owner"
	TypeMsgUpdatePauser                        = "update_pauser"
	TypeMsgUpdateSignatureThreshold            = "update_signature_threshold"
	TypeMsgUpdateTokenController               = "update_token_controller"
)

// Ensure messages implement sdk.Msg interface
var (
	_ sdk.Msg = &MsgDepositForBurn{}
	_ sdk.Msg = &MsgDepositForBurnWithCaller{}
	_ sdk.Msg = &MsgReceiveMessage{}
	_ sdk.Msg = &MsgSendMessage{}
	_ sdk.Msg = &MsgSendMessageWithCaller{}
	_ sdk.Msg = &MsgReplaceDepositForBurn{}
	_ sdk.Msg = &MsgReplaceMessage{}
	_ sdk.Msg = &MsgAcceptOwner{}
	_ sdk.Msg = &MsgAddRemoteTokenMessenger{}
	_ sdk.Msg = &MsgDisableAttester{}
	_ sdk.Msg = &MsgEnableAttester{}
	_ sdk.Msg = &MsgLinkTokenPair{}
	_ sdk.Msg = &MsgPauseBurningAndMinting{}
	_ sdk.Msg = &MsgPauseSendingAndReceivingMessages{}
	_ sdk.Msg = &MsgRemoveRemoteTokenMessenger{}
	_ sdk.Msg = &MsgSetMaxBurnAmountPerMessage{}
	_ sdk.Msg = &MsgUnlinkTokenPair{}
	_ sdk.Msg = &MsgUnpauseBurningAndMinting{}
	_ sdk.Msg = &MsgUnpauseSendingAndReceivingMessages{}
	_ sdk.Msg = &MsgUpdateAttesterManager{}
	_ sdk.Msg = &MsgUpdateMaxMessageBodySize{}
	_ sdk.Msg = &MsgUpdateOwner{}
	_ sdk.Msg = &MsgUpdatePauser{}
	_ sdk.Msg = &MsgUpdateSignatureThreshold{}
	_ sdk.Msg = &MsgUpdateTokenController{}
)

// MsgDepositForBurn
func NewMsgDepositForBurn(from string, amount sdk.Coin, destinationDomain uint32, mintRecipient []byte, burnToken string) *MsgDepositForBurn {
	return &MsgDepositForBurn{
		From:              from,
		Amount:            amount,
		DestinationDomain: destinationDomain,
		MintRecipient:     mintRecipient,
		BurnToken:         burnToken,
	}
}

func (msg *MsgDepositForBurn) Route() string { return RouterKey }
func (msg *MsgDepositForBurn) Type() string  { return TypeMsgDepositForBurn }

func (msg *MsgDepositForBurn) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgDepositForBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositForBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount must be positive")
	}

	if len(msg.MintRecipient) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "mint recipient must be %d bytes", AddressSize)
	}

	if IsZeroBytes(msg.MintRecipient) {
		return sdkerrors.Wrap(ErrInvalidAddress, "mint recipient cannot be zero address")
	}

	if msg.BurnToken == "" {
		return sdkerrors.Wrap(ErrInvalidToken, "burn token cannot be empty")
	}

	return nil
}

// MsgDepositForBurnWithCaller
func NewMsgDepositForBurnWithCaller(from string, amount sdk.Coin, destinationDomain uint32, mintRecipient []byte, burnToken string, destinationCaller []byte) *MsgDepositForBurnWithCaller {
	return &MsgDepositForBurnWithCaller{
		From:              from,
		Amount:            amount,
		DestinationDomain: destinationDomain,
		MintRecipient:     mintRecipient,
		BurnToken:         burnToken,
		DestinationCaller: destinationCaller,
	}
}

func (msg *MsgDepositForBurnWithCaller) Route() string { return RouterKey }
func (msg *MsgDepositForBurnWithCaller) Type() string  { return TypeMsgDepositForBurnWithCaller }

func (msg *MsgDepositForBurnWithCaller) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgDepositForBurnWithCaller) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositForBurnWithCaller) ValidateBasic() error {
	// Validate base deposit for burn fields
	baseMsg := &MsgDepositForBurn{
		From:              msg.From,
		Amount:            msg.Amount,
		DestinationDomain: msg.DestinationDomain,
		MintRecipient:     msg.MintRecipient,
		BurnToken:         msg.BurnToken,
	}
	if err := baseMsg.ValidateBasic(); err != nil {
		return err
	}

	if len(msg.DestinationCaller) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "destination caller must be %d bytes", AddressSize)
	}

	if IsZeroBytes(msg.DestinationCaller) {
		return sdkerrors.Wrap(ErrInvalidAddress, "destination caller cannot be zero address")
	}

	return nil
}

// MsgReceiveMessage
func NewMsgReceiveMessage(from string, message []byte, attestation []byte) *MsgReceiveMessage {
	return &MsgReceiveMessage{
		From:        from,
		Message:     message,
		Attestation: attestation,
	}
}

func (msg *MsgReceiveMessage) Route() string { return RouterKey }
func (msg *MsgReceiveMessage) Type() string  { return TypeMsgReceiveMessage }

func (msg *MsgReceiveMessage) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgReceiveMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReceiveMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if len(msg.Message) == 0 {
		return sdkerrors.Wrap(ErrInvalidMessage, "message cannot be empty")
	}

	if len(msg.Attestation) == 0 {
		return sdkerrors.Wrap(ErrInvalidAttestation, "attestation cannot be empty")
	}

	// Validate message can be decoded
	_, err = DecodeMessage(msg.Message)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidMessage, err.Error())
	}

	return nil
}

// MsgSendMessage
func NewMsgSendMessage(from string, destinationDomain uint32, recipient []byte, messageBody []byte) *MsgSendMessage {
	return &MsgSendMessage{
		From:              from,
		DestinationDomain: destinationDomain,
		Recipient:         recipient,
		MessageBody:       messageBody,
	}
}

func (msg *MsgSendMessage) Route() string { return RouterKey }
func (msg *MsgSendMessage) Type() string  { return TypeMsgSendMessage }

func (msg *MsgSendMessage) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgSendMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if len(msg.Recipient) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "recipient must be %d bytes", AddressSize)
	}

	if IsZeroBytes(msg.Recipient) {
		return sdkerrors.Wrap(ErrInvalidAddress, "recipient cannot be zero address")
	}

	return nil
}

// MsgSendMessageWithCaller
func NewMsgSendMessageWithCaller(from string, destinationDomain uint32, recipient []byte, messageBody []byte, destinationCaller []byte) *MsgSendMessageWithCaller {
	return &MsgSendMessageWithCaller{
		From:              from,
		DestinationDomain: destinationDomain,
		Recipient:         recipient,
		MessageBody:       messageBody,
		DestinationCaller: destinationCaller,
	}
}

func (msg *MsgSendMessageWithCaller) Route() string { return RouterKey }
func (msg *MsgSendMessageWithCaller) Type() string  { return TypeMsgSendMessageWithCaller }

func (msg *MsgSendMessageWithCaller) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgSendMessageWithCaller) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendMessageWithCaller) ValidateBasic() error {
	// Validate base send message fields
	baseMsg := &MsgSendMessage{
		From:              msg.From,
		DestinationDomain: msg.DestinationDomain,
		Recipient:         msg.Recipient,
		MessageBody:       msg.MessageBody,
	}
	if err := baseMsg.ValidateBasic(); err != nil {
		return err
	}

	if len(msg.DestinationCaller) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "destination caller must be %d bytes", AddressSize)
	}

	if IsZeroBytes(msg.DestinationCaller) {
		return sdkerrors.Wrap(ErrInvalidAddress, "destination caller cannot be zero address")
	}

	return nil
}

// MsgReplaceDepositForBurn
func NewMsgReplaceDepositForBurn(from string, originalMessage []byte, originalAttestation []byte, newDestinationCaller []byte, newMintRecipient []byte) *MsgReplaceDepositForBurn {
	return &MsgReplaceDepositForBurn{
		From:                 from,
		OriginalMessage:      originalMessage,
		OriginalAttestation:  originalAttestation,
		NewDestinationCaller: newDestinationCaller,
		NewMintRecipient:     newMintRecipient,
	}
}

func (msg *MsgReplaceDepositForBurn) Route() string { return RouterKey }
func (msg *MsgReplaceDepositForBurn) Type() string  { return TypeMsgReplaceDepositForBurn }

func (msg *MsgReplaceDepositForBurn) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgReplaceDepositForBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReplaceDepositForBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if len(msg.OriginalMessage) == 0 {
		return sdkerrors.Wrap(ErrInvalidMessage, "original message cannot be empty")
	}

	if len(msg.OriginalAttestation) == 0 {
		return sdkerrors.Wrap(ErrInvalidAttestation, "original attestation cannot be empty")
	}

	if len(msg.NewMintRecipient) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "new mint recipient must be %d bytes", AddressSize)
	}

	if len(msg.NewDestinationCaller) != 0 && len(msg.NewDestinationCaller) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "new destination caller must be empty or %d bytes", AddressSize)
	}

	return nil
}

// MsgReplaceMessage
func NewMsgReplaceMessage(from string, originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller []byte) *MsgReplaceMessage {
	return &MsgReplaceMessage{
		From:                 from,
		OriginalMessage:      originalMessage,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       newMessageBody,
		NewDestinationCaller: newDestinationCaller,
	}
}

func (msg *MsgReplaceMessage) Route() string { return RouterKey }
func (msg *MsgReplaceMessage) Type() string  { return TypeMsgReplaceMessage }

func (msg *MsgReplaceMessage) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgReplaceMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReplaceMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if len(msg.OriginalMessage) == 0 {
		return sdkerrors.Wrap(ErrInvalidMessage, "original message cannot be empty")
	}

	if len(msg.OriginalAttestation) == 0 {
		return sdkerrors.Wrap(ErrInvalidAttestation, "original attestation cannot be empty")
	}

	if len(msg.NewDestinationCaller) != 0 && len(msg.NewDestinationCaller) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "new destination caller must be empty or %d bytes", AddressSize)
	}

	return nil
}

// Admin messages implementation follows...

// MsgAcceptOwner
func NewMsgAcceptOwner(from string) *MsgAcceptOwner {
	return &MsgAcceptOwner{From: from}
}

func (msg *MsgAcceptOwner) Route() string { return RouterKey }
func (msg *MsgAcceptOwner) Type() string  { return TypeMsgAcceptOwner }

func (msg *MsgAcceptOwner) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgAcceptOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}

// Additional message implementations for admin functions...
// MsgUpdateOwner
func NewMsgUpdateOwner(from string, newOwner string) *MsgUpdateOwner {
	return &MsgUpdateOwner{
		From:     from,
		NewOwner: newOwner,
	}
}

func (msg *MsgUpdateOwner) Route() string { return RouterKey }
func (msg *MsgUpdateOwner) Type() string  { return TypeMsgUpdateOwner }

func (msg *MsgUpdateOwner) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdateOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new owner address (%s)", err)
	}

	return nil
}

// MsgLinkTokenPair
func NewMsgLinkTokenPair(from string, localToken string, remoteToken []byte, remoteDomain uint32) *MsgLinkTokenPair {
	return &MsgLinkTokenPair{
		From:         from,
		LocalToken:   localToken,
		RemoteToken:  remoteToken,
		RemoteDomain: remoteDomain,
	}
}

func (msg *MsgLinkTokenPair) Route() string { return RouterKey }
func (msg *MsgLinkTokenPair) Type() string  { return TypeMsgLinkTokenPair }

func (msg *MsgLinkTokenPair) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgLinkTokenPair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLinkTokenPair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if msg.LocalToken == "" {
		return sdkerrors.Wrap(ErrInvalidToken, "local token cannot be empty")
	}

	if len(msg.RemoteToken) != AddressSize {
		return errorsmod.Wrapf(ErrInvalidAddressLength, "remote token must be %d bytes", AddressSize)
	}

	return nil
}

// MsgEnableAttester
func NewMsgEnableAttester(from string, attester string) *MsgEnableAttester {
	return &MsgEnableAttester{
		From:     from,
		Attester: attester,
	}
}

func (msg *MsgEnableAttester) Route() string { return RouterKey }
func (msg *MsgEnableAttester) Type() string  { return TypeMsgEnableAttester }

func (msg *MsgEnableAttester) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgEnableAttester) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableAttester) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Attester)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid attester address (%s)", err)
	}

	return nil
}

// MsgDisableAttester
func NewMsgDisableAttester(from string, attester string) *MsgDisableAttester {
	return &MsgDisableAttester{
		From:     from,
		Attester: attester,
	}
}

func (msg *MsgDisableAttester) Route() string { return RouterKey }
func (msg *MsgDisableAttester) Type() string  { return TypeMsgDisableAttester }

func (msg *MsgDisableAttester) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgDisableAttester) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableAttester) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Attester)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid attester address (%s)", err)
	}

	return nil
}

// Pause/Unpause messages
func NewMsgPauseBurningAndMinting(from string) *MsgPauseBurningAndMinting {
	return &MsgPauseBurningAndMinting{From: from}
}

func (msg *MsgPauseBurningAndMinting) Route() string { return RouterKey }
func (msg *MsgPauseBurningAndMinting) Type() string  { return TypeMsgPauseBurningAndMinting }

func (msg *MsgPauseBurningAndMinting) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgPauseBurningAndMinting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPauseBurningAndMinting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}

func NewMsgUnpauseBurningAndMinting(from string) *MsgUnpauseBurningAndMinting {
	return &MsgUnpauseBurningAndMinting{From: from}
}

func (msg *MsgUnpauseBurningAndMinting) Route() string { return RouterKey }
func (msg *MsgUnpauseBurningAndMinting) Type() string  { return TypeMsgUnpauseBurningAndMinting }

func (msg *MsgUnpauseBurningAndMinting) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUnpauseBurningAndMinting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnpauseBurningAndMinting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}