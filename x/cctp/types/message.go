package types

import (
	"encoding/binary"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Message represents the core CCTP message structure
type Message struct {
	Version           uint32 // Message format version
	SourceDomain      uint32 // Domain of home chain
	DestinationDomain uint32 // Domain of destination chain
	Nonce             uint64 // Destination-specific nonce
	Sender            []byte // Address of sender on source chain as bytes32
	Recipient         []byte // Address of recipient on destination chain as bytes32
	DestinationCaller []byte // Address permitted to call receiveMessage() on destination chain, or 0 if any caller is valid
	MessageBody       []byte // Application-specific message body
}

// BurnMessage represents a burn message for token transfers
type BurnMessage struct {
	Version     uint32 // Burn message body version
	BurnToken   []byte // Address of burned token on source domain as bytes32
	MintRecipient []byte // Address to receive minted tokens on destination domain as bytes32
	Amount      sdk.Int // Amount of burned tokens
	MessageSender []byte // Address of caller of depositForBurn on source domain as bytes32
}

// MessageDomains contains source and destination domain information
type MessageDomains struct {
	SourceDomain      uint32
	DestinationDomain uint32
	Version           uint32
}

// Constants for message encoding
const (
	MessageHeaderSize    = 116 // Size of message header in bytes
	BurnMessageBodySize = 116 // Size of burn message body in bytes
	MessageBodyVersion  = 0   // Current message body version
	MessageVersion      = 0   // Current message version
	AddressSize         = 32  // Size of address in bytes
	NobleChainDomain    = 4   // Noble chain domain ID
)

// NewMessage creates a new CCTP message
func NewMessage(
	version uint32,
	sourceDomain uint32,
	destinationDomain uint32,
	nonce uint64,
	sender []byte,
	recipient []byte,
	destinationCaller []byte,
	messageBody []byte,
) *Message {
	return &Message{
		Version:           version,
		SourceDomain:      sourceDomain,
		DestinationDomain: destinationDomain,
		Nonce:             nonce,
		Sender:            sender,
		Recipient:         recipient,
		DestinationCaller: destinationCaller,
		MessageBody:       messageBody,
	}
}

// NewBurnMessage creates a new burn message
func NewBurnMessage(
	version uint32,
	burnToken []byte,
	mintRecipient []byte,
	amount sdk.Int,
	messageSender []byte,
) *BurnMessage {
	return &BurnMessage{
		Version:       version,
		BurnToken:     burnToken,
		MintRecipient: mintRecipient,
		Amount:        amount,
		MessageSender: messageSender,
	}
}

// Validate validates the message structure
func (m *Message) Validate() error {
	if m.Version != MessageVersion {
		return sdkerrors.Wrapf(ErrInvalidMessageVersion, "expected %d, got %d", MessageVersion, m.Version)
	}

	if len(m.Sender) != AddressSize {
		return sdkerrors.Wrapf(ErrInvalidAddressLength, "sender address must be %d bytes", AddressSize)
	}

	if len(m.Recipient) != AddressSize {
		return sdkerrors.Wrapf(ErrInvalidAddressLength, "recipient address must be %d bytes", AddressSize)
	}

	if len(m.DestinationCaller) != 0 && len(m.DestinationCaller) != AddressSize {
		return sdkerrors.Wrapf(ErrInvalidAddressLength, "destination caller must be empty or %d bytes", AddressSize)
	}

	return nil
}

// Validate validates the burn message structure
func (bm *BurnMessage) Validate() error {
	if bm.Version != MessageBodyVersion {
		return sdkerrors.Wrapf(ErrInvalidMessageBodyVersion, "expected %d, got %d", MessageBodyVersion, bm.Version)
	}

	if len(bm.BurnToken) != AddressSize {
		return sdkerrors.Wrapf(ErrInvalidAddressLength, "burn token address must be %d bytes", AddressSize)
	}

	if len(bm.MintRecipient) != AddressSize {
		return sdkerrors.Wrapf(ErrInvalidAddressLength, "mint recipient address must be %d bytes", AddressSize)
	}

	if len(bm.MessageSender) != AddressSize {
		return sdkerrors.Wrapf(ErrInvalidAddressLength, "message sender address must be %d bytes", AddressSize)
	}

	if bm.Amount.IsNil() || bm.Amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAmount, "amount must be positive")
	}

	return nil
}

// Encode serializes the message to bytes
func (m *Message) Encode() []byte {
	buf := make([]byte, MessageHeaderSize+len(m.MessageBody))
	
	binary.BigEndian.PutUint32(buf[0:4], m.Version)
	binary.BigEndian.PutUint32(buf[4:8], m.SourceDomain)
	binary.BigEndian.PutUint32(buf[8:12], m.DestinationDomain)
	binary.BigEndian.PutUint64(buf[12:20], m.Nonce)
	
	copy(buf[20:52], m.Sender)
	copy(buf[52:84], m.Recipient)
	copy(buf[84:116], m.DestinationCaller)
	
	copy(buf[MessageHeaderSize:], m.MessageBody)
	
	return buf
}

