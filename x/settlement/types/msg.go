package types

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// mustGetSigner safely returns the signer address.
// ValidateBasic is called before GetSigners in the SDK, so this should not panic.
// However, we still handle errors gracefully by returning an empty slice.
func mustGetSigner(bech32 string) []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(bech32)
	if err != nil {
		// Return empty - ValidateBasic should have caught this
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// blacklistedHosts contains hosts that are not allowed for webhook URLs
var blacklistedHosts = []string{
	"localhost",
	"127.0.0.1",
	"0.0.0.0",
	"::1",
	"169.254.", // Link-local
	"10.",      // Private network
	"172.16.",  // Private network
	"172.17.",
	"172.18.",
	"172.19.",
	"172.20.",
	"172.21.",
	"172.22.",
	"172.23.",
	"172.24.",
	"172.25.",
	"172.26.",
	"172.27.",
	"172.28.",
	"172.29.",
	"172.30.",
	"172.31.",
	"192.168.", // Private network
}

// ValidateWebhookURL validates a webhook URL for security
func ValidateWebhookURL(webhookURL string) error {
	if webhookURL == "" {
		return nil // Empty URL is allowed (optional)
	}

	parsed, err := url.Parse(webhookURL)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidWebhookURL, "failed to parse URL: %v", err)
	}

	// Must be HTTPS
	if parsed.Scheme != "https" {
		return errorsmod.Wrap(ErrWebhookURLNotHTTPS, "webhook URL must use HTTPS scheme")
	}

	// Check against blacklisted hosts
	host := strings.ToLower(parsed.Hostname())
	for _, blacklisted := range blacklistedHosts {
		if strings.HasPrefix(host, blacklisted) || host == blacklisted {
			return errorsmod.Wrapf(ErrWebhookURLBlacklisted, "host %s is not allowed", host)
		}
	}

	// Must have a valid host
	if host == "" {
		return errorsmod.Wrap(ErrInvalidWebhookURL, "webhook URL must have a valid host")
	}

	return nil
}

const (
	TypeMsgInstantTransfer   = "instant_transfer"
	TypeMsgCreateEscrow      = "create_escrow"
	TypeMsgReleaseEscrow     = "release_escrow"
	TypeMsgRefundEscrow      = "refund_escrow"
	TypeMsgCreateBatch       = "create_batch"
	TypeMsgSettleBatch       = "settle_batch"
	TypeMsgOpenChannel       = "open_channel"
	TypeMsgCloseChannel      = "close_channel"
	TypeMsgClaimChannel      = "claim_channel"
	TypeMsgRegisterMerchant  = "register_merchant"
	TypeMsgUpdateMerchant    = "update_merchant"
)

var (
	_ sdk.Msg = (*MsgInstantTransfer)(nil)
	_ sdk.Msg = (*MsgCreateEscrow)(nil)
	_ sdk.Msg = (*MsgReleaseEscrow)(nil)
	_ sdk.Msg = (*MsgRefundEscrow)(nil)
	_ sdk.Msg = (*MsgCreateBatch)(nil)
	_ sdk.Msg = (*MsgSettleBatch)(nil)
	_ sdk.Msg = (*MsgOpenChannel)(nil)
	_ sdk.Msg = (*MsgCloseChannel)(nil)
	_ sdk.Msg = (*MsgClaimChannel)(nil)
	_ sdk.Msg = (*MsgRegisterMerchant)(nil)
	_ sdk.Msg = (*MsgUpdateMerchant)(nil)
)

// MsgInstantTransfer - instant stablecoin transfer with immediate settlement
type MsgInstantTransfer struct {
	Sender    string   `json:"sender" yaml:"sender"`
	Recipient string   `json:"recipient" yaml:"recipient"`
	Amount    sdk.Coin `json:"amount" yaml:"amount"`
	Reference string   `json:"reference" yaml:"reference"`
	Metadata  string   `json:"metadata" yaml:"metadata"`
}

func (m *MsgInstantTransfer) Reset() { *m = MsgInstantTransfer{} }
func (m *MsgInstantTransfer) String() string {
	return fmt.Sprintf("MsgInstantTransfer{%s->%s %s ref:%s}", m.Sender, m.Recipient, m.Amount.String(), m.Reference)
}
func (*MsgInstantTransfer) ProtoMessage() {}

