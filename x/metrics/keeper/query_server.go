package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/metrics/types"
)

var _ types.QueryServer = Keeper{}

// SystemMetrics returns the current system metrics
func (k Keeper) SystemMetrics(goCtx context.Context, req *types.QuerySystemMetricsRequest) (*types.QuerySystemMetricsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	metrics := k.GetSystemMetrics(ctx)

	return &types.QuerySystemMetricsResponse{
		Metrics: &metrics,
	}, nil
}

// Counter returns a specific counter value
func (k Keeper) Counter(goCtx context.Context, req *types.QueryCounterRequest) (*types.QueryCounterResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	counter := k.GetCounter(ctx, req.Name, req.Labels)

	return &types.QueryCounterResponse{
		Counter: &counter,
	}, nil
}

// Gauge returns a specific gauge value
func (k Keeper) Gauge(goCtx context.Context, req *types.QueryGaugeRequest) (*types.QueryGaugeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	gauge := k.GetGauge(ctx, req.Name, req.Labels)

	return &types.QueryGaugeResponse{
		Gauge: &gauge,
	}, nil
}

// ModuleHealth returns health status for a specific module
func (k Keeper) ModuleHealth(goCtx context.Context, req *types.QueryModuleHealthRequest) (*types.QueryModuleHealthResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	metrics := k.GetSystemMetrics(ctx)

	if metrics.ModuleHealth == nil {
		return &types.QueryModuleHealthResponse{}, nil
	}

	if req.Module != "" {
		health, exists := metrics.ModuleHealth[req.Module]
		if !exists {
			return nil, fmt.Errorf("module %s not found", req.Module)
		}
		return &types.QueryModuleHealthResponse{
			Health: map[string]types.ModuleHealth{req.Module: health},
		}, nil
	}

	return &types.QueryModuleHealthResponse{
		Health: metrics.ModuleHealth,
	}, nil
}

// Alerts returns all triggered alerts or active alert configs
func (k Keeper) Alerts(goCtx context.Context, req *types.QueryAlertsRequest) (*types.QueryAlertsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if req.Active {
		// Return currently triggered alerts
		alerts := k.CheckAlerts(ctx)
		return &types.QueryAlertsResponse{
			Alerts: alerts,
		}, nil
	}

	// Return alert configurations
	configs := k.GetAlertConfigs(ctx)
	return &types.QueryAlertsResponse{
		AlertConfigs: configs,
	}, nil
}

