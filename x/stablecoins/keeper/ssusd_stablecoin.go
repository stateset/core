package keeper

import (
	"context"
	"fmt"
	"math"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	"github.com/stateset/core/x/stablecoins/types"
)

// SSUSDStablecoinEngine provides specialized functionality for the ssUSD stablecoin
type SSUSDStablecoinEngine struct {
	keeper                *Keeper
	pegMaintainer         *SSUSDPegMaintainer
	liquidityManager      *SSUSDLiquidityManager
	collateralManager     *SSUSDCollateralManager
	yieldOptimizer        *SSUSDYieldOptimizer
	riskManager          *SSUSDRiskManager
	crossChainBridge     *SSUSDCrossChainBridge
	rebaseController     *SSUSDRebaseController
}

// SSUSDPegMaintainer maintains the USD peg for ssUSD
type SSUSDPegMaintainer struct {
	targetPrice           sdkmath.LegacyDec                    `json:"target_price"`
	priceToleranceBPS     int64                     `json:"price_tolerance_bps"` // basis points
	rebalanceThresholdBPS int64                     `json:"rebalance_threshold_bps"`
	priceFeeds           map[string]*SSUSDPriceFeed `json:"price_feeds"`
	lastRebalance        time.Time                  `json:"last_rebalance"`
	rebalanceHistory     []SSUSDRebalanceEvent      `json:"rebalance_history"`
	emergencyMode        bool                       `json:"emergency_mode"`
}

// SSUSDPriceFeed represents a price feed source for ssUSD
type SSUSDPriceFeed struct {
	Provider         string    `json:"provider"`         // "chainlink", "coingecko", "binance", etc.
	Asset           string    `json:"asset"`            // "USD", "USDC", "USDT"
	Price           sdkmath.LegacyDec   `json:"price"`
	Weight          sdkmath.LegacyDec   `json:"weight"`           // Weight in price calculation
	LastUpdate      time.Time `json:"last_update"`
	IsActive        bool      `json:"is_active"`
	DeviationLimit  sdkmath.LegacyDec   `json:"deviation_limit"`  // Max allowed deviation
	UpdateFrequency int64     `json:"update_frequency"` // Update frequency in seconds
}

// SSUSDRebalanceEvent tracks rebalancing operations
type SSUSDRebalanceEvent struct {
	Timestamp       time.Time `json:"timestamp"`
	PriceBefore     sdkmath.LegacyDec   `json:"price_before"`
	PriceAfter      sdkmath.LegacyDec   `json:"price_after"`
	Action          string    `json:"action"`          // "mint", "burn", "adjust_rates"
	AmountAdjusted  sdkmath.Int   `json:"amount_adjusted"`
	TriggerReason   string    `json:"trigger_reason"`
	Success         bool      `json:"success"`
	GasUsed         uint64    `json:"gas_used"`
}

// SSUSDLiquidityManager manages liquidity pools and incentives
type SSUSDLiquidityManager struct {
	pools               map[string]*SSUSDLiquidityPool `json:"pools"`
	totalLiquidity      sdk.Coins                      `json:"total_liquidity"`
	rewardsPools        map[string]*SSUSDRewardsPool   `json:"rewards_pools"`
	liquidityProviders  map[string]*SSUSDLPPosition    `json:"liquidity_providers"`
	autoCompounding     bool                           `json:"auto_compounding"`
	feeDistribution     *SSUSDFeeDistribution          `json:"fee_distribution"`
}

// SSUSDLiquidityPool represents a liquidity pool for ssUSD
type SSUSDLiquidityPool struct {
	ID              string     `json:"id"`
	TokenPair       string     `json:"token_pair"`      // e.g., "ssUSD/USDC"
	TotalLiquidity  sdk.Coins  `json:"total_liquidity"`
	PoolShare       sdkmath.LegacyDec    `json:"pool_share"`      // Share of total ssUSD liquidity
	APY             sdkmath.LegacyDec    `json:"apy"`
	TradingFee      sdkmath.LegacyDec    `json:"trading_fee"`     // Trading fee percentage
	RewardMultiplier sdkmath.LegacyDec   `json:"reward_multiplier"`
	IsActive        bool       `json:"is_active"`
	CreatedAt       time.Time  `json:"created_at"`
	LastUpdate      time.Time  `json:"last_update"`
}

// SSUSDLPPosition represents a liquidity provider position
type SSUSDLPPosition struct {
	UserAddress       string    `json:"user_address"`
	PoolID           string    `json:"pool_id"`
	LPTokens         sdkmath.Int   `json:"lp_tokens"`
	ProvidedLiquidity sdk.Coins `json:"provided_liquidity"`
	AccruedRewards   sdk.Coins `json:"accrued_rewards"`
	LastRewardClaim  time.Time `json:"last_reward_claim"`
	EntryPrice       sdkmath.LegacyDec   `json:"entry_price"`
	ImpermanentLoss  sdkmath.LegacyDec   `json:"impermanent_loss"`
}

// SSUSDRewardsPool manages rewards distribution
type SSUSDRewardsPool struct {
	PoolID            string     `json:"pool_id"`
	RewardTokens      []string   `json:"reward_tokens"`
	RewardPerSecond   sdk.Coins  `json:"reward_per_second"`
	TotalDistributed  sdk.Coins  `json:"total_distributed"`
	StartTime         time.Time  `json:"start_time"`
	EndTime           time.Time  `json:"end_time"`
	EmissionSchedule  []sdkmath.LegacyDec  `json:"emission_schedule"`
}

// SSUSDCollateralManager handles collateral backing ssUSD
type SSUSDCollateralManager struct {
	collateralTypes    map[string]*SSUSDCollateralType `json:"collateral_types"`
	totalCollateral    sdk.Coins                       `json:"total_collateral"`
	collateralRatio    sdkmath.LegacyDec                         `json:"collateral_ratio"`
	liquidationEngine  *SSUSDLiquidationEngine         `json:"liquidation_engine"`
	diversificationTarget map[string]sdkmath.LegacyDec           `json:"diversification_target"`
}

// SSUSDCollateralType defines a type of collateral backing ssUSD
type SSUSDCollateralType struct {
	Denom              string    `json:"denom"`
	LTV                sdkmath.LegacyDec   `json:"ltv"`                // Loan-to-value ratio
	LiquidationThreshold sdkmath.LegacyDec `json:"liquidation_threshold"`
	LiquidationPenalty sdkmath.LegacyDec   `json:"liquidation_penalty"`
	MaxAllocation      sdkmath.LegacyDec   `json:"max_allocation"`     // Max % of total collateral
	CurrentAllocation  sdkmath.LegacyDec   `json:"current_allocation"`
	PriceVolatility    sdkmath.LegacyDec   `json:"price_volatility"`
	IsActive          bool      `json:"is_active"`
	RiskWeight        sdkmath.LegacyDec   `json:"risk_weight"`
}

// SSUSDLiquidationEngine handles liquidations
type SSUSDLiquidationEngine struct {
	liquidationQueue    []SSUSDLiquidationPosition `json:"liquidation_queue"`
	liquidationRewards  sdkmath.LegacyDec                    `json:"liquidation_rewards"`
	gracePeriod         time.Duration              `json:"grace_period"`
	auctionDuration     time.Duration              `json:"auction_duration"`
}

// SSUSDLiquidationPosition represents a position being liquidated
type SSUSDLiquidationPosition struct {
	PositionID       string    `json:"position_id"`
	Owner           string    `json:"owner"`
	CollateralAmount sdk.Coins `json:"collateral_amount"`
	DebtAmount      sdkmath.Int   `json:"debt_amount"`
	LiquidationPrice sdkmath.LegacyDec  `json:"liquidation_price"`
	AuctionStartTime time.Time `json:"auction_start_time"`
	Status          string    `json:"status"` // "pending", "active", "completed"
}

