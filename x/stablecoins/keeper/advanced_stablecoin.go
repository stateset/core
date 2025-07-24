package keeper

import (
	"context"
	"fmt"
	"math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/stablecoins/types"
)

// AdvancedStablecoinEngine provides sophisticated stablecoin management
type AdvancedStablecoinEngine struct {
	keeper           *Keeper
	priceOracle      *PriceOracle
	stabilityPool    *StabilityPool
	yieldFarming     *YieldFarmingEngine
	algorithmicPeg   *AlgorithmicPeg
	crossChainBridge *CrossChainBridge
	riskEngine       *RiskManagementEngine
}

// PriceOracle handles price feeds and stability mechanisms
type PriceOracle struct {
	priceFeeders     map[string]PriceFeed
	priceHistory     []PricePoint
	volatilityTracker *VolatilityTracker
	lastUpdate       time.Time
}

// PriceFeed represents a price data source
type PriceFeed struct {
	Source      string    `json:"source"`
	Symbol      string    `json:"symbol"`
	Price       sdk.Dec   `json:"price"`
	Confidence  float64   `json:"confidence"`
	LastUpdate  time.Time `json:"last_update"`
	IsActive    bool      `json:"is_active"`
}

// PricePoint represents a historical price point
type PricePoint struct {
	Timestamp time.Time `json:"timestamp"`
	Price     sdk.Dec   `json:"price"`
	Volume    sdk.Int   `json:"volume"`
	Source    string    `json:"source"`
}

// VolatilityTracker monitors price volatility
type VolatilityTracker struct {
	ShortTermVol  float64   `json:"short_term_vol"`
	LongTermVol   float64   `json:"long_term_vol"`
	VaR95         sdk.Dec   `json:"var_95"`
	LastUpdate    time.Time `json:"last_update"`
}

// StabilityPool manages collateral and stability mechanisms
type StabilityPool struct {
	TotalCollateral    sdk.Coins              `json:"total_collateral"`
	CollateralRatio    sdk.Dec                `json:"collateral_ratio"`
	MinCollateralRatio sdk.Dec                `json:"min_collateral_ratio"`
	StabilityFee       sdk.Dec                `json:"stability_fee"`
	LiquidationPenalty sdk.Dec                `json:"liquidation_penalty"`
	ReserveFactors     map[string]sdk.Dec     `json:"reserve_factors"`
	PositionsByUser    map[string][]Position  `json:"positions_by_user"`
}

// Position represents a user's collateralized debt position
type Position struct {
	ID                string    `json:"id"`
	Owner             string    `json:"owner"`
	CollateralType    string    `json:"collateral_type"`
	CollateralAmount  sdk.Int   `json:"collateral_amount"`
	DebtAmount        sdk.Int   `json:"debt_amount"`
	CollateralRatio   sdk.Dec   `json:"collateral_ratio"`
	LastUpdate        time.Time `json:"last_update"`
	LiquidationPrice  sdk.Dec   `json:"liquidation_price"`
	IsLiquidatable    bool      `json:"is_liquidatable"`
}

// YieldFarmingEngine manages yield farming and liquidity incentives
type YieldFarmingEngine struct {
	ActivePools      map[string]*LiquidityPool `json:"active_pools"`
	RewardSchedules  map[string]*RewardSchedule `json:"reward_schedules"`
	TotalRewardsPool sdk.Coins                 `json:"total_rewards_pool"`
	StakingPositions map[string][]StakingPosition `json:"staking_positions"`
}

// LiquidityPool represents a liquidity mining pool
type LiquidityPool struct {
	ID                string    `json:"id"`
	TokenPair         string    `json:"token_pair"`
	TotalLiquidity    sdk.Coins `json:"total_liquidity"`
	APY               sdk.Dec   `json:"apy"`
	TotalStaked       sdk.Int   `json:"total_staked"`
	RewardTokens      []string  `json:"reward_tokens"`
	PoolStartTime     time.Time `json:"pool_start_time"`
	PoolEndTime       time.Time `json:"pool_end_time"`
	IsActive          bool      `json:"is_active"`
}

