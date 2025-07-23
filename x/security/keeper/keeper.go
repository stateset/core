package keeper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/security/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Security Rules Management

func (k Keeper) SetSecurityRule(ctx sdk.Context, rule types.SecurityRule) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityRule/"))
	b := k.cdc.MustMarshal(&rule)
	store.Set([]byte(rule.ID), b)
}

func (k Keeper) GetSecurityRule(ctx sdk.Context, id string) (val types.SecurityRule, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityRule/"))
	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveSecurityRule(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityRule/"))
	store.Delete([]byte(id))
}

func (k Keeper) GetAllSecurityRules(ctx sdk.Context) (list []types.SecurityRule) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityRule/"))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityRule
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Security Alerts Management

func (k Keeper) SetSecurityAlert(ctx sdk.Context, alert types.SecurityAlert) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityAlert/"))
	b := k.cdc.MustMarshal(&alert)
	store.Set([]byte(alert.ID), b)
}

func (k Keeper) GetSecurityAlert(ctx sdk.Context, id string) (val types.SecurityAlert, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityAlert/"))
	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllSecurityAlerts(ctx sdk.Context) (list []types.SecurityAlert) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("SecurityAlert/"))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityAlert
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Risk Profile Management

func (k Keeper) SetRiskProfile(ctx sdk.Context, profile types.RiskProfile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("RiskProfile/"))
	b := k.cdc.MustMarshal(&profile)
	store.Set([]byte(profile.Address), b)
}

func (k Keeper) GetRiskProfile(ctx sdk.Context, address string) (val types.RiskProfile, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("RiskProfile/"))
	b := store.Get([]byte(address))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllRiskProfiles(ctx sdk.Context) (list []types.RiskProfile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("RiskProfile/"))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RiskProfile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Transaction Monitoring

func (k Keeper) SetTransactionMonitor(ctx sdk.Context, monitor types.TransactionMonitor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("TransactionMonitor/"))
	b := k.cdc.MustMarshal(&monitor)
	store.Set([]byte(monitor.TransactionHash), b)
}

func (k Keeper) GetTransactionMonitor(ctx sdk.Context, txHash string) (val types.TransactionMonitor, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("TransactionMonitor/"))
	b := store.Get([]byte(txHash))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Security Operations

// PerformSecurityChecks runs security checks on each block
func (k Keeper) PerformSecurityChecks(ctx sdk.Context) {
	k.Logger(ctx).Info("Performing security checks", "height", ctx.BlockHeight())
	
	// Get all active security rules
	rules := k.GetAllSecurityRules(ctx)
	
	for _, rule := range rules {
		if rule.Enabled {
			k.executeSecurityRule(ctx, rule)
		}
	}
	
	// Update risk profiles based on recent transactions
	k.updateRiskProfiles(ctx)
}

// executeSecurityRule executes a specific security rule
func (k Keeper) executeSecurityRule(ctx sdk.Context, rule types.SecurityRule) {
	k.Logger(ctx).Debug("Executing security rule", "rule", rule.Name)
	
	// Parse conditions and check against current state
	// This is a simplified implementation - in production, you'd have a more
	// sophisticated rule engine
	
	switch rule.Type {
	case types.AlertTypeVelocity:
		k.checkVelocityRules(ctx, rule)
	case types.AlertTypeFraud:
		k.checkFraudRules(ctx, rule)
	case types.AlertTypeCompliance:
		k.checkComplianceRules(ctx, rule)
	}
}

// checkVelocityRules checks for velocity-based anomalies
func (k Keeper) checkVelocityRules(ctx sdk.Context, rule types.SecurityRule) {
	// Implementation for velocity checking
	// This would analyze transaction patterns, frequency, and amounts
}

// checkFraudRules checks for fraud patterns
func (k Keeper) checkFraudRules(ctx sdk.Context, rule types.SecurityRule) {
	// Implementation for fraud detection
	// This would use ML algorithms and pattern recognition
}

// checkComplianceRules checks compliance requirements
func (k Keeper) checkComplianceRules(ctx sdk.Context, rule types.SecurityRule) {
	// Implementation for compliance checking
	// This would verify KYC, AML, and regulatory requirements
}

// updateRiskProfiles updates risk profiles for active addresses
func (k Keeper) updateRiskProfiles(ctx sdk.Context) {
	// Get recent transactions and update risk scores
	// This is where you'd implement risk scoring algorithms
}

// CreateAlert creates a new security alert
func (k Keeper) CreateAlert(ctx sdk.Context, ruleID, alertType, title, description, txID, address string, amount sdk.Coins, severity int32) string {
	alertID := fmt.Sprintf("%d-%s", ctx.BlockHeight(), strconv.FormatInt(time.Now().UnixNano(), 36))
	
	alert := types.SecurityAlert{
		ID:            alertID,
		RuleID:        ruleID,
		Type:          alertType,
		Severity:      severity,
		Title:         title,
		Description:   description,
		TransactionID: txID,
		Address:       address,
		Amount:        amount,
		Status:        "pending",
		CreatedAt:     ctx.BlockTime(),
	}
	
	k.SetSecurityAlert(ctx, alert)
	
	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"security_alert_created",
			sdk.NewAttribute("alert_id", alertID),
			sdk.NewAttribute("type", alertType),
			sdk.NewAttribute("severity", fmt.Sprintf("%d", severity)),
			sdk.NewAttribute("address", address),
		),
	)
	
	return alertID
}

// CalculateRiskScore calculates a risk score for an address
func (k Keeper) CalculateRiskScore(ctx sdk.Context, address string) int32 {
	// Implement risk scoring algorithm
	// This would consider:
	// - Transaction history
	// - Velocity patterns
	// - Network connections
	// - Compliance status
	
	profile, found := k.GetRiskProfile(ctx, address)
	if !found {
		// Create new risk profile for address
		profile = types.RiskProfile{
			Address:         address,
			RiskScore:       50, // Default medium risk
			VelocityScore:   50,
			PatternScore:    50,
			ComplianceScore: 50,
			UpdatedAt:       ctx.BlockTime(),
		}
	}
	
	// Update scores based on recent activity
	// ... risk calculation logic ...
	
	k.SetRiskProfile(ctx, profile)
	return profile.GetOverallRiskScore()
}

// Genesis functions
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Initialize default security rules
	defaultRules := k.getDefaultSecurityRules()
	for _, rule := range defaultRules {
		k.SetSecurityRule(ctx, rule)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	genesis := types.DefaultGenesis()
	// Export current state
	return *genesis
}

func (k Keeper) getDefaultSecurityRules() []types.SecurityRule {
	return []types.SecurityRule{
		{
			ID:       "velocity-001",
			Name:     "High Velocity Transactions",
			Type:     types.AlertTypeVelocity,
			Enabled:  true,
			Severity: 3,
			Conditions: `{"max_transactions_per_hour": 100, "max_amount_per_hour": "1000000"}`,
			Actions:    `{"alert": true, "block": false}`,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:       "fraud-001",
			Name:     "Suspicious Pattern Detection",
			Type:     types.AlertTypeFraud,
			Enabled:  true,
			Severity: 4,
			Conditions: `{"unusual_patterns": true, "risk_threshold": 80}`,
			Actions:    `{"alert": true, "block": true}`,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
}