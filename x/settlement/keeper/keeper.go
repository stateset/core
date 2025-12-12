package keeper

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/settlement/types"
)

// Keeper handles settlement operations
type Keeper struct {
	storeKey      storetypes.StoreKey
	cdc           codec.BinaryCodec
	bankKeeper    types.BankKeeper
	compKeeper    types.ComplianceKeeper
	accountKeeper types.AccountKeeper
	authority     string
	params        types.Params
}

// NewKeeper creates a new settlement keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	compKeeper types.ComplianceKeeper,
	accountKeeper types.AccountKeeper,
	authority string,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		bankKeeper:    bankKeeper,
		compKeeper:    compKeeper,
		accountKeeper: accountKeeper,
		authority:     authority,
		params:        types.DefaultParams(),
	}
}

// GetParams returns the current module parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if len(bz) == 0 {
		return types.DefaultParams()
	}
	var params types.Params
	types.ModuleCdc.MustUnmarshalJSON(bz, &params)
	return params
}

// SetParams sets the module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz := types.ModuleCdc.MustMarshalJSON(&params)
	store.Set(types.ParamsKey, bz)
	return nil
}

// GetAuthority returns the module authority address
func (k Keeper) GetAuthority() string {
	return k.authority
}

// ============================================================================
// Settlement ID Management
// ============================================================================

func (k Keeper) getNextSettlementID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextSettlementIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextSettlementID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextSettlementIDKey, bz)
}

func (k Keeper) getNextBatchID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextBatchIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextBatchID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextBatchIDKey, bz)
}

func (k Keeper) getNextChannelID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextChannelIDKey)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextChannelID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextChannelIDKey, bz)
}

// ============================================================================
// Settlement Operations
// ============================================================================

func (k Keeper) storeSettlement(ctx sdk.Context, settlement types.Settlement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SettlementKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&settlement)
	store.Set(mustWriteUint64(settlement.Id), bz)
}

// GetSettlement retrieves a settlement by ID
func (k Keeper) GetSettlement(ctx sdk.Context, id uint64) (types.Settlement, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SettlementKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.Settlement{}, false
	}
	var settlement types.Settlement
	types.ModuleCdc.MustUnmarshalJSON(bz, &settlement)
	return settlement, true
}

// InstantTransfer performs an instant stablecoin transfer with settlement
func (k Keeper) InstantTransfer(ctx sdk.Context, sender, recipient string, amount sdk.Coin, reference, metadata string) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return 0, types.ErrInvalidSettlement
	}
	recipientAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return 0, types.ErrInvalidRecipient
	}
	if senderAddr.Equals(recipientAddr) {
		return 0, types.ErrInvalidRecipient.Wrap("sender and recipient must be different")
	}

	// Check compliance for both parties
	if err := k.compKeeper.AssertCompliant(wrappedCtx, senderAddr); err != nil {
		return 0, types.ErrComplianceCheckFailed
	}
	if err := k.compKeeper.AssertCompliant(wrappedCtx, recipientAddr); err != nil {
		return 0, types.ErrComplianceCheckFailed
	}

	// Validate amount against params
	params := k.GetParams(ctx)
	if amount.IsLT(params.MinSettlementAmount) {
		return 0, types.ErrSettlementTooSmall
	}
	if amount.IsGTE(params.MaxSettlementAmount) {
		return 0, types.ErrSettlementTooLarge
	}

	// Check sender balance
	balance := k.bankKeeper.GetBalance(wrappedCtx, senderAddr, amount.Denom)
	if balance.IsLT(amount) {
		return 0, types.ErrInsufficientFunds
	}

	// Calculate fee using params default or merchant-specific rate
	fee := k.calculateFee(ctx, amount, recipient)
	netAmount := sdk.NewCoin(amount.Denom, amount.Amount.Sub(fee.Amount))

	// Transfer: sender -> module (escrow)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, senderAddr, types.ModuleAccountName, sdk.NewCoins(amount)); err != nil {
		return 0, err
	}

	// Transfer: module -> recipient (net amount)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, recipientAddr, sdk.NewCoins(netAmount)); err != nil {
		return 0, err
	}

	// Transfer fee to fee collector (if configured) or burn it
	if fee.IsPositive() {
		if err := k.collectFee(ctx, fee); err != nil {
			return 0, err
		}
	}

	// Create settlement record
	nextID := k.getNextSettlementID(ctx)
	settlement := types.Settlement{
		Id:            nextID,
		Type:          types.SettlementTypeInstant,
		Sender:        sender,
		Recipient:     recipient,
		Amount:        amount,
		Fee:           fee,
		NetAmount:     netAmount,
		Status:        types.SettlementStatusCompleted,
		Reference:     reference,
		Metadata:      metadata,
		CreatedHeight: ctx.BlockHeight(),
		CreatedTime:   ctx.BlockTime(),
		SettledHeight: ctx.BlockHeight(),
		SettledTime:   ctx.BlockTime(),
	}

	k.storeSettlement(ctx, settlement)
	k.setNextSettlementID(ctx, nextID+1)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInstantTransfer,
			sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", settlement.Id)),
			sdk.NewAttribute(types.AttributeKeySender, sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyFee, fee.String()),
			sdk.NewAttribute(types.AttributeKeyReference, reference),
		),
	)

	return settlement.Id, nil
}