// RewardSchedule defines how rewards are distributed
type RewardSchedule struct {
	PoolID           string    `json:"pool_id"`
	RewardPerBlock   sdk.Coins `json:"reward_per_block"`
	StartBlock       int64     `json:"start_block"`
	EndBlock         int64     `json:"end_block"`
	TotalRewards     sdk.Coins `json:"total_rewards"`
	DistributedRewards sdk.Coins `json:"distributed_rewards"`
}

// StakingPosition represents a user's staking position
type StakingPosition struct {
	ID               string    `json:"id"`
	User             string    `json:"user"`
	PoolID           string    `json:"pool_id"`
	StakedAmount     sdk.Int   `json:"staked_amount"`
	RewardDebt       sdk.Int   `json:"reward_debt"`
	PendingRewards   sdk.Coins `json:"pending_rewards"`
	LastClaimTime    time.Time `json:"last_claim_time"`
	StakingStartTime time.Time `json:"staking_start_time"`
}

// AlgorithmicPeg manages algorithmic price stability
type AlgorithmicPeg struct {
	TargetPrice       sdk.Dec            `json:"target_price"`
	CurrentPrice      sdk.Dec            `json:"current_price"`
	PriceDeviation    sdk.Dec            `json:"price_deviation"`
	RebaseHistory     []RebaseEvent      `json:"rebase_history"`
	StabilityActions  []StabilityAction  `json:"stability_actions"`
	Controller        *PIDController     `json:"controller"`
}

// RebaseEvent represents a rebase operation
type RebaseEvent struct {
	Timestamp       time.Time `json:"timestamp"`
	PriceBefore     sdk.Dec   `json:"price_before"`
	PriceAfter      sdk.Dec   `json:"price_after"`
	SupplyChange    sdk.Dec   `json:"supply_change"`
	RebaseType      string    `json:"rebase_type"` // "expansion", "contraction"
}

// StabilityAction represents an action taken to maintain stability
type StabilityAction struct {
	Timestamp   time.Time `json:"timestamp"`
	ActionType  string    `json:"action_type"` // "mint", "burn", "adjust_rate"
	Amount      sdk.Int   `json:"amount"`
	Reason      string    `json:"reason"`
	Success     bool      `json:"success"`
}

// PIDController implements a PID controller for algorithmic stability
type PIDController struct {
	Kp           float64   `json:"kp"`           // Proportional gain
	Ki           float64   `json:"ki"`           // Integral gain
	Kd           float64   `json:"kd"`           // Derivative gain
	Integral     float64   `json:"integral"`     // Integral term
	LastError    float64   `json:"last_error"`   // Last error for derivative
	LastUpdate   time.Time `json:"last_update"`
}

// CrossChainBridge handles cross-chain stablecoin operations
type CrossChainBridge struct {
	SupportedChains    map[string]*ChainConfig    `json:"supported_chains"`
	PendingTransfers   map[string]*CrossChainTx   `json:"pending_transfers"`
	BridgeLiquidity    map[string]sdk.Coins       `json:"bridge_liquidity"`
	TransferFees       map[string]sdk.Dec         `json:"transfer_fees"`
	DailyLimits        map[string]sdk.Int         `json:"daily_limits"`
	TotalVolumeToday   map[string]sdk.Int         `json:"total_volume_today"`
}

// ChainConfig represents configuration for a supported chain
type ChainConfig struct {
	ChainID          string    `json:"chain_id"`
	ChainName        string    `json:"chain_name"`
	BridgeContract   string    `json:"bridge_contract"`
	MinConfirmations int       `json:"min_confirmations"`
	IsActive         bool      `json:"is_active"`
	LastUpdate       time.Time `json:"last_update"`
}

