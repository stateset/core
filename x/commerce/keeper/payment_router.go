package keeper

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/commerce/types"
)

// PaymentRouter handles intelligent payment routing
type PaymentRouter struct {
	keeper *Keeper
}

// NewPaymentRouter creates a new payment router
func NewPaymentRouter(keeper *Keeper) *PaymentRouter {
	return &PaymentRouter{
		keeper: keeper,
	}
}

// RouteScore represents a scored payment route
type RouteScore struct {
	Route       types.PaymentRoute
	Score       float64
	CostScore   float64
	TimeScore   float64
	ReliabilityScore float64
}

// FindOptimalRoute finds the optimal payment route using AI/ML algorithms
func (pr *PaymentRouter) FindOptimalRoute(ctx sdk.Context, paymentInfo types.PaymentInfo) (types.PaymentRoute, error) {
	// Get all possible routes
	routes := pr.generateRoutes(ctx, paymentInfo)
	if len(routes) == 0 {
		return types.PaymentRoute{}, fmt.Errorf("no viable payment routes found")
	}

	// Score each route
	scoredRoutes := make([]RouteScore, len(routes))
	for i, route := range routes {
		score := pr.scoreRoute(ctx, route, paymentInfo)
		scoredRoutes[i] = RouteScore{
			Route:           route,
			Score:           score.OverallScore,
			CostScore:       score.CostScore,
			TimeScore:       score.TimeScore,
			ReliabilityScore: score.ReliabilityScore,
		}
	}

	// Sort by score (highest first)
	sort.Slice(scoredRoutes, func(i, j int) bool {
		return scoredRoutes[i].Score > scoredRoutes[j].Score
	})

	bestRoute := scoredRoutes[0].Route
	
	// Add backup route
	if len(scoredRoutes) > 1 {
		bestRoute.Backup = &scoredRoutes[1].Route
	}

	return bestRoute, nil
}

// generateRoutes generates all possible payment routes
func (pr *PaymentRouter) generateRoutes(ctx sdk.Context, paymentInfo types.PaymentInfo) []types.PaymentRoute {
	var routes []types.PaymentRoute

	// Direct route
	if directRoute := pr.generateDirectRoute(ctx, paymentInfo); directRoute != nil {
		routes = append(routes, *directRoute)
	}

	// Multi-hop routes
	multiHopRoutes := pr.generateMultiHopRoutes(ctx, paymentInfo)
	routes = append(routes, multiHopRoutes...)

	// Bridge routes (for cross-border payments)
	if paymentInfo.CrossBorderInfo != nil {
		bridgeRoutes := pr.generateBridgeRoutes(ctx, paymentInfo)
		routes = append(routes, bridgeRoutes...)
	}

	// Stablecoin conversion routes
	stablecoinRoutes := pr.generateStablecoinRoutes(ctx, paymentInfo)
	routes = append(routes, stablecoinRoutes...)

	return routes
}

// generateDirectRoute creates a direct payment route
func (pr *PaymentRouter) generateDirectRoute(ctx sdk.Context, paymentInfo types.PaymentInfo) *types.PaymentRoute {
	// Check if direct transfer is possible (same chain, compatible currencies)
	if paymentInfo.CrossBorderInfo == nil || 
	   paymentInfo.CrossBorderInfo.SourceCountry == paymentInfo.CrossBorderInfo.DestinationCountry {
		
		estimatedTime := time.Duration(6 * time.Second) // Average block time
		
		return &types.PaymentRoute{
			Type:          types.PaymentRouteDirect,
			Hops:         []types.PaymentHop{}, // No intermediary hops
			EstimatedTime: estimatedTime,
			Confidence:   0.95, // High confidence for direct transfers
		}
	}
	
	return nil
}

