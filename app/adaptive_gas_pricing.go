package app

import (
	"math"
	"math/big"
	"sort"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
)

// AdaptiveGasPricer implements dynamic gas pricing based on network conditions
type AdaptiveGasPricer struct {
	config              *GasPricingConfig
	networkMonitor      *NetworkMonitor
	demandPredictor     *DemandPredictor
	priceCalculator     *PriceCalculator
	gasPriceHistory     *GasPriceHistory
	congestionDetector  *CongestionDetector
	priorityEngine      *PriorityEngine
	metrics             *GasPricingMetrics
	currentBaseFee      sdk.Dec
	mu                  sync.RWMutex
}

// GasPricingConfig contains configuration for adaptive gas pricing
type GasPricingConfig struct {
	BaseGasPrice          sdk.Dec
	MinGasPrice           sdk.Dec
	MaxGasPrice           sdk.Dec
	TargetBlockUtilization float64
	BaseFeeChangeRate     float64
	PriorityFeeMultiplier float64
	CongestionThreshold   float64
	HistoryBlocks         int
	PredictionWindow      int
	AdaptationSpeed       float64
	VolatilityDamping     float64
	PriceUpdateInterval   time.Duration
}

// NetworkMonitor tracks network conditions for gas pricing
type NetworkMonitor struct {
	blockUtilization    []float64
	transactionCounts   []int
	averageGasPrices    []sdk.Dec
	mempoolSizes        []int
	validatorCounts     []int
	networkLatency      []time.Duration
	congestionLevel     float64
	utilizationTrend    float64
	mu                  sync.RWMutex
}

// DemandPredictor predicts future gas demand using various algorithms
type DemandPredictor struct {
	historicalData    *HistoricalData
	seasonalPatterns  map[string]*SeasonalPattern
	trendAnalyzer     *TrendAnalyzer
	mlPredictor       *MLPredictor
	predictionCache   map[string]*DemandPrediction
	accuracy          float64
	mu                sync.RWMutex
}

// PriceCalculator calculates optimal gas prices
type PriceCalculator struct {
	eip1559Calculator   *EIP1559Calculator
	auctionCalculator   *AuctionCalculator
	demandCalculator    *DemandCalculator
	fairnessCalculator  *FairnessCalculator
	strategyWeights     map[string]float64
	currentStrategy     PricingStrategy
}

// GasPriceHistory maintains historical gas price data
type GasPriceHistory struct {
	prices      []GasPricePoint
	maxHistory  int
	avgWindow   int
	volatility  float64
	trend       float64
	mu          sync.RWMutex
}

// CongestionDetector detects and measures network congestion
type CongestionDetector struct {
	congestionMetrics   *CongestionMetrics
	thresholds          *CongestionThresholds
	alertSystem         *AlertSystem
	adaptationTriggers  map[string]bool
	severityLevel       CongestionSeverity
}

// PriorityEngine manages transaction prioritization
type PriorityEngine struct {
	priorityQueues     map[string]*PriorityQueue
	dynamicWeights     map[string]float64
	fairnessEnforcer   *FairnessEnforcer
	antiMEVMeasures    *AntiMEVMeasures
	config             *PriorityConfig
}

// Data structures
type GasPricePoint struct {
	Timestamp     time.Time
	BasePrice     sdk.Dec
	PriorityPrice sdk.Dec
	BlockNumber   int64
	Utilization   float64
	TxCount       int
}

type HistoricalData struct {
	Prices        []GasPricePoint
	Utilization   []float64
	TxCounts      []int
	TimeOfDay     []int
	DayOfWeek     []int
	BlockIntervals []time.Duration
}

type SeasonalPattern struct {
	Pattern     []float64
	Confidence  float64
	Period      time.Duration
	LastUpdate  time.Time
}

type TrendAnalyzer struct {
	shortTermTrend  float64
	mediumTermTrend float64
	longTermTrend   float64
	momentum        float64
	volatility      float64
}

