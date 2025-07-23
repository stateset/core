package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PerformanceMonitor provides comprehensive blockchain performance monitoring
type PerformanceMonitor struct {
	mu                sync.RWMutex
	metrics          *Metrics
	alerts           []Alert
	healthStatus     HealthStatus
	recommendations  []Recommendation
	startTime        time.Time
	lastBlockTime    time.Time
	config          *MonitorConfig
}

// Metrics contains all blockchain performance metrics
type Metrics struct {
	// Block metrics
	BlockHeight          int64         `json:"block_height"`
	BlockTime           time.Duration `json:"block_time"`
	AverageBlockTime    time.Duration `json:"average_block_time"`
	BlockSize           int64         `json:"block_size"`
	AverageBlockSize    int64         `json:"average_block_size"`
	
	// Transaction metrics
	TxCount             int64         `json:"tx_count"`
	TxPerSecond         float64       `json:"tx_per_second"`
	AverageTxPerBlock   float64       `json:"average_tx_per_block"`
	TxSuccessRate       float64       `json:"tx_success_rate"`
	AverageTxFee        sdk.Coins     `json:"average_tx_fee"`
	
	// Network metrics
	PeerCount           int           `json:"peer_count"`
	ValidatorCount      int           `json:"validator_count"`
	ActiveValidators    int           `json:"active_validators"`
	NetworkLatency      time.Duration `json:"network_latency"`
	
	// Resource metrics
	CPUUsage            float64       `json:"cpu_usage"`
	MemoryUsage         int64         `json:"memory_usage"`
	DiskUsage           int64         `json:"disk_usage"`
	NetworkIO           NetworkIO     `json:"network_io"`
	
	// Business metrics
	TotalAgreements     int64         `json:"total_agreements"`
	TotalInvoices       int64         `json:"total_invoices"`
	TotalLoans          int64         `json:"total_loans"`
	TotalPurchaseOrders int64         `json:"total_purchase_orders"`
	TotalValue          sdk.Coins     `json:"total_value"`
	
	// Security metrics
	SecurityAlerts      int64         `json:"security_alerts"`
	HighRiskTxCount     int64         `json:"high_risk_tx_count"`
	ComplianceScore     float64       `json:"compliance_score"`
	
	// Performance indicators
	ThroughputScore     float64       `json:"throughput_score"`
	LatencyScore        float64       `json:"latency_score"`
	StabilityScore      float64       `json:"stability_score"`
	OverallScore        float64       `json:"overall_score"`
	
	LastUpdated         time.Time     `json:"last_updated"`
}

// NetworkIO represents network input/output metrics
type NetworkIO struct {
	BytesIn    int64 `json:"bytes_in"`
	BytesOut   int64 `json:"bytes_out"`
	PacketsIn  int64 `json:"packets_in"`
	PacketsOut int64 `json:"packets_out"`
}

// Alert represents a performance alert
type Alert struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Severity    string    `json:"severity"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Metric      string    `json:"metric"`
	Value       float64   `json:"value"`
	Threshold   float64   `json:"threshold"`
	CreatedAt   time.Time `json:"created_at"`
	Resolved    bool      `json:"resolved"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
}

// HealthStatus represents the overall health of the blockchain
type HealthStatus struct {
	Status           string    `json:"status"`           // healthy, warning, critical
	OverallScore     float64   `json:"overall_score"`    // 0-100
	ComponentScores  map[string]float64 `json:"component_scores"`
	LastChecked      time.Time `json:"last_checked"`
	Uptime           time.Duration `json:"uptime"`
	Issues           []string  `json:"issues,omitempty"`
}