// generateMultiHopRoutes creates multi-hop payment routes through intermediaries
func (pr *PaymentRouter) generateMultiHopRoutes(ctx sdk.Context, paymentInfo types.PaymentInfo) []types.PaymentRoute {
	var routes []types.PaymentRoute

	// Find potential intermediaries based on liquidity and reputation
	intermediaries := pr.findIntermediaries(ctx, paymentInfo)

	for _, intermediary := range intermediaries {
		// Create route through this intermediary
		hops := []types.PaymentHop{
			{
				From:         "", // Will be filled by caller
				To:           intermediary.Address,
				Amount:       paymentInfo.Amount,
				Fee:          pr.calculateHopFee(ctx, intermediary, paymentInfo.Amount),
				Network:      "stateset",
				Protocol:     "cosmos-sdk",
				TimeEstimate: time.Duration(12 * time.Second),
			},
			{
				From:         intermediary.Address,
				To:           "", // Will be filled by caller
				Amount:       paymentInfo.Amount,
				Fee:          pr.calculateHopFee(ctx, intermediary, paymentInfo.Amount),
				Network:      "stateset",
				Protocol:     "cosmos-sdk",
				TimeEstimate: time.Duration(12 * time.Second),
			},
		}

		route := types.PaymentRoute{
			Type:          types.PaymentRouteMultiHop,
			Hops:         hops,
			EstimatedTime: time.Duration(30 * time.Second),
			Confidence:   intermediary.ReliabilityScore,
		}

		routes = append(routes, route)
	}

	return routes
}

// generateBridgeRoutes creates cross-chain bridge routes
func (pr *PaymentRouter) generateBridgeRoutes(ctx sdk.Context, paymentInfo types.PaymentInfo) []types.PaymentRoute {
	var routes []types.PaymentRoute

	if paymentInfo.CrossBorderInfo == nil {
		return routes
	}

	// CCTP bridge route
	cctpRoute := types.PaymentRoute{
		Type:          types.PaymentRouteBridge,
		Hops:         pr.generateCCTPHops(ctx, paymentInfo),
		EstimatedTime: time.Duration(5 * time.Minute), // Cross-chain finality
		Confidence:   0.90, // High confidence for established bridges
	}
	routes = append(routes, cctpRoute)

	// IBC bridge route
	ibcRoute := types.PaymentRoute{
		Type:          types.PaymentRouteBridge,
		Hops:         pr.generateIBCHops(ctx, paymentInfo),
		EstimatedTime: time.Duration(2 * time.Minute), // IBC finality
		Confidence:   0.85,
	}
	routes = append(routes, ibcRoute)

	return routes
}

// generateStablecoinRoutes creates routes using stablecoin conversions
func (pr *PaymentRouter) generateStablecoinRoutes(ctx sdk.Context, paymentInfo types.PaymentInfo) []types.PaymentRoute {
	var routes []types.PaymentRoute

	// Get available stablecoins
	stablecoins := pr.getAvailableStablecoins(ctx)

	for _, stablecoin := range stablecoins {
		// Check if conversion is beneficial
		if pr.isStablecoinConversionBeneficial(ctx, paymentInfo, stablecoin) {
			hops := []types.PaymentHop{
				{
					From:         "", // Source
					To:           stablecoin.IssuerAddress,
					Amount:       paymentInfo.Amount,
					Fee:          pr.calculateConversionFee(ctx, paymentInfo.Amount, stablecoin),
					Network:      "stateset",
					Protocol:     "stablecoin-swap",
					TimeEstimate: time.Duration(15 * time.Second),
				},
				{
					From:         stablecoin.IssuerAddress,
					To:           "", // Destination
					Amount:       pr.calculateStablecoinAmount(ctx, paymentInfo.Amount, stablecoin),
					Fee:          pr.calculateConversionFee(ctx, paymentInfo.Amount, stablecoin),
					Network:      "stateset",
					Protocol:     "stablecoin-swap",
					TimeEstimate: time.Duration(15 * time.Second),
				},
			}

			route := types.PaymentRoute{
				Type:          types.PaymentRouteOptimized,
				Hops:         hops,
				EstimatedTime: time.Duration(45 * time.Second),
				Confidence:   stablecoin.ReliabilityScore,
			}

			routes = append(routes, route)
		}
	}

	return routes
}

// RouteScore represents the scoring metrics for a route
type RouteScoreDetails struct {
	OverallScore     float64
	CostScore        float64
	TimeScore        float64
	ReliabilityScore float64
}