// CreateEscrow creates an escrow settlement that holds funds until released
// expirationSeconds specifies the escrow duration in seconds (0 uses default from params)
func (k Keeper) CreateEscrow(ctx sdk.Context, sender, recipient string, amount sdk.Coin, reference, metadata string, expirationSeconds int64) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return 0, types.ErrInvalidSettlement
	}
	recipientAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return 0, types.ErrInvalidRecipient
	}
	if senderAddr.Equals(recipientAddr) {
		return 0, types.ErrInvalidRecipient.Wrap("sender and recipient must be different")
	}

	// Check compliance
	if err := k.compKeeper.AssertCompliant(wrappedCtx, senderAddr); err != nil {
		return 0, types.ErrComplianceCheckFailed
	}
	if err := k.compKeeper.AssertCompliant(wrappedCtx, recipientAddr); err != nil {
		return 0, types.ErrComplianceCheckFailed
	}

	// Validate amount against params
	params := k.GetParams(ctx)
	if amount.IsLT(params.MinSettlementAmount) {
		return 0, types.ErrSettlementTooSmall
	}
	if amount.IsGTE(params.MaxSettlementAmount) {
		return 0, types.ErrSettlementTooLarge
	}

	// Validate escrow expiration
	if expirationSeconds <= 0 {
		expirationSeconds = params.DefaultEscrowExpiration
	}
	if expirationSeconds > params.MaxEscrowExpiration {
		return 0, types.ErrInvalidEscrowExpiration
	}

	// Check balance
	balance := k.bankKeeper.GetBalance(wrappedCtx, senderAddr, amount.Denom)
	if balance.IsLT(amount) {
		return 0, types.ErrInsufficientFunds
	}

	// Transfer to escrow
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, senderAddr, types.ModuleAccountName, sdk.NewCoins(amount)); err != nil {
		return 0, err
	}

	// Calculate fee
	fee := k.calculateFee(ctx, amount, recipient)
	netAmount := sdk.NewCoin(amount.Denom, amount.Amount.Sub(fee.Amount))

	// Calculate expiration time
	expiresAt := ctx.BlockTime().Add(time.Duration(expirationSeconds) * time.Second)

	nextID := k.getNextSettlementID(ctx)
	settlement := types.Settlement{
		Id:            nextID,
		Type:          types.SettlementTypeEscrow,
		Sender:        sender,
		Recipient:     recipient,
		Amount:        amount,
		Fee:           fee,
		NetAmount:     netAmount,
		Status:        types.SettlementStatusPending,
		Reference:     reference,
		Metadata:      metadata,
		CreatedHeight: ctx.BlockHeight(),
		CreatedTime:   ctx.BlockTime(),
		ExpiresAt:     expiresAt,
	}

	k.storeSettlement(ctx, settlement)
	k.setNextSettlementID(ctx, nextID+1)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSettlementCreated,
			sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", settlement.Id)),
			sdk.NewAttribute(types.AttributeKeySender, sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyStatus, string(types.SettlementStatusPending)),
		),
	)

	return settlement.Id, nil
}

// ReleaseEscrow releases escrowed funds to the recipient
func (k Keeper) ReleaseEscrow(ctx sdk.Context, settlementId uint64, sender sdk.AccAddress) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	settlement, found := k.GetSettlement(ctx, settlementId)
	if !found {
		return types.ErrSettlementNotFound
	}

	if settlement.Status == types.SettlementStatusCompleted {
		return types.ErrSettlementCompleted
	}
	if settlement.Status == types.SettlementStatusCancelled || settlement.Status == types.SettlementStatusRefunded {
		return types.ErrSettlementCancelled
	}

	// Only the original sender can release
	expectedSender, err := sdk.AccAddressFromBech32(settlement.Sender)
	if err != nil {
		return types.ErrInvalidSettlement
	}
	if !expectedSender.Equals(sender) {
		return types.ErrUnauthorized
	}

	recipientAddr, err := sdk.AccAddressFromBech32(settlement.Recipient)
	if err != nil {
		return types.ErrInvalidRecipient
	}

	// Compliance check
	if err := k.compKeeper.AssertCompliant(wrappedCtx, recipientAddr); err != nil {
		return types.ErrComplianceCheckFailed
	}

	// Transfer net amount to recipient
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, recipientAddr, sdk.NewCoins(settlement.NetAmount)); err != nil {
		return err
	}

	// Collect the fee
	if settlement.Fee.IsPositive() {
		if err := k.collectFee(ctx, settlement.Fee); err != nil {
			return err
		}
	}

	// Update settlement
	settlement.Status = types.SettlementStatusCompleted
	settlement.SettledHeight = ctx.BlockHeight()
	settlement.SettledTime = ctx.BlockTime()
	k.storeSettlement(ctx, settlement)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSettlementCompleted,
			sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", settlement.Id)),
			sdk.NewAttribute(types.AttributeKeyRecipient, settlement.Recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, settlement.NetAmount.String()),
			sdk.NewAttribute(types.AttributeKeyFee, settlement.Fee.String()),
		),
	)

	return nil
}