func NewMsgInstantTransfer(sender, recipient string, amount sdk.Coin, reference, metadata string) *MsgInstantTransfer {
	return &MsgInstantTransfer{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Reference: reference,
		Metadata:  metadata,
	}
}

func (m MsgInstantTransfer) Route() string { return RouterKey }
func (m MsgInstantTransfer) Type() string  { return TypeMsgInstantTransfer }

func (m MsgInstantTransfer) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Amount.Denom != StablecoinDenom {
		return errorsmod.Wrapf(ErrInvalidDenom, "expected %s, got %s", StablecoinDenom, m.Amount.Denom)
	}
	return nil
}

func (m MsgInstantTransfer) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Sender)
}

func (m MsgInstantTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgInstantTransferResponse struct {
	SettlementId uint64 `json:"settlement_id"`
	TxHash       string `json:"tx_hash"`
}

// MsgCreateEscrow - create an escrow settlement with delayed release
type MsgCreateEscrow struct {
	Sender    string        `json:"sender" yaml:"sender"`
	Recipient string        `json:"recipient" yaml:"recipient"`
	Amount    sdk.Coin      `json:"amount" yaml:"amount"`
	Reference string        `json:"reference" yaml:"reference"`
	Metadata  string        `json:"metadata" yaml:"metadata"`
	ExpiresIn time.Duration `json:"expires_in" yaml:"expires_in"`
}

func (m *MsgCreateEscrow) Reset() { *m = MsgCreateEscrow{} }
func (m *MsgCreateEscrow) String() string {
	return fmt.Sprintf("MsgCreateEscrow{%s->%s %s expires:%v}", m.Sender, m.Recipient, m.Amount.String(), m.ExpiresIn)
}
func (*MsgCreateEscrow) ProtoMessage() {}

func NewMsgCreateEscrow(sender, recipient string, amount sdk.Coin, reference, metadata string, expiresIn time.Duration) *MsgCreateEscrow {
	return &MsgCreateEscrow{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Reference: reference,
		Metadata:  metadata,
		ExpiresIn: expiresIn,
	}
}

func (m MsgCreateEscrow) Route() string { return RouterKey }
func (m MsgCreateEscrow) Type() string  { return TypeMsgCreateEscrow }

func (m MsgCreateEscrow) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Amount.Denom != StablecoinDenom {
		return errorsmod.Wrapf(ErrInvalidDenom, "expected %s, got %s", StablecoinDenom, m.Amount.Denom)
	}
	if m.ExpiresIn <= 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "expiration must be positive")
	}
	return nil
}

func (m MsgCreateEscrow) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Sender)
}