// scoreRoute scores a payment route based on cost, time, and reliability
func (pr *PaymentRouter) scoreRoute(ctx sdk.Context, route types.PaymentRoute, paymentInfo types.PaymentInfo) RouteScoreDetails {
	// Calculate cost score (lower cost = higher score)
	totalCost := pr.calculateRouteCost(ctx, route, paymentInfo)
	costScore := pr.calculateCostScore(totalCost, paymentInfo.Amount)

	// Calculate time score (faster = higher score)
	timeScore := pr.calculateTimeScore(route.EstimatedTime)

	// Reliability score is already in the route
	reliabilityScore := route.Confidence

	// Weight the scores
	costWeight := 0.4
	timeWeight := 0.3
	reliabilityWeight := 0.3

	overallScore := (costScore * costWeight) + 
	               (timeScore * timeWeight) + 
	               (reliabilityScore * reliabilityWeight)

	return RouteScoreDetails{
		OverallScore:     overallScore,
		CostScore:        costScore,
		TimeScore:        timeScore,
		ReliabilityScore: reliabilityScore,
	}
}

// calculateRouteCost calculates the total cost of a route
func (pr *PaymentRouter) calculateRouteCost(ctx sdk.Context, route types.PaymentRoute, paymentInfo types.PaymentInfo) sdk.Coins {
	totalCost := sdk.NewCoins()

	for _, hop := range route.Hops {
		totalCost = totalCost.Add(hop.Fee...)
	}

	// Add network fees based on route type
	switch route.Type {
	case types.PaymentRouteBridge:
		// Bridge fees are typically higher
		bridgeFee := sdk.NewCoins(sdk.NewCoin("ustate", sdk.NewInt(1000)))
		totalCost = totalCost.Add(bridgeFee...)
	case types.PaymentRouteOptimized:
		// Optimization fees
		optimizationFee := sdk.NewCoins(sdk.NewCoin("ustate", sdk.NewInt(500)))
		totalCost = totalCost.Add(optimizationFee...)
	}

	return totalCost
}

// calculateCostScore converts cost to a score (0-1, higher is better)
func (pr *PaymentRouter) calculateCostScore(cost sdk.Coins, amount sdk.Coins) float64 {
	if cost.IsZero() {
		return 1.0
	}

	// Calculate cost as percentage of amount
	if len(amount) == 0 || len(cost) == 0 {
		return 0.5 // Default score
	}

	costRatio := float64(cost[0].Amount.Int64()) / float64(amount[0].Amount.Int64())
	
	// Score inversely proportional to cost ratio
	// 0% cost = 1.0 score, 10% cost = 0.0 score
	score := math.Max(0, 1.0 - (costRatio * 10))
	return score
}

// calculateTimeScore converts time to a score (0-1, faster is better)
func (pr *PaymentRouter) calculateTimeScore(estimatedTime time.Duration) float64 {
	// Benchmark: 1 minute = 1.0 score, 10 minutes = 0.0 score
	minutes := estimatedTime.Minutes()
	score := math.Max(0, 1.0 - (minutes-1)/9)
	return score
}

// Supporting structures and methods

type Intermediary struct {
	Address           string
	ReliabilityScore  float64
	LiquidityScore    float64
	FeeRate          sdk.Dec
}

type StablecoinInfo struct {
	Denom             string
	IssuerAddress     string
	ReliabilityScore  float64
	ExchangeRate      sdk.Dec
}

// findIntermediaries finds potential payment intermediaries
func (pr *PaymentRouter) findIntermediaries(ctx sdk.Context, paymentInfo types.PaymentInfo) []Intermediary {
	// This would query a registry of trusted intermediaries
	// For now, return some mock data
	return []Intermediary{
		{
			Address:          "stateset1intermediary1...",
			ReliabilityScore: 0.95,
			LiquidityScore:   0.90,
			FeeRate:         sdk.NewDecWithPrec(1, 3), // 0.1%
		},
		{
			Address:          "stateset1intermediary2...",
			ReliabilityScore: 0.88,
			LiquidityScore:   0.85,
			FeeRate:         sdk.NewDecWithPrec(15, 4), // 0.15%
		},
	}
}