// SSUSDYieldOptimizer optimizes yield for ssUSD holders
type SSUSDYieldOptimizer struct {
	strategies         map[string]*SSUSDYieldStrategy `json:"strategies"`
	autoCompoundPools  map[string]bool               `json:"auto_compound_pools"`
	yieldReserveFund   sdk.Coins                     `json:"yield_reserve_fund"`
	distributionRules  *SSUSDYieldDistribution       `json:"distribution_rules"`
}

// SSUSDYieldStrategy defines a yield generation strategy
type SSUSDYieldStrategy struct {
	ID                string    `json:"id"`
	Name             string    `json:"name"`
	TargetAPY        sdkmath.LegacyDec   `json:"target_apy"`
	RiskLevel        string    `json:"risk_level"`  // "low", "medium", "high"
	AllocatedFunds   sdk.Coins `json:"allocated_funds"`
	CurrentAPY       sdkmath.LegacyDec   `json:"current_apy"`
	Strategy         string    `json:"strategy"`    // "lending", "staking", "liquidity_mining"
	IsActive         bool      `json:"is_active"`
	LastOptimization time.Time `json:"last_optimization"`
}

// SSUSDYieldDistribution defines how yields are distributed
type SSUSDYieldDistribution struct {
	HolderRewards    sdkmath.LegacyDec `json:"holder_rewards"`     // % to ssUSD holders
	LPRewards        sdkmath.LegacyDec `json:"lp_rewards"`         // % to liquidity providers
	ProtocolReserve  sdkmath.LegacyDec `json:"protocol_reserve"`   // % to protocol reserve
	BuybackBurn      sdkmath.LegacyDec `json:"buyback_burn"`       // % for buyback and burn
}

// SSUSDRiskManager manages risk for the ssUSD ecosystem
type SSUSDRiskManager struct {
	riskMetrics        *SSUSDRiskMetrics        `json:"risk_metrics"`
	stresstTestResults []SSUSDStressTestResult  `json:"stress_test_results"`
	contingencyPlans   map[string]*SSUSDContingencyPlan `json:"contingency_plans"`
	insuranceFund      sdk.Coins                `json:"insurance_fund"`
}

// SSUSDRiskMetrics tracks various risk metrics
type SSUSDRiskMetrics struct {
	OverallRiskScore      float64   `json:"overall_risk_score"`
	LiquidityRisk         float64   `json:"liquidity_risk"`
	CollateralRisk        float64   `json:"collateral_risk"`
	PegRisk              float64   `json:"peg_risk"`
	ConcentrationRisk     float64   `json:"concentration_risk"`
	LastStressTest       time.Time `json:"last_stress_test"`
	VaR95                sdkmath.LegacyDec   `json:"var_95"`
	ExpectedShortfall    sdkmath.LegacyDec   `json:"expected_shortfall"`
}

// SSUSDStressTestResult represents results of stress testing
type SSUSDStressTestResult struct {
	TestID            string            `json:"test_id"`
	TestDate          time.Time         `json:"test_date"`
	Scenario          string            `json:"scenario"`
	PriceShock        sdkmath.LegacyDec           `json:"price_shock"`
	LiquidityImpact   sdkmath.LegacyDec           `json:"liquidity_impact"`
	CollateralImpact  sdkmath.LegacyDec           `json:"collateral_impact"`
	PegMaintenance    bool              `json:"peg_maintenance"`
	RecoveryTime      time.Duration     `json:"recovery_time"`
	TestResults       map[string]string `json:"test_results"`
}

// SSUSDContingencyPlan defines emergency response plans
type SSUSDContingencyPlan struct {
	PlanID           string            `json:"plan_id"`
	TriggerCondition string            `json:"trigger_condition"`
	Actions          []string          `json:"actions"`
	AutoExecute      bool              `json:"auto_execute"`
	RequiredVotes    int               `json:"required_votes"`
	IsActive         bool              `json:"is_active"`
}

// SSUSDCrossChainBridge handles cross-chain ssUSD operations
type SSUSDCrossChainBridge struct {
	supportedChains    map[string]*SSUSDChainConfig  `json:"supported_chains"`
	bridgeReserves     map[string]sdk.Coins          `json:"bridge_reserves"`
	pendingTransfers   map[string]*SSUSDCrossChainTx `json:"pending_transfers"`
	dailyLimits        map[string]sdkmath.Int            `json:"daily_limits"`
	transferFees       map[string]sdkmath.LegacyDec            `json:"transfer_fees"`
}

// SSUSDChainConfig represents configuration for a supported chain
type SSUSDChainConfig struct {
	ChainID          string    `json:"chain_id"`
	ChainName        string    `json:"chain_name"`
	BridgeContract   string    `json:"bridge_contract"`
	SSUSDContract    string    `json:"ssusd_contract"`
	MinConfirmations int       `json:"min_confirmations"`
	MaxTransferAmount sdkmath.Int  `json:"max_transfer_amount"`
	IsActive         bool      `json:"is_active"`
	LastUpdate       time.Time `json:"last_update"`
}

// SSUSDCrossChainTx represents a cross-chain ssUSD transaction
type SSUSDCrossChainTx struct {
	ID               string    `json:"id"`
	FromChain        string    `json:"from_chain"`
	ToChain          string    `json:"to_chain"`
	FromAddress      string    `json:"from_address"`
	ToAddress        string    `json:"to_address"`
	Amount           sdkmath.Int   `json:"amount"`
	Fee              sdkmath.Int   `json:"fee"`
	Status           string    `json:"status"` // "pending", "confirmed", "failed"
	TxHash           string    `json:"tx_hash"`
	Confirmations    int       `json:"confirmations"`
	InitiatedAt      time.Time `json:"initiated_at"`
	CompletedAt      time.Time `json:"completed_at"`
}

// SSUSDRebaseController handles supply adjustments
type SSUSDRebaseController struct {
	rebaseEnabled     bool                   `json:"rebase_enabled"`
	rebaseFrequency   time.Duration         `json:"rebase_frequency"`
	maxRebaseAmount   sdkmath.LegacyDec               `json:"max_rebase_amount"`
	rebaseThreshold   sdkmath.LegacyDec               `json:"rebase_threshold"`
	lastRebase        time.Time             `json:"last_rebase"`
	rebaseHistory     []SSUSDRebaseEvent    `json:"rebase_history"`
	supplyCap         sdkmath.Int               `json:"supply_cap"`
}

// SSUSDRebaseEvent tracks rebase operations
type SSUSDRebaseEvent struct {
	Timestamp         time.Time `json:"timestamp"`
	SupplyBefore      sdkmath.Int   `json:"supply_before"`
	SupplyAfter       sdkmath.Int   `json:"supply_after"`
	RebasePercentage  sdkmath.LegacyDec   `json:"rebase_percentage"`
	TriggerPrice      sdkmath.LegacyDec   `json:"trigger_price"`
	TriggerReason     string    `json:"trigger_reason"`
	Success           bool      `json:"success"`
}

// SSUSDFeeDistribution manages fee distribution
type SSUSDFeeDistribution struct {
	CollectedFees     sdk.Coins            `json:"collected_fees"`
	DistributionRules map[string]sdkmath.LegacyDec   `json:"distribution_rules"`
	LastDistribution  time.Time            `json:"last_distribution"`
	PendingDistribution sdk.Coins          `json:"pending_distribution"`
}

// SSUSDReserveType represents the types of reserves backing ssUSD
type SSUSDReserveType int

const (
	USCash SSUSDReserveType = iota      // U.S. Dollar Cash (FDIC-insured deposits)
	TreasuryBills                       // Treasury Bills (≤93 days maturity)
	GovernmentMMFs                      // Government-only money market funds
	OvernightRepos                      // Tri-party repo agreements
)

// SSUSDConservativeReserve represents the conservative reserve composition
type SSUSDConservativeReserve struct {
	CashReserves    SSUSDCashReserve    `json:"cash_reserves"`    // 10%
	TreasuryBills   SSUSDTreasuryBills  `json:"treasury_bills"`   // 70%
	GovernmentMMFs  SSUSDGovernmentMMFs `json:"government_mmfs"`  // 15%
	OvernightRepos  SSUSDOvernightRepos `json:"overnight_repos"`  // 5%
	TotalValue      sdkmath.LegacyDec             `json:"total_value"`
	LastUpdate      time.Time           `json:"last_update"`
}