// PrometheusMetrics exports metrics in Prometheus format
func (k Keeper) PrometheusMetrics(goCtx context.Context, req *types.QueryPrometheusMetricsRequest) (*types.QueryPrometheusMetricsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	metrics := k.GetSystemMetrics(ctx)

	var output strings.Builder

	// System metrics
	output.WriteString("# HELP stateset_block_height Current block height\n")
	output.WriteString("# TYPE stateset_block_height gauge\n")
	output.WriteString(fmt.Sprintf("stateset_block_height %d\n", metrics.LastBlockHeight))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_average_block_time Average block time in seconds\n")
	output.WriteString("# TYPE stateset_average_block_time gauge\n")
	output.WriteString(fmt.Sprintf("stateset_average_block_time %.3f\n", metrics.AverageBlockTime.Seconds()))
	output.WriteString("\n")

	// Settlement metrics
	output.WriteString("# HELP stateset_total_settlements Total number of settlements\n")
	output.WriteString("# TYPE stateset_total_settlements counter\n")
	output.WriteString(fmt.Sprintf("stateset_total_settlements %d\n", metrics.TotalSettlements))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_total_settlement_volume Total settlement volume\n")
	output.WriteString("# TYPE stateset_total_settlement_volume counter\n")
	output.WriteString(fmt.Sprintf("stateset_total_settlement_volume %s\n", metrics.TotalSettlementVolume.String()))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_active_escrows Number of active escrows\n")
	output.WriteString("# TYPE stateset_active_escrows gauge\n")
	output.WriteString(fmt.Sprintf("stateset_active_escrows %d\n", metrics.ActiveEscrows))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_active_channels Number of active payment channels\n")
	output.WriteString("# TYPE stateset_active_channels gauge\n")
	output.WriteString(fmt.Sprintf("stateset_active_channels %d\n", metrics.ActiveChannels))
	output.WriteString("\n")

	// Economic metrics
	output.WriteString("# HELP stateset_total_collateral_value Total collateral value\n")
	output.WriteString("# TYPE stateset_total_collateral_value gauge\n")
	output.WriteString(fmt.Sprintf("stateset_total_collateral_value %s\n", metrics.TotalCollateralValue.String()))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_total_debt_value Total debt value\n")
	output.WriteString("# TYPE stateset_total_debt_value gauge\n")
	output.WriteString(fmt.Sprintf("stateset_total_debt_value %s\n", metrics.TotalDebtValue.String()))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_system_collateral_ratio System-wide collateral ratio\n")
	output.WriteString("# TYPE stateset_system_collateral_ratio gauge\n")
	output.WriteString(fmt.Sprintf("stateset_system_collateral_ratio %s\n", metrics.SystemCollateralRatio.String()))
	output.WriteString("\n")

	// Security metrics
	output.WriteString("# HELP stateset_circuit_trips Total circuit breaker trips\n")
	output.WriteString("# TYPE stateset_circuit_trips counter\n")
	output.WriteString(fmt.Sprintf("stateset_circuit_trips %d\n", metrics.CircuitTrips))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_rate_limit_hits Total rate limit hits\n")
	output.WriteString("# TYPE stateset_rate_limit_hits counter\n")
	output.WriteString(fmt.Sprintf("stateset_rate_limit_hits %d\n", metrics.RateLimitHits))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_compliance_blocks Total compliance blocks\n")
	output.WriteString("# TYPE stateset_compliance_blocks counter\n")
	output.WriteString(fmt.Sprintf("stateset_compliance_blocks %d\n", metrics.ComplianceBlocks))
	output.WriteString("\n")

	// Oracle metrics
	output.WriteString("# HELP stateset_prices_updated Total price updates\n")
	output.WriteString("# TYPE stateset_prices_updated counter\n")
	output.WriteString(fmt.Sprintf("stateset_prices_updated %d\n", metrics.PricesUpdated))
	output.WriteString("\n")

	output.WriteString("# HELP stateset_stale_price_count Number of stale prices\n")
	output.WriteString("# TYPE stateset_stale_price_count gauge\n")
	output.WriteString(fmt.Sprintf("stateset_stale_price_count %d\n", metrics.StalePriceCount))
	output.WriteString("\n")

	// Module health
	if metrics.ModuleHealth != nil {
		output.WriteString("# HELP stateset_module_error_rate Module error rate\n")
		output.WriteString("# TYPE stateset_module_error_rate gauge\n")
		for module, health := range metrics.ModuleHealth {
			output.WriteString(fmt.Sprintf("stateset_module_error_rate{module=\"%s\"} %.6f\n", module, health.ErrorRate))
		}
		output.WriteString("\n")

		output.WriteString("# HELP stateset_module_latency Module average latency in milliseconds\n")
		output.WriteString("# TYPE stateset_module_latency gauge\n")
		for module, health := range metrics.ModuleHealth {
			output.WriteString(fmt.Sprintf("stateset_module_latency{module=\"%s\"} %.2f\n", module, health.Latency))
		}
		output.WriteString("\n")

		output.WriteString("# HELP stateset_module_transactions Total transactions per module\n")
		output.WriteString("# TYPE stateset_module_transactions counter\n")
		for module, health := range metrics.ModuleHealth {
			output.WriteString(fmt.Sprintf("stateset_module_transactions{module=\"%s\"} %d\n", module, health.Transactions))
		}
		output.WriteString("\n")

		output.WriteString("# HELP stateset_module_healthy Module health status (1=healthy, 0=unhealthy)\n")
		output.WriteString("# TYPE stateset_module_healthy gauge\n")
		for module, health := range metrics.ModuleHealth {
			healthValue := 0
			if health.Status == "healthy" {
				healthValue = 1
			}
			output.WriteString(fmt.Sprintf("stateset_module_healthy{module=\"%s\",status=\"%s\"} %d\n", module, health.Status, healthValue))
		}
		output.WriteString("\n")
	}

	return &types.QueryPrometheusMetricsResponse{
		Metrics: output.String(),
	}, nil
}