// RefundEscrow refunds escrowed funds back to the sender
func (k Keeper) RefundEscrow(ctx sdk.Context, settlementId uint64, recipient sdk.AccAddress, reason string) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	settlement, found := k.GetSettlement(ctx, settlementId)
	if !found {
		return types.ErrSettlementNotFound
	}

	if settlement.Status == types.SettlementStatusCompleted {
		return types.ErrSettlementCompleted
	}
	if settlement.Status == types.SettlementStatusRefunded {
		return types.ErrSettlementCancelled
	}

	// Only the recipient can initiate refund
	expectedRecipient, err := sdk.AccAddressFromBech32(settlement.Recipient)
	if err != nil {
		return types.ErrInvalidRecipient
	}
	if !expectedRecipient.Equals(recipient) {
		return types.ErrUnauthorized
	}

	senderAddr, err := sdk.AccAddressFromBech32(settlement.Sender)
	if err != nil {
		return types.ErrInvalidSettlement
	}

	// Compliance check
	if err := k.compKeeper.AssertCompliant(wrappedCtx, senderAddr); err != nil {
		return types.ErrComplianceCheckFailed
	}

	// Refund full amount (including fee) back to sender
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, senderAddr, sdk.NewCoins(settlement.Amount)); err != nil {
		return err
	}

	// Update settlement
	settlement.Status = types.SettlementStatusRefunded
	settlement.Metadata = reason
	k.storeSettlement(ctx, settlement)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSettlementRefunded,
			sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", settlement.Id)),
			sdk.NewAttribute(types.AttributeKeySender, settlement.Sender),
			sdk.NewAttribute(types.AttributeKeyAmount, settlement.Amount.String()),
		),
	)

	return nil
}

// ============================================================================
// Batch Operations
// ============================================================================

func (k Keeper) storeBatch(ctx sdk.Context, batch types.BatchSettlement) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BatchKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&batch)
	store.Set(mustWriteUint64(batch.Id), bz)
}

// GetBatch retrieves a batch by ID
func (k Keeper) GetBatch(ctx sdk.Context, id uint64) (types.BatchSettlement, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BatchKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.BatchSettlement{}, false
	}
	var batch types.BatchSettlement
	types.ModuleCdc.MustUnmarshalJSON(bz, &batch)
	return batch, true
}

// CreateBatch creates a batch settlement for multiple payments
func (k Keeper) CreateBatch(ctx sdk.Context, merchant string, senders []string, amounts []sdk.Coin, references []string) (uint64, []uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	merchantAddr, err := sdk.AccAddressFromBech32(merchant)
	if err != nil {
		return 0, nil, types.ErrInvalidRecipient
	}

	// Check merchant compliance
	if err := k.compKeeper.AssertCompliant(wrappedCtx, merchantAddr); err != nil {
		return 0, nil, types.ErrComplianceCheckFailed
	}

	settlementIds := make([]uint64, len(senders))
	totalAmount := sdk.NewCoin(types.StablecoinDenom, sdkmath.ZeroInt())
	totalFees := sdk.NewCoin(types.StablecoinDenom, sdkmath.ZeroInt())

	// Create individual settlements and transfer funds
	for i, sender := range senders {
		senderAddr, err := sdk.AccAddressFromBech32(sender)
		if err != nil {
			return 0, nil, types.ErrInvalidSettlement
		}

		// Check compliance
		if err := k.compKeeper.AssertCompliant(wrappedCtx, senderAddr); err != nil {
			return 0, nil, types.ErrComplianceCheckFailed
		}

		// Check balance
		balance := k.bankKeeper.GetBalance(wrappedCtx, senderAddr, amounts[i].Denom)
		if balance.IsLT(amounts[i]) {
			return 0, nil, types.ErrInsufficientFunds
		}

		// Transfer to module
		if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, senderAddr, types.ModuleAccountName, sdk.NewCoins(amounts[i])); err != nil {
			return 0, nil, err
		}

		fee := k.calculateFee(ctx, amounts[i], merchant)
		netAmount := sdk.NewCoin(amounts[i].Denom, amounts[i].Amount.Sub(fee.Amount))

		// Create settlement record
		nextID := k.getNextSettlementID(ctx)
		settlement := types.Settlement{
			Id:            nextID,
			Type:          types.SettlementTypeBatch,
			Sender:        sender,
			Recipient:     merchant,
			Amount:        amounts[i],
			Fee:           fee,
			NetAmount:     netAmount,
			Status:        types.SettlementStatusPending,
			Reference:     references[i],
			CreatedHeight: ctx.BlockHeight(),
			CreatedTime:   ctx.BlockTime(),
		}

		k.storeSettlement(ctx, settlement)
		k.setNextSettlementID(ctx, nextID+1)

		settlementIds[i] = nextID
		totalAmount = sdk.NewCoin(totalAmount.Denom, totalAmount.Amount.Add(amounts[i].Amount))
		totalFees = sdk.NewCoin(totalFees.Denom, totalFees.Amount.Add(fee.Amount))
	}

	netTotal := sdk.NewCoin(totalAmount.Denom, totalAmount.Amount.Sub(totalFees.Amount))

	// Create batch record
	batchID := k.getNextBatchID(ctx)
	batch := types.BatchSettlement{
		Id:            batchID,
		Merchant:      merchant,
		SettlementIds: settlementIds,
		TotalAmount:   totalAmount,
		TotalFees:     totalFees,
		NetAmount:     netTotal,
		Count:         uint64(len(senders)),
		Status:        types.SettlementStatusPending,
		CreatedHeight: ctx.BlockHeight(),
		CreatedTime:   ctx.BlockTime(),
	}

	k.storeBatch(ctx, batch)
	k.setNextBatchID(ctx, batchID+1)

	// Update settlements with batch ID
	for _, sid := range settlementIds {
		settlement, _ := k.GetSettlement(ctx, sid)
		settlement.BatchId = batchID
		k.storeSettlement(ctx, settlement)
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBatchCreated,
			sdk.NewAttribute(types.AttributeKeyBatchID, fmt.Sprintf("%d", batchID)),
			sdk.NewAttribute(types.AttributeKeyMerchant, merchant),
			sdk.NewAttribute(types.AttributeKeyAmount, totalAmount.String()),
		),
	)

	return batchID, settlementIds, nil
}

