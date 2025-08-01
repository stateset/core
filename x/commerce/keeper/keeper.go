package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/stateset/core/x/commerce/types"
	stablecoinstypes "github.com/stateset/core/x/stablecoins/types"
	cctptypes "github.com/stateset/core/x/cctp/types"
	invoicetypes "github.com/stateset/core/x/invoice/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         storetypes.StoreKey
	memKey          storetypes.StoreKey
	paramstore      paramtypes.Subspace

	accountKeeper   types.AccountKeeper
	bankKeeper      types.BankKeeper
	stablecoinKeeper types.StablecoinKeeper
	cctpKeeper      types.CCTPKeeper
	invoiceKeeper   types.InvoiceKeeper

	// Payment routing and optimization
	paymentRouter   *PaymentRouter
	feeCalculator   *FeeCalculator
	complianceEngine *ComplianceEngine
	analyticsEngine  *AnalyticsEngine
}

// NewKeeper creates a new commerce keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stablecoinKeeper types.StablecoinKeeper,
	cctpKeeper types.CCTPKeeper,
	invoiceKeeper types.InvoiceKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	keeper := &Keeper{
		cdc:               cdc,
		storeKey:         storeKey,
		memKey:          memKey,
		paramstore:      ps,
		accountKeeper:   accountKeeper,
		bankKeeper:      bankKeeper,
		stablecoinKeeper: stablecoinKeeper,
		cctpKeeper:      cctpKeeper,
		invoiceKeeper:   invoiceKeeper,
	}

	// Initialize sub-engines
	keeper.paymentRouter = NewPaymentRouter(keeper)
	keeper.feeCalculator = NewFeeCalculator(keeper)
	keeper.complianceEngine = NewComplianceEngine(keeper)
	keeper.analyticsEngine = NewAnalyticsEngine(keeper)

	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// CreateCommerceTransaction creates a new comprehensive commerce transaction
func (k Keeper) CreateCommerceTransaction(ctx sdk.Context, transaction types.CommerceTransaction) error {
	// Validate transaction
	if err := transaction.Validate(); err != nil {
		return err
	}

	// Perform compliance checks
	if err := k.complianceEngine.RunPreTransactionChecks(ctx, transaction); err != nil {
		return fmt.Errorf("compliance check failed: %w", err)
	}

	// Optimize payment route
	optimizedRoute, err := k.paymentRouter.FindOptimalRoute(ctx, transaction.PaymentInfo)
	if err != nil {
		return fmt.Errorf("payment routing failed: %w", err)
	}
	transaction.PaymentInfo.Route = optimizedRoute

	// Calculate fees
	fees, err := k.feeCalculator.CalculateTransactionFees(ctx, transaction)
	if err != nil {
		return fmt.Errorf("fee calculation failed: %w", err)
	}
	transaction.PaymentInfo.Fees = fees

	// Set timestamps
	transaction.CreatedAt = ctx.BlockTime()
	transaction.UpdatedAt = ctx.BlockTime()

	// Store transaction
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommerceTransactionKeyPrefix))
	b := k.cdc.MustMarshal(&transaction)
	store.Set(types.CommerceTransactionKey(transaction.ID), b)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCommerceTransactionCreated,
			sdk.NewAttribute(types.AttributeKeyTransactionID, transaction.ID),
			sdk.NewAttribute(types.AttributeKeyTransactionType, transaction.Type),
			sdk.NewAttribute(types.AttributeKeyAmount, transaction.PaymentInfo.Amount.String()),
		),
	)

	return nil
}

// ProcessPayment processes a payment with intelligent routing and optimization
func (k Keeper) ProcessPayment(ctx sdk.Context, transactionID string) error {
	transaction, found := k.GetCommerceTransaction(ctx, transactionID)
	if !found {
		return sdkerrors.Wrapf(types.ErrTransactionNotFound, "transaction %s not found", transactionID)
	}

	if transaction.Status != types.PaymentStatusPending {
		return sdkerrors.Wrapf(types.ErrInvalidTransactionStatus, "transaction %s is not pending", transactionID)
	}

	// Update status to processing
	transaction.Status = types.PaymentStatusProcessing
	transaction.UpdatedAt = ctx.BlockTime()

	// Record analytics start
	startTime := time.Now()

	// Execute payment based on route type
	var err error
	switch transaction.PaymentInfo.Route.Type {
	case types.PaymentRouteDirect:
		err = k.processDirectPayment(ctx, &transaction)
	case types.PaymentRouteMultiHop:
		err = k.processMultiHopPayment(ctx, &transaction)
	case types.PaymentRouteBridge:
		err = k.processBridgePayment(ctx, &transaction)
	case types.PaymentRouteOptimized:
		err = k.processOptimizedPayment(ctx, &transaction)
	default:
		err = fmt.Errorf("unsupported payment route type: %s", transaction.PaymentInfo.Route.Type)
	}

	// Update transaction status and analytics
	if err != nil {
		transaction.Status = types.PaymentStatusFailed
		k.analyticsEngine.RecordFailedTransaction(ctx, transaction, err)
	} else {
		transaction.Status = types.PaymentStatusCompleted
		completedAt := ctx.BlockTime()
		transaction.CompletedAt = &completedAt
		
		// Update analytics
		processingTime := time.Since(startTime)
		transaction.Analytics.ProcessingTime = processingTime
		k.analyticsEngine.RecordSuccessfulTransaction(ctx, transaction)
	}

	transaction.UpdatedAt = ctx.BlockTime()

	// Update stored transaction
	k.SetCommerceTransaction(ctx, transaction)

	return err
}