type MLPredictor struct {
	model           interface{} // Placeholder for ML model
	features        []string
	lastPrediction  *DemandPrediction
	accuracy        float64
	trainingData    [][]float64
	isEnabled       bool
}

type DemandPrediction struct {
	PredictedDemand   float64
	Confidence        float64
	TimeHorizon       time.Duration
	Factors           map[string]float64
	Timestamp         time.Time
}

type EIP1559Calculator struct {
	targetUtilization float64
	maxBaseFeeChange  float64
}

type AuctionCalculator struct {
	reservePrice    sdk.Dec
	incrementSize   sdk.Dec
	auctionDuration time.Duration
}

type DemandCalculator struct {
	elasticity        float64
	demandCurve       []DemandPoint
	priceFunction     func(demand float64) sdk.Dec
}

type FairnessCalculator struct {
	giniThreshold     float64
	maxWaitTime       time.Duration
	priorityDecayRate float64
}

type DemandPoint struct {
	Price  sdk.Dec
	Demand float64
}

type CongestionMetrics struct {
	MempoolSize       int
	BlockUtilization  float64
	AverageWaitTime   time.Duration
	TxRejectionRate   float64
	ValidatorLoad     float64
}

type CongestionThresholds struct {
	Low      float64
	Medium   float64
	High     float64
	Critical float64
}

type AlertSystem struct {
	alerts     []CongestionAlert
	callbacks  []AlertCallback
	isEnabled  bool
}

type CongestionAlert struct {
	Level     CongestionSeverity
	Message   string
	Timestamp time.Time
	Metrics   *CongestionMetrics
}

type AlertCallback func(alert CongestionAlert)

type PriorityQueue struct {
	transactions []PriorityTransaction
	weights      map[string]float64
	mu           sync.Mutex
}

type PriorityTransaction struct {
	Tx           sdk.Tx
	Priority     float64
	GasPrice     sdk.Dec
	Timestamp    time.Time
	WaitTime     time.Duration
	Category     string
}

type FairnessEnforcer struct {
	maxWaitTime       time.Duration
	priorityDecay     float64
	antiStarvation    bool
	fairnessMetrics   *FairnessMetrics
}

type AntiMEVMeasures struct {
	commitReveal      bool
	randomDelay       bool
	batchAuction      bool
	fairSequencing    bool
	mevProtection     float64
}

type FairnessMetrics struct {
	GiniCoefficient   float64
	WaitTimeVariance  float64
	ThroughputFairness float64
}

type PriorityConfig struct {
	MaxPriorityLevels int
	DecayRate         float64
	FairnessWeight    float64
	AntiMEVEnabled    bool
}

type GasPricingMetrics struct {
	TotalTransactions    uint64
	AverageGasPrice      sdk.Dec
	PriceVolatility      float64
	CongestionEvents     uint64
	PredictionAccuracy   float64
	RevenuePrediction    sdk.Dec
	UserSatisfaction     float64
	mu                   sync.RWMutex
}

// Enums
type PricingStrategy int

const (
	EIP1559Strategy PricingStrategy = iota
	AuctionStrategy
	DemandStrategy
	HybridStrategy
	MLStrategy
)

type CongestionSeverity int

const (
	CongestionNone CongestionSeverity = iota
	CongestionLow
	CongestionMedium
	CongestionHigh
	CongestionCritical
)

// NewAdaptiveGasPricer creates a new adaptive gas pricer
func NewAdaptiveGasPricer(config *GasPricingConfig) *AdaptiveGasPricer {
	if config == nil {
		config = DefaultGasPricingConfig()
	}

	pricer := &AdaptiveGasPricer{
		config:              config,
		networkMonitor:      NewNetworkMonitor(config.HistoryBlocks),
		demandPredictor:     NewDemandPredictor(config.PredictionWindow),
		priceCalculator:     NewPriceCalculator(),
		gasPriceHistory:     NewGasPriceHistory(config.HistoryBlocks),
		congestionDetector:  NewCongestionDetector(),
		priorityEngine:      NewPriorityEngine(),
		metrics:             NewGasPricingMetrics(),
		currentBaseFee:      config.BaseGasPrice,
	}

	// Start background processes
	go pricer.priceUpdateLoop()
	go pricer.networkMonitoringLoop()
	go pricer.congestionDetectionLoop()

	return pricer
}