// SettleBatch settles all payments in a batch
func (k Keeper) SettleBatch(ctx sdk.Context, batchId uint64, authority string) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	// Verify authority
	if authority != k.GetAuthority() {
		return types.ErrUnauthorized
	}

	batch, found := k.GetBatch(ctx, batchId)
	if !found {
		return types.ErrBatchNotFound
	}

	if batch.Status == types.SettlementStatusCompleted {
		return types.ErrBatchAlreadySettled
	}

	merchantAddr, err := sdk.AccAddressFromBech32(batch.Merchant)
	if err != nil {
		return types.ErrInvalidRecipient
	}

	// Compliance check
	if err := k.compKeeper.AssertCompliant(wrappedCtx, merchantAddr); err != nil {
		return types.ErrComplianceCheckFailed
	}

	// Transfer net amount to merchant
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, merchantAddr, sdk.NewCoins(batch.NetAmount)); err != nil {
		return err
	}

	// Collect the batch fees
	if batch.TotalFees.IsPositive() {
		if err := k.collectFee(ctx, batch.TotalFees); err != nil {
			return err
		}
	}

	// Update all settlements in batch
	for _, sid := range batch.SettlementIds {
		settlement, found := k.GetSettlement(ctx, sid)
		if found {
			settlement.Status = types.SettlementStatusCompleted
			settlement.SettledHeight = ctx.BlockHeight()
			settlement.SettledTime = ctx.BlockTime()
			k.storeSettlement(ctx, settlement)
		}
	}

	// Update batch
	batch.Status = types.SettlementStatusCompleted
	batch.SettledHeight = ctx.BlockHeight()
	batch.SettledTime = ctx.BlockTime()
	k.storeBatch(ctx, batch)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBatchSettled,
			sdk.NewAttribute(types.AttributeKeyBatchID, fmt.Sprintf("%d", batchId)),
			sdk.NewAttribute(types.AttributeKeyMerchant, batch.Merchant),
			sdk.NewAttribute(types.AttributeKeyAmount, batch.NetAmount.String()),
			sdk.NewAttribute(types.AttributeKeyFee, batch.TotalFees.String()),
		),
	)

	return nil
}

// ============================================================================
// Payment Channel Operations
// ============================================================================

func (k Keeper) storeChannel(ctx sdk.Context, channel types.PaymentChannel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChannelKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&channel)
	store.Set(mustWriteUint64(channel.Id), bz)
}

// GetChannel retrieves a payment channel by ID
func (k Keeper) GetChannel(ctx sdk.Context, id uint64) (types.PaymentChannel, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChannelKeyPrefix)
	bz := store.Get(mustWriteUint64(id))
	if len(bz) == 0 {
		return types.PaymentChannel{}, false
	}
	var channel types.PaymentChannel
	types.ModuleCdc.MustUnmarshalJSON(bz, &channel)
	return channel, true
}

// OpenChannel opens a new payment channel
func (k Keeper) OpenChannel(ctx sdk.Context, sender, recipient string, deposit sdk.Coin, expiresInBlocks int64) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return 0, types.ErrInvalidSettlement
	}
	recipientAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return 0, types.ErrInvalidRecipient
	}
	if senderAddr.Equals(recipientAddr) {
		return 0, types.ErrInvalidRecipient.Wrap("sender and recipient must be different")
	}

	// Check compliance
	if err := k.compKeeper.AssertCompliant(wrappedCtx, senderAddr); err != nil {
		return 0, types.ErrComplianceCheckFailed
	}
	if err := k.compKeeper.AssertCompliant(wrappedCtx, recipientAddr); err != nil {
		return 0, types.ErrComplianceCheckFailed
	}

	// Check balance
	balance := k.bankKeeper.GetBalance(wrappedCtx, senderAddr, deposit.Denom)
	if balance.IsLT(deposit) {
		return 0, types.ErrInsufficientFunds
	}

	// Transfer deposit to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, senderAddr, types.ModuleAccountName, sdk.NewCoins(deposit)); err != nil {
		return 0, err
	}

	channelID := k.getNextChannelID(ctx)
	channel := types.PaymentChannel{
		Id:              channelID,
		Sender:          sender,
		Recipient:       recipient,
		Deposit:         deposit,
		Spent:           sdk.NewCoin(deposit.Denom, sdkmath.ZeroInt()),
		Balance:         deposit,
		IsOpen:          true,
		OpenedHeight:    ctx.BlockHeight(),
		OpenedTime:      ctx.BlockTime(),
		ExpiresAtHeight: ctx.BlockHeight() + expiresInBlocks,
		Nonce:           0,
	}

	k.storeChannel(ctx, channel)
	k.setNextChannelID(ctx, channelID+1)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChannelOpened,
			sdk.NewAttribute(types.AttributeKeyChannelID, fmt.Sprintf("%d", channelID)),
			sdk.NewAttribute(types.AttributeKeySender, sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient),
			sdk.NewAttribute(types.AttributeKeyAmount, deposit.String()),
		),
	)

	return channelID, nil
}