// CrossChainTx represents a cross-chain transaction
type CrossChainTx struct {
	ID               string    `json:"id"`
	FromChain        string    `json:"from_chain"`
	ToChain          string    `json:"to_chain"`
	FromAddress      string    `json:"from_address"`
	ToAddress        string    `json:"to_address"`
	Amount           sdk.Int   `json:"amount"`
	Fee              sdk.Int   `json:"fee"`
	Status           string    `json:"status"` // "pending", "confirmed", "failed"
	TxHash           string    `json:"tx_hash"`
	Timestamp        time.Time `json:"timestamp"`
	Confirmations    int       `json:"confirmations"`
}

// RiskManagementEngine handles risk assessment and management
type RiskManagementEngine struct {
	RiskThresholds    map[string]sdk.Dec        `json:"risk_thresholds"`
	SystemRiskScore   float64                   `json:"system_risk_score"`
	UserRiskProfiles  map[string]*RiskProfile   `json:"user_risk_profiles"`
	RiskMetrics       *SystemRiskMetrics        `json:"risk_metrics"`
}

// RiskProfile represents a user's risk assessment
type RiskProfile struct {
	UserAddress      string    `json:"user_address"`
	RiskScore        float64   `json:"risk_score"`
	PositionSize     sdk.Int   `json:"position_size"`
	LeverageRatio    sdk.Dec   `json:"leverage_ratio"`
	VaR              sdk.Dec   `json:"var"`
	LastUpdate       time.Time `json:"last_update"`
	RiskCategory     string    `json:"risk_category"` // "low", "medium", "high"
}

// SystemRiskMetrics tracks overall system risk
type SystemRiskMetrics struct {
	TotalSystemCollateral  sdk.Int   `json:"total_system_collateral"`
	TotalSystemDebt        sdk.Int   `json:"total_system_debt"`
	AverageCollateralRatio sdk.Dec   `json:"average_collateral_ratio"`
	SystemUtilizationRate  sdk.Dec   `json:"system_utilization_rate"`
	ConcentrationRisk      float64   `json:"concentration_risk"`
	LiquidityRisk          float64   `json:"liquidity_risk"`
	LastUpdate             time.Time `json:"last_update"`
}

// NewAdvancedStablecoinEngine creates a new advanced stablecoin engine
func NewAdvancedStablecoinEngine(keeper *Keeper) *AdvancedStablecoinEngine {
	return &AdvancedStablecoinEngine{
		keeper: keeper,
		priceOracle: &PriceOracle{
			priceFeeders: make(map[string]PriceFeed),
			priceHistory: []PricePoint{},
			volatilityTracker: &VolatilityTracker{},
		},
		stabilityPool: &StabilityPool{
			ReserveFactors:  make(map[string]sdk.Dec),
			PositionsByUser: make(map[string][]Position),
		},
		yieldFarming: &YieldFarmingEngine{
			ActivePools:      make(map[string]*LiquidityPool),
			RewardSchedules:  make(map[string]*RewardSchedule),
			StakingPositions: make(map[string][]StakingPosition),
		},
		algorithmicPeg: &AlgorithmicPeg{
			Controller: &PIDController{
				Kp: 0.1,
				Ki: 0.05,
				Kd: 0.02,
			},
		},
		crossChainBridge: &CrossChainBridge{
			SupportedChains:  make(map[string]*ChainConfig),
			PendingTransfers: make(map[string]*CrossChainTx),
			BridgeLiquidity:  make(map[string]sdk.Coins),
			TransferFees:     make(map[string]sdk.Dec),
			DailyLimits:      make(map[string]sdk.Int),
			TotalVolumeToday: make(map[string]sdk.Int),
		},
		riskEngine: &RiskManagementEngine{
			RiskThresholds:   make(map[string]sdk.Dec),
			UserRiskProfiles: make(map[string]*RiskProfile),
			RiskMetrics:      &SystemRiskMetrics{},
		},
	}
}