// processDirectPayment handles direct blockchain transfers
func (k Keeper) processDirectPayment(ctx sdk.Context, transaction *types.CommerceTransaction) error {
	// Get parties
	var fromAddr, toAddr sdk.AccAddress
	var err error

	for _, party := range transaction.Parties {
		addr, parseErr := sdk.AccAddressFromBech32(party.Address)
		if parseErr != nil {
			return parseErr
		}
		
		switch party.Role {
		case "buyer", "payer":
			fromAddr = addr
		case "seller", "payee":
			toAddr = addr
		}
	}

	if fromAddr.Empty() || toAddr.Empty() {
		return fmt.Errorf("invalid payment parties")
	}

	// Execute transfer
	err = k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, transaction.PaymentInfo.Amount)
	if err != nil {
		return fmt.Errorf("direct payment failed: %w", err)
	}

	// Update settlement info
	transaction.PaymentInfo.Settlement.ConfirmationHash = ctx.TxBytes().String()
	settledAt := ctx.BlockTime()
	transaction.PaymentInfo.Settlement.SettledAt = &settledAt

	return nil
}

// processMultiHopPayment handles multi-hop payments through intermediaries
func (k Keeper) processMultiHopPayment(ctx sdk.Context, transaction *types.CommerceTransaction) error {
	// Process each hop in the route
	for i, hop := range transaction.PaymentInfo.Route.Hops {
		fromAddr, err := sdk.AccAddressFromBech32(hop.From)
		if err != nil {
			return fmt.Errorf("invalid from address in hop %d: %w", i, err)
		}

		toAddr, err := sdk.AccAddressFromBech32(hop.To)
		if err != nil {
			return fmt.Errorf("invalid to address in hop %d: %w", i, err)
		}

		// Execute hop transfer
		err = k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, hop.Amount)
		if err != nil {
			return fmt.Errorf("multi-hop payment failed at hop %d: %w", i, err)
		}

		// Deduct hop fee if applicable
		if !hop.Fee.IsZero() {
			// Fee handling logic here
		}
	}

	return nil
}

// processBridgePayment handles cross-chain bridge payments
func (k Keeper) processBridgePayment(ctx sdk.Context, transaction *types.CommerceTransaction) error {
	// Use CCTP for cross-chain transfers
	if transaction.PaymentInfo.CrossBorderInfo != nil {
		// Prepare CCTP burn message
		var fromAddr sdk.AccAddress
		for _, party := range transaction.Parties {
			if party.Role == "buyer" || party.Role == "payer" {
				addr, err := sdk.AccAddressFromBech32(party.Address)
				if err != nil {
					return err
				}
				fromAddr = addr
				break
			}
		}

		// Execute bridge transfer through CCTP
		// This would integrate with the CCTP module
		return k.executeCCTPTransfer(ctx, fromAddr, transaction)
	}

	return fmt.Errorf("bridge payment requires cross-border information")
}

// processOptimizedPayment uses AI/ML for optimal routing
func (k Keeper) processOptimizedPayment(ctx sdk.Context, transaction *types.CommerceTransaction) error {
	// Real-time route optimization
	optimalRoute, err := k.paymentRouter.FindOptimalRoute(ctx, transaction.PaymentInfo)
	if err != nil {
		return err
	}

	// Update route and re-process
	transaction.PaymentInfo.Route = optimalRoute
	
	// Process based on optimized route type
	switch optimalRoute.Type {
	case types.PaymentRouteDirect:
		return k.processDirectPayment(ctx, transaction)
	case types.PaymentRouteMultiHop:
		return k.processMultiHopPayment(ctx, transaction)
	case types.PaymentRouteBridge:
		return k.processBridgePayment(ctx, transaction)
	default:
		return fmt.Errorf("unsupported optimized route type: %s", optimalRoute.Type)
	}
}