// Decode deserializes bytes to a message
func DecodeMessage(data []byte) (*Message, error) {
	if len(data) < MessageHeaderSize {
		return nil, sdkerrors.Wrap(ErrInvalidMessageLength, "message too short")
	}

	m := &Message{
		Version:           binary.BigEndian.Uint32(data[0:4]),
		SourceDomain:      binary.BigEndian.Uint32(data[4:8]),
		DestinationDomain: binary.BigEndian.Uint32(data[8:12]),
		Nonce:             binary.BigEndian.Uint64(data[12:20]),
		Sender:            make([]byte, AddressSize),
		Recipient:         make([]byte, AddressSize),
		DestinationCaller: make([]byte, AddressSize),
		MessageBody:       make([]byte, len(data)-MessageHeaderSize),
	}

	copy(m.Sender, data[20:52])
	copy(m.Recipient, data[52:84])
	copy(m.DestinationCaller, data[84:116])
	copy(m.MessageBody, data[MessageHeaderSize:])

	return m, nil
}

// Encode serializes the burn message to bytes
func (bm *BurnMessage) Encode() []byte {
	buf := make([]byte, BurnMessageBodySize)
	
	binary.BigEndian.PutUint32(buf[0:4], bm.Version)
	copy(buf[4:36], bm.BurnToken)
	copy(buf[36:68], bm.MintRecipient)
	
	// Amount is encoded as bytes32
	amountBytes := bm.Amount.BigInt().Bytes()
	if len(amountBytes) > 32 {
		// This shouldn't happen in practice with reasonable token amounts
		amountBytes = amountBytes[len(amountBytes)-32:]
	}
	copy(buf[68+32-len(amountBytes):100], amountBytes)
	
	copy(buf[100:132], bm.MessageSender)
	
	return buf
}

// DecodeBurnMessage deserializes bytes to a burn message
func DecodeBurnMessage(data []byte) (*BurnMessage, error) {
	if len(data) < BurnMessageBodySize {
		return nil, sdkerrors.Wrap(ErrInvalidMessageLength, "burn message too short")
	}

	bm := &BurnMessage{
		Version:       binary.BigEndian.Uint32(data[0:4]),
		BurnToken:     make([]byte, AddressSize),
		MintRecipient: make([]byte, AddressSize),
		MessageSender: make([]byte, AddressSize),
	}

	copy(bm.BurnToken, data[4:36])
	copy(bm.MintRecipient, data[36:68])
	copy(bm.MessageSender, data[100:132])

	// Decode amount from bytes32
	amountBytes := data[68:100]
	bm.Amount = sdk.NewIntFromBigInt(sdk.NewIntFromBigInt(sdk.ZeroInt().BigInt()).SetBytes(amountBytes).BigInt())

	return bm, nil
}

// Hash calculates the keccak256 hash of the message
func (m *Message) Hash() []byte {
	encoded := m.Encode()
	return Keccak256Hash(encoded)
}

// GetMessageHash returns the hash of the message for attestation
func (m *Message) GetMessageHash() []byte {
	return m.Hash()
}

// String returns a string representation of the message
func (m *Message) String() string {
	return fmt.Sprintf(`Message{
		Version: %d,
		SourceDomain: %d,
		DestinationDomain: %d,
		Nonce: %d,
		Sender: %x,
		Recipient: %x,
		DestinationCaller: %x,
		MessageBodyLength: %d
	}`, m.Version, m.SourceDomain, m.DestinationDomain, m.Nonce, 
		m.Sender, m.Recipient, m.DestinationCaller, len(m.MessageBody))
}

// String returns a string representation of the burn message
func (bm *BurnMessage) String() string {
	return fmt.Sprintf(`BurnMessage{
		Version: %d,
		BurnToken: %x,
		MintRecipient: %x,
		Amount: %s,
		MessageSender: %x
	}`, bm.Version, bm.BurnToken, bm.MintRecipient, bm.Amount.String(), bm.MessageSender)
}

// IsBurnMessage checks if the message body is a valid burn message
func (m *Message) IsBurnMessage() bool {
	return len(m.MessageBody) == BurnMessageBodySize
}

// GetBurnMessage decodes the message body as a burn message
func (m *Message) GetBurnMessage() (*BurnMessage, error) {
	if !m.IsBurnMessage() {
		return nil, sdkerrors.Wrap(ErrInvalidMessageBody, "message body is not a burn message")
	}

	return DecodeBurnMessage(m.MessageBody)
}

// GetNonceString returns the nonce as a string for use in events
func (m *Message) GetNonceString() string {
	return strconv.FormatUint(m.Nonce, 10)
}

// GetSourceDomainString returns the source domain as a string
func (m *Message) GetSourceDomainString() string {
	return strconv.FormatUint(uint64(m.SourceDomain), 10)
}

// GetDestinationDomainString returns the destination domain as a string
func (m *Message) GetDestinationDomainString() string {
	return strconv.FormatUint(uint64(m.DestinationDomain), 10)
}

// IsDestinationCaller checks if the given address is the destination caller
func (m *Message) IsDestinationCaller(caller sdk.AccAddress) bool {
	if len(m.DestinationCaller) == 0 {
		return true // Any caller is allowed
	}
	
	// Convert caller address to bytes32 format for comparison
	callerBytes32 := PadAddressTo32Bytes(caller.Bytes())
	return BytesEqual(m.DestinationCaller, callerBytes32)
}

// HasDestinationCaller returns true if the message has a specific destination caller
func (m *Message) HasDestinationCaller() bool {
	return len(m.DestinationCaller) == AddressSize && !IsZeroBytes(m.DestinationCaller)
}

// ValidateMessageDomains validates that the message domains are correct for this chain
func (m *Message) ValidateMessageDomains() error {
	if m.DestinationDomain != NobleChainDomain {
		return sdkerrors.Wrapf(ErrInvalidDestinationDomain, 
			"message destination domain %d does not match Noble domain %d", 
			m.DestinationDomain, NobleChainDomain)
	}
	return nil
}