// DefaultGasPricingConfig returns default gas pricing configuration
func DefaultGasPricingConfig() *GasPricingConfig {
	return &GasPricingConfig{
		BaseGasPrice:           sdk.NewDecWithPrec(1, 6),    // 0.000001
		MinGasPrice:            sdk.NewDecWithPrec(1, 9),    // 0.000000001
		MaxGasPrice:            sdk.NewDecWithPrec(1, 3),    // 0.001
		TargetBlockUtilization: 0.5,                         // 50%
		BaseFeeChangeRate:      0.125,                       // 12.5%
		PriorityFeeMultiplier:  2.0,
		CongestionThreshold:    0.8,                         // 80%
		HistoryBlocks:          100,
		PredictionWindow:       10,
		AdaptationSpeed:        0.1,
		VolatilityDamping:      0.8,
		PriceUpdateInterval:    5 * time.Second,
	}
}

// CalculateGasPrice calculates the current gas price based on network conditions
func (agp *AdaptiveGasPricer) CalculateGasPrice(ctx sdk.Context, tx sdk.Tx) (*GasPriceResult, error) {
	agp.mu.RLock()
	defer agp.mu.RUnlock()

	// Get current network state
	networkState := agp.networkMonitor.GetCurrentState()
	
	// Predict future demand
	demandPrediction, err := agp.demandPredictor.PredictDemand(networkState)
	if err != nil {
		return nil, err
	}

	// Calculate base fee using EIP-1559 style mechanism
	baseFee := agp.calculateBaseFee(networkState)
	
	// Calculate priority fee based on transaction characteristics
	priorityFee := agp.calculatePriorityFee(tx, networkState, demandPrediction)
	
	// Apply congestion adjustments
	congestionMultiplier := agp.congestionDetector.GetCongestionMultiplier()
	adjustedBaseFee := baseFee.Mul(sdk.NewDec(int64(congestionMultiplier * 1000)).Quo(sdk.NewDec(1000)))
	
	// Ensure prices are within bounds
	totalGasPrice := adjustedBaseFee.Add(priorityFee)
	totalGasPrice = sdk.MaxDec(totalGasPrice, agp.config.MinGasPrice)
	totalGasPrice = sdk.MinDec(totalGasPrice, agp.config.MaxGasPrice)

	result := &GasPriceResult{
		BaseFee:            adjustedBaseFee,
		PriorityFee:        priorityFee,
		TotalPrice:         totalGasPrice,
		CongestionLevel:    agp.congestionDetector.severityLevel,
		PredictedDemand:    demandPrediction.PredictedDemand,
		Recommendation:     agp.generateRecommendation(totalGasPrice, networkState),
		MaxGasEstimate:     agp.estimateMaxGas(tx),
		PriceValidUntil:    time.Now().Add(agp.config.PriceUpdateInterval),
	}

	// Update metrics
	agp.metrics.RecordGasPriceCalculation(result)

	return result, nil
}