// UpdatePrice updates the stablecoin price and triggers stability mechanisms
func (ase *AdvancedStablecoinEngine) UpdatePrice(ctx sdk.Context, newPrice sdk.Dec, source string) error {
	// Update price oracle
	ase.priceOracle.priceFeeders[source] = PriceFeed{
		Source:     source,
		Price:      newPrice,
		Confidence: 0.95, // High confidence for this example
		LastUpdate: time.Now(),
		IsActive:   true,
	}

	// Calculate weighted average price
	weightedPrice := ase.calculateWeightedPrice()
	
	// Update algorithmic peg
	ase.algorithmicPeg.CurrentPrice = weightedPrice
	
	// Calculate price deviation
	deviation := weightedPrice.Sub(ase.algorithmicPeg.TargetPrice).Quo(ase.algorithmicPeg.TargetPrice)
	ase.algorithmicPeg.PriceDeviation = deviation

	// Trigger stability actions if needed
	if deviation.Abs().GT(sdk.NewDecWithPrec(5, 2)) { // 5% deviation threshold
		return ase.triggerStabilityAction(ctx, deviation)
	}

	// Update volatility metrics
	ase.updateVolatilityMetrics(weightedPrice)
	
	// Update risk metrics
	ase.updateRiskMetrics(ctx)

	return nil
}

// CreateLiquidityPool creates a new liquidity mining pool
func (ase *AdvancedStablecoinEngine) CreateLiquidityPool(ctx sdk.Context, poolConfig *LiquidityPool) error {
	// Validate pool configuration
	if poolConfig.ID == "" || poolConfig.TokenPair == "" {
		return fmt.Errorf("invalid pool configuration")
	}

	// Check if pool already exists
	if _, exists := ase.yieldFarming.ActivePools[poolConfig.ID]; exists {
		return fmt.Errorf("pool already exists: %s", poolConfig.ID)
	}

	// Initialize pool
	poolConfig.IsActive = true
	poolConfig.PoolStartTime = time.Now()
	poolConfig.TotalStaked = sdk.ZeroInt()
	
	// Add to active pools
	ase.yieldFarming.ActivePools[poolConfig.ID] = poolConfig

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"liquidity_pool_created",
			sdk.NewAttribute("pool_id", poolConfig.ID),
			sdk.NewAttribute("token_pair", poolConfig.TokenPair),
			sdk.NewAttribute("apy", poolConfig.APY.String()),
		),
	)

	return nil
}

// StakeTokens stakes tokens in a liquidity pool
func (ase *AdvancedStablecoinEngine) StakeTokens(ctx sdk.Context, userAddress string, poolID string, amount sdk.Int) error {
	// Get pool
	pool, exists := ase.yieldFarming.ActivePools[poolID]
	if !exists || !pool.IsActive {
		return fmt.Errorf("pool not found or inactive: %s", poolID)
	}

	// Create staking position
	position := StakingPosition{
		ID:               fmt.Sprintf("%s-%s-%d", userAddress, poolID, time.Now().Unix()),
		User:             userAddress,
		PoolID:           poolID,
		StakedAmount:     amount,
		RewardDebt:       sdk.ZeroInt(),
		PendingRewards:   sdk.NewCoins(),
		LastClaimTime:    time.Now(),
		StakingStartTime: time.Now(),
	}

	// Add to user positions
	if ase.yieldFarming.StakingPositions[userAddress] == nil {
		ase.yieldFarming.StakingPositions[userAddress] = []StakingPosition{}
	}
	ase.yieldFarming.StakingPositions[userAddress] = append(
		ase.yieldFarming.StakingPositions[userAddress], 
		position,
	)

	// Update pool totals
	pool.TotalStaked = pool.TotalStaked.Add(amount)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"tokens_staked",
			sdk.NewAttribute("user", userAddress),
			sdk.NewAttribute("pool_id", poolID),
			sdk.NewAttribute("amount", amount.String()),
		),
	)

	return nil
}