// calculateHopFee calculates the fee for a payment hop
func (pr *PaymentRouter) calculateHopFee(ctx sdk.Context, intermediary Intermediary, amount sdk.Coins) sdk.Coins {
	if len(amount) == 0 {
		return sdk.NewCoins()
	}

	feeAmount := amount[0].Amount.ToDec().Mul(intermediary.FeeRate).TruncateInt()
	return sdk.NewCoins(sdk.NewCoin(amount[0].Denom, feeAmount))
}

// generateCCTPHops generates hops for CCTP bridge transfers
func (pr *PaymentRouter) generateCCTPHops(ctx sdk.Context, paymentInfo types.PaymentInfo) []types.PaymentHop {
	return []types.PaymentHop{
		{
			From:         "", // Source address
			To:           "cctp-bridge",
			Amount:       paymentInfo.Amount,
			Fee:          sdk.NewCoins(sdk.NewCoin("ustate", sdk.NewInt(1000))),
			Network:      "ethereum",
			Protocol:     "cctp",
			TimeEstimate: time.Duration(5 * time.Minute),
		},
	}
}

// generateIBCHops generates hops for IBC transfers
func (pr *PaymentRouter) generateIBCHops(ctx sdk.Context, paymentInfo types.PaymentInfo) []types.PaymentHop {
	return []types.PaymentHop{
		{
			From:         "", // Source address
			To:           "ibc-channel",
			Amount:       paymentInfo.Amount,
			Fee:          sdk.NewCoins(sdk.NewCoin("ustate", sdk.NewInt(500))),
			Network:      "cosmos-hub",
			Protocol:     "ibc",
			TimeEstimate: time.Duration(2 * time.Minute),
		},
	}
}

// getAvailableStablecoins returns available stablecoins for routing
func (pr *PaymentRouter) getAvailableStablecoins(ctx sdk.Context) []StablecoinInfo {
	// This would query the stablecoins module
	return []StablecoinInfo{
		{
			Denom:            "usdx",
			IssuerAddress:    "stateset1usdx_issuer...",
			ReliabilityScore: 0.98,
			ExchangeRate:     sdk.OneDec(),
		},
		{
			Denom:            "eurx",
			IssuerAddress:    "stateset1eurx_issuer...",
			ReliabilityScore: 0.96,
			ExchangeRate:     sdk.NewDecWithPrec(85, 2), // 0.85 USD per EUR
		},
	}
}

// isStablecoinConversionBeneficial checks if using a stablecoin improves the route
func (pr *PaymentRouter) isStablecoinConversionBeneficial(ctx sdk.Context, paymentInfo types.PaymentInfo, stablecoin StablecoinInfo) bool {
	// Check if conversion reduces costs or improves reliability
	// This would involve complex calculations based on exchange rates, fees, etc.
	return true // Simplified for now
}

// calculateConversionFee calculates the fee for stablecoin conversion
func (pr *PaymentRouter) calculateConversionFee(ctx sdk.Context, amount sdk.Coins, stablecoin StablecoinInfo) sdk.Coins {
	if len(amount) == 0 {
		return sdk.NewCoins()
	}

	// 0.1% conversion fee
	feeRate := sdk.NewDecWithPrec(1, 3)
	feeAmount := amount[0].Amount.ToDec().Mul(feeRate).TruncateInt()
	return sdk.NewCoins(sdk.NewCoin("ustate", feeAmount))
}

// calculateStablecoinAmount calculates the equivalent amount in stablecoin
func (pr *PaymentRouter) calculateStablecoinAmount(ctx sdk.Context, amount sdk.Coins, stablecoin StablecoinInfo) sdk.Coins {
	if len(amount) == 0 {
		return sdk.NewCoins()
	}

	convertedAmount := amount[0].Amount.ToDec().Mul(stablecoin.ExchangeRate).TruncateInt()
	return sdk.NewCoins(sdk.NewCoin(stablecoin.Denom, convertedAmount))
}