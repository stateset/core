package types

import (
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
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// blacklistedHosts contains hosts that are not allowed for webhook URLs.
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

// ValidateWebhookURL validates a webhook URL for security.
func ValidateWebhookURL(webhookURL string) error {
	if webhookURL == "" {
		return nil // Empty URL is allowed (optional)
	}

	parsed, err := url.Parse(webhookURL)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidWebhookURL, "failed to parse URL: %v", err)
	}

	if parsed.Scheme != "https" {
		return errorsmod.Wrap(ErrWebhookURLNotHTTPS, "webhook URL must use HTTPS scheme")
	}

	host := strings.ToLower(parsed.Hostname())
	for _, blacklisted := range blacklistedHosts {
		if strings.HasPrefix(host, blacklisted) || host == blacklisted {
			return errorsmod.Wrapf(ErrWebhookURLBlacklisted, "host %s is not allowed", host)
		}
	}

	if host == "" {
		return errorsmod.Wrap(ErrInvalidWebhookURL, "webhook URL must have a valid host")
	}

	return nil
}

func NewMsgInstantTransfer(sender, recipient string, amount sdk.Coin, reference, metadata string) *MsgInstantTransfer {
	return &MsgInstantTransfer{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Reference: reference,
		Metadata:  metadata,
	}
}

func (m MsgInstantTransfer) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if m.Sender == m.Recipient {
		return errorsmod.Wrap(ErrInvalidRecipient, "sender and recipient must be different")
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

func (m MsgCreateEscrow) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if m.Sender == m.Recipient {
		return errorsmod.Wrap(ErrInvalidRecipient, "sender and recipient must be different")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Amount.Denom != StablecoinDenom {
		return errorsmod.Wrapf(ErrInvalidDenom, "expected %s, got %s", StablecoinDenom, m.Amount.Denom)
	}
	if m.ExpiresIn < 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "expiration cannot be negative")
	}
	return nil
}

func (m MsgCreateEscrow) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Sender)
}

func NewMsgReleaseEscrow(sender string, settlementId uint64) *MsgReleaseEscrow {
	return &MsgReleaseEscrow{Sender: sender, SettlementId: settlementId}
}

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

func NewMsgRefundEscrow(recipient string, settlementId uint64, reason string) *MsgRefundEscrow {
	return &MsgRefundEscrow{Recipient: recipient, SettlementId: settlementId, Reason: reason}
}

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

func NewMsgCreateBatch(authority, merchant string, senders []string, amounts []sdk.Coin, references []string) *MsgCreateBatch {
	return &MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant,
		Senders:    senders,
		Amounts:    amounts,
		References: references,
	}
}

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

func NewMsgSettleBatch(authority string, batchId uint64) *MsgSettleBatch {
	return &MsgSettleBatch{Authority: authority, BatchId: batchId}
}

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

func NewMsgOpenChannel(sender, recipient string, deposit sdk.Coin, expiresInBlocks int64) *MsgOpenChannel {
	return &MsgOpenChannel{
		Sender:          sender,
		Recipient:       recipient,
		Deposit:         deposit,
		ExpiresInBlocks: expiresInBlocks,
	}
}

func (m MsgOpenChannel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid sender address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid recipient address")
	}
	if m.Sender == m.Recipient {
		return errorsmod.Wrap(ErrInvalidRecipient, "sender and recipient must be different")
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

func NewMsgCloseChannel(closer string, channelId uint64) *MsgCloseChannel {
	return &MsgCloseChannel{Closer: closer, ChannelId: channelId}
}

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

func NewMsgClaimChannel(recipient string, channelId uint64, amount sdk.Coin, nonce uint64, signature string) *MsgClaimChannel {
	return &MsgClaimChannel{
		Recipient: recipient,
		ChannelId: channelId,
		Amount:    amount,
		Nonce:     nonce,
		Signature: signature,
	}
}

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
	if len(m.Signature) < 64 {
		return errorsmod.Wrap(ErrInvalidSignature, "signature too short")
	}
	return nil
}

func (m MsgClaimChannel) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Recipient)
}

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
	if err := ValidateWebhookURL(m.WebhookUrl); err != nil {
		return err
	}
	return nil
}

func (m MsgRegisterMerchant) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Authority)
}

func NewMsgUpdateMerchant(authority, merchant string) *MsgUpdateMerchant {
	return &MsgUpdateMerchant{Authority: authority, Merchant: merchant}
}

func (m MsgUpdateMerchant) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid authority address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Merchant); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid merchant address")
	}
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

func NewMsgInstantCheckout(customer, merchant string, amount sdk.Coin, orderRef string, useEscrow bool, items []CheckoutItem, metadata string) *MsgInstantCheckout {
	return &MsgInstantCheckout{
		Customer:       customer,
		Merchant:       merchant,
		Amount:         amount,
		OrderReference: orderRef,
		UseEscrow:      useEscrow,
		Items:          items,
		Metadata:       metadata,
	}
}

func (m MsgInstantCheckout) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Customer); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid customer address")
	}
	if _, err := sdk.AccAddressFromBech32(m.Merchant); err != nil {
		return errorsmod.Wrap(ErrInvalidRecipient, "invalid merchant address")
	}
	if m.Customer == m.Merchant {
		return errorsmod.Wrap(ErrInvalidSettlement, "customer and merchant cannot be the same")
	}
	if !m.Amount.IsValid() || m.Amount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	if m.Amount.Denom != StablecoinDenom {
		return errorsmod.Wrapf(ErrInvalidDenom, "expected %s, got %s", StablecoinDenom, m.Amount.Denom)
	}
	if m.OrderReference == "" {
		return errorsmod.Wrap(ErrInvalidSettlement, "order reference is required")
	}
	return nil
}

func (m MsgInstantCheckout) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Customer)
}

func NewMsgPartialRefund(authority string, settlementId uint64, refundAmount sdk.Coin, reason string) *MsgPartialRefund {
	return &MsgPartialRefund{
		Authority:    authority,
		SettlementId: settlementId,
		RefundAmount: refundAmount,
		Reason:       reason,
	}
}

func (m MsgPartialRefund) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidSettlement, "invalid authority address")
	}
	if m.SettlementId == 0 {
		return errorsmod.Wrap(ErrInvalidSettlement, "settlement id required")
	}
	if !m.RefundAmount.IsValid() || m.RefundAmount.IsZero() {
		return errorsmod.Wrap(ErrInvalidAmount, "refund amount must be positive")
	}
	if m.Reason == "" {
		return errorsmod.Wrap(ErrInvalidSettlement, "refund reason is required")
	}
	return nil
}

func (m MsgPartialRefund) GetSigners() []sdk.AccAddress {
	return mustGetSigner(m.Authority)
}