// CalculateRewards calculates pending rewards for a user
func (ase *AdvancedStablecoinEngine) CalculateRewards(ctx sdk.Context, userAddress string, poolID string) (sdk.Coins, error) {
	userPositions := ase.yieldFarming.StakingPositions[userAddress]
	if userPositions == nil {
		return sdk.NewCoins(), nil
	}

	totalRewards := sdk.NewCoins()
	
	for _, position := range userPositions {
		if position.PoolID == poolID {
			// Get reward schedule
			schedule, exists := ase.yieldFarming.RewardSchedules[poolID]
			if !exists {
				continue
			}

			// Calculate time-based rewards
			stakingDuration := time.Since(position.StakingStartTime)
			stakingBlocks := int64(stakingDuration.Seconds() / 6) // Assuming 6-second blocks
			
			// Calculate rewards per block
			if stakingBlocks > 0 {
				rewards := schedule.RewardPerBlock.MulInt(sdk.NewInt(stakingBlocks))
				totalRewards = totalRewards.Add(rewards...)
			}
		}
	}

	return totalRewards, nil
}

// ExecuteCrossChainTransfer initiates a cross-chain stablecoin transfer
func (ase *AdvancedStablecoinEngine) ExecuteCrossChainTransfer(ctx sdk.Context, fromAddress, toAddress, toChain string, amount sdk.Int) error {
	// Validate destination chain
	chainConfig, exists := ase.crossChainBridge.SupportedChains[toChain]
	if !exists || !chainConfig.IsActive {
		return fmt.Errorf("unsupported or inactive destination chain: %s", toChain)
	}

	// Check daily limits
	todayVolume := ase.crossChainBridge.TotalVolumeToday[toChain]
	dailyLimit := ase.crossChainBridge.DailyLimits[toChain]
	
	if todayVolume.Add(amount).GT(dailyLimit) {
		return fmt.Errorf("daily transfer limit exceeded for chain %s", toChain)
	}

	// Calculate transfer fee
	fee := amount.MulRaw(int64(ase.crossChainBridge.TransferFees[toChain].MulInt64(10000).TruncateInt64())).QuoRaw(10000)
	
	// Create cross-chain transaction
	crossChainTx := &CrossChainTx{
		ID:               fmt.Sprintf("tx-%s-%d", toChain, time.Now().Unix()),
		FromChain:        "stateset",
		ToChain:          toChain,
		FromAddress:      fromAddress,
		ToAddress:        toAddress,
		Amount:           amount,
		Fee:              fee,
		Status:           "pending",
		Timestamp:        time.Now(),
		Confirmations:    0,
	}

	// Add to pending transfers
	ase.crossChainBridge.PendingTransfers[crossChainTx.ID] = crossChainTx

	// Update daily volume
	ase.crossChainBridge.TotalVolumeToday[toChain] = todayVolume.Add(amount)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"cross_chain_transfer_initiated",
			sdk.NewAttribute("tx_id", crossChainTx.ID),
			sdk.NewAttribute("from_address", fromAddress),
			sdk.NewAttribute("to_address", toAddress),
			sdk.NewAttribute("to_chain", toChain),
			sdk.NewAttribute("amount", amount.String()),
			sdk.NewAttribute("fee", fee.String()),
		),
	)

	return nil
}