// ClaimChannel allows the recipient to claim funds from a channel
// The signature must be signed by the channel sender authorizing the payment
func (k Keeper) ClaimChannel(ctx sdk.Context, channelId uint64, recipient sdk.AccAddress, amount sdk.Coin, nonce uint64, signature string) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	channel, found := k.GetChannel(ctx, channelId)
	if !found {
		return types.ErrChannelNotFound
	}

	if !channel.IsOpen {
		return types.ErrChannelClosed
	}

	// Check recipient
	expectedRecipient, err := sdk.AccAddressFromBech32(channel.Recipient)
	if err != nil {
		return types.ErrInvalidRecipient
	}
	if !expectedRecipient.Equals(recipient) {
		return types.ErrUnauthorized
	}

	// Check nonce - must be strictly greater than current
	if nonce <= channel.Nonce {
		return types.ErrInvalidNonce
	}

	// Check balance
	if channel.Balance.IsLT(amount) {
		return types.ErrChannelInsufficientBalance
	}

	// Verify signature from sender authorizing this payment
	// Message format: channelId || recipient || amount || nonce
	if err := k.verifyChannelSignature(ctx, channel, recipient, amount, nonce, signature); err != nil {
		return err
	}

	// Compliance check
	if err := k.compKeeper.AssertCompliant(wrappedCtx, recipient); err != nil {
		return types.ErrComplianceCheckFailed
	}

	// Transfer to recipient
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, recipient, sdk.NewCoins(amount)); err != nil {
		return err
	}

	// Update channel
	channel.Spent = sdk.NewCoin(channel.Spent.Denom, channel.Spent.Amount.Add(amount.Amount))
	channel.Balance = sdk.NewCoin(channel.Balance.Denom, channel.Balance.Amount.Sub(amount.Amount))
	channel.Nonce = nonce
	k.storeChannel(ctx, channel)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChannelUpdated,
			sdk.NewAttribute(types.AttributeKeyChannelID, fmt.Sprintf("%d", channelId)),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	)

	return nil
}

// verifyChannelSignature verifies that the sender has signed the payment authorization
func (k Keeper) verifyChannelSignature(ctx sdk.Context, channel types.PaymentChannel, recipient sdk.AccAddress, amount sdk.Coin, nonce uint64, signature string) error {
	// Decode hex signature
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return types.ErrInvalidSignature
	}

	// Verify signature length (secp256k1 signatures are 64 bytes)
	if len(sigBytes) < 64 {
		return types.ErrInvalidSignature
	}

	// Construct the message that should have been signed
	// Format: "channel_claim:{channelId}:{recipient}:{amount}:{nonce}"
	msg := fmt.Sprintf("channel_claim:%d:%s:%s:%d", channel.Id, recipient.String(), amount.String(), nonce)
	msgBytes := []byte(msg)

	// Get sender's address
	senderAddr, err := sdk.AccAddressFromBech32(channel.Sender)
	if err != nil {
		return types.ErrInvalidSettlement
	}

	// Get sender's public key from account keeper if available
	wrappedCtx := sdk.WrapSDKContext(ctx)
	if k.accountKeeper != nil {
		pubKey, err := k.accountKeeper.GetPubKey(wrappedCtx, senderAddr)
		if err == nil && pubKey != nil {
			if !pubKey.VerifySignature(msgBytes, sigBytes[:64]) {
				return types.ErrSignatureVerificationFailed
			}
			return nil
		}
		ctx.Logger().Debug("account keeper pubkey not available, rejecting claim", "sender", senderAddr.String())
	}

	return types.ErrInvalidSignature
}

// CloseChannel closes a payment channel
func (k Keeper) CloseChannel(ctx sdk.Context, channelId uint64, closer sdk.AccAddress) (sdk.Coin, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	channel, found := k.GetChannel(ctx, channelId)
	if !found {
		return sdk.Coin{}, types.ErrChannelNotFound
	}

	if !channel.IsOpen {
		return sdk.Coin{}, types.ErrChannelClosed
	}

	senderAddr, err := sdk.AccAddressFromBech32(channel.Sender)
	if err != nil {
		return sdk.Coin{}, types.ErrInvalidSettlement
	}

	// Only sender can close, and only after expiration
	if !senderAddr.Equals(closer) {
		return sdk.Coin{}, types.ErrUnauthorized
	}

	// Channel can only be closed AFTER it has expired (block height >= expiration height)
	if ctx.BlockHeight() < channel.ExpiresAtHeight {
		return sdk.Coin{}, types.ErrChannelNotExpired
	}

	// Return remaining balance to sender
	if channel.Balance.IsPositive() {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, senderAddr, sdk.NewCoins(channel.Balance)); err != nil {
			return sdk.Coin{}, err
		}
	}

	finalBalance := channel.Balance

	// Close channel
	channel.IsOpen = false
	channel.ClosedHeight = ctx.BlockHeight()
	channel.ClosedTime = ctx.BlockTime()
	channel.Balance = sdk.NewCoin(channel.Deposit.Denom, sdkmath.ZeroInt())
	k.storeChannel(ctx, channel)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChannelClosed,
			sdk.NewAttribute(types.AttributeKeyChannelID, fmt.Sprintf("%d", channelId)),
			sdk.NewAttribute(types.AttributeKeyAmount, finalBalance.String()),
		),
	)

	return finalBalance, nil
}

// ============================================================================
// Merchant Operations
// ============================================================================

func (k Keeper) storeMerchant(ctx sdk.Context, merchant types.MerchantConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MerchantKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&merchant)
	store.Set([]byte(merchant.Address), bz)
}

// GetMerchant retrieves a merchant configuration
func (k Keeper) GetMerchant(ctx sdk.Context, address string) (types.MerchantConfig, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MerchantKeyPrefix)
	bz := store.Get([]byte(address))
	if len(bz) == 0 {
		return types.MerchantConfig{}, false
	}
	var merchant types.MerchantConfig
	types.ModuleCdc.MustUnmarshalJSON(bz, &merchant)
	return merchant, true
}

// RegisterMerchant registers a new merchant
func (k Keeper) RegisterMerchant(ctx sdk.Context, config types.MerchantConfig) error {
	config.IsActive = true
	config.RegisteredAt = ctx.BlockTime()
	k.storeMerchant(ctx, config)
	return nil
}

// UpdateMerchant updates a merchant configuration
func (k Keeper) UpdateMerchant(ctx sdk.Context, address string, updates map[string]interface{}) error {
	merchant, found := k.GetMerchant(ctx, address)
	if !found {
		return types.ErrMerchantNotFound
	}

	// Apply updates
	for key, value := range updates {
		switch key {
		case "name":
			if name, ok := value.(string); ok && name != "" {
				merchant.Name = name
			}
		case "fee_rate_bps":
			if rate, ok := value.(uint32); ok {
				merchant.FeeRateBps = rate
			}
		case "batch_enabled":
			if enabled, ok := value.(bool); ok {
				merchant.BatchEnabled = enabled
			}
		case "is_active":
			if active, ok := value.(bool); ok {
				merchant.IsActive = active
			}
		}
	}

	k.storeMerchant(ctx, merchant)
	return nil
}

// ============================================================================
// Fee Calculation and Collection
// ============================================================================

func (k Keeper) calculateFee(ctx sdk.Context, amount sdk.Coin, recipient string) sdk.Coin {
	// Get fee rate from params
	params := k.GetParams(ctx)
	feeRateBps := params.DefaultFeeRateBps

	// Check if recipient is a registered merchant with custom fee rate
	merchant, found := k.GetMerchant(ctx, recipient)
	if found && merchant.IsActive && merchant.FeeRateBps > 0 {
		feeRateBps = merchant.FeeRateBps
	}

	// Calculate fee: amount * feeRateBps / 10000
	feeAmount := amount.Amount.Mul(sdkmath.NewInt(int64(feeRateBps))).Quo(sdkmath.NewInt(10000))
	return sdk.NewCoin(amount.Denom, feeAmount)
}

// collectFee transfers the fee from the module account to the fee collector
func (k Keeper) collectFee(ctx sdk.Context, fee sdk.Coin) error {
	if !fee.IsPositive() {
		return nil
	}

	wrappedCtx := sdk.WrapSDKContext(ctx)
	params := k.GetParams(ctx)

	// If fee collector is configured, send to that address
	if params.FeeCollector != "" {
		feeCollectorAddr, err := sdk.AccAddressFromBech32(params.FeeCollector)
		if err != nil {
			return types.ErrInvalidFeeCollector
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, feeCollectorAddr, sdk.NewCoins(fee)); err != nil {
			return err
		}

		// Emit fee collection event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeFeeCollected,
				sdk.NewAttribute(types.AttributeKeyAmount, fee.String()),
				sdk.NewAttribute(types.AttributeKeyRecipient, params.FeeCollector),
			),
		)
	}
	// If no fee collector configured, fees remain in module account (can be withdrawn by governance)

	return nil
}

// ============================================================================
// Iterators
// ============================================================================

// IterateSettlements iterates over all settlements
func (k Keeper) IterateSettlements(ctx sdk.Context, cb func(types.Settlement) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SettlementKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var settlement types.Settlement
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &settlement)
		if cb(settlement) {
			break
		}
	}
}

// IterateBatches iterates over all batches
func (k Keeper) IterateBatches(ctx sdk.Context, cb func(types.BatchSettlement) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BatchKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var batch types.BatchSettlement
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &batch)
		if cb(batch) {
			break
		}
	}
}

// IterateChannels iterates over all channels
func (k Keeper) IterateChannels(ctx sdk.Context, cb func(types.PaymentChannel) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChannelKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var channel types.PaymentChannel
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &channel)
		if cb(channel) {
			break
		}
	}
}

// IterateMerchants iterates over all merchants
func (k Keeper) IterateMerchants(ctx sdk.Context, cb func(types.MerchantConfig) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MerchantKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var merchant types.MerchantConfig
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &merchant)
		if cb(merchant) {
			break
		}
	}
}

// ============================================================================
// Genesis
// ============================================================================

// InitGenesis initializes the settlement module's genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}

	// Store params
	if err := k.SetParams(ctx, state.Params); err != nil {
		panic(fmt.Sprintf("failed to set params during genesis: %v", err))
	}

	k.setNextSettlementID(ctx, state.NextSettlementId)
	k.setNextBatchID(ctx, state.NextBatchId)
	k.setNextChannelID(ctx, state.NextChannelId)

	for _, settlement := range state.Settlements {
		k.storeSettlement(ctx, settlement)
	}
	for _, batch := range state.Batches {
		k.storeBatch(ctx, batch)
	}
	for _, channel := range state.Channels {
		k.storeChannel(ctx, channel)
	}
	for _, merchant := range state.Merchants {
		k.storeMerchant(ctx, merchant)
	}
}

// ExportGenesis exports the settlement module's genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Params = k.GetParams(ctx)
	state.NextSettlementId = k.getNextSettlementID(ctx)
	state.NextBatchId = k.getNextBatchID(ctx)
	state.NextChannelId = k.getNextChannelID(ctx)

	k.IterateSettlements(ctx, func(s types.Settlement) bool {
		state.Settlements = append(state.Settlements, s)
		return false
	})
	k.IterateBatches(ctx, func(b types.BatchSettlement) bool {
		state.Batches = append(state.Batches, b)
		return false
	})
	k.IterateChannels(ctx, func(c types.PaymentChannel) bool {
		state.Channels = append(state.Channels, c)
		return false
	})
	k.IterateMerchants(ctx, func(m types.MerchantConfig) bool {
		state.Merchants = append(state.Merchants, m)
		return false
	})

	return state
}

// ============================================================================
// EndBlock Processing
// ============================================================================

// ProcessExpiredEscrows handles expired escrow settlements
// When an escrow expires without being released, funds are returned to sender
func (k Keeper) ProcessExpiredEscrows(ctx sdk.Context) {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	currentTime := ctx.BlockTime()

	k.IterateSettlements(ctx, func(s types.Settlement) bool {
		// Only process pending escrow settlements that have expired
		if s.Type != types.SettlementTypeEscrow {
			return false
		}
		if s.Status != types.SettlementStatusPending {
			return false
		}
		if s.ExpiresAt.IsZero() || currentTime.Before(s.ExpiresAt) {
			return false
		}

		// Escrow has expired - refund to sender
		senderAddr, err := sdk.AccAddressFromBech32(s.Sender)
		if err != nil {
			return false // Skip invalid - should not happen
		}

		// Refund the full amount to sender
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, senderAddr, sdk.NewCoins(s.Amount)); err != nil {
			// Log error but continue processing
			ctx.Logger().Error("failed to refund expired escrow", "settlement_id", s.Id, "error", err)
			return false
		}

		// Update settlement status
		s.Status = types.SettlementStatusCancelled
		s.Metadata = "expired: auto-refunded to sender"
		k.storeSettlement(ctx, s)

		// Emit event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEscrowExpired,
				sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", s.Id)),
				sdk.NewAttribute(types.AttributeKeySender, s.Sender),
				sdk.NewAttribute(types.AttributeKeyAmount, s.Amount.String()),
			),
		)

		return false // Continue iteration
	})
}

// ProcessExpiredChannels handles expired payment channels
// Note: This doesn't auto-close channels, it just marks them for close eligibility
// Actual closing requires sender to submit a CloseChannel transaction
func (k Keeper) ProcessExpiredChannels(ctx sdk.Context) {
	currentHeight := ctx.BlockHeight()

	k.IterateChannels(ctx, func(c types.PaymentChannel) bool {
		// Only check open channels
		if !c.IsOpen {
			return false
		}

		// Check if channel has expired
		if currentHeight < c.ExpiresAtHeight {
			return false
		}

		// Emit event that channel is now closeable
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeChannelExpired,
				sdk.NewAttribute(types.AttributeKeyChannelID, fmt.Sprintf("%d", c.Id)),
				sdk.NewAttribute(types.AttributeKeySender, c.Sender),
				sdk.NewAttribute(types.AttributeKeyRecipient, c.Recipient),
				sdk.NewAttribute(types.AttributeKeyAmount, c.Balance.String()),
			),
		)

		return false // Continue iteration
	})
}

// ============================================================================
// Instant Checkout (Streamlined Ecommerce)
// ============================================================================

// InstantCheckout performs a streamlined checkout combining compliance check and payment
// This is optimized for ecommerce use cases where speed and simplicity are critical
func (k Keeper) InstantCheckout(ctx sdk.Context, customer, merchant string, amount sdk.Coin, orderRef string, useEscrow bool, metadata string) (uint64, sdk.Coin, sdk.Coin, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	customerAddr, err := sdk.AccAddressFromBech32(customer)
	if err != nil {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrInvalidSettlement
	}
	merchantAddr, err := sdk.AccAddressFromBech32(merchant)
	if err != nil {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrInvalidRecipient
	}

	// Batch compliance checks for efficiency
	if err := k.compKeeper.AssertCompliant(wrappedCtx, customerAddr); err != nil {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrComplianceCheckFailed
	}
	if err := k.compKeeper.AssertCompliant(wrappedCtx, merchantAddr); err != nil {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrComplianceCheckFailed
	}

	// Validate amount
	params := k.GetParams(ctx)
	if amount.IsLT(params.MinSettlementAmount) {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrSettlementTooSmall
	}
	if amount.IsGTE(params.MaxSettlementAmount) {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrSettlementTooLarge
	}

	// Check balance
	balance := k.bankKeeper.GetBalance(wrappedCtx, customerAddr, amount.Denom)
	if balance.IsLT(amount) {
		return 0, sdk.Coin{}, sdk.Coin{}, types.ErrInsufficientFunds
	}

	// Calculate fee
	fee := k.calculateFee(ctx, amount, merchant)
	netAmount := sdk.NewCoin(amount.Denom, amount.Amount.Sub(fee.Amount))

	var settlementId uint64

	if useEscrow {
		// Create escrow for buyer protection
		settlementId, err = k.CreateEscrow(ctx, customer, merchant, amount, orderRef, metadata, params.DefaultEscrowExpiration)
		if err != nil {
			return 0, sdk.Coin{}, sdk.Coin{}, err
		}
	} else {
		// Instant transfer for speed
		settlementId, err = k.InstantTransfer(ctx, customer, merchant, amount, orderRef, metadata)
		if err != nil {
			return 0, sdk.Coin{}, sdk.Coin{}, err
		}
	}

	// Emit checkout event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInstantCheckout,
			sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", settlementId)),
			sdk.NewAttribute(types.AttributeKeySender, customer),
			sdk.NewAttribute(types.AttributeKeyRecipient, merchant),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyFee, fee.String()),
			sdk.NewAttribute(types.AttributeKeyReference, orderRef),
			sdk.NewAttribute("use_escrow", fmt.Sprintf("%t", useEscrow)),
		),
	)

	return settlementId, netAmount, fee, nil
}

// PartialRefund processes a partial refund for a completed settlement
func (k Keeper) PartialRefund(ctx sdk.Context, authority string, settlementId uint64, refundAmount sdk.Coin, reason string) (sdk.Coin, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	// Verify authority (merchant or module authority)
	settlement, found := k.GetSettlement(ctx, settlementId)
	if !found {
		return sdk.Coin{}, types.ErrSettlementNotFound
	}

	// Only the merchant or module authority can issue refunds
	if authority != settlement.Recipient && authority != k.GetAuthority() {
		return sdk.Coin{}, types.ErrUnauthorized
	}

	// Validate settlement status - can only refund completed settlements
	if settlement.Status != types.SettlementStatusCompleted {
		return sdk.Coin{}, types.ErrSettlementNotCompleted
	}

	// Validate refund amount
	if refundAmount.Amount.GT(settlement.NetAmount.Amount) {
		return sdk.Coin{}, types.ErrRefundTooLarge
	}

	// Get addresses
	merchantAddr, err := sdk.AccAddressFromBech32(settlement.Recipient)
	if err != nil {
		return sdk.Coin{}, types.ErrInvalidRecipient
	}
	customerAddr, err := sdk.AccAddressFromBech32(settlement.Sender)
	if err != nil {
		return sdk.Coin{}, types.ErrInvalidSettlement
	}

	// Transfer refund from merchant to customer
	if err := k.bankKeeper.SendCoins(wrappedCtx, merchantAddr, customerAddr, sdk.NewCoins(refundAmount)); err != nil {
		return sdk.Coin{}, types.ErrInsufficientFunds
	}

	// Calculate remaining amount
	remainingAmount := sdk.NewCoin(settlement.NetAmount.Denom, settlement.NetAmount.Amount.Sub(refundAmount.Amount))

	// Update settlement status if fully refunded
	if remainingAmount.IsZero() {
		settlement.Status = types.SettlementStatusRefunded
	}
	settlement.Metadata = fmt.Sprintf("partial_refund: %s - %s", refundAmount.String(), reason)
	k.storeSettlement(ctx, settlement)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePartialRefund,
			sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", settlementId)),
			sdk.NewAttribute(types.AttributeKeyAmount, refundAmount.String()),
			sdk.NewAttribute("remaining", remainingAmount.String()),
			sdk.NewAttribute("reason", reason),
		),
	)

	return remainingAmount, nil
}

// ============================================================================
// Helpers
// ============================================================================

func mustWriteUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
