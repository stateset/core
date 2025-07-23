package types

import (
	"time"
)

const (
	// ModuleName defines the module name
	ModuleName = "analytics"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for analytics
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_analytics"

	// MetricKeyPrefix is the prefix to retrieve all metrics
	MetricKeyPrefix = "metric"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// PerformanceMetric represents a blockchain performance metric
type PerformanceMetric struct {
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
	BlockHeight int64   `json:"block_height"`
	Tags      map[string]string `json:"tags"`
}

// BlockMetrics contains metrics for a specific block
type BlockMetrics struct {
	Height           int64   `json:"height"`
	TransactionCount int32   `json:"transaction_count"`
	BlockSize        int64   `json:"block_size"`
	BlockTime        float64 `json:"block_time"`
	GasUsed          int64   `json:"gas_used"`
	GasLimit         int64   `json:"gas_limit"`
	ValidatorCount   int32   `json:"validator_count"`
}

// SystemMetrics contains system-level metrics
type SystemMetrics struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	NetworkIO   NetworkIOMetrics `json:"network_io"`
}

// NetworkIOMetrics contains network I/O metrics
type NetworkIOMetrics struct {
	BytesIn  int64 `json:"bytes_in"`
	BytesOut int64 `json:"bytes_out"`
}

// BusinessMetrics contains business-specific metrics
type BusinessMetrics struct {
	TotalInvoices     int64   `json:"total_invoices"`
	TotalAgreements   int64   `json:"total_agreements"`
	TotalLoanValue    float64 `json:"total_loan_value"`
	ActivePurchaseOrders int64 `json:"active_purchase_orders"`
	SecurityAlerts    int32   `json:"security_alerts"`
}