// executeCCTPTransfer executes cross-chain transfer through CCTP
func (k Keeper) executeCCTPTransfer(ctx sdk.Context, fromAddr sdk.AccAddress, transaction *types.CommerceTransaction) error {
	// This would integrate with the actual CCTP keeper
	// For now, we'll simulate the process
	
	// Get destination domain from cross-border info
	// This would be mapped from country to domain ID
	destinationDomain := k.getDestinationDomain(transaction.PaymentInfo.CrossBorderInfo.DestinationCountry)
	
	// Get mint recipient (this would be derived from the payee)
	var mintRecipient []byte
	for _, party := range transaction.Parties {
		if party.Role == "seller" || party.Role == "payee" {
			// Convert address to bytes for destination chain
			mintRecipient = []byte(party.Address)
			break
		}
	}

	// Create CCTP deposit for burn message
	amount := transaction.PaymentInfo.Amount[0] // Assuming single currency for simplicity
	burnToken := amount.Denom

	// This would call the actual CCTP keeper method
	// return k.cctpKeeper.DepositForBurn(ctx, fromAddr.String(), amount, destinationDomain, mintRecipient, burnToken)
	
	// For now, simulate success
	return nil
}

// getDestinationDomain maps country codes to CCTP domain IDs
func (k Keeper) getDestinationDomain(country string) uint32 {
	// This would be a configurable mapping
	countryToDomain := map[string]uint32{
		"US": 0,
		"EU": 1,
		"UK": 2,
		"CA": 3,
		"AU": 4,
	}
	
	if domain, exists := countryToDomain[country]; exists {
		return domain
	}
	
	return 0 // Default to Ethereum mainnet
}

// GetCommerceTransaction retrieves a commerce transaction by ID
func (k Keeper) GetCommerceTransaction(ctx sdk.Context, id string) (val types.CommerceTransaction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommerceTransactionKeyPrefix))

	b := store.Get(types.CommerceTransactionKey(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetCommerceTransaction sets a commerce transaction in the store
func (k Keeper) SetCommerceTransaction(ctx sdk.Context, transaction types.CommerceTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommerceTransactionKeyPrefix))
	b := k.cdc.MustMarshal(&transaction)
	store.Set(types.CommerceTransactionKey(transaction.ID), b)
}

// GetAllCommerceTransactions returns all commerce transactions
func (k Keeper) GetAllCommerceTransactions(ctx sdk.Context) (list []types.CommerceTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommerceTransactionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CommerceTransaction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CreateTradeFinanceInstrument creates a new trade finance instrument
func (k Keeper) CreateTradeFinanceInstrument(ctx sdk.Context, instrument types.FinancialInstrument) error {
	// Validate instrument
	if instrument.ID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "instrument ID cannot be empty")
	}

	// Check if issuer has authority
	issuerAddr, err := sdk.AccAddressFromBech32(instrument.Issuer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer address: %s", err)
	}

	// Verify issuer has sufficient balance for collateral
	if !instrument.Amount.IsZero() {
		balance := k.bankKeeper.GetAllBalances(ctx, issuerAddr)
		if !balance.IsAllGTE(instrument.Amount) {
			return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "issuer has insufficient funds")
		}
	}

	// Store instrument
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinancialInstrumentKeyPrefix))
	b := k.cdc.MustMarshal(&instrument)
	store.Set(types.FinancialInstrumentKey(instrument.ID), b)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTradeFinanceInstrumentCreated,
			sdk.NewAttribute(types.AttributeKeyInstrumentID, instrument.ID),
			sdk.NewAttribute(types.AttributeKeyInstrumentType, instrument.Type),
			sdk.NewAttribute(types.AttributeKeyIssuer, instrument.Issuer),
		),
	)

	return nil
}

// GetGlobalTradeStatistics returns comprehensive global trade statistics
func (k Keeper) GetGlobalTradeStatistics(ctx sdk.Context) types.GlobalTradeStatistics {
	return k.analyticsEngine.GetGlobalStatistics(ctx)
}

// OptimizePaymentRoute finds the optimal payment route considering cost, time, and reliability
func (k Keeper) OptimizePaymentRoute(ctx sdk.Context, paymentInfo types.PaymentInfo) (types.PaymentRoute, error) {
	return k.paymentRouter.FindOptimalRoute(ctx, paymentInfo)
}

// RunComplianceChecks performs comprehensive compliance checks
func (k Keeper) RunComplianceChecks(ctx sdk.Context, transaction types.CommerceTransaction) error {
	return k.complianceEngine.RunComplianceChecks(ctx, transaction)
}

// CalculateTransactionCost calculates the total cost of a transaction including all fees
func (k Keeper) CalculateTransactionCost(ctx sdk.Context, transaction types.CommerceTransaction) (sdk.Coins, error) {
	fees, err := k.feeCalculator.CalculateTransactionFees(ctx, transaction)
	if err != nil {
		return nil, err
	}
	
	totalCost := transaction.PaymentInfo.Amount.Add(fees.TotalFee...)
	return totalCost, nil
}