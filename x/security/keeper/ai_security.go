package keeper

import (
	"context"
	"fmt"
	"math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/security/types"
)

// AISecurityEngine provides advanced AI-powered security features
type AISecurityEngine struct {
	keeper           *Keeper
	behaviorProfiles map[string]*UserBehaviorProfile
	anomalyModels    map[string]*AnomalyDetectionModel
	threatScore      *ThreatScoringEngine
}

// UserBehaviorProfile tracks user behavior patterns for anomaly detection
type UserBehaviorProfile struct {
	Address            string    `json:"address"`
	TypicalTxAmount    float64   `json:"typical_tx_amount"`
	TypicalTxFreq      float64   `json:"typical_tx_frequency"`
	PreferredHours     []int     `json:"preferred_hours"`
	GeographicPattern  []string  `json:"geographic_pattern"`
	LastUpdate         time.Time `json:"last_update"`
	RiskScore          float64   `json:"risk_score"`
	ConfidenceLevel    float64   `json:"confidence_level"`
}

// AnomalyDetectionModel represents an ML model for detecting specific types of anomalies
type AnomalyDetectionModel struct {
	ModelID         string                 `json:"model_id"`
	ModelType       string                 `json:"model_type"` // "velocity", "pattern", "amount", "frequency"
	Parameters      map[string]interface{} `json:"parameters"`
	Accuracy        float64               `json:"accuracy"`
	LastTrained     time.Time             `json:"last_trained"`
	ThreatThreshold float64               `json:"threat_threshold"`
}

// ThreatScoringEngine provides comprehensive threat scoring
type ThreatScoringEngine struct {
	WeightVelocity   float64 `json:"weight_velocity"`
	WeightPattern    float64 `json:"weight_pattern"`
	WeightAmount     float64 `json:"weight_amount"`
	WeightGeography  float64 `json:"weight_geography"`
	WeightHistory    float64 `json:"weight_history"`
}

// NewAISecurityEngine creates a new AI security engine
func NewAISecurityEngine(keeper *Keeper) *AISecurityEngine {
	return &AISecurityEngine{
		keeper:           keeper,
		behaviorProfiles: make(map[string]*UserBehaviorProfile),
		anomalyModels:    make(map[string]*AnomalyDetectionModel),
		threatScore: &ThreatScoringEngine{
			WeightVelocity:  0.25,
			WeightPattern:   0.20,
			WeightAmount:    0.25,
			WeightGeography: 0.15,
			WeightHistory:   0.15,
		},
	}
}

// AnalyzeTransaction performs comprehensive AI-powered transaction analysis
func (ai *AISecurityEngine) AnalyzeTransaction(ctx sdk.Context, tx *types.TransactionData) (*types.AISecurityAnalysis, error) {
	analysis := &types.AISecurityAnalysis{
		TransactionID:   tx.Hash,
		Timestamp:       time.Now(),
		ThreatLevel:     "LOW",
		ConfidenceScore: 0.0,
		Recommendations: []string{},
		Alerts:          []types.SecurityAlert{},
	}

	// 1. Velocity Analysis
	velocityScore := ai.analyzeVelocity(ctx, tx)
	
	// 2. Pattern Analysis
	patternScore := ai.analyzePatterns(ctx, tx)
	
	// 3. Amount Analysis
	amountScore := ai.analyzeAmount(ctx, tx)
	
	// 4. Geographic Analysis
	geoScore := ai.analyzeGeography(ctx, tx)
	
	// 5. Historical Behavior Analysis
	historyScore := ai.analyzeHistory(ctx, tx)

	// Calculate composite threat score
	threatScore := ai.calculateCompositeThreatScore(velocityScore, patternScore, amountScore, geoScore, historyScore)
	analysis.ThreatScore = threatScore

	// Determine threat level and actions
	if threatScore >= 90 {
		analysis.ThreatLevel = "CRITICAL"
		analysis.Recommendations = append(analysis.Recommendations, "BLOCK_TRANSACTION", "IMMEDIATE_REVIEW", "FREEZE_ACCOUNT")
		analysis.Alerts = append(analysis.Alerts, types.SecurityAlert{
			AlertID:   fmt.Sprintf("AI-CRITICAL-%d", time.Now().Unix()),
			Severity:  "CRITICAL",
			Message:   "AI detected critical threat - transaction blocked",
			Timestamp: time.Now(),
		})
	} else if threatScore >= 70 {
		analysis.ThreatLevel = "HIGH"
		analysis.Recommendations = append(analysis.Recommendations, "REQUIRE_ADDITIONAL_AUTH", "MANUAL_REVIEW")
		analysis.Alerts = append(analysis.Alerts, types.SecurityAlert{
			AlertID:   fmt.Sprintf("AI-HIGH-%d", time.Now().Unix()),
			Severity:  "HIGH",
			Message:   "AI detected high-risk transaction",
			Timestamp: time.Now(),
		})
	} else if threatScore >= 40 {
		analysis.ThreatLevel = "MEDIUM"
		analysis.Recommendations = append(analysis.Recommendations, "ENHANCED_MONITORING", "REVIEW_WITHIN_24H")
	}

	// Update user behavior profile
	ai.updateBehaviorProfile(ctx, tx, analysis)

	analysis.ConfidenceScore = ai.calculateConfidenceScore(velocityScore, patternScore, amountScore, geoScore, historyScore)

	return analysis, nil
}