// Recommendation represents a performance optimization recommendation
type Recommendation struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Priority    string    `json:"priority"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Impact      string    `json:"impact"`
	Effort      string    `json:"effort"`
	Actions     []string  `json:"actions"`
	CreatedAt   time.Time `json:"created_at"`
	Applied     bool      `json:"applied"`
}

// MonitorConfig contains configuration for the performance monitor
type MonitorConfig struct {
	MetricsInterval     time.Duration `json:"metrics_interval"`
	AlertThresholds     map[string]float64 `json:"alert_thresholds"`
	EnablePrometheus    bool          `json:"enable_prometheus"`
	PrometheusPort      int           `json:"prometheus_port"`
	HTTPPort            int           `json:"http_port"`
	LogLevel            string        `json:"log_level"`
}

// Prometheus metrics
var (
	blockHeightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_block_height",
		Help: "Current block height",
	})
	
	blockTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_block_time_seconds",
		Help: "Block time in seconds",
	})
	
	txCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_tx_count_total",
		Help: "Total transaction count",
	})
	
	txPerSecondGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_tx_per_second",
		Help: "Transactions per second",
	})
	
	cpuUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_cpu_usage_percent",
		Help: "CPU usage percentage",
	})
	
	memoryUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_memory_usage_bytes",
		Help: "Memory usage in bytes",
	})
	
	performanceScoreGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "stateset_performance_score",
		Help: "Overall performance score (0-100)",
	})
)

// NewPerformanceMonitor creates a new performance monitor
func NewPerformanceMonitor(config *MonitorConfig) *PerformanceMonitor {
	if config == nil {
		config = &MonitorConfig{
			MetricsInterval:  time.Second * 30,
			EnablePrometheus: true,
			PrometheusPort:  9090,
			HTTPPort:        8080,
			LogLevel:        "info",
			AlertThresholds: map[string]float64{
				"block_time":      10.0,  // seconds
				"cpu_usage":       80.0,  // percent
				"memory_usage":    85.0,  // percent
				"tx_success_rate": 95.0,  // percent
				"performance_score": 70.0, // minimum score
			},
		}
	}

	pm := &PerformanceMonitor{
		metrics:         &Metrics{},
		alerts:          make([]Alert, 0),
		startTime:       time.Now(),
		config:          config,
		recommendations: make([]Recommendation, 0),
		healthStatus: HealthStatus{
			Status:          "healthy",
			ComponentScores: make(map[string]float64),
		},
	}

	// Register Prometheus metrics
	if config.EnablePrometheus {
		prometheus.MustRegister(
			blockHeightGauge,
			blockTimeGauge,
			txCountGauge,
			txPerSecondGauge,
			cpuUsageGauge,
			memoryUsageGauge,
			performanceScoreGauge,
		)
	}

	return pm
}

// Start begins the performance monitoring
func (pm *PerformanceMonitor) Start(ctx context.Context) error {
	log.Printf("Starting Performance Monitor...")

	// Start metrics collection
	go pm.metricsCollectionLoop(ctx)

	// Start health monitoring
	go pm.healthMonitoringLoop(ctx)

	// Start alert processing
	go pm.alertProcessingLoop(ctx)

	// Start recommendation engine
	go pm.recommendationEngineLoop(ctx)

	// Start HTTP server for metrics API
	go pm.startHTTPServer()

	// Start Prometheus server if enabled
	if pm.config.EnablePrometheus {
		go pm.startPrometheusServer()
	}

	log.Printf("Performance Monitor started successfully")
	return nil
}

// metricsCollectionLoop continuously collects performance metrics
func (pm *PerformanceMonitor) metricsCollectionLoop(ctx context.Context) {
	ticker := time.NewTicker(pm.config.MetricsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pm.collectMetrics()
		}
	}
}

// collectMetrics collects all performance metrics
func (pm *PerformanceMonitor) collectMetrics() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Collect system metrics
	pm.collectSystemMetrics()
	
	// Collect blockchain metrics
	pm.collectBlockchainMetrics()
	
	// Collect business metrics
	pm.collectBusinessMetrics()
	
	// Calculate performance scores
	pm.calculatePerformanceScores()
	
	// Update Prometheus metrics
	pm.updatePrometheusMetrics()
	
	pm.metrics.LastUpdated = time.Now()
}

// collectSystemMetrics collects system resource metrics
func (pm *PerformanceMonitor) collectSystemMetrics() {
	// Get CPU usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	pm.metrics.CPUUsage = pm.getCPUUsage()
	pm.metrics.MemoryUsage = int64(m.Alloc)
	pm.metrics.DiskUsage = pm.getDiskUsage()
	pm.metrics.NetworkIO = pm.getNetworkIO()
}

// collectBlockchainMetrics collects blockchain-specific metrics
func (pm *PerformanceMonitor) collectBlockchainMetrics() {
	// This would integrate with the actual blockchain state
	// For demonstration, using mock values
	
	currentTime := time.Now()
	if !pm.lastBlockTime.IsZero() {
		pm.metrics.BlockTime = currentTime.Sub(pm.lastBlockTime)
	}
	pm.lastBlockTime = currentTime
	
	// Calculate average block time (simplified)
	pm.metrics.AverageBlockTime = time.Second * 7 // Example: 7 seconds
	
	// Calculate TPS (simplified)
	if pm.metrics.BlockTime > 0 {
		pm.metrics.TxPerSecond = float64(pm.metrics.TxCount) / pm.metrics.BlockTime.Seconds()
	}
}

// collectBusinessMetrics collects business logic metrics
func (pm *PerformanceMonitor) collectBusinessMetrics() {
	// This would query the actual modules for their metrics
	// For demonstration, using incremental values
	
	pm.metrics.TotalAgreements++
	pm.metrics.TotalInvoices++
	pm.metrics.TotalLoans++
	pm.metrics.TotalPurchaseOrders++
}

// calculatePerformanceScores calculates various performance scores
func (pm *PerformanceMonitor) calculatePerformanceScores() {
	// Throughput score (based on TPS and block time)
	maxTPS := 1000.0 // theoretical maximum
	pm.metrics.ThroughputScore = math.Min(100, (pm.metrics.TxPerSecond/maxTPS)*100)
	
	// Latency score (based on block time)
	targetBlockTime := 6.0 // seconds
	latencyRatio := float64(pm.metrics.BlockTime.Seconds()) / targetBlockTime
	pm.metrics.LatencyScore = math.Max(0, 100-(latencyRatio-1)*50)
	
	// Stability score (based on resource usage)
	cpuPenalty := math.Max(0, pm.metrics.CPUUsage-70) * 0.5
	memoryPenalty := math.Max(0, float64(pm.metrics.MemoryUsage)/(1024*1024*1024)-2) * 10
	pm.metrics.StabilityScore = math.Max(0, 100-cpuPenalty-memoryPenalty)
	
	// Overall score (weighted average)
	pm.metrics.OverallScore = (pm.metrics.ThroughputScore*0.3 + 
		pm.metrics.LatencyScore*0.3 + 
		pm.metrics.StabilityScore*0.4)
}

// updatePrometheusMetrics updates Prometheus metrics
func (pm *PerformanceMonitor) updatePrometheusMetrics() {
	if !pm.config.EnablePrometheus {
		return
	}
	
	blockHeightGauge.Set(float64(pm.metrics.BlockHeight))
	blockTimeGauge.Set(pm.metrics.BlockTime.Seconds())
	txCountGauge.Set(float64(pm.metrics.TxCount))
	txPerSecondGauge.Set(pm.metrics.TxPerSecond)
	cpuUsageGauge.Set(pm.metrics.CPUUsage)
	memoryUsageGauge.Set(float64(pm.metrics.MemoryUsage))
	performanceScoreGauge.Set(pm.metrics.OverallScore)
}

// healthMonitoringLoop continuously monitors system health
func (pm *PerformanceMonitor) healthMonitoringLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pm.checkSystemHealth()
		}
	}
}

// checkSystemHealth evaluates overall system health
func (pm *PerformanceMonitor) checkSystemHealth() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.healthStatus.LastChecked = time.Now()
	pm.healthStatus.Uptime = time.Since(pm.startTime)
	pm.healthStatus.Issues = make([]string, 0)

	// Component health scores
	pm.healthStatus.ComponentScores["throughput"] = pm.metrics.ThroughputScore
	pm.healthStatus.ComponentScores["latency"] = pm.metrics.LatencyScore
	pm.healthStatus.ComponentScores["stability"] = pm.metrics.StabilityScore
	pm.healthStatus.ComponentScores["security"] = pm.metrics.ComplianceScore

	// Calculate overall health
	totalScore := 0.0
	for _, score := range pm.healthStatus.ComponentScores {
		totalScore += score
	}
	pm.healthStatus.OverallScore = totalScore / float64(len(pm.healthStatus.ComponentScores))

	// Determine status
	if pm.healthStatus.OverallScore >= 80 {
		pm.healthStatus.Status = "healthy"
	} else if pm.healthStatus.OverallScore >= 60 {
		pm.healthStatus.Status = "warning"
		pm.healthStatus.Issues = append(pm.healthStatus.Issues, "Performance below optimal")
	} else {
		pm.healthStatus.Status = "critical"
		pm.healthStatus.Issues = append(pm.healthStatus.Issues, "Critical performance issues detected")
	}

	// Check individual thresholds
	if pm.metrics.CPUUsage > pm.config.AlertThresholds["cpu_usage"] {
		pm.healthStatus.Issues = append(pm.healthStatus.Issues, "High CPU usage")
	}
	if float64(pm.metrics.MemoryUsage)/(1024*1024*1024) > pm.config.AlertThresholds["memory_usage"] {
		pm.healthStatus.Issues = append(pm.healthStatus.Issues, "High memory usage")
	}
}

// alertProcessingLoop processes and manages alerts
func (pm *PerformanceMonitor) alertProcessingLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pm.processAlerts()
		}
	}
}

// processAlerts checks conditions and generates alerts
func (pm *PerformanceMonitor) processAlerts() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Check CPU usage
	if pm.metrics.CPUUsage > pm.config.AlertThresholds["cpu_usage"] {
		pm.createAlert("cpu_usage", "warning", "High CPU Usage", 
			fmt.Sprintf("CPU usage is %.2f%%, above threshold of %.2f%%", 
				pm.metrics.CPUUsage, pm.config.AlertThresholds["cpu_usage"]),
			pm.metrics.CPUUsage, pm.config.AlertThresholds["cpu_usage"])
	}

	// Check performance score
	if pm.metrics.OverallScore < pm.config.AlertThresholds["performance_score"] {
		pm.createAlert("performance_score", "critical", "Low Performance Score",
			fmt.Sprintf("Overall performance score is %.2f, below threshold of %.2f",
				pm.metrics.OverallScore, pm.config.AlertThresholds["performance_score"]),
			pm.metrics.OverallScore, pm.config.AlertThresholds["performance_score"])
	}

	// Auto-resolve old alerts
	for i := range pm.alerts {
		if !pm.alerts[i].Resolved && time.Since(pm.alerts[i].CreatedAt) > time.Hour {
			pm.alerts[i].Resolved = true
			now := time.Now()
			pm.alerts[i].ResolvedAt = &now
		}
	}
}

// createAlert creates a new alert if one doesn't already exist
func (pm *PerformanceMonitor) createAlert(alertType, severity, title, description string, value, threshold float64) {
	// Check if alert already exists
	for _, alert := range pm.alerts {
		if alert.Type == alertType && !alert.Resolved {
			return // Alert already exists
		}
	}

	alert := Alert{
		ID:          fmt.Sprintf("%s-%d", alertType, time.Now().Unix()),
		Type:        alertType,
		Severity:    severity,
		Title:       title,
		Description: description,
		Metric:      alertType,
		Value:       value,
		Threshold:   threshold,
		CreatedAt:   time.Now(),
		Resolved:    false,
	}

	pm.alerts = append(pm.alerts, alert)
	log.Printf("Alert created: %s - %s", alert.Severity, alert.Title)
}

// recommendationEngineLoop generates performance recommendations
func (pm *PerformanceMonitor) recommendationEngineLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Hour) // Check hourly
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pm.generateRecommendations()
		}
	}
}

// generateRecommendations analyzes metrics and generates optimization recommendations
func (pm *PerformanceMonitor) generateRecommendations() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Clear old recommendations
	pm.recommendations = make([]Recommendation, 0)

	// High CPU usage recommendation
	if pm.metrics.CPUUsage > 80 {
		pm.recommendations = append(pm.recommendations, Recommendation{
			ID:          "cpu-optimization",
			Type:        "performance",
			Priority:    "high",
			Title:       "Optimize CPU Usage",
			Description: "CPU usage is consistently high, consider optimizing transaction processing",
			Impact:      "high",
			Effort:      "medium",
			Actions: []string{
				"Review transaction validation logic",
				"Implement batch processing",
				"Consider hardware upgrade",
			},
			CreatedAt: time.Now(),
		})
	}

	// Low throughput recommendation
	if pm.metrics.ThroughputScore < 50 {
		pm.recommendations = append(pm.recommendations, Recommendation{
			ID:          "throughput-improvement",
			Type:        "performance",
			Priority:    "high",
			Title:       "Improve Transaction Throughput",
			Description: "Transaction throughput is below optimal levels",
			Impact:      "high",
			Effort:      "high",
			Actions: []string{
				"Optimize consensus parameters",
				"Implement parallel transaction processing",
				"Review gas limits and fees",
			},
			CreatedAt: time.Now(),
		})
	}

	// Security improvement recommendation
	if pm.metrics.SecurityAlerts > 10 {
		pm.recommendations = append(pm.recommendations, Recommendation{
			ID:          "security-enhancement",
			Type:        "security",
			Priority:    "critical",
			Title:       "Address Security Concerns",
			Description: "Multiple security alerts detected",
			Impact:      "critical",
			Effort:      "medium",
			Actions: []string{
				"Review security rules",
				"Implement stricter validation",
				"Audit suspicious transactions",
			},
			CreatedAt: time.Now(),
		})
	}
}

// HTTP API endpoints
func (pm *PerformanceMonitor) startHTTPServer() {
	r := mux.NewRouter()
	
	r.HandleFunc("/metrics", pm.handleMetrics).Methods("GET")
	r.HandleFunc("/health", pm.handleHealth).Methods("GET")
	r.HandleFunc("/alerts", pm.handleAlerts).Methods("GET")
	r.HandleFunc("/recommendations", pm.handleRecommendations).Methods("GET")
	r.HandleFunc("/status", pm.handleStatus).Methods("GET")

	log.Printf("Starting HTTP server on port %d", pm.config.HTTPPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", pm.config.HTTPPort), r))
}

func (pm *PerformanceMonitor) handleMetrics(w http.ResponseWriter, r *http.Request) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pm.metrics)
}

func (pm *PerformanceMonitor) handleHealth(w http.ResponseWriter, r *http.Request) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pm.healthStatus)
}

func (pm *PerformanceMonitor) handleAlerts(w http.ResponseWriter, r *http.Request) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pm.alerts)
}

func (pm *PerformanceMonitor) handleRecommendations(w http.ResponseWriter, r *http.Request) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pm.recommendations)
}

func (pm *PerformanceMonitor) handleStatus(w http.ResponseWriter, r *http.Request) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	status := map[string]interface{}{
		"health":          pm.healthStatus,
		"metrics_summary": map[string]interface{}{
			"overall_score":   pm.metrics.OverallScore,
			"throughput":      pm.metrics.TxPerSecond,
			"block_time":      pm.metrics.BlockTime.Seconds(),
			"cpu_usage":       pm.metrics.CPUUsage,
		},
		"active_alerts":   len(pm.getActiveAlerts()),
		"recommendations": len(pm.recommendations),
		"uptime":          time.Since(pm.startTime).String(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (pm *PerformanceMonitor) startPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting Prometheus server on port %d", pm.config.PrometheusPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", pm.config.PrometheusPort), nil))
}

// Helper functions
func (pm *PerformanceMonitor) getActiveAlerts() []Alert {
	var active []Alert
	for _, alert := range pm.alerts {
		if !alert.Resolved {
			active = append(active, alert)
		}
	}
	return active
}

func (pm *PerformanceMonitor) getCPUUsage() float64 {
	// Simplified CPU usage calculation
	// In production, this would use proper system monitoring
	return float64(runtime.NumGoroutine()) / 100.0 * 20.0
}

func (pm *PerformanceMonitor) getDiskUsage() int64 {
	// Simplified disk usage calculation
	return 1024 * 1024 * 1024 * 50 // 50GB
}

func (pm *PerformanceMonitor) getNetworkIO() NetworkIO {
	// Simplified network I/O calculation
	return NetworkIO{
		BytesIn:    1024 * 1024 * 100, // 100MB
		BytesOut:   1024 * 1024 * 80,  // 80MB
		PacketsIn:  10000,
		PacketsOut: 8000,
	}
}

// GetMetrics returns current metrics (thread-safe)
func (pm *PerformanceMonitor) GetMetrics() Metrics {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return *pm.metrics
}

// GetHealthStatus returns current health status (thread-safe)
func (pm *PerformanceMonitor) GetHealthStatus() HealthStatus {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.healthStatus
}