// calculateBaseFee calculates the base fee using EIP-1559 style mechanism
func (agp *AdaptiveGasPricer) calculateBaseFee(networkState *NetworkState) sdk.Dec {
	targetUtilization := agp.config.TargetBlockUtilization
	currentUtilization := networkState.BlockUtilization
	
	// Calculate utilization ratio
	utilizationRatio := currentUtilization / targetUtilization
	
	// Apply exponential scaling for base fee adjustment
	if utilizationRatio > 1 {
		// Increase base fee exponentially as utilization exceeds target
		multiplier := math.Pow(1+agp.config.BaseFeeChangeRate, utilizationRatio-1)
		return agp.currentBaseFee.Mul(sdk.NewDecFromBigInt(big.NewInt(int64(multiplier * 1000000)))).Quo(sdk.NewDec(1000000))
	} else {
		// Decrease base fee when utilization is below target
		multiplier := math.Pow(1-agp.config.BaseFeeChangeRate, 1-utilizationRatio)
		return agp.currentBaseFee.Mul(sdk.NewDecFromBigInt(big.NewInt(int64(multiplier * 1000000)))).Quo(sdk.NewDec(1000000))
	}
}

// calculatePriorityFee calculates the priority fee for a transaction
func (agp *AdaptiveGasPricer) calculatePriorityFee(tx sdk.Tx, networkState *NetworkState, prediction *DemandPrediction) sdk.Dec {
	basePriorityFee := agp.currentBaseFee.Mul(sdk.NewDecFromBigInt(big.NewInt(int64(agp.config.PriorityFeeMultiplier * 1000)))).Quo(sdk.NewDec(1000))
	
	// Adjust based on transaction characteristics
	txPriority := agp.calculateTransactionPriority(tx)
	priorityMultiplier := 1.0 + (txPriority-0.5)*0.5 // Scale between 0.75 and 1.25
	
	// Adjust based on predicted demand
	demandMultiplier := 1.0 + prediction.PredictedDemand*0.2
	
	// Adjust based on mempool competition
	competitionMultiplier := 1.0 + float64(networkState.MempoolSize)/1000.0*0.1
	
	totalMultiplier := priorityMultiplier * demandMultiplier * competitionMultiplier
	
	return basePriorityFee.Mul(sdk.NewDecFromBigInt(big.NewInt(int64(totalMultiplier * 1000000)))).Quo(sdk.NewDec(1000000))
}

// calculateTransactionPriority calculates priority score for a transaction
func (agp *AdaptiveGasPricer) calculateTransactionPriority(tx sdk.Tx) float64 {
	priority := 0.5 // Base priority
	
	// Analyze transaction type
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		msgType := sdk.MsgTypeURL(msg)
		switch msgType {
		case "/cosmos.bank.v1beta1.MsgSend":
			priority += 0.1 // Basic transfers get slight boost
		case "/cosmos.staking.v1beta1.MsgDelegate":
			priority += 0.2 // Staking operations get higher priority
		case "/cosmos.gov.v1beta1.MsgSubmitProposal":
			priority += 0.3 // Governance gets high priority
		default:
			priority += 0.0 // Default priority
		}
	}
	
	// Analyze fee
	if feeTx, ok := tx.(sdk.FeeTx); ok {
		fees := feeTx.GetFee()
		gas := feeTx.GetGas()
		if gas > 0 && !fees.IsZero() {
			// Higher fee ratio increases priority
			feeRatio := float64(fees.AmountOf("stake").Int64()) / float64(gas)
			priority += feeRatio * 0.1
		}
	}
	
	// Clamp priority between 0 and 1
	if priority > 1.0 {
		priority = 1.0
	}
	if priority < 0.0 {
		priority = 0.0
	}
	
	return priority
}

// estimateMaxGas estimates maximum gas for a transaction
func (agp *AdaptiveGasPricer) estimateMaxGas(tx sdk.Tx) uint64 {
	if feeTx, ok := tx.(sdk.FeeTx); ok {
		return feeTx.GetGas()
	}
	
	// Default gas estimate based on transaction type
	msgs := tx.GetMsgs()
	totalGas := uint64(21000) // Base gas
	
	for _, msg := range msgs {
		msgType := sdk.MsgTypeURL(msg)
		switch msgType {
		case "/cosmos.bank.v1beta1.MsgSend":
			totalGas += 25000
		case "/cosmos.staking.v1beta1.MsgDelegate":
			totalGas += 50000
		case "/cosmos.gov.v1beta1.MsgSubmitProposal":
			totalGas += 100000
		default:
			totalGas += 30000
		}
	}
	
	return totalGas
}