// SSUSDCashReserve represents FDIC-insured deposits
type SSUSDCashReserve struct {
	Amount          sdkmath.LegacyDec   `json:"amount"`
	Allocation      sdkmath.LegacyDec   `json:"allocation"`      // 10%
	BankDeposits    []SSUSDBankDeposit `json:"bank_deposits"`
	FDICInsured     bool      `json:"fdic_insured"`
	RiskLevel       string    `json:"risk_level"`      // "minimal"
}

// SSUSDBankDeposit represents a deposit at a regulated bank
type SSUSDBankDeposit struct {
	BankName        string    `json:"bank_name"`
	RoutingNumber   string    `json:"routing_number"`
	AccountNumber   string    `json:"account_number"`
	Amount          sdkmath.LegacyDec   `json:"amount"`
	InterestRate    sdkmath.LegacyDec   `json:"interest_rate"`
	FDICInsured     bool      `json:"fdic_insured"`
	LastUpdate      time.Time `json:"last_update"`
}

// SSUSDTreasuryBills represents U.S. T-Bills with ≤93 days maturity
type SSUSDTreasuryBills struct {
	Amount          sdkmath.LegacyDec     `json:"amount"`
	Allocation      sdkmath.LegacyDec     `json:"allocation"`    // 70%
	TBills          []SSUSDTBill `json:"t_bills"`
	AverageMaturity int64       `json:"average_maturity"` // days
	RiskLevel       string      `json:"risk_level"`       // "minimal"
}

// SSUSDTBill represents a single Treasury Bill
type SSUSDTBill struct {
	CUSIP           string    `json:"cusip"`
	FaceValue       sdkmath.LegacyDec   `json:"face_value"`
	PurchasePrice   sdkmath.LegacyDec   `json:"purchase_price"`
	MaturityDate    time.Time `json:"maturity_date"`
	YieldRate       sdkmath.LegacyDec   `json:"yield_rate"`
	DaysToMaturity  int64     `json:"days_to_maturity"`
}

// SSUSDGovernmentMMFs represents government-only money market funds
type SSUSDGovernmentMMFs struct {
	Amount          sdkmath.LegacyDec      `json:"amount"`
	Allocation      sdkmath.LegacyDec      `json:"allocation"`    // 15%
	MMFFunds        []SSUSDMMFFund `json:"mmf_funds"`
	WAM             int64        `json:"wam"`           // Weighted Average Maturity
	RiskLevel       string       `json:"risk_level"`    // "minimal"
}

// SSUSDMMFFund represents a government money market fund
type SSUSDMMFFund struct {
	FundName        string    `json:"fund_name"`
	FundSymbol      string    `json:"fund_symbol"`
	SharesHeld      sdkmath.LegacyDec   `json:"shares_held"`
	NAVPerShare     sdkmath.LegacyDec   `json:"nav_per_share"`
	YieldRate       sdkmath.LegacyDec   `json:"yield_rate"`
	GovernmentOnly  bool      `json:"government_only"`
	LastUpdate      time.Time `json:"last_update"`
}

// SSUSDOvernightRepos represents tri-party repo agreements
type SSUSDOvernightRepos struct {
	Amount          sdkmath.LegacyDec     `json:"amount"`
	Allocation      sdkmath.LegacyDec     `json:"allocation"`    // 5%
	RepoAgreements  []SSUSDRepo `json:"repo_agreements"`
	AverageRate     sdkmath.LegacyDec     `json:"average_rate"`
	RiskLevel       string      `json:"risk_level"`    // "low"
}

// SSUSDRepo represents a single repo agreement
type SSUSDRepo struct {
	CounterpartyID  string    `json:"counterparty_id"`
	Principal       sdkmath.LegacyDec   `json:"principal"`
	CollateralType  string    `json:"collateral_type"`
	CollateralValue sdkmath.LegacyDec   `json:"collateral_value"`
	RepoRate        sdkmath.LegacyDec   `json:"repo_rate"`
	MaturityDate    time.Time `json:"maturity_date"`
	TriParty        bool      `json:"tri_party"`
}

// SSUSDIssueRequest represents a request to issue (mint) ssUSD
type SSUSDIssueRequest struct {
	Requester       string    `json:"requester"`
	Amount          sdkmath.Int   `json:"amount"`
	ReservePayment  sdk.Coins `json:"reserve_payment"`  // Payment in reserve assets
	RequestTime     time.Time `json:"request_time"`
}

// SSUSDRedeemRequest represents a request to redeem (burn) ssUSD
type SSUSDRedeemRequest struct {
	Requester       string    `json:"requester"`
	SSUSDAmount     sdkmath.Int   `json:"ssusd_amount"`
	PreferredAsset  string    `json:"preferred_asset"`  // Preferred reserve asset for redemption
	RequestTime     time.Time `json:"request_time"`
}

// NewSSUSDStablecoinEngine creates a new ssUSD stablecoin engine
func NewSSUSDStablecoinEngine(keeper *Keeper) *SSUSDStablecoinEngine {
	return &SSUSDStablecoinEngine{
		keeper: keeper,
		pegMaintainer: &SSUSDPegMaintainer{
			targetPrice:           sdkmath.LegacyOneDec(), // $1.00
			priceToleranceBPS:     50,           // 0.5%
			rebalanceThresholdBPS: 100,          // 1%
			priceFeeds:           make(map[string]*SSUSDPriceFeed),
			emergencyMode:        false,
		},
		liquidityManager: &SSUSDLiquidityManager{
			pools:              make(map[string]*SSUSDLiquidityPool),
			rewardsPools:       make(map[string]*SSUSDRewardsPool),
			liquidityProviders: make(map[string]*SSUSDLPPosition),
			autoCompounding:    true,
		},
		collateralManager: &SSUSDCollateralManager{
			collateralTypes:    make(map[string]*SSUSDCollateralType),
			diversificationTarget: make(map[string]sdkmath.LegacyDec),
		},
		yieldOptimizer: &SSUSDYieldOptimizer{
			strategies:        make(map[string]*SSUSDYieldStrategy),
			autoCompoundPools: make(map[string]bool),
		},
		riskManager: &SSUSDRiskManager{
			riskMetrics:      &SSUSDRiskMetrics{},
			contingencyPlans: make(map[string]*SSUSDContingencyPlan),
		},
		crossChainBridge: &SSUSDCrossChainBridge{
			supportedChains:  make(map[string]*SSUSDChainConfig),
			bridgeReserves:   make(map[string]sdk.Coins),
			pendingTransfers: make(map[string]*SSUSDCrossChainTx),
			dailyLimits:      make(map[string]sdkmath.Int),
			transferFees:     make(map[string]sdkmath.LegacyDec),
		},
		rebaseController: &SSUSDRebaseController{
			rebaseEnabled:    true,
			rebaseFrequency:  24 * time.Hour,
			maxRebaseAmount:  sdkmath.LegacyNewDecWithPrec(5, 2), // 5%
			rebaseThreshold:  sdkmath.LegacyNewDecWithPrec(1, 2), // 1%
		},
	}
}