func (m MsgCreateEscrow) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCreateEscrowResponse struct {
	SettlementId uint64    `json:"settlement_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// MsgReleaseEscrow - release escrowed funds to the recipient
type MsgReleaseEscrow struct {
	Sender       string `json:"sender" yaml:"sender"`
	SettlementId uint64 `json:"settlement_id" yaml:"settlement_id"`
}

func (m *MsgReleaseEscrow) Reset() { *m = MsgReleaseEscrow{} }
func (m *MsgReleaseEscrow) String() string {
	return fmt.Sprintf("MsgReleaseEscrow{%s %d}", m.Sender, m.SettlementId)
}
func (*MsgReleaseEscrow) ProtoMessage() {}

func NewMsgReleaseEscrow(sender string, settlementId uint64) *MsgReleaseEscrow {
	return &MsgReleaseEscrow{Sender: sender, SettlementId: settlementId}
}

func (m MsgReleaseEscrow) Route() string { return RouterKey }
func (m MsgReleaseEscrow) Type() string  { return TypeMsgReleaseEscrow }

func (m MsgReleaseEscrow) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if m.SettlementId == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "settlement id required")
	}
	return nil
}

func (m MsgReleaseEscrow) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Sender)
}

func (m MsgReleaseEscrow) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgReleaseEscrowResponse struct{}

// MsgRefundEscrow - refund escrowed funds back to sender (before expiration or on dispute)
type MsgRefundEscrow struct {
	Recipient    string `json:"recipient" yaml:"recipient"`
	SettlementId uint64 `json:"settlement_id" yaml:"settlement_id"`
	Reason       string `json:"reason" yaml:"reason"`
}

func (m *MsgRefundEscrow) Reset() { *m = MsgRefundEscrow{} }
func (m *MsgRefundEscrow) String() string {
	return fmt.Sprintf("MsgRefundEscrow{%s %d}", m.Recipient, m.SettlementId)
}
func (*MsgRefundEscrow) ProtoMessage() {}

func NewMsgRefundEscrow(recipient string, settlementId uint64, reason string) *MsgRefundEscrow {
	return &MsgRefundEscrow{Recipient: recipient, SettlementId: settlementId, Reason: reason}
}

func (m MsgRefundEscrow) Route() string { return RouterKey }
func (m MsgRefundEscrow) Type() string  { return TypeMsgRefundEscrow }

func (m MsgRefundEscrow) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid recipient address")
	}
	if m.SettlementId == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "settlement id required")
	}
	return nil
}

func (m MsgRefundEscrow) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Recipient)
}

func (m MsgRefundEscrow) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgRefundEscrowResponse struct{}

// MsgCreateBatch - create a batch settlement for multiple payments to a merchant
type MsgCreateBatch struct {
	Authority string   `json:"authority" yaml:"authority"`
	Merchant  string   `json:"merchant" yaml:"merchant"`
	Senders   []string `json:"senders" yaml:"senders"`
	Amounts   []sdk.Coin `json:"amounts" yaml:"amounts"`
	References []string `json:"references" yaml:"references"`
}

func (m *MsgCreateBatch) Reset() { *m = MsgCreateBatch{} }
func (m *MsgCreateBatch) String() string {
	return fmt.Sprintf("MsgCreateBatch{%s count:%d}", m.Merchant, len(m.Senders))
}
func (*MsgCreateBatch) ProtoMessage() {}

func NewMsgCreateBatch(authority, merchant string, senders []string, amounts []sdk.Coin, references []string) *MsgCreateBatch {
	return &MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant,
		Senders:    senders,
		Amounts:    amounts,
		References: references,
	}
}

func (m MsgCreateBatch) Route() string { return RouterKey }
func (m MsgCreateBatch) Type() string  { return TypeMsgCreateBatch }

func (m MsgCreateBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid authority address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Merchant); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid merchant address")
	}
	if len(m.Senders) == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "at least one sender required")
	}
	if len(m.Senders) != len(m.Amounts) || len(m.Senders) != len(m.References) {
		return errorsmod.Wrap(ErrInvalidSettlement, "senders, amounts, and references must have same length")
	}
	for i, sender := range m.Senders {
		if _, err := sdk.AccAddressFromBech32(sender); err != nil {
			return errorsmod.Wrapf(ErrInvalidSettlement, "invalid sender address at index %d", i)
		}
		if !m.Amounts[i].IsValid() || m.Amounts[i].IsZero() {
			return errorsmod.Wrapf(ErrInvalidAmount, "invalid amount at index %d", i)
		}
	}
	return nil
}

func (m MsgCreateBatch) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Authority)
}

func (m MsgCreateBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCreateBatchResponse struct {
	BatchId       uint64   `json:"batch_id"`
	SettlementIds []uint64 `json:"settlement_ids"`
}

// MsgSettleBatch - settle a batch of payments
type MsgSettleBatch struct {
	Authority string `json:"authority" yaml:"authority"`
	BatchId   uint64 `json:"batch_id" yaml:"batch_id"`
}

func (m *MsgSettleBatch) Reset() { *m = MsgSettleBatch{} }
func (m *MsgSettleBatch) String() string {
	return fmt.Sprintf("MsgSettleBatch{%d}", m.BatchId)
}
func (*MsgSettleBatch) ProtoMessage() {}

func NewMsgSettleBatch(authority string, batchId uint64) *MsgSettleBatch {
	return &MsgSettleBatch{Authority: authority, BatchId: batchId}
}

func (m MsgSettleBatch) Route() string { return RouterKey }
func (m MsgSettleBatch) Type() string  { return TypeMsgSettleBatch }

func (m MsgSettleBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid authority address")
	}
	if m.BatchId == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "batch id required")
	}
	return nil
}

func (m MsgSettleBatch) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Authority)
}

func (m MsgSettleBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgSettleBatchResponse struct {
	TotalAmount sdk.Coin `json:"total_amount"`
	TotalFees   sdk.Coin `json:"total_fees"`
	NetAmount   sdk.Coin `json:"net_amount"`
}

// MsgOpenChannel - open a payment channel
type MsgOpenChannel struct {
	Sender          string   `json:"sender" yaml:"sender"`
	Recipient       string   `json:"recipient" yaml:"recipient"`
	Deposit         sdk.Coin `json:"deposit" yaml:"deposit"`
	ExpiresInBlocks int64    `json:"expires_in_blocks" yaml:"expires_in_blocks"`
}

func (m *MsgOpenChannel) Reset() { *m = MsgOpenChannel{} }
func (m *MsgOpenChannel) String() string {
	return fmt.Sprintf("MsgOpenChannel{%s->%s %s}", m.Sender, m.Recipient, m.Deposit.String())
}
func (*MsgOpenChannel) ProtoMessage() {}

func NewMsgOpenChannel(sender, recipient string, deposit sdk.Coin, expiresInBlocks int64) *MsgOpenChannel {
	return &MsgOpenChannel{
		Sender:          sender,
		Recipient:       recipient,
		Deposit:         deposit,
		ExpiresInBlocks: expiresInBlocks,
	}
}

func (m MsgOpenChannel) Route() string { return RouterKey }
func (m MsgOpenChannel) Type() string  { return TypeMsgOpenChannel }

func (m MsgOpenChannel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if !m.Deposit.IsValid() || m.Deposit.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "deposit must be positive")
	}
	if m.Deposit.Denom != StablecoinDenom {
		return errorsmod.Wrapf(ErrInvalidDenom, "expected %s, got %s", StablecoinDenom, m.Deposit.Denom)
	}
	if m.ExpiresInBlocks <= 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "expiration blocks must be positive")
	}
	return nil
}

func (m MsgOpenChannel) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Sender)
}

func (m MsgOpenChannel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgOpenChannelResponse struct {
	ChannelId       uint64 `json:"channel_id"`
	ExpiresAtHeight int64  `json:"expires_at_height"`
}

// MsgCloseChannel - close a payment channel (sender can close after expiration)
type MsgCloseChannel struct {
	Closer    string `json:"closer" yaml:"closer"`
	ChannelId uint64 `json:"channel_id" yaml:"channel_id"`
}

func (m *MsgCloseChannel) Reset() { *m = MsgCloseChannel{} }
func (m *MsgCloseChannel) String() string {
	return fmt.Sprintf("MsgCloseChannel{%s %d}", m.Closer, m.ChannelId)
}
func (*MsgCloseChannel) ProtoMessage() {}

func NewMsgCloseChannel(closer string, channelId uint64) *MsgCloseChannel {
	return &MsgCloseChannel{Closer: closer, ChannelId: channelId}
}

func (m MsgCloseChannel) Route() string { return RouterKey }
func (m MsgCloseChannel) Type() string  { return TypeMsgCloseChannel }

func (m MsgCloseChannel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Closer); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid closer address")
	}
	if m.ChannelId == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "channel id required")
	}
	return nil
}

func (m MsgCloseChannel) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Closer)
}

func (m MsgCloseChannel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgCloseChannelResponse struct {
	FinalBalance sdk.Coin `json:"final_balance"`
}

// MsgClaimChannel - claim funds from a payment channel (recipient claims)
type MsgClaimChannel struct {
	Recipient string   `json:"recipient" yaml:"recipient"`
	ChannelId uint64   `json:"channel_id" yaml:"channel_id"`
	Amount    sdk.Coin `json:"amount" yaml:"amount"`
	Nonce     uint64   `json:"nonce" yaml:"nonce"`
	Signature string   `json:"signature" yaml:"signature"`
}

func (m *MsgClaimChannel) Reset() { *m = MsgClaimChannel{} }
func (m *MsgClaimChannel) String() string {
	return fmt.Sprintf("MsgClaimChannel{%s %d %s}", m.Recipient, m.ChannelId, m.Amount.String())
}
func (*MsgClaimChannel) ProtoMessage() {}

func NewMsgClaimChannel(recipient string, channelId uint64, amount sdk.Coin, nonce uint64, signature string) *MsgClaimChannel {
	return &MsgClaimChannel{
		Recipient: recipient,
		ChannelId: channelId,
		Amount:    amount,
		Nonce:     nonce,
		Signature: signature,
	}
}

func (m MsgClaimChannel) Route() string { return RouterKey }
func (m MsgClaimChannel) Type() string  { return TypeMsgClaimChannel }

func (m MsgClaimChannel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid recipient address")
	}
	if m.ChannelId == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "channel id required")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Nonce == 0 {
		return errorsmod.Wrap(ErrInvalidNonce, "nonce must be greater than zero")
	}
	if m.Signature == "" {
		return errorsmod.Wrap(ErrInvalidSignature, "signature is required")
	}
	// Validate signature format (hex-encoded)
	if len(m.Signature) < 64 {
		return errorsmod.Wrap(ErrInvalidSignature, "signature too short")
	}
	return nil
}

func (m MsgClaimChannel) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Recipient)
}

func (m MsgClaimChannel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgClaimChannelResponse struct {
	AmountClaimed sdk.Coin `json:"amount_claimed"`
	RemainingBalance sdk.Coin `json:"remaining_balance"`
}

// MsgRegisterMerchant - register a merchant for settlements
type MsgRegisterMerchant struct {
	Authority       string   `json:"authority" yaml:"authority"`
	Merchant        string   `json:"merchant" yaml:"merchant"`
	Name            string   `json:"name" yaml:"name"`
	FeeRateBps      uint32   `json:"fee_rate_bps" yaml:"fee_rate_bps"`
	MinSettlement   sdk.Coin `json:"min_settlement" yaml:"min_settlement"`
	MaxSettlement   sdk.Coin `json:"max_settlement" yaml:"max_settlement"`
	BatchEnabled    bool     `json:"batch_enabled" yaml:"batch_enabled"`
	BatchThreshold  sdk.Coin `json:"batch_threshold" yaml:"batch_threshold"`
	WebhookUrl      string   `json:"webhook_url" yaml:"webhook_url"`
}

func (m *MsgRegisterMerchant) Reset() { *m = MsgRegisterMerchant{} }
func (m *MsgRegisterMerchant) String() string {
	return fmt.Sprintf("MsgRegisterMerchant{%s %s}", m.Merchant, m.Name)
}
func (*MsgRegisterMerchant) ProtoMessage() {}

func NewMsgRegisterMerchant(authority, merchant, name string, feeRateBps uint32, minSettlement, maxSettlement sdk.Coin, batchEnabled bool, batchThreshold sdk.Coin, webhookUrl string) *MsgRegisterMerchant {
	return &MsgRegisterMerchant{
		Authority:      authority,
		Merchant:       merchant,
		Name:           name,
		FeeRateBps:     feeRateBps,
		MinSettlement:  minSettlement,
		MaxSettlement:  maxSettlement,
		BatchEnabled:   batchEnabled,
		BatchThreshold: batchThreshold,
		WebhookUrl:     webhookUrl,
	}
}

func (m MsgRegisterMerchant) Route() string { return RouterKey }
func (m MsgRegisterMerchant) Type() string  { return TypeMsgRegisterMerchant }

func (m MsgRegisterMerchant) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid authority address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Merchant); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid merchant address")
	}
	if len(m.Name) == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "merchant name required")
	}
	if m.FeeRateBps > 10000 {
		return errorsmod.Wrap(ErrInvalidSettlement, "fee rate must be <= 10000 bps (100%)")
	}
	// Validate webhook URL if provided
	if err := ValidateWebhookURL(m.WebhookUrl); err != nil {
		return err
	}
	return nil
}

func (m MsgRegisterMerchant) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Authority)
}

func (m MsgRegisterMerchant) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgRegisterMerchantResponse struct{}

// MsgUpdateMerchant - update merchant configuration
type MsgUpdateMerchant struct {
	Authority       string   `json:"authority" yaml:"authority"`
	Merchant        string   `json:"merchant" yaml:"merchant"`
	Name            string   `json:"name,omitempty" yaml:"name,omitempty"`
	FeeRateBps      uint32   `json:"fee_rate_bps,omitempty" yaml:"fee_rate_bps,omitempty"`
	MinSettlement   sdk.Coin `json:"min_settlement,omitempty" yaml:"min_settlement,omitempty"`
	MaxSettlement   sdk.Coin `json:"max_settlement,omitempty" yaml:"max_settlement,omitempty"`
	BatchEnabled    *bool    `json:"batch_enabled,omitempty" yaml:"batch_enabled,omitempty"`
	BatchThreshold  sdk.Coin `json:"batch_threshold,omitempty" yaml:"batch_threshold,omitempty"`
	IsActive        *bool    `json:"is_active,omitempty" yaml:"is_active,omitempty"`
	WebhookUrl      string   `json:"webhook_url,omitempty" yaml:"webhook_url,omitempty"`
}

func (m *MsgUpdateMerchant) Reset() { *m = MsgUpdateMerchant{} }
func (m *MsgUpdateMerchant) String() string {
	return fmt.Sprintf("MsgUpdateMerchant{%s}", m.Merchant)
}
func (*MsgUpdateMerchant) ProtoMessage() {}

func NewMsgUpdateMerchant(authority, merchant string) *MsgUpdateMerchant {
	return &MsgUpdateMerchant{Authority: authority, Merchant: merchant}
}

func (m MsgUpdateMerchant) Route() string { return RouterKey }
func (m MsgUpdateMerchant) Type() string  { return TypeMsgUpdateMerchant }

func (m MsgUpdateMerchant) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid authority address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Merchant); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid merchant address")
	}
	// Validate webhook URL if provided (non-empty means update)
	if m.WebhookUrl != "" {
		if err := ValidateWebhookURL(m.WebhookUrl); err != nil {
			return err
		}
	}
	if m.FeeRateBps > 10000 {
		return errorsmod.Wrap(ErrInvalidSettlement, "fee rate must be <= 10000 bps (100%)")
	}
	return nil
}

func (m MsgUpdateMerchant) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Authority)
}

func (m MsgUpdateMerchant) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

type MsgUpdateMerchantResponse struct{}

// MsgServer interface
type MsgServer interface {
	InstantTransfer(ctx context.Context, msg *MsgInstantTransfer) (*MsgInstantTransferResponse, error)
	CreateEscrow(ctx context.Context, msg *MsgCreateEscrow) (*MsgCreateEscrowResponse, error)
	ReleaseEscrow(ctx context.Context, msg *MsgReleaseEscrow) (*MsgReleaseEscrowResponse, error)
	RefundEscrow(ctx context.Context, msg *MsgRefundEscrow) (*MsgRefundEscrowResponse, error)
	CreateBatch(ctx context.Context, msg *MsgCreateBatch) (*MsgCreateBatchResponse, error)
	SettleBatch(ctx context.Context, msg *MsgSettleBatch) (*MsgSettleBatchResponse, error)
	OpenChannel(ctx context.Context, msg *MsgOpenChannel) (*MsgOpenChannelResponse, error)
	CloseChannel(ctx context.Context, msg *MsgCloseChannel) (*MsgCloseChannelResponse, error)
	ClaimChannel(ctx context.Context, msg *MsgClaimChannel) (*MsgClaimChannelResponse, error)
	RegisterMerchant(ctx context.Context, msg *MsgRegisterMerchant) (*MsgRegisterMerchantResponse, error)
	UpdateMerchant(ctx context.Context, msg *MsgUpdateMerchant) (*MsgUpdateMerchantResponse, error)
}

// _Msg_serviceDesc is the gRPC service descriptor
var _Msg_serviceDesc = struct {
	ServiceName string
	HandlerType interface{}
	Methods     []struct {
		MethodName string
		Handler    interface{}
	}
	Streams  []struct{}
	Metadata string
}{
	ServiceName: "stateset.settlement.Msg",
}