// generateRecommendation generates gas price recommendation
func (agp *AdaptiveGasPricer) generateRecommendation(gasPrice sdk.Dec, networkState *NetworkState) GasPriceRecommendation {
	if networkState.BlockUtilization < 0.3 {
		return RecommendationEconomy
	} else if networkState.BlockUtilization < 0.7 {
		return RecommendationStandard
	} else if networkState.BlockUtilization < 0.9 {
		return RecommendationFast
	} else {
		return RecommendationUrgent
	}
}

// priceUpdateLoop continuously updates gas prices
func (agp *AdaptiveGasPricer) priceUpdateLoop() {
	ticker := time.NewTicker(agp.config.PriceUpdateInterval)
	defer ticker.Stop()
	
	for range ticker.C {
		agp.updateBaseFee()
	}
}

// networkMonitoringLoop monitors network conditions
func (agp *AdaptiveGasPricer) networkMonitoringLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		agp.networkMonitor.UpdateMetrics()
	}
}

// congestionDetectionLoop detects network congestion
func (agp *AdaptiveGasPricer) congestionDetectionLoop() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		agp.congestionDetector.UpdateCongestionLevel()
	}
}

// updateBaseFee updates the base fee based on current conditions
func (agp *AdaptiveGasPricer) updateBaseFee() {
	agp.mu.Lock()
	defer agp.mu.Unlock()
	
	networkState := agp.networkMonitor.GetCurrentState()
	newBaseFee := agp.calculateBaseFee(networkState)
	
	// Apply volatility damping
	dampingFactor := agp.config.VolatilityDamping
	agp.currentBaseFee = agp.currentBaseFee.Mul(sdk.NewDecFromBigInt(big.NewInt(int64(dampingFactor*1000)))).Quo(sdk.NewDec(1000)).Add(
		newBaseFee.Mul(sdk.NewDecFromBigInt(big.NewInt(int64((1-dampingFactor)*1000)))).Quo(sdk.NewDec(1000)))
	
	// Record price point
	agp.gasPriceHistory.AddPricePoint(GasPricePoint{
		Timestamp:   time.Now(),
		BasePrice:   agp.currentBaseFee,
		BlockNumber: 0, // Would be actual block number in implementation
		Utilization: networkState.BlockUtilization,
		TxCount:     networkState.TxCount,
	})
}

// Additional data structures and helper types

type GasPriceResult struct {
	BaseFee            sdk.Dec
	PriorityFee        sdk.Dec
	TotalPrice         sdk.Dec
	CongestionLevel    CongestionSeverity
	PredictedDemand    float64
	Recommendation     GasPriceRecommendation
	MaxGasEstimate     uint64
	PriceValidUntil    time.Time
}

type GasPriceRecommendation int

const (
	RecommendationEconomy GasPriceRecommendation = iota
	RecommendationStandard
	RecommendationFast
	RecommendationUrgent
)

type NetworkState struct {
	BlockUtilization float64
	TxCount          int
	MempoolSize      int
	AverageGasPrice  sdk.Dec
	ValidatorCount   int
	NetworkLatency   time.Duration
	Timestamp        time.Time
}

// Helper constructors and methods
func NewNetworkMonitor(historySize int) *NetworkMonitor {
	return &NetworkMonitor{
		blockUtilization:  make([]float64, 0, historySize),
		transactionCounts: make([]int, 0, historySize),
		averageGasPrices:  make([]sdk.Dec, 0, historySize),
		mempoolSizes:      make([]int, 0, historySize),
		validatorCounts:   make([]int, 0, historySize),
		networkLatency:    make([]time.Duration, 0, historySize),
	}
}