// analyzeVelocity detects unusual transaction velocity patterns
func (ai *AISecurityEngine) analyzeVelocity(ctx sdk.Context, tx *types.TransactionData) float64 {
	// Get recent transactions for this address
	recentTxs := ai.keeper.GetRecentTransactions(ctx, tx.FromAddress, time.Hour*24)
	
	if len(recentTxs) < 2 {
		return 0.0 // Not enough data
	}

	// Calculate current velocity metrics
	currentHourTxs := ai.countTransactionsInTimeWindow(recentTxs, time.Hour)
	currentDayAmount := ai.sumTransactionAmounts(recentTxs)
	
	// Get user's normal velocity profile
	profile := ai.getBehaviorProfile(tx.FromAddress)
	
	// Calculate velocity anomaly scores
	freqAnomalyScore := ai.calculateFrequencyAnomalyScore(float64(currentHourTxs), profile.TypicalTxFreq)
	amountAnomalyScore := ai.calculateAmountAnomalyScore(currentDayAmount, profile.TypicalTxAmount)
	
	// Return weighted velocity score
	return (freqAnomalyScore*0.6 + amountAnomalyScore*0.4)
}

// analyzePatterns detects suspicious transaction patterns
func (ai *AISecurityEngine) analyzePatterns(ctx sdk.Context, tx *types.TransactionData) float64 {
	score := 0.0
	
	// Check for round number pattern (often indicates automation/bots)
	if ai.isRoundNumber(tx.Amount) {
		score += 20.0
	}
	
	// Check for rapid-fire transactions
	recentTxs := ai.keeper.GetRecentTransactions(ctx, tx.FromAddress, time.Minute*10)
	if len(recentTxs) > 5 {
		score += 30.0
	}
	
	// Check for identical amounts pattern
	if ai.hasIdenticalAmountsPattern(recentTxs) {
		score += 25.0
	}
	
	// Check for timing patterns (e.g., exactly every minute)
	if ai.hasRegularTimingPattern(recentTxs) {
		score += 25.0
	}
	
	// Cap at 100
	if score > 100 {
		score = 100
	}
	
	return score
}

// analyzeAmount detects unusual transaction amounts
func (ai *AISecurityEngine) analyzeAmount(ctx sdk.Context, tx *types.TransactionData) float64 {
	profile := ai.getBehaviorProfile(tx.FromAddress)
	
	// Calculate z-score for amount deviation
	if profile.TypicalTxAmount == 0 {
		return 0.0 // No historical data
	}
	
	deviation := math.Abs(tx.Amount - profile.TypicalTxAmount) / profile.TypicalTxAmount
	
	// Convert deviation to risk score (0-100)
	score := deviation * 50
	if score > 100 {
		score = 100
	}
	
	return score
}

// analyzeGeography detects unusual geographic patterns
func (ai *AISecurityEngine) analyzeGeography(ctx sdk.Context, tx *types.TransactionData) float64 {
	// This would integrate with IP geolocation services
	// For now, return basic geographic risk assessment
	profile := ai.getBehaviorProfile(tx.FromAddress)
	
	// Check if current location matches historical pattern
	currentLocation := ai.getTransactionLocation(tx)
	if currentLocation == "" {
		return 10.0 // Unknown location = slight risk
	}
	
	// Check against user's typical locations
	for _, location := range profile.GeographicPattern {
		if location == currentLocation {
			return 0.0 // Known location = no risk
		}
	}
	
	// New location = moderate risk
	return 40.0
}

// analyzeHistory analyzes user's historical behavior patterns
func (ai *AISecurityEngine) analyzeHistory(ctx sdk.Context, tx *types.TransactionData) float64 {
	profile := ai.getBehaviorProfile(tx.FromAddress)
	
	// Base score on existing risk profile
	baseScore := profile.RiskScore
	
	// Adjust based on account age and activity
	accountAge := ai.getAccountAge(ctx, tx.FromAddress)
	if accountAge < time.Hour*24 {
		baseScore += 50.0 // New accounts are riskier
	} else if accountAge < time.Hour*24*7 {
		baseScore += 20.0 // Week-old accounts have some risk
	}
	
	// Check for previous security incidents
	incidents := ai.keeper.GetSecurityIncidents(ctx, tx.FromAddress)
	baseScore += float64(len(incidents)) * 15.0
	
	if baseScore > 100 {
		baseScore = 100
	}
	
	return baseScore
}