// AssessUserRisk evaluates the risk profile of a user
func (ase *AdvancedStablecoinEngine) AssessUserRisk(ctx sdk.Context, userAddress string) (*RiskProfile, error) {
	// Get user's positions
	positions := ase.stabilityPool.PositionsByUser[userAddress]
	if positions == nil {
		return &RiskProfile{
			UserAddress:   userAddress,
			RiskScore:     0.0,
			RiskCategory:  "low",
			LastUpdate:    time.Now(),
		}, nil
	}

	totalPositionSize := sdk.ZeroInt()
	totalCollateral := sdk.ZeroInt()
	totalDebt := sdk.ZeroInt()
	
	for _, position := range positions {
		totalPositionSize = totalPositionSize.Add(position.CollateralAmount)
		totalCollateral = totalCollateral.Add(position.CollateralAmount)
		totalDebt = totalDebt.Add(position.DebtAmount)
	}

	// Calculate leverage ratio
	leverageRatio := sdk.NewDecFromInt(totalDebt).Quo(sdk.NewDecFromInt(totalCollateral))
	
	// Calculate risk score (0-100)
	riskScore := 0.0
	
	// Factor 1: Leverage (40% weight)
	leverageRisk := math.Min(leverageRatio.MustFloat64()*20, 40)
	
	// Factor 2: Position size relative to total system (30% weight)
	systemCollateral := ase.riskEngine.RiskMetrics.TotalSystemCollateral
	positionSizeRisk := 0.0
	if systemCollateral.GT(sdk.ZeroInt()) {
		positionWeight := sdk.NewDecFromInt(totalPositionSize).Quo(sdk.NewDecFromInt(systemCollateral))
		positionSizeRisk = math.Min(positionWeight.MustFloat64()*300, 30) // High concentration = high risk
	}
	
	// Factor 3: Volatility risk (30% weight)
	volatilityRisk := ase.priceOracle.volatilityTracker.ShortTermVol * 30
	
	riskScore = leverageRisk + positionSizeRisk + volatilityRisk
	
	// Determine risk category
	riskCategory := "low"
	if riskScore > 70 {
		riskCategory = "high"
	} else if riskScore > 40 {
		riskCategory = "medium"
	}

	profile := &RiskProfile{
		UserAddress:   userAddress,
		RiskScore:     riskScore,
		PositionSize:  totalPositionSize,
		LeverageRatio: leverageRatio,
		VaR:           sdk.NewDecWithPrec(int64(riskScore*100), 4), // Simplified VaR calculation
		LastUpdate:    time.Now(),
		RiskCategory:  riskCategory,
	}

	// Store risk profile
	ase.riskEngine.UserRiskProfiles[userAddress] = profile

	return profile, nil
}

// Helper functions

func (ase *AdvancedStablecoinEngine) calculateWeightedPrice() sdk.Dec {
	if len(ase.priceOracle.priceFeeders) == 0 {
		return sdk.OneDec() // Default to $1.00
	}

	totalWeight := 0.0
	weightedSum := sdk.ZeroDec()

	for _, feed := range ase.priceOracle.priceFeeders {
		if feed.IsActive && time.Since(feed.LastUpdate) < time.Hour {
			weight := feed.Confidence
			totalWeight += weight
			weightedSum = weightedSum.Add(feed.Price.MulInt64(int64(weight * 1000000)).QuoInt64(1000000))
		}
	}

	if totalWeight == 0 {
		return sdk.OneDec()
	}

	return weightedSum.QuoInt64(int64(totalWeight))
}

func (ase *AdvancedStablecoinEngine) triggerStabilityAction(ctx sdk.Context, deviation sdk.Dec) error {
	// Use PID controller to determine action
	currentError := deviation.MustFloat64()
	
	controller := ase.algorithmicPeg.Controller
	dt := time.Since(controller.LastUpdate).Seconds()
	
	if dt == 0 {
		dt = 1 // Prevent division by zero
	}

	// Calculate PID terms
	proportional := controller.Kp * currentError
	controller.Integral += currentError * dt
	integral := controller.Ki * controller.Integral
	derivative := controller.Kd * (currentError - controller.LastError) / dt
	
	// Calculate control output
	controlOutput := proportional + integral + derivative
	
	// Determine action based on control output
	action := &StabilityAction{
		Timestamp: time.Now(),
		Reason:    fmt.Sprintf("Price deviation: %.4f%%", deviation.MustFloat64()*100),
	}

	if controlOutput > 0.02 { // Expansion needed
		action.ActionType = "mint"
		action.Amount = sdk.NewInt(int64(math.Abs(controlOutput) * 1000000)) // Scale appropriately
	} else if controlOutput < -0.02 { // Contraction needed
		action.ActionType = "burn"
		action.Amount = sdk.NewInt(int64(math.Abs(controlOutput) * 1000000))
	} else {
		action.ActionType = "adjust_rate"
		action.Amount = sdk.ZeroInt()
	}

	// Execute the action (implementation would depend on specific mechanisms)
	action.Success = true // Assume success for this example

	// Update controller state
	controller.LastError = currentError
	controller.LastUpdate = time.Now()

	// Record stability action
	ase.algorithmicPeg.StabilityActions = append(ase.algorithmicPeg.StabilityActions, *action)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"stability_action_executed",
			sdk.NewAttribute("action_type", action.ActionType),
			sdk.NewAttribute("amount", action.Amount.String()),
			sdk.NewAttribute("reason", action.Reason),
		),
	)

	return nil
}

func (ase *AdvancedStablecoinEngine) updateVolatilityMetrics(currentPrice sdk.Dec) {
	// Add price point to history
	pricePoint := PricePoint{
		Timestamp: time.Now(),
		Price:     currentPrice,
		Volume:    sdk.ZeroInt(), // Would be populated with actual volume data
		Source:    "weighted_average",
	}
	
	ase.priceOracle.priceHistory = append(ase.priceOracle.priceHistory, pricePoint)
	
	// Keep only last 1000 points for efficiency
	if len(ase.priceOracle.priceHistory) > 1000 {
		ase.priceOracle.priceHistory = ase.priceOracle.priceHistory[1:]
	}

	// Calculate short-term volatility (last 24 hours)
	shortTermPrices := []float64{}
	cutoff := time.Now().Add(-24 * time.Hour)
	
	for _, point := range ase.priceOracle.priceHistory {
		if point.Timestamp.After(cutoff) {
			shortTermPrices = append(shortTermPrices, point.Price.MustFloat64())
		}
	}
	
	if len(shortTermPrices) > 1 {
		ase.priceOracle.volatilityTracker.ShortTermVol = calculateVolatility(shortTermPrices)
	}

	// Calculate long-term volatility (last 30 days)
	longTermPrices := []float64{}
	longTermCutoff := time.Now().Add(-30 * 24 * time.Hour)
	
	for _, point := range ase.priceOracle.priceHistory {
		if point.Timestamp.After(longTermCutoff) {
			longTermPrices = append(longTermPrices, point.Price.MustFloat64())
		}
	}
	
	if len(longTermPrices) > 1 {
		ase.priceOracle.volatilityTracker.LongTermVol = calculateVolatility(longTermPrices)
	}

	ase.priceOracle.volatilityTracker.LastUpdate = time.Now()
}

func (ase *AdvancedStablecoinEngine) updateRiskMetrics(ctx sdk.Context) {
	totalCollateral := sdk.ZeroInt()
	totalDebt := sdk.ZeroInt()
	positionCount := 0

	// Aggregate across all users
	for _, positions := range ase.stabilityPool.PositionsByUser {
		for _, position := range positions {
			totalCollateral = totalCollateral.Add(position.CollateralAmount)
			totalDebt = totalDebt.Add(position.DebtAmount)
			positionCount++
		}
	}

	metrics := ase.riskEngine.RiskMetrics
	metrics.TotalSystemCollateral = totalCollateral
	metrics.TotalSystemDebt = totalDebt
	
	if totalCollateral.GT(sdk.ZeroInt()) {
		metrics.AverageCollateralRatio = sdk.NewDecFromInt(totalCollateral).Quo(sdk.NewDecFromInt(totalDebt))
		metrics.SystemUtilizationRate = sdk.NewDecFromInt(totalDebt).Quo(sdk.NewDecFromInt(totalCollateral))
	}

	metrics.LastUpdate = time.Now()
}

func calculateVolatility(prices []float64) float64 {
	if len(prices) < 2 {
		return 0.0
	}

	// Calculate returns
	returns := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		returns[i-1] = math.Log(prices[i] / prices[i-1])
	}

	// Calculate standard deviation of returns
	mean := 0.0
	for _, ret := range returns {
		mean += ret
	}
	mean /= float64(len(returns))

	variance := 0.0
	for _, ret := range returns {
		variance += math.Pow(ret-mean, 2)
	}
	variance /= float64(len(returns) - 1)

	// Annualized volatility (assuming daily data)
	return math.Sqrt(variance * 365)
}