func (nm *NetworkMonitor) UpdateMetrics() {
	// Implementation would update metrics from actual network state
}

func (nm *NetworkMonitor) GetCurrentState() *NetworkState {
	nm.mu.RLock()
	defer nm.mu.RUnlock()
	
	return &NetworkState{
		BlockUtilization: nm.getAverageUtilization(),
		TxCount:          nm.getAverageTxCount(),
		MempoolSize:      nm.getCurrentMempoolSize(),
		ValidatorCount:   nm.getCurrentValidatorCount(),
		Timestamp:        time.Now(),
	}
}

func (nm *NetworkMonitor) getAverageUtilization() float64 {
	if len(nm.blockUtilization) == 0 {
		return 0.5 // Default
	}
	
	sum := 0.0
	for _, util := range nm.blockUtilization {
		sum += util
	}
	return sum / float64(len(nm.blockUtilization))
}

func (nm *NetworkMonitor) getAverageTxCount() int {
	if len(nm.transactionCounts) == 0 {
		return 100 // Default
	}
	
	sum := 0
	for _, count := range nm.transactionCounts {
		sum += count
	}
	return sum / len(nm.transactionCounts)
}

func (nm *NetworkMonitor) getCurrentMempoolSize() int {
	if len(nm.mempoolSizes) == 0 {
		return 1000 // Default
	}
	return nm.mempoolSizes[len(nm.mempoolSizes)-1]
}

func (nm *NetworkMonitor) getCurrentValidatorCount() int {
	if len(nm.validatorCounts) == 0 {
		return 100 // Default
	}
	return nm.validatorCounts[len(nm.validatorCounts)-1]
}

func NewDemandPredictor(window int) *DemandPredictor {
	return &DemandPredictor{
		historicalData:   &HistoricalData{},
		seasonalPatterns: make(map[string]*SeasonalPattern),
		trendAnalyzer:    &TrendAnalyzer{},
		mlPredictor:      &MLPredictor{isEnabled: false},
		predictionCache:  make(map[string]*DemandPrediction),
		accuracy:         0.75, // Initial accuracy estimate
	}
}

func (dp *DemandPredictor) PredictDemand(networkState *NetworkState) (*DemandPrediction, error) {
	// Simple prediction based on current utilization and trends
	prediction := &DemandPrediction{
		PredictedDemand: networkState.BlockUtilization * 1.1, // Slight increase prediction
		Confidence:      dp.accuracy,
		TimeHorizon:     5 * time.Minute,
		Factors:         make(map[string]float64),
		Timestamp:       time.Now(),
	}
	
	prediction.Factors["utilization"] = networkState.BlockUtilization
	prediction.Factors["mempool_size"] = float64(networkState.MempoolSize) / 1000.0
	prediction.Factors["validator_count"] = float64(networkState.ValidatorCount) / 100.0
	
	return prediction, nil
}

func NewPriceCalculator() *PriceCalculator {
	return &PriceCalculator{
		eip1559Calculator:  &EIP1559Calculator{targetUtilization: 0.5, maxBaseFeeChange: 0.125},
		auctionCalculator:  &AuctionCalculator{},
		demandCalculator:   &DemandCalculator{elasticity: 0.5},
		fairnessCalculator: &FairnessCalculator{giniThreshold: 0.4, maxWaitTime: 5 * time.Minute},
		strategyWeights:    map[string]float64{"eip1559": 0.7, "demand": 0.2, "fairness": 0.1},
		currentStrategy:    EIP1559Strategy,
	}
}

func NewGasPriceHistory(maxHistory int) *GasPriceHistory {
	return &GasPriceHistory{
		prices:     make([]GasPricePoint, 0, maxHistory),
		maxHistory: maxHistory,
		avgWindow:  10,
	}
}

func (gph *GasPriceHistory) AddPricePoint(point GasPricePoint) {
	gph.mu.Lock()
	defer gph.mu.Unlock()
	
	gph.prices = append(gph.prices, point)
	if len(gph.prices) > gph.maxHistory {
		gph.prices = gph.prices[1:]
	}
	
	gph.updateStatistics()
}