// InitializeSSUSD initializes the ssUSD stablecoin with enhanced features
func (engine *SSUSDStablecoinEngine) InitializeSSUSD(ctx sdk.Context) error {
	// Create the ssUSD stablecoin if it doesn't exist
	_, found := engine.keeper.GetStablecoin(ctx, "ssusd")
	if !found {
		// Define ssUSD with enhanced configuration
		pegInfo := &types.PegInfo{
			TargetAsset:          "USD",
			TargetPrice:          sdkmath.LegacyOneDec(),
			PriceTolerance:       sdkmath.LegacyNewDecWithPrec(5, 3), // 0.5%
			OracleSources:        []string{"chainlink", "band", "internal"},
			RebalancingFrequency: "daily",
		}

		reserveInfo := &types.ReserveInfo{
			ReserveAssets: []*types.ReserveAsset{
				{
					Denom:  "us_cash_token",        // U.S. Dollar Cash (FDIC-insured deposits)
					Amount: sdkmath.ZeroInt(),
					Weight: sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
					Price:  sdkmath.LegacyOneDec(),
				},
				{
					Denom:  "treasury_bill_token",  // Treasury Bills (≤93 days maturity)
					Amount: sdkmath.ZeroInt(),
					Weight: sdkmath.LegacyNewDecWithPrec(70, 2), // 70%
					Price:  sdkmath.LegacyOneDec(),
				},
				{
					Denom:  "mmf_token",           // Government-only money market funds
					Amount: sdkmath.ZeroInt(),
					Weight: sdkmath.LegacyNewDecWithPrec(15, 2), // 15%
					Price:  sdkmath.LegacyOneDec(),
				},
				{
					Denom:  "repo_token",          // Tri-party repo agreements
					Amount: sdkmath.ZeroInt(),
					Weight: sdkmath.LegacyNewDecWithPrec(5, 2), // 5%
					Price:  sdkmath.LegacyOneDec(),
				},
			},
			TotalReserveValue: sdkmath.LegacyZeroDec(),
			ReserveRatio:      sdkmath.LegacyNewDecWithPrec(100, 2), // 100% for 1:1 backing
			MinReserveRatio:   sdkmath.LegacyNewDecWithPrec(100, 2), // 100% minimum for stable backing
		}

		feeInfo := &types.FeeInfo{
			MintFee:        sdkmath.LegacyNewDecWithPrec(1, 3),  // 0.1%
			BurnFee:        sdkmath.LegacyNewDecWithPrec(1, 3),  // 0.1%
			TransferFee:    sdkmath.LegacyNewDecWithPrec(5, 4),  // 0.05%
			RedemptionFee:  sdkmath.LegacyNewDecWithPrec(2, 3),  // 0.2%
			FeeRecipient:   "", // Module account
		}

		accessControl := &types.AccessControlInfo{
			WhitelistEnabled: false,
			BlacklistEnabled: true,
			Whitelist:       []string{},
			Blacklist:       []string{},
			KycRequirement:  "none",
		}

		// Create ssUSD with advanced features
		stablecoin := types.NewStablecoin(
			"ssusd",
			"Stateset USD",
			"ssUSD",
			6,
			"Stateset USD stablecoin backed by diversified reserves with yield generation",
			"", // issuer will be set to module account
			"", // admin will be set to governance
			sdk.NewInt(1000000000000000), // 1B ssUSD max supply
			pegInfo,
			reserveInfo,
			"hybrid", // algorithmic + collateralized
			feeInfo,
			accessControl,
			`{
				"features": ["yield_bearing", "cross_chain", "algorithmic_peg", "liquidity_mining"],
				"version": "2.0",
				"launch_date": "2024-01-01",
				"audits": ["certik", "hacken"],
				"insurance": true
			}`,
		)

		// Set module account as issuer
		stablecoin.Issuer = engine.keeper.accountKeeper.GetModuleAddress("stablecoins").String()
		stablecoin.Admin = "stateset1gov"  // Governance address
		stablecoin.Active = true

		// Save the stablecoin
		engine.keeper.SetStablecoin(ctx, stablecoin)

		// Initialize price feeds
		engine.initializePriceFeeds(ctx)

		// Initialize collateral types
		engine.initializeCollateralTypes(ctx)

		// Initialize liquidity pools
		engine.initializeLiquidityPools(ctx)

		// Initialize yield strategies
		engine.initializeYieldStrategies(ctx)

		// Emit creation event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"ssusd_initialized",
				sdk.NewAttribute("denom", "ssusd"),
				sdk.NewAttribute("max_supply", stablecoin.MaxSupply.String()),
				sdk.NewAttribute("stability_mechanism", "hybrid"),
				sdk.NewAttribute("features", "yield_bearing,cross_chain,algorithmic_peg"),
			),
		)
	}

	return nil
}

// initializePriceFeeds sets up price feeds for ssUSD
func (engine *SSUSDStablecoinEngine) initializePriceFeeds(ctx sdk.Context) {
	priceFeeds := map[string]*SSUSDPriceFeed{
		"chainlink_usd": {
			Provider:        "chainlink",
			Asset:          "USD",
			Price:          sdkmath.LegacyOneDec(),
			Weight:         sdkmath.LegacyNewDecWithPrec(40, 2), // 40%
			IsActive:       true,
			DeviationLimit: sdkmath.LegacyNewDecWithPrec(2, 2), // 2%
			UpdateFrequency: 300, // 5 minutes
		},
		"band_usd": {
			Provider:        "band",
			Asset:          "USD",
			Price:          sdkmath.LegacyOneDec(),
			Weight:         sdkmath.LegacyNewDecWithPrec(30, 2), // 30%
			IsActive:       true,
			DeviationLimit: sdkmath.LegacyNewDecWithPrec(2, 2), // 2%
			UpdateFrequency: 300, // 5 minutes
		},
		"internal_twap": {
			Provider:        "internal",
			Asset:          "USD",
			Price:          sdkmath.LegacyOneDec(),
			Weight:         sdkmath.LegacyNewDecWithPrec(30, 2), // 30%
			IsActive:       true,
			DeviationLimit: sdkmath.LegacyNewDecWithPrec(3, 2), // 3%
			UpdateFrequency: 60, // 1 minute
		},
	}

	engine.pegMaintainer.priceFeeds = priceFeeds
}

// initializeCollateralTypes sets up supported collateral types
func (engine *SSUSDStablecoinEngine) initializeCollateralTypes(ctx sdk.Context) {
	collateralTypes := map[string]*SSUSDCollateralType{
		"uusdc": {
			Denom:               "uusdc",
			LTV:                 sdkmath.LegacyNewDecWithPrec(90, 2), // 90%
			LiquidationThreshold: sdkmath.LegacyNewDecWithPrec(95, 2), // 95%
			LiquidationPenalty:  sdkmath.LegacyNewDecWithPrec(5, 2),  // 5%
			MaxAllocation:       sdkmath.LegacyNewDecWithPrec(40, 2), // 40%
			PriceVolatility:     sdkmath.LegacyNewDecWithPrec(1, 2),  // 1%
			IsActive:           true,
			RiskWeight:         sdkmath.LegacyNewDecWithPrec(10, 2),  // 10%
		},
		"uusdt": {
			Denom:               "uusdt",
			LTV:                 sdkmath.LegacyNewDecWithPrec(90, 2), // 90%
			LiquidationThreshold: sdkmath.LegacyNewDecWithPrec(95, 2), // 95%
			LiquidationPenalty:  sdkmath.LegacyNewDecWithPrec(5, 2),  // 5%
			MaxAllocation:       sdkmath.LegacyNewDecWithPrec(30, 2), // 30%
			PriceVolatility:     sdkmath.LegacyNewDecWithPrec(1, 2),  // 1%
			IsActive:           true,
			RiskWeight:         sdkmath.LegacyNewDecWithPrec(10, 2),  // 10%
		},
		"uatom": {
			Denom:               "uatom",
			LTV:                 sdkmath.LegacyNewDecWithPrec(70, 2), // 70%
			LiquidationThreshold: sdkmath.LegacyNewDecWithPrec(80, 2), // 80%
			LiquidationPenalty:  sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
			MaxAllocation:       sdkmath.LegacyNewDecWithPrec(20, 2), // 20%
			PriceVolatility:     sdkmath.LegacyNewDecWithPrec(20, 2), // 20%
			IsActive:           true,
			RiskWeight:         sdkmath.LegacyNewDecWithPrec(30, 2),  // 30%
		},
		"ustake": {
			Denom:               "ustake",
			LTV:                 sdkmath.LegacyNewDecWithPrec(60, 2), // 60%
			LiquidationThreshold: sdkmath.LegacyNewDecWithPrec(75, 2), // 75%
			LiquidationPenalty:  sdkmath.LegacyNewDecWithPrec(15, 2), // 15%
			MaxAllocation:       sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
			PriceVolatility:     sdkmath.LegacyNewDecWithPrec(30, 2), // 30%
			IsActive:           true,
			RiskWeight:         sdkmath.LegacyNewDecWithPrec(40, 2),  // 40%
		},
	}

	engine.collateralManager.collateralTypes = collateralTypes
	
	// Set diversification targets
	engine.collateralManager.diversificationTarget = map[string]sdkmath.LegacyDec{
		"uusdc":  sdkmath.LegacyNewDecWithPrec(40, 2), // 40%
		"uusdt":  sdkmath.LegacyNewDecWithPrec(30, 2), // 30%
		"uatom":  sdkmath.LegacyNewDecWithPrec(20, 2), // 20%
		"ustake": sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
	}
}