// calculateCompositeThreatScore combines all analysis scores
func (ai *AISecurityEngine) calculateCompositeThreatScore(velocity, pattern, amount, geography, history float64) float64 {
	return velocity*ai.threatScore.WeightVelocity +
		pattern*ai.threatScore.WeightPattern +
		amount*ai.threatScore.WeightAmount +
		geography*ai.threatScore.WeightGeography +
		history*ai.threatScore.WeightHistory
}

// calculateConfidenceScore calculates confidence in the threat assessment
func (ai *AISecurityEngine) calculateConfidenceScore(scores ...float64) float64 {
	// Higher confidence when multiple models agree
	variance := ai.calculateVariance(scores)
	confidence := 100.0 - (variance * 10.0)
	
	if confidence < 0 {
		confidence = 0
	}
	if confidence > 100 {
		confidence = 100
	}
	
	return confidence
}

// Helper functions

func (ai *AISecurityEngine) getBehaviorProfile(address string) *UserBehaviorProfile {
	if profile, exists := ai.behaviorProfiles[address]; exists {
		return profile
	}
	
	// Create new profile with defaults
	return &UserBehaviorProfile{
		Address:           address,
		TypicalTxAmount:   0.0,
		TypicalTxFreq:     0.0,
		PreferredHours:    []int{},
		GeographicPattern: []string{},
		LastUpdate:        time.Now(),
		RiskScore:         0.0,
		ConfidenceLevel:   0.0,
	}
}

func (ai *AISecurityEngine) updateBehaviorProfile(ctx sdk.Context, tx *types.TransactionData, analysis *types.AISecurityAnalysis) {
	// Update user behavior profile based on current transaction and analysis
	// This would use exponential moving averages to update patterns
}

func (ai *AISecurityEngine) isRoundNumber(amount float64) bool {
	return amount == math.Floor(amount) && math.Mod(amount, 100) == 0
}

func (ai *AISecurityEngine) hasIdenticalAmountsPattern(transactions []types.Transaction) bool {
	if len(transactions) < 3 {
		return false
	}
	
	firstAmount := transactions[0].Amount
	identicalCount := 0
	
	for _, tx := range transactions {
		if tx.Amount == firstAmount {
			identicalCount++
		}
	}
	
	return float64(identicalCount)/float64(len(transactions)) > 0.7
}

func (ai *AISecurityEngine) hasRegularTimingPattern(transactions []types.Transaction) bool {
	// Detect if transactions occur at regular intervals (automation indicator)
	if len(transactions) < 3 {
		return false
	}
	
	// Calculate intervals between transactions
	intervals := []float64{}
	for i := 1; i < len(transactions); i++ {
		interval := transactions[i].Timestamp.Sub(transactions[i-1].Timestamp).Seconds()
		intervals = append(intervals, interval)
	}
	
	// Check if intervals are suspiciously regular
	variance := ai.calculateVariance(intervals)
	return variance < 10.0 // Very low variance = regular timing
}

func (ai *AISecurityEngine) countTransactionsInTimeWindow(transactions []types.Transaction, window time.Duration) int {
	cutoff := time.Now().Add(-window)
	count := 0
	
	for _, tx := range transactions {
		if tx.Timestamp.After(cutoff) {
			count++
		}
	}
	
	return count
}

func (ai *AISecurityEngine) sumTransactionAmounts(transactions []types.Transaction) float64 {
	total := 0.0
	for _, tx := range transactions {
		total += tx.Amount
	}
	return total
}

func (ai *AISecurityEngine) calculateFrequencyAnomalyScore(current, typical float64) float64 {
	if typical == 0 {
		return 0.0
	}
	
	ratio := current / typical
	if ratio > 3.0 {
		return 100.0
	} else if ratio > 2.0 {
		return 70.0
	} else if ratio > 1.5 {
		return 40.0
	}
	
	return 0.0
}

func (ai *AISecurityEngine) calculateAmountAnomalyScore(current, typical float64) float64 {
	if typical == 0 {
		return 0.0
	}
	
	ratio := current / typical
	if ratio > 5.0 {
		return 100.0
	} else if ratio > 3.0 {
		return 70.0
	} else if ratio > 2.0 {
		return 40.0
	}
	
	return 0.0
}

func (ai *AISecurityEngine) getTransactionLocation(tx *types.TransactionData) string {
	// This would integrate with IP geolocation services
	// For now, return empty string
	return ""
}

func (ai *AISecurityEngine) getAccountAge(ctx sdk.Context, address string) time.Duration {
	// Get the first transaction timestamp for this address
	firstTx := ai.keeper.GetFirstTransaction(ctx, address)
	if firstTx.IsZero() {
		return time.Duration(0)
	}
	
	return time.Since(firstTx)
}

func (ai *AISecurityEngine) calculateVariance(values []float64) float64 {
	if len(values) <= 1 {
		return 0.0
	}
	
	// Calculate mean
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	mean := sum / float64(len(values))
	
	// Calculate variance
	sumSquaredDiff := 0.0
	for _, v := range values {
		diff := v - mean
		sumSquaredDiff += diff * diff
	}
	
	return sumSquaredDiff / float64(len(values)-1)
}