func (gph *GasPriceHistory) updateStatistics() {
	if len(gph.prices) < 2 {
		return
	}
	
	// Calculate volatility
	prices := make([]float64, len(gph.prices))
	for i, point := range gph.prices {
		prices[i], _ = point.BasePrice.Float64()
	}
	
	mean := 0.0
	for _, price := range prices {
		mean += price
	}
	mean /= float64(len(prices))
	
	variance := 0.0
	for _, price := range prices {
		variance += (price - mean) * (price - mean)
	}
	variance /= float64(len(prices))
	
	gph.volatility = math.Sqrt(variance) / mean
}

func NewCongestionDetector() *CongestionDetector {
	return &CongestionDetector{
		congestionMetrics: &CongestionMetrics{},
		thresholds: &CongestionThresholds{
			Low:      0.3,
			Medium:   0.6,
			High:     0.8,
			Critical: 0.95,
		},
		alertSystem:        &AlertSystem{isEnabled: true},
		adaptationTriggers: make(map[string]bool),
		severityLevel:      CongestionNone,
	}
}

func (cd *CongestionDetector) UpdateCongestionLevel() {
	// Implementation would update congestion metrics and severity
}

func (cd *CongestionDetector) GetCongestionMultiplier() float64 {
	switch cd.severityLevel {
	case CongestionNone:
		return 1.0
	case CongestionLow:
		return 1.1
	case CongestionMedium:
		return 1.3
	case CongestionHigh:
		return 1.6
	case CongestionCritical:
		return 2.0
	default:
		return 1.0
	}
}

func NewPriorityEngine() *PriorityEngine {
	return &PriorityEngine{
		priorityQueues:   make(map[string]*PriorityQueue),
		dynamicWeights:   make(map[string]float64),
		fairnessEnforcer: &FairnessEnforcer{maxWaitTime: 5 * time.Minute, priorityDecay: 0.1, antiStarvation: true},
		antiMEVMeasures:  &AntiMEVMeasures{commitReveal: true, randomDelay: true, mevProtection: 0.8},
		config:           &PriorityConfig{MaxPriorityLevels: 10, DecayRate: 0.1, FairnessWeight: 0.3, AntiMEVEnabled: true},
	}
}

func NewGasPricingMetrics() *GasPricingMetrics {
	return &GasPricingMetrics{}
}

func (gpm *GasPricingMetrics) RecordGasPriceCalculation(result *GasPriceResult) {
	gpm.mu.Lock()
	defer gpm.mu.Unlock()
	
	gpm.TotalTransactions++
	
	// Update average gas price (exponential moving average)
	if gpm.AverageGasPrice.IsZero() {
		gpm.AverageGasPrice = result.TotalPrice
	} else {
		// 90% old value + 10% new value
		gpm.AverageGasPrice = gpm.AverageGasPrice.Mul(sdk.NewDecWithPrec(9, 1)).Add(result.TotalPrice.Mul(sdk.NewDecWithPrec(1, 1)))
	}
}

// GetCurrentGasPrice returns the current gas price recommendation
func (agp *AdaptiveGasPricer) GetCurrentGasPrice() sdk.Dec {
	agp.mu.RLock()
	defer agp.mu.RUnlock()
	return agp.currentBaseFee
}

// GetGasPricingMetrics returns current gas pricing metrics
func (agp *AdaptiveGasPricer) GetGasPricingMetrics() *GasPricingMetrics {
	return agp.metrics
}

// SetGasPricingStrategy sets the gas pricing strategy
func (agp *AdaptiveGasPricer) SetGasPricingStrategy(strategy PricingStrategy) {
	agp.mu.Lock()
	defer agp.mu.Unlock()
	agp.priceCalculator.currentStrategy = strategy
}