// initializeLiquidityPools sets up initial liquidity pools
func (engine *SSUSDStablecoinEngine) initializeLiquidityPools(ctx sdk.Context) {
	pools := map[string]*SSUSDLiquidityPool{
		"ssusd_usdc": {
			ID:               "ssusd_usdc",
			TokenPair:        "ssUSD/USDC",
			TotalLiquidity:   sdk.NewCoins(),
			PoolShare:        sdkmath.LegacyNewDecWithPrec(40, 2), // 40%
			APY:              sdkmath.LegacyNewDecWithPrec(8, 2),  // 8%
			TradingFee:       sdkmath.LegacyNewDecWithPrec(3, 3),  // 0.3%
			RewardMultiplier: sdkmath.LegacyNewDecWithPrec(120, 2), // 1.2x
			IsActive:         true,
			CreatedAt:        time.Now(),
		},
		"ssusd_usdt": {
			ID:               "ssusd_usdt",
			TokenPair:        "ssUSD/USDT",
			TotalLiquidity:   sdk.NewCoins(),
			PoolShare:        sdkmath.LegacyNewDecWithPrec(30, 2), // 30%
			APY:              sdkmath.LegacyNewDecWithPrec(7, 2),  // 7%
			TradingFee:       sdkmath.LegacyNewDecWithPrec(3, 3),  // 0.3%
			RewardMultiplier: sdkmath.LegacyNewDecWithPrec(110, 2), // 1.1x
			IsActive:         true,
			CreatedAt:        time.Now(),
		},
		"ssusd_atom": {
			ID:               "ssusd_atom",
			TokenPair:        "ssUSD/ATOM",
			TotalLiquidity:   sdk.NewCoins(),
			PoolShare:        sdkmath.LegacyNewDecWithPrec(20, 2), // 20%
			APY:              sdkmath.LegacyNewDecWithPrec(12, 2), // 12%
			TradingFee:       sdkmath.LegacyNewDecWithPrec(5, 3),  // 0.5%
			RewardMultiplier: sdkmath.LegacyNewDecWithPrec(150, 2), // 1.5x
			IsActive:         true,
			CreatedAt:        time.Now(),
		},
		"ssusd_stake": {
			ID:               "ssusd_stake",
			TokenPair:        "ssUSD/STAKE",
			TotalLiquidity:   sdk.NewCoins(),
			PoolShare:        sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
			APY:              sdkmath.LegacyNewDecWithPrec(15, 2), // 15%
			TradingFee:       sdkmath.LegacyNewDecWithPrec(5, 3),  // 0.5%
			RewardMultiplier: sdkmath.LegacyNewDecWithPrec(180, 2), // 1.8x
			IsActive:         true,
			CreatedAt:        time.Now(),
		},
	}

	engine.liquidityManager.pools = pools
}

// initializeYieldStrategies sets up yield generation strategies
func (engine *SSUSDStablecoinEngine) initializeYieldStrategies(ctx sdk.Context) {
	strategies := map[string]*SSUSDYieldStrategy{
		"stable_lending": {
			ID:               "stable_lending",
			Name:            "Stable Asset Lending",
			TargetAPY:       sdkmath.LegacyNewDecWithPrec(6, 2), // 6%
			RiskLevel:       "low",
			AllocatedFunds:  sdk.NewCoins(),
			Strategy:        "lending",
			IsActive:        true,
			LastOptimization: time.Now(),
		},
		"liquidity_mining": {
			ID:               "liquidity_mining",
			Name:            "Liquidity Mining",
			TargetAPY:       sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
			RiskLevel:       "medium",
			AllocatedFunds:  sdk.NewCoins(),
			Strategy:        "liquidity_mining",
			IsActive:        true,
			LastOptimization: time.Now(),
		},
		"cross_chain_yield": {
			ID:               "cross_chain_yield",
			Name:            "Cross-Chain Yield Farming",
			TargetAPY:       sdkmath.LegacyNewDecWithPrec(15, 2), // 15%
			RiskLevel:       "high",
			AllocatedFunds:  sdk.NewCoins(),
			Strategy:        "staking",
			IsActive:        true,
			LastOptimization: time.Now(),
		},
	}

	engine.yieldOptimizer.strategies = strategies

	// Set yield distribution rules
	engine.yieldOptimizer.distributionRules = &SSUSDYieldDistribution{
		HolderRewards:   sdkmath.LegacyNewDecWithPrec(60, 2), // 60%
		LPRewards:       sdkmath.LegacyNewDecWithPrec(25, 2), // 25%
		ProtocolReserve: sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
		BuybackBurn:     sdkmath.LegacyNewDecWithPrec(5, 2),  // 5%
	}
}

// GetSSUSDPrice returns the current price of ssUSD
func (engine *SSUSDStablecoinEngine) GetSSUSDPrice(ctx sdk.Context) sdkmath.LegacyDec {
	if engine.pegMaintainer.emergencyMode {
		return engine.pegMaintainer.targetPrice
	}

	totalWeight := sdkmath.LegacyZeroDec()
	weightedPrice := sdkmath.LegacyZeroDec()

	for _, feed := range engine.pegMaintainer.priceFeeds {
		if feed.IsActive && time.Since(feed.LastUpdate) < time.Duration(feed.UpdateFrequency)*time.Second {
			totalWeight = totalWeight.Add(feed.Weight)
			weightedPrice = weightedPrice.Add(feed.Price.Mul(feed.Weight))
		}
	}

	if totalWeight.IsZero() {
		return engine.pegMaintainer.targetPrice
	}

	return weightedPrice.Quo(totalWeight)
}

// UpdateSSUSDPrice updates the price feeds for ssUSD
func (engine *SSUSDStablecoinEngine) UpdateSSUSDPrice(ctx sdk.Context, provider string, price sdkmath.LegacyDec) error {
	feed, exists := engine.pegMaintainer.priceFeeds[provider]
	if !exists {
		return errorsmod.Wrapf(types.ErrInvalidPriceFeed, "price feed not found: %s", provider)
	}

	// Check for price deviation
	currentPrice := engine.GetSSUSDPrice(ctx)
	deviation := price.Sub(currentPrice).Quo(currentPrice).Abs()
	
	if deviation.GT(feed.DeviationLimit) {
		return errorsmod.Wrapf(types.ErrPriceDeviationTooHigh, 
			"price deviation %.4f%% exceeds limit %.4f%%", 
			deviation.MulInt64(100).MustFloat64(), 
			feed.DeviationLimit.MulInt64(100).MustFloat64())
	}

	// Update the price feed
	feed.Price = price
	feed.LastUpdate = time.Now()

	// Check if rebalancing is needed
	newPrice := engine.GetSSUSDPrice(ctx)
	priceDeviation := newPrice.Sub(engine.pegMaintainer.targetPrice).Quo(engine.pegMaintainer.targetPrice).Abs()
	
	if priceDeviation.GT(sdkmath.LegacyNewDecWithPrec(engine.pegMaintainer.rebalanceThresholdBPS, 4)) {
		return engine.triggerRebalance(ctx, newPrice, priceDeviation)
	}

	// Emit price update event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ssusd_price_updated",
			sdk.NewAttribute("provider", provider),
			sdk.NewAttribute("price", price.String()),
			sdk.NewAttribute("weighted_price", newPrice.String()),
		),
	)

	return nil
}

// triggerRebalance triggers a rebalancing operation
func (engine *SSUSDStablecoinEngine) triggerRebalance(ctx sdk.Context, currentPrice, deviation sdkmath.LegacyDec) error {
	// Prevent frequent rebalancing
	if time.Since(engine.pegMaintainer.lastRebalance) < time.Hour {
		return nil
	}

	action := "none"
	var amountAdjusted sdkmath.Int

	// Determine rebalancing action based on price deviation
	if currentPrice.GT(engine.pegMaintainer.targetPrice) {
		// Price too high, increase supply
		action = "mint"
		// Calculate mint amount based on deviation (simplified)
		supplyIncrease := deviation.Mul(sdkmath.LegacyNewDec(1000000)) // Scale factor
		amountAdjusted = sdkmath.NewIntFromBigInt(supplyIncrease.BigInt())
	} else {
		// Price too low, decrease supply
		action = "burn"
		// Calculate burn amount based on deviation
		supplyDecrease := deviation.Mul(sdkmath.LegacyNewDec(1000000)) // Scale factor
		amountAdjusted = sdkmath.NewIntFromBigInt(supplyDecrease.BigInt())
	}

	// Record rebalance event
	rebalanceEvent := SSUSDRebalanceEvent{
		Timestamp:      time.Now(),
		PriceBefore:    currentPrice,
		PriceAfter:     engine.pegMaintainer.targetPrice, // Target
		Action:         action,
		AmountAdjusted: amountAdjusted,
		TriggerReason:  fmt.Sprintf("Price deviation: %.4f%%", deviation.MulInt64(100).MustFloat64()),
		Success:        true,
	}

	engine.pegMaintainer.rebalanceHistory = append(engine.pegMaintainer.rebalanceHistory, rebalanceEvent)
	engine.pegMaintainer.lastRebalance = time.Now()

	// Emit rebalance event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ssusd_rebalanced",
			sdk.NewAttribute("action", action),
			sdk.NewAttribute("amount", amountAdjusted.String()),
			sdk.NewAttribute("price_before", currentPrice.String()),
			sdk.NewAttribute("deviation", deviation.String()),
		),
	)

	return nil
}

// OptimizeYield optimizes yield generation for ssUSD
func (engine *SSUSDStablecoinEngine) OptimizeYield(ctx sdk.Context) error {
	totalYield := sdkmath.LegacyZeroDec()
	
	for strategyID, strategy := range engine.yieldOptimizer.strategies {
		if !strategy.IsActive {
			continue
		}

		// Calculate current yield for strategy
		// This would involve actual yield calculation logic
		currentYield := strategy.CurrentAPY
		totalYield = totalYield.Add(currentYield)

		// Update last optimization time
		strategy.LastOptimization = time.Now()
		engine.yieldOptimizer.strategies[strategyID] = strategy
	}

	// Distribute yields according to rules
	return engine.distributeYields(ctx, totalYield)
}

// distributeYields distributes generated yields
func (engine *SSUSDStablecoinEngine) distributeYields(ctx sdk.Context, totalYield sdkmath.LegacyDec) error {
	rules := engine.yieldOptimizer.distributionRules
	
	// Calculate distribution amounts
	holderAmount := totalYield.Mul(rules.HolderRewards)
	lpAmount := totalYield.Mul(rules.LPRewards)
	reserveAmount := totalYield.Mul(rules.ProtocolReserve)
	buybackAmount := totalYield.Mul(rules.BuybackBurn)

	// Emit yield distribution event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ssusd_yield_distributed",
			sdk.NewAttribute("total_yield", totalYield.String()),
			sdk.NewAttribute("holder_rewards", holderAmount.String()),
			sdk.NewAttribute("lp_rewards", lpAmount.String()),
			sdk.NewAttribute("protocol_reserve", reserveAmount.String()),
			sdk.NewAttribute("buyback_burn", buybackAmount.String()),
		),
	)

	return nil
}

// GetSSUSDMetrics returns comprehensive metrics for ssUSD
func (engine *SSUSDStablecoinEngine) GetSSUSDMetrics(ctx sdk.Context) (*SSUSDMetrics, error) {
	stablecoin, found := engine.keeper.GetStablecoin(ctx, "ssusd")
	if !found {
		return nil, sdkerrors.Wrap(types.ErrStablecoinNotFound, "ssusd")
	}

	currentPrice := engine.GetSSUSDPrice(ctx)
	
	return &SSUSDMetrics{
		CurrentPrice:      currentPrice,
		TargetPrice:       engine.pegMaintainer.targetPrice,
		TotalSupply:       stablecoin.TotalSupply,
		CollateralRatio:   engine.collateralManager.collateralRatio,
		TotalLiquidity:    engine.liquidityManager.totalLiquidity,
		AverageAPY:        engine.calculateAverageAPY(),
		RiskScore:         engine.riskManager.riskMetrics.OverallRiskScore,
		ActivePools:       int64(len(engine.liquidityManager.pools)),
		CrossChainSupport: int64(len(engine.crossChainBridge.supportedChains)),
		LastRebalance:     engine.pegMaintainer.lastRebalance,
	}, nil
}

// SSUSDMetrics represents comprehensive metrics for ssUSD
type SSUSDMetrics struct {
	CurrentPrice      sdkmath.LegacyDec     `json:"current_price"`
	TargetPrice       sdkmath.LegacyDec     `json:"target_price"`
	TotalSupply       sdkmath.Int     `json:"total_supply"`
	CollateralRatio   sdkmath.LegacyDec     `json:"collateral_ratio"`
	TotalLiquidity    sdk.Coins   `json:"total_liquidity"`
	AverageAPY        sdkmath.LegacyDec     `json:"average_apy"`
	RiskScore         float64     `json:"risk_score"`
	ActivePools       int64       `json:"active_pools"`
	CrossChainSupport int64       `json:"cross_chain_support"`
	LastRebalance     time.Time   `json:"last_rebalance"`
}

// calculateAverageAPY calculates the weighted average APY
func (engine *SSUSDStablecoinEngine) calculateAverageAPY() sdkmath.LegacyDec {
	totalWeight := sdkmath.LegacyZeroDec()
	weightedAPY := sdkmath.LegacyZeroDec()

	for _, pool := range engine.liquidityManager.pools {
		if pool.IsActive {
			weight := pool.PoolShare
			totalWeight = totalWeight.Add(weight)
			weightedAPY = weightedAPY.Add(pool.APY.Mul(weight))
		}
	}

	if totalWeight.IsZero() {
		return sdkmath.LegacyZeroDec()
	}

	return weightedAPY.Quo(totalWeight)
}

// IssueSSUSD issues new ssUSD tokens backed 1:1 by reserves
func (engine *SSUSDStablecoinEngine) IssueSSUSD(ctx sdk.Context, request SSUSDIssueRequest) error {
	// Validate the stablecoin exists and is active
	stablecoin, found := engine.keeper.GetStablecoin(ctx, "ssusd")
	if !found {
		return sdkerrors.Wrap(types.ErrStablecoinNotFound, "ssusd")
	}

	if !stablecoin.Active {
		return sdkerrors.Wrap(types.ErrStablecoinInactive, "ssusd")
	}

	// Calculate USD value of reserve payment
	reserveValue, err := engine.calculateReserveValue(ctx, request.ReservePayment)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to calculate reserve value")
	}

	// Ensure 1:1 backing - reserve value must equal or exceed ssUSD amount
	ssUSDValue := sdkmath.LegacyNewDecFromInt(request.Amount).QuoInt64(1000000) // Convert from micro units
	if reserveValue.LT(ssUSDValue) {
		return errorsmod.Wrapf(types.ErrInsufficientCollateral, 
			"reserve value %s insufficient for ssUSD amount %s", 
			reserveValue.String(), ssUSDValue.String())
	}

	// Validate reserve composition against target allocations
	err = engine.validateReserveComposition(ctx, request.ReservePayment)
	if err != nil {
		return sdkerrors.Wrap(err, "reserve composition validation failed")
	}

	// Transfer reserve assets to module account
	requesterAddr, err := sdk.AccAddressFromBech32(request.Requester)
	if err != nil {
		return err
	}

	err = engine.keeper.bankKeeper.SendCoinsFromAccountToModule(
		ctx, requesterAddr, types.ModuleName, request.ReservePayment)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to transfer reserve assets")
	}

	// Update reserve composition
	err = engine.updateReserveComposition(ctx, request.ReservePayment, true)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to update reserve composition")
	}

	// Mint ssUSD tokens
	mintCoins := sdk.NewCoins(sdk.NewCoin("ssusd", request.Amount))
	err = engine.keeper.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to mint ssUSD")
	}

	// Send ssUSD to requester
	err = engine.keeper.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, requesterAddr, mintCoins)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to send ssUSD to requester")
	}

	// Update stablecoin total supply
	stablecoin.TotalSupply = stablecoin.TotalSupply.Add(request.Amount)
	engine.keeper.SetStablecoin(ctx, stablecoin)

	// Update reserve info
	stablecoin.ReserveInfo.TotalReserveValue = stablecoin.ReserveInfo.TotalReserveValue.Add(reserveValue)
	engine.keeper.SetStablecoin(ctx, stablecoin)

	// Emit issue event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ssusd_issued",
			sdk.NewAttribute("requester", request.Requester),
			sdk.NewAttribute("ssusd_amount", request.Amount.String()),
			sdk.NewAttribute("reserve_value", reserveValue.String()),
			sdk.NewAttribute("reserve_payment", request.ReservePayment.String()),
			sdk.NewAttribute("new_total_supply", stablecoin.TotalSupply.String()),
		),
	)

	return nil
}

// RedeemSSUSD redeems ssUSD tokens for underlying reserves
func (engine *SSUSDStablecoinEngine) RedeemSSUSD(ctx sdk.Context, request SSUSDRedeemRequest) error {
	// Validate the stablecoin exists and is active
	stablecoin, found := engine.keeper.GetStablecoin(ctx, "ssusd")
	if !found {
		return sdkerrors.Wrap(types.ErrStablecoinNotFound, "ssusd")
	}

	if !stablecoin.Active {
		return sdkerrors.Wrap(types.ErrStablecoinInactive, "ssusd")
	}

	// Validate requester has sufficient ssUSD balance
	requesterAddr, err := sdk.AccAddressFromBech32(request.Requester)
	if err != nil {
		return err
	}

	balance := engine.keeper.bankKeeper.GetBalance(ctx, requesterAddr, "ssusd")
	if balance.Amount.LT(request.SSUSDAmount) {
		return sdkerrors.Wrap(types.ErrInsufficientFunds, "insufficient ssUSD balance")
	}

	// Calculate reserve value to redeem (1:1 backing)
	redeemValue := sdkmath.LegacyNewDecFromInt(request.SSUSDAmount).QuoInt64(1000000) // Convert from micro units

	// Calculate reserve assets to redeem based on current composition
	reserveAssets, err := engine.calculateRedemptionAssets(ctx, redeemValue, request.PreferredAsset)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to calculate redemption assets")
	}

	// Burn ssUSD tokens
	burnCoins := sdk.NewCoins(sdk.NewCoin("ssusd", request.SSUSDAmount))
	err = engine.keeper.bankKeeper.SendCoinsFromAccountToModule(
		ctx, requesterAddr, types.ModuleName, burnCoins)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to transfer ssUSD for burning")
	}

	err = engine.keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to burn ssUSD")
	}

	// Send reserve assets to requester
	err = engine.keeper.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, requesterAddr, reserveAssets)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to send reserve assets")
	}

	// Update reserve composition
	err = engine.updateReserveComposition(ctx, reserveAssets, false)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to update reserve composition")
	}

	// Update stablecoin total supply and reserve value
	stablecoin.TotalSupply = stablecoin.TotalSupply.Sub(request.SSUSDAmount)
	stablecoin.ReserveInfo.TotalReserveValue = stablecoin.ReserveInfo.TotalReserveValue.Sub(redeemValue)
	engine.keeper.SetStablecoin(ctx, stablecoin)

	// Emit redemption event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"ssusd_redeemed",
			sdk.NewAttribute("requester", request.Requester),
			sdk.NewAttribute("ssusd_amount", request.SSUSDAmount.String()),
			sdk.NewAttribute("redeemed_value", redeemValue.String()),
			sdk.NewAttribute("redeemed_assets", reserveAssets.String()),
			sdk.NewAttribute("new_total_supply", stablecoin.TotalSupply.String()),
		),
	)

	return nil
}

// calculateReserveValue calculates the USD value of reserve assets
func (engine *SSUSDStablecoinEngine) calculateReserveValue(ctx sdk.Context, reserves sdk.Coins) (sdkmath.LegacyDec, error) {
	totalValue := sdkmath.LegacyZeroDec()

	for _, coin := range reserves {
		// Get price for each asset
		price, err := engine.getAssetPrice(ctx, coin.Denom)
		if err != nil {
			return sdkmath.LegacyZeroDec(), err
		}

		// Calculate value: amount * price
		assetValue := sdkmath.LegacyNewDecFromInt(coin.Amount).Mul(price)
		totalValue = totalValue.Add(assetValue)
	}

	return totalValue, nil
}

// validateReserveComposition validates that reserve payment aligns with target allocations
func (engine *SSUSDStablecoinEngine) validateReserveComposition(ctx sdk.Context, payment sdk.Coins) error {
	// Get current reserve composition
	reserves, err := engine.GetConservativeReserves(ctx)
	if err != nil {
		return err
	}

	// Calculate total value after adding payment
	paymentValue, err := engine.calculateReserveValue(ctx, payment)
	if err != nil {
		return err
	}

	newTotalValue := reserves.TotalValue.Add(paymentValue)

	// Define target allocations based on conservative composition
	targetAllocations := map[string]sdkmath.LegacyDec{
		"us_cash":         sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
		"treasury_bills":  sdkmath.LegacyNewDecWithPrec(70, 2), // 70%
		"government_mmfs": sdkmath.LegacyNewDecWithPrec(15, 2), // 15%
		"overnight_repos": sdkmath.LegacyNewDecWithPrec(5, 2),  // 5%
	}

	// Check if payment maintains proper allocation ratios
	// This is a simplified check - in production, you'd implement more sophisticated validation
	for denom, coin := range payment {
		assetType := engine.getAssetType(denom)
		targetAllocation := targetAllocations[assetType]
		
		if targetAllocation.IsZero() {
			return errorsmod.Wrapf(types.ErrInvalidAsset, 
				"asset %s not allowed in conservative reserve composition", denom)
		}
	}

	return nil
}

// updateReserveComposition updates the reserve composition after issue/redeem
func (engine *SSUSDStablecoinEngine) updateReserveComposition(ctx sdk.Context, assets sdk.Coins, isAddition bool) error {
	reserves, err := engine.GetConservativeReserves(ctx)
	if err != nil {
		return err
	}

	for _, coin := range assets {
		assetValue, err := engine.getAssetPrice(ctx, coin.Denom)
		if err != nil {
			return err
		}

		coinValue := sdkmath.LegacyNewDecFromInt(coin.Amount).Mul(assetValue)
		assetType := engine.getAssetType(coin.Denom)

		// Update the appropriate reserve component
		switch assetType {
		case "us_cash":
			if isAddition {
				reserves.CashReserves.Amount = reserves.CashReserves.Amount.Add(coinValue)
			} else {
				reserves.CashReserves.Amount = reserves.CashReserves.Amount.Sub(coinValue)
			}
		case "treasury_bills":
			if isAddition {
				reserves.TreasuryBills.Amount = reserves.TreasuryBills.Amount.Add(coinValue)
			} else {
				reserves.TreasuryBills.Amount = reserves.TreasuryBills.Amount.Sub(coinValue)
			}
		case "government_mmfs":
			if isAddition {
				reserves.GovernmentMMFs.Amount = reserves.GovernmentMMFs.Amount.Add(coinValue)
			} else {
				reserves.GovernmentMMFs.Amount = reserves.GovernmentMMFs.Amount.Sub(coinValue)
			}
		case "overnight_repos":
			if isAddition {
				reserves.OvernightRepos.Amount = reserves.OvernightRepos.Amount.Add(coinValue)
			} else {
				reserves.OvernightRepos.Amount = reserves.OvernightRepos.Amount.Sub(coinValue)
			}
		}
	}

	// Recalculate total value
	reserves.TotalValue = reserves.CashReserves.Amount.
		Add(reserves.TreasuryBills.Amount).
		Add(reserves.GovernmentMMFs.Amount).
		Add(reserves.OvernightRepos.Amount)
	
	reserves.LastUpdate = time.Now()

	return engine.SetConservativeReserves(ctx, reserves)
}

// calculateRedemptionAssets calculates which reserve assets to redeem
func (engine *SSUSDStablecoinEngine) calculateRedemptionAssets(ctx sdk.Context, redeemValue sdkmath.LegacyDec, preferredAsset string) (sdk.Coins, error) {
	reserves, err := engine.GetConservativeReserves(ctx)
	if err != nil {
		return nil, err
	}

	var redemptionAssets sdk.Coins

	// If preferred asset is specified and available, prioritize it
	if preferredAsset != "" {
		assetType := engine.getAssetType(preferredAsset)
		var availableAmount sdkmath.LegacyDec

		switch assetType {
		case "us_cash":
			availableAmount = reserves.CashReserves.Amount
		case "treasury_bills":
			availableAmount = reserves.TreasuryBills.Amount
		case "government_mmfs":
			availableAmount = reserves.GovernmentMMFs.Amount
		case "overnight_repos":
			availableAmount = reserves.OvernightRepos.Amount
		}

		if availableAmount.GTE(redeemValue) {
			// Can fulfill entirely with preferred asset
			price, err := engine.getAssetPrice(ctx, preferredAsset)
			if err != nil {
				return nil, err
			}
			amount := redeemValue.Quo(price)
			redemptionAssets = sdk.NewCoins(sdk.NewCoin(preferredAsset, sdkmath.NewIntFromBigInt(amount.BigInt())))
			return redemptionAssets, nil
		}
	}

	// Redeem proportionally based on current composition
	remainingValue := redeemValue

	// Calculate proportions
	if reserves.TotalValue.IsZero() {
		return nil, sdkerrors.Wrap(types.ErrInsufficientCollateral, "no reserves available for redemption")
	}

	// Redeem from each reserve type proportionally
	assetTypes := []struct {
		denom  string
		amount sdkmath.LegacyDec
	}{
		{"us_cash_token", reserves.CashReserves.Amount},
		{"treasury_bill_token", reserves.TreasuryBills.Amount},
		{"mmf_token", reserves.GovernmentMMFs.Amount},
		{"repo_token", reserves.OvernightRepos.Amount},
	}

	for _, asset := range assetTypes {
		if asset.amount.IsZero() {
			continue
		}

		proportion := asset.amount.Quo(reserves.TotalValue)
		assetRedeemValue := redeemValue.Mul(proportion)

		if assetRedeemValue.GT(remainingValue) {
			assetRedeemValue = remainingValue
		}

		price, err := engine.getAssetPrice(ctx, asset.denom)
		if err != nil {
			return nil, err
		}

		amount := assetRedeemValue.Quo(price)
		if amount.GT(sdkmath.LegacyZeroDec()) {
			coin := sdk.NewCoin(asset.denom, sdkmath.NewIntFromBigInt(amount.BigInt()))
			redemptionAssets = redemptionAssets.Add(coin)
			remainingValue = remainingValue.Sub(assetRedeemValue)
		}

		if remainingValue.LTE(sdkmath.LegacyZeroDec()) {
			break
		}
	}

	return redemptionAssets, nil
}

// getAssetPrice gets the current price of an asset in USD
func (engine *SSUSDStablecoinEngine) getAssetPrice(ctx sdk.Context, denom string) (sdkmath.LegacyDec, error) {
	// This would integrate with your price oracle system
	// For now, returning simplified prices
	switch denom {
	case "us_cash_token", "treasury_bill_token", "mmf_token":
		return sdkmath.LegacyOneDec(), nil // $1.00 for USD-denominated assets
	case "repo_token":
		return sdkmath.LegacyOneDec(), nil // $1.00 for repo agreements
	default:
		// For other assets, you'd query the price oracle
		return sdkmath.LegacyOneDec(), nil
	}
}

// getAssetType maps denomination to asset type
func (engine *SSUSDStablecoinEngine) getAssetType(denom string) string {
	switch denom {
	case "us_cash_token":
		return "us_cash"
	case "treasury_bill_token":
		return "treasury_bills"
	case "mmf_token":
		return "government_mmfs"
	case "repo_token":
		return "overnight_repos"
	default:
		return "unknown"
	}
}

// GetConservativeReserves gets the current conservative reserve composition
func (engine *SSUSDStablecoinEngine) GetConservativeReserves(ctx sdk.Context) (*SSUSDConservativeReserve, error) {
	store := ctx.KVStore(engine.keeper.storeKey)
	key := []byte("ssusd_conservative_reserves")
	
	bz := store.Get(key)
	if bz == nil {
		// Initialize with default values
		return &SSUSDConservativeReserve{
			CashReserves: SSUSDCashReserve{
				Amount:      sdkmath.LegacyZeroDec(),
				Allocation:  sdkmath.LegacyNewDecWithPrec(10, 2), // 10%
				FDICInsured: true,
				RiskLevel:   "minimal",
			},
			TreasuryBills: SSUSDTreasuryBills{
				Amount:          sdkmath.LegacyZeroDec(),
				Allocation:      sdkmath.LegacyNewDecWithPrec(70, 2), // 70%
				AverageMaturity: 45, // 45 days average
				RiskLevel:       "minimal",
			},
			GovernmentMMFs: SSUSDGovernmentMMFs{
				Amount:     sdkmath.LegacyZeroDec(),
				Allocation: sdkmath.LegacyNewDecWithPrec(15, 2), // 15%
				WAM:        30, // 30 days weighted average maturity
				RiskLevel:  "minimal",
			},
			OvernightRepos: SSUSDOvernightRepos{
				Amount:      sdkmath.LegacyZeroDec(),
				Allocation:  sdkmath.LegacyNewDecWithPrec(5, 2), // 5%
				AverageRate: sdkmath.LegacyNewDecWithPrec(525, 4), // 5.25%
				RiskLevel:   "low",
			},
			TotalValue: sdkmath.LegacyZeroDec(),
			LastUpdate: time.Now(),
		}, nil
	}

	var reserves SSUSDConservativeReserve
	err := engine.keeper.cdc.Unmarshal(bz, &reserves)
	if err != nil {
		return nil, err
	}

	return &reserves, nil
}

// SetConservativeReserves sets the conservative reserve composition
func (engine *SSUSDStablecoinEngine) SetConservativeReserves(ctx sdk.Context, reserves *SSUSDConservativeReserve) error {
	store := ctx.KVStore(engine.keeper.storeKey)
	key := []byte("ssusd_conservative_reserves")
	
	bz, err := engine.keeper.cdc.Marshal(reserves)
	if err != nil {
		return err
	}
	
	store.Set(key, bz)
	return nil
}