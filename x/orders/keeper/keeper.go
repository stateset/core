package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/orders/types"
)

// Keeper manages order state.
type Keeper struct {
	storeKey         storetypes.StoreKey
	cdc              codec.BinaryCodec
	authority        string
	bankKeeper       types.BankKeeper
	complianceKeeper types.ComplianceKeeper
	settlementKeeper types.SettlementKeeper
	accountKeeper    types.AccountKeeper
}

// NewKeeper creates a new orders keeper.
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	authority string,
	bank types.BankKeeper,
	compliance types.ComplianceKeeper,
	settlement types.SettlementKeeper,
	account types.AccountKeeper,
) Keeper {
	return Keeper{
		storeKey:         key,
		cdc:              cdc,
		authority:        authority,
		bankKeeper:       bank,
		complianceKeeper: compliance,
		settlementKeeper: settlement,
		accountKeeper:    account,
	}
}

// GetAuthority returns the module authority address.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// GetParams returns module parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.ParamsKey) {
		return types.DefaultParams()
	}
	var params types.Params
	types.ModuleCdc.MustUnmarshalJSON(store.Get(types.ParamsKey), &params)
	return params
}

// SetParams stores module parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ParamsKey, types.ModuleCdc.MustMarshalJSON(&params))
	return nil
}

// ============================================================================
// Order ID Management
// ============================================================================

func (k Keeper) getNextOrderID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.NextOrderIDKey) {
		return 1
	}
	return binary.BigEndian.Uint64(store.Get(types.NextOrderIDKey))
}

func (k Keeper) setNextOrderID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(types.NextOrderIDKey, bz)
}

func (k Keeper) getNextDisputeID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.DisputeKeyPrefix)
	if len(bz) == 0 {
		return 1
	}
	return binary.BigEndian.Uint64(bz) + 1
}

func mustBz(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// ============================================================================
// Order Storage
// ============================================================================

func (k Keeper) setOrder(ctx sdk.Context, order types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderKeyPrefix)
	store.Set(mustBz(order.Id), types.ModuleCdc.MustMarshalJSON(&order))

	// Update indexes
	k.indexOrderByCustomer(ctx, order)
	k.indexOrderByMerchant(ctx, order)
	k.indexOrderByStatus(ctx, order)
}

func (k Keeper) indexOrderByCustomer(ctx sdk.Context, order types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderByCustomerKeyPrefix)
	key := append([]byte(order.Customer), mustBz(order.Id)...)
	store.Set(key, mustBz(order.Id))
}

func (k Keeper) indexOrderByMerchant(ctx sdk.Context, order types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderByMerchantKeyPrefix)
	key := append([]byte(order.Merchant), mustBz(order.Id)...)
	store.Set(key, mustBz(order.Id))
}

func (k Keeper) indexOrderByStatus(ctx sdk.Context, order types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderByStatusKeyPrefix)
	key := append([]byte(order.Status), mustBz(order.Id)...)
	store.Set(key, mustBz(order.Id))
}

// GetOrder retrieves an order by ID.
func (k Keeper) GetOrder(ctx sdk.Context, id uint64) (types.Order, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.Order{}, false
	}
	var order types.Order
	types.ModuleCdc.MustUnmarshalJSON(bz, &order)
	return order, true
}

// ============================================================================
// Order Operations
// ============================================================================

// CreateOrder creates a new order.
func (k Keeper) CreateOrder(ctx sdk.Context, customer, merchant string, items []types.OrderItem, shippingInfo types.ShippingInfo, metadata string) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)

	customerAddr, err := sdk.AccAddressFromBech32(customer)
	if err != nil {
		return 0, types.ErrInvalidCustomer
	}
	merchantAddr, err := sdk.AccAddressFromBech32(merchant)
	if err != nil {
		return 0, types.ErrInvalidMerchant
	}

	// Compliance checks
	if err := k.complianceKeeper.AssertCompliant(wrappedCtx, customerAddr); err != nil {
		return 0, types.ErrComplianceFailed
	}
	if err := k.complianceKeeper.AssertCompliant(wrappedCtx, merchantAddr); err != nil {
		return 0, types.ErrComplianceFailed
	}

	params := k.GetParams(ctx)

	// Calculate totals
	subtotal := sdkmath.ZeroInt()
	for i := range items {
		itemTotal := items[i].UnitPrice.Amount.Mul(sdkmath.NewInt(int64(items[i].Quantity)))
		items[i].TotalPrice = sdk.NewCoin(items[i].UnitPrice.Denom, itemTotal)
		subtotal = subtotal.Add(itemTotal)
	}

	// Create order
	now := ctx.BlockTime()
	orderId := k.getNextOrderID(ctx)

	order := types.Order{
		Id:           orderId,
		Customer:     customer,
		Merchant:     merchant,
		Status:       types.OrderStatusPending,
		Items:        items,
		Subtotal:     sdk.NewCoin(params.StablecoinDenom, subtotal),
		ShippingCost: sdk.NewCoin(params.StablecoinDenom, sdkmath.ZeroInt()),
		TaxAmount:    sdk.NewCoin(params.StablecoinDenom, sdkmath.ZeroInt()),
		DiscountAmount: sdk.NewCoin(params.StablecoinDenom, sdkmath.ZeroInt()),
		TotalAmount:  sdk.NewCoin(params.StablecoinDenom, subtotal),
		PaymentInfo: types.PaymentInfo{
			Status: types.PaymentStatusPending,
		},
		ShippingInfo: shippingInfo,
		Metadata:     metadata,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiresAt:    now.Add(time.Duration(params.DefaultOrderExpiration) * time.Second),
	}

	// Validate amount limits
	if order.TotalAmount.IsLT(params.MinOrderAmount) {
		return 0, types.ErrInvalidAmount
	}
	if order.TotalAmount.IsGTE(params.MaxOrderAmount) {
		return 0, types.ErrInvalidAmount
	}

	k.setOrder(ctx, order)
	k.setNextOrderID(ctx, orderId+1)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_created",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("customer", customer),
			sdk.NewAttribute("merchant", merchant),
			sdk.NewAttribute("total_amount", order.TotalAmount.String()),
		),
	)

	return orderId, nil
}

// ConfirmOrder confirms an order by the merchant.
func (k Keeper) ConfirmOrder(ctx sdk.Context, merchant string, orderId uint64) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	if order.Merchant != merchant {
		return types.ErrUnauthorized
	}

	if !order.IsValidTransition(types.OrderStatusConfirmed) {
		return types.ErrInvalidTransition
	}

	order.Status = types.OrderStatusConfirmed
	order.UpdatedAt = ctx.BlockTime()
	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_confirmed",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("merchant", merchant),
		),
	)

	return nil
}

// PayOrder pays for an order using stablecoin.
func (k Keeper) PayOrder(ctx sdk.Context, customer string, orderId uint64, amount sdk.Coin, useEscrow bool) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	if order.Customer != customer {
		return types.ErrUnauthorized
	}

	if order.Status != types.OrderStatusConfirmed {
		return types.ErrInvalidStatus
	}

	if !amount.IsGTE(order.TotalAmount) {
		return types.ErrInvalidAmount
	}

	params := k.GetParams(ctx)
	reference := fmt.Sprintf("order_%d", orderId)

	var settlementId uint64
	var err error

	if useEscrow {
		// Use escrow for buyer protection
		settlementId, err = k.settlementKeeper.CreateEscrow(
			ctx,
			customer,
			order.Merchant,
			order.TotalAmount,
			reference,
			order.Metadata,
			params.DefaultEscrowExpiration,
		)
		if err != nil {
			return types.ErrSettlementFailed
		}
		order.PaymentInfo.Method = "escrow"
		order.PaymentInfo.EscrowId = settlementId
	} else {
		// Instant transfer to merchant
		settlementId, err = k.settlementKeeper.InstantTransfer(
			ctx,
			customer,
			order.Merchant,
			order.TotalAmount,
			reference,
			order.Metadata,
		)
		if err != nil {
			return types.ErrSettlementFailed
		}
		order.PaymentInfo.Method = "instant"
	}

	order.Status = types.OrderStatusPaid
	order.PaymentInfo.Status = types.PaymentStatusCaptured
	order.PaymentInfo.SettlementId = settlementId
	order.PaymentInfo.PaidAmount = order.TotalAmount
	order.PaymentInfo.PaidAt = ctx.BlockTime()
	order.PaidAt = ctx.BlockTime()
	order.UpdatedAt = ctx.BlockTime()
	order.SettlementId = settlementId

	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_paid",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("customer", customer),
			sdk.NewAttribute("amount", order.TotalAmount.String()),
			sdk.NewAttribute("settlement_id", fmt.Sprintf("%d", settlementId)),
			sdk.NewAttribute("use_escrow", fmt.Sprintf("%t", useEscrow)),
		),
	)

	return nil
}

// ShipOrder marks an order as shipped.
func (k Keeper) ShipOrder(ctx sdk.Context, merchant string, orderId uint64, carrier, trackingNumber string) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	if order.Merchant != merchant {
		return types.ErrUnauthorized
	}

	if !order.IsValidTransition(types.OrderStatusShipped) {
		return types.ErrInvalidTransition
	}

	order.Status = types.OrderStatusShipped
	order.ShippingInfo.Carrier = carrier
	order.ShippingInfo.TrackingNumber = trackingNumber
	order.ShippedAt = ctx.BlockTime()
	order.UpdatedAt = ctx.BlockTime()

	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_shipped",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("merchant", merchant),
			sdk.NewAttribute("carrier", carrier),
			sdk.NewAttribute("tracking_number", trackingNumber),
		),
	)

	return nil
}

// DeliverOrder marks an order as delivered.
func (k Keeper) DeliverOrder(ctx sdk.Context, signer string, orderId uint64) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	// Either customer or merchant can mark as delivered
	if signer != order.Customer && signer != order.Merchant {
		return types.ErrUnauthorized
	}

	if !order.IsValidTransition(types.OrderStatusDelivered) {
		return types.ErrInvalidTransition
	}

	order.Status = types.OrderStatusDelivered
	order.DeliveredAt = ctx.BlockTime()
	order.ShippingInfo.ActualDelivery = ctx.BlockTime()
	order.UpdatedAt = ctx.BlockTime()

	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_delivered",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("delivered_by", signer),
		),
	)

	return nil
}

// CompleteOrder marks an order as completed and releases escrow.
func (k Keeper) CompleteOrder(ctx sdk.Context, customer string, orderId uint64) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	if order.Customer != customer {
		return types.ErrUnauthorized
	}

	if !order.IsValidTransition(types.OrderStatusCompleted) {
		return types.ErrInvalidTransition
	}

	// Release escrow if applicable
	if order.PaymentInfo.Method == "escrow" && order.PaymentInfo.EscrowId > 0 {
		customerAddr, _ := sdk.AccAddressFromBech32(customer)
		if err := k.settlementKeeper.ReleaseEscrow(ctx, order.PaymentInfo.EscrowId, customerAddr); err != nil {
			return types.ErrSettlementFailed
		}
		order.PaymentInfo.Status = types.PaymentStatusReleased
	}

	order.Status = types.OrderStatusCompleted
	order.CompletedAt = ctx.BlockTime()
	order.UpdatedAt = ctx.BlockTime()

	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_completed",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("customer", customer),
		),
	)

	return nil
}

// CancelOrder cancels an unpaid order.
func (k Keeper) CancelOrder(ctx sdk.Context, signer string, orderId uint64, reason string) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	// Customer or merchant can cancel
	if signer != order.Customer && signer != order.Merchant {
		return types.ErrUnauthorized
	}

	if !order.IsValidTransition(types.OrderStatusCancelled) {
		return types.ErrCannotCancel
	}

	order.Status = types.OrderStatusCancelled
	order.Metadata = fmt.Sprintf("cancelled: %s", reason)
	order.UpdatedAt = ctx.BlockTime()

	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_cancelled",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("cancelled_by", signer),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// RefundOrder processes a refund for an order.
func (k Keeper) RefundOrder(ctx sdk.Context, merchant string, orderId uint64, refundAmount sdk.Coin, reason string, fullRefund bool) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return types.ErrOrderNotFound
	}

	if order.Merchant != merchant {
		return types.ErrUnauthorized
	}

	if !order.CanBeRefunded() {
		return types.ErrCannotRefund
	}

	// Determine refund amount
	if fullRefund {
		refundAmount = order.TotalAmount
	}

	// If using escrow, refund from escrow
	if order.PaymentInfo.Method == "escrow" && order.PaymentInfo.EscrowId > 0 {
		merchantAddr, _ := sdk.AccAddressFromBech32(merchant)
		if err := k.settlementKeeper.RefundEscrow(ctx, order.PaymentInfo.EscrowId, merchantAddr, reason); err != nil {
			return types.ErrSettlementFailed
		}
	} else {
		// For instant payments, merchant must transfer back
		merchantAddr, _ := sdk.AccAddressFromBech32(merchant)
		customerAddr, _ := sdk.AccAddressFromBech32(order.Customer)
		if err := k.bankKeeper.SendCoins(sdk.WrapSDKContext(ctx), merchantAddr, customerAddr, sdk.NewCoins(refundAmount)); err != nil {
			return types.ErrInsufficientFunds
		}
	}

	order.Status = types.OrderStatusRefunded
	order.PaymentInfo.Status = types.PaymentStatusRefunded
	order.PaymentInfo.RefundedAmount = refundAmount
	order.UpdatedAt = ctx.BlockTime()
	order.Metadata = fmt.Sprintf("refunded: %s", reason)

	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"order_refunded",
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("merchant", merchant),
			sdk.NewAttribute("refund_amount", refundAmount.String()),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// ============================================================================
// Dispute Operations
// ============================================================================

func (k Keeper) setDispute(ctx sdk.Context, dispute types.Dispute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DisputeKeyPrefix)
	store.Set(mustBz(dispute.Id), types.ModuleCdc.MustMarshalJSON(&dispute))
}

// GetDispute retrieves a dispute by ID.
func (k Keeper) GetDispute(ctx sdk.Context, id uint64) (types.Dispute, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DisputeKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.Dispute{}, false
	}
	var dispute types.Dispute
	types.ModuleCdc.MustUnmarshalJSON(bz, &dispute)
	return dispute, true
}

// OpenDispute opens a dispute on an order.
func (k Keeper) OpenDispute(ctx sdk.Context, customer string, orderId uint64, reason, description string, evidence []string) (uint64, error) {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return 0, types.ErrOrderNotFound
	}

	if order.Customer != customer {
		return 0, types.ErrUnauthorized
	}

	if !order.CanBeDisputed() {
		return 0, types.ErrCannotDispute
	}

	if order.DisputeId != 0 {
		return 0, types.ErrDisputeAlreadyExists
	}

	disputeId := k.getNextDisputeID(ctx)

	dispute := types.Dispute{
		Id:          disputeId,
		OrderId:     orderId,
		Customer:    customer,
		Merchant:    order.Merchant,
		Reason:      reason,
		Description: description,
		Evidence:    evidence,
		Status:      types.DisputeStatusOpen,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Amount:      order.TotalAmount,
	}

	k.setDispute(ctx, dispute)

	// Update order
	order.Status = types.OrderStatusDisputed
	order.DisputeId = disputeId
	order.UpdatedAt = ctx.BlockTime()
	k.setOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"dispute_opened",
			sdk.NewAttribute("dispute_id", fmt.Sprintf("%d", disputeId)),
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute("customer", customer),
			sdk.NewAttribute("reason", reason),
		),
	)

	return disputeId, nil
}

// ResolveDispute resolves a dispute (by authority).
func (k Keeper) ResolveDispute(ctx sdk.Context, authority string, disputeId uint64, resolution string, refundAmount sdk.Coin, toCustomer bool) error {
	if authority != k.authority {
		return types.ErrUnauthorized
	}

	dispute, found := k.GetDispute(ctx, disputeId)
	if !found {
		return types.ErrDisputeNotFound
	}

	order, found := k.GetOrder(ctx, dispute.OrderId)
	if !found {
		return types.ErrOrderNotFound
	}

	// Process resolution
	if toCustomer && refundAmount.IsPositive() {
		// Refund to customer
		if order.PaymentInfo.Method == "escrow" && order.PaymentInfo.EscrowId > 0 {
			merchantAddr, _ := sdk.AccAddressFromBech32(order.Merchant)
			if err := k.settlementKeeper.RefundEscrow(ctx, order.PaymentInfo.EscrowId, merchantAddr, resolution); err != nil {
				return types.ErrSettlementFailed
			}
		}
		order.Status = types.OrderStatusRefunded
		order.PaymentInfo.Status = types.PaymentStatusRefunded
		order.PaymentInfo.RefundedAmount = refundAmount
	} else {
		// Release to merchant
		if order.PaymentInfo.Method == "escrow" && order.PaymentInfo.EscrowId > 0 {
			customerAddr, _ := sdk.AccAddressFromBech32(order.Customer)
			if err := k.settlementKeeper.ReleaseEscrow(ctx, order.PaymentInfo.EscrowId, customerAddr); err != nil {
				return types.ErrSettlementFailed
			}
		}
		order.Status = types.OrderStatusCompleted
		order.PaymentInfo.Status = types.PaymentStatusReleased
	}

	order.UpdatedAt = ctx.BlockTime()
	k.setOrder(ctx, order)

	dispute.Status = types.DisputeStatusResolved
	dispute.Resolution = resolution
	dispute.ResolvedBy = authority
	dispute.ResolvedAt = ctx.BlockTime()
	dispute.UpdatedAt = ctx.BlockTime()
	k.setDispute(ctx, dispute)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"dispute_resolved",
			sdk.NewAttribute("dispute_id", fmt.Sprintf("%d", disputeId)),
			sdk.NewAttribute("order_id", fmt.Sprintf("%d", dispute.OrderId)),
			sdk.NewAttribute("resolution", resolution),
			sdk.NewAttribute("to_customer", fmt.Sprintf("%t", toCustomer)),
		),
	)

	return nil
}

// ============================================================================
// Iterators
// ============================================================================

// IterateOrders iterates over all orders.
func (k Keeper) IterateOrders(ctx sdk.Context, cb func(types.Order) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var order types.Order
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &order)
		if cb(order) {
			break
		}
	}
}

// IterateDisputes iterates over all disputes.
func (k Keeper) IterateDisputes(ctx sdk.Context, cb func(types.Dispute) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DisputeKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var dispute types.Dispute
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &dispute)
		if cb(dispute) {
			break
		}
	}
}

// ============================================================================
// Genesis
// ============================================================================

// InitGenesis initializes the orders module's genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	if err := k.SetParams(ctx, state.Params); err != nil {
		panic(fmt.Sprintf("failed to set params: %v", err))
	}

	k.setNextOrderID(ctx, state.NextOrderId)

	for _, order := range state.Orders {
		k.setOrder(ctx, order)
	}
	for _, dispute := range state.Disputes {
		k.setDispute(ctx, dispute)
	}
}

// ExportGenesis exports the orders module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Params = k.GetParams(ctx)
	state.NextOrderId = k.getNextOrderID(ctx)

	k.IterateOrders(ctx, func(order types.Order) bool {
		state.Orders = append(state.Orders, order)
		return false
	})

	k.IterateDisputes(ctx, func(dispute types.Dispute) bool {
		state.Disputes = append(state.Disputes, dispute)
		return false
	})

	return state
}

// ============================================================================
// EndBlock Processing
// ============================================================================

// ProcessExpiredOrders handles expired unpaid orders.
func (k Keeper) ProcessExpiredOrders(ctx sdk.Context) {
	currentTime := ctx.BlockTime()

	k.IterateOrders(ctx, func(order types.Order) bool {
		// Only process pending or confirmed orders that have expired
		if order.Status != types.OrderStatusPending && order.Status != types.OrderStatusConfirmed {
			return false
		}

		if order.ExpiresAt.IsZero() || currentTime.Before(order.ExpiresAt) {
			return false
		}

		// Cancel expired order
		order.Status = types.OrderStatusCancelled
		order.Metadata = "expired: auto-cancelled"
		order.UpdatedAt = currentTime
		k.setOrder(ctx, order)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"order_expired",
				sdk.NewAttribute("order_id", fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute("customer", order.Customer),
			),
		)

		return false
	})
}

// ProcessAutoCompleteOrders handles auto-completing delivered orders.
func (k Keeper) ProcessAutoCompleteOrders(ctx sdk.Context) {
	params := k.GetParams(ctx)
	if !params.AutoCompleteAfterDelivery {
		return
	}

	currentTime := ctx.BlockTime()
	autoCompleteWindow := time.Duration(params.AutoCompleteWindow) * time.Second

	k.IterateOrders(ctx, func(order types.Order) bool {
		// Only process delivered orders
		if order.Status != types.OrderStatusDelivered {
			return false
		}

		// Check if auto-complete window has passed
		autoCompleteTime := order.DeliveredAt.Add(autoCompleteWindow)
		if currentTime.Before(autoCompleteTime) {
			return false
		}

		// Auto-complete the order
		if order.PaymentInfo.Method == "escrow" && order.PaymentInfo.EscrowId > 0 {
			customerAddr, _ := sdk.AccAddressFromBech32(order.Customer)
			if err := k.settlementKeeper.ReleaseEscrow(ctx, order.PaymentInfo.EscrowId, customerAddr); err != nil {
				ctx.Logger().Error("failed to release escrow for auto-complete", "order_id", order.Id, "error", err)
				return false
			}
			order.PaymentInfo.Status = types.PaymentStatusReleased
		}

		order.Status = types.OrderStatusCompleted
		order.CompletedAt = currentTime
		order.UpdatedAt = currentTime
		order.Metadata = "auto-completed after delivery window"
		k.setOrder(ctx, order)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"order_auto_completed",
				sdk.NewAttribute("order_id", fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute("customer", order.Customer),
			),
		)

		return false
	})
}
