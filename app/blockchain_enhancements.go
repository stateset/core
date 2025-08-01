package app

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	tmcfg "github.com/cometbft/cometbft/config"
)

// BlockchainEnhancements contains all performance optimizations
type BlockchainEnhancements struct {
	StateCache         *CacheWrapper
	PerformanceMonitor *PerformanceMonitor
	BatchProcessor     *BatchProcessor
	MemoryPool         *MemoryPool
	Config             *EnhancedConfig
}

// EnhancedConfig contains optimized blockchain configuration
type EnhancedConfig struct {
	// Consensus settings
	BlockTimeTarget       time.Duration
	MaxBlockSize          int64
	MaxGasPerBlock        int64
	MaxTxSize             int64
	MaxTxGasWanted        uint64
	
	// Cache settings
	StateCacheSize        int
	StateCacheTTL         time.Duration
	QueryCacheSize        int
	
	// Batch processing settings
	BatchProcessingEnabled bool
	MaxBatchSize          int
	BatchTimeout          time.Duration
	BatchWorkers          int
	
	// Mempool settings
	MempoolSize           int
	MempoolCacheSize      int
	MinGasPrice           sdk.DecCoins
	
	// Monitoring settings
	MetricsEnabled        bool
	PrometheusPort        string
	
	// Network settings
	MaxPeers              int
	PeerGossipInterval    time.Duration
	
	// Validation settings
	SignatureVerification bool
	TxValidationCache     bool
}

// DefaultEnhancedConfig returns a default enhanced configuration
func DefaultEnhancedConfig() *EnhancedConfig {
	return &EnhancedConfig{
		// Consensus settings
		BlockTimeTarget:       2 * time.Second,
		MaxBlockSize:          10 * 1024 * 1024, // 10MB
		MaxGasPerBlock:        50_000_000,       // 50M gas units
		MaxTxSize:             1024 * 1024,      // 1MB
		MaxTxGasWanted:        10_000_000,       // 10M gas per tx
		
		// Cache settings
		StateCacheSize:        10000,
		StateCacheTTL:         5 * time.Minute,
		QueryCacheSize:        5000,
		
		// Batch processing settings
		BatchProcessingEnabled: true,
		MaxBatchSize:          100,
		BatchTimeout:          100 * time.Millisecond,
		BatchWorkers:          4,
		
		// Mempool settings
		MempoolSize:           10000,
		MempoolCacheSize:      100000,
		MinGasPrice:           sdk.NewDecCoins(sdk.NewDecCoin("state", sdk.NewInt(1000))),
		
		// Monitoring settings
		MetricsEnabled:        true,
		PrometheusPort:        ":26660",
		
		// Network settings
		MaxPeers:              100,
		PeerGossipInterval:    50 * time.Millisecond,
		
		// Validation settings
		SignatureVerification: true,
		TxValidationCache:     true,
	}
}

// InitializeBlockchainEnhancements sets up all performance optimizations
func InitializeBlockchainEnhancements(config *EnhancedConfig) (*BlockchainEnhancements, error) {
	// Initialize state cache
	cacheWrapper := NewCacheWrapper()
	
	// Initialize performance monitor
	perfMonitor := NewPerformanceMonitor(config.MetricsEnabled)
	
	// Initialize batch processor (placeholder - needs actual handlers)
	var batchProcessor *BatchProcessor
	if config.BatchProcessingEnabled {
		batchProcessor = NewBatchProcessor(
			nil, // ante handler will be set later
			nil, // deliver handler will be set later
			config.MaxBatchSize,
			config.BatchTimeout,
			config.BatchWorkers,
		)
	}
	
	// Initialize memory pool
	memoryPool := NewMemoryPool(
		config.MempoolSize,
		config.MaxTxSize,
		config.MinGasPrice,
	)
	
	return &BlockchainEnhancements{
		StateCache:         cacheWrapper,
		PerformanceMonitor: perfMonitor,
		BatchProcessor:     batchProcessor,
		MemoryPool:         memoryPool,
		Config:             config,
	}, nil
}

// SetupEnhancedAnteHandler creates an optimized ante handler
func (be *BlockchainEnhancements) SetupEnhancedAnteHandler(
	accountKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	feegrantKeeper feegrantkeeper.Keeper,
	govKeeper govkeeper.Keeper,
	signModeHandler sdk.SignModeHandler,
	ibcKeeper *ibckeeper.Keeper,
	paramsKeeper paramtypes.Keeper,
) (sdk.AnteHandler, error) {
	
	handlerOptions := HandlerOptions{
		AccountKeeper:    accountKeeper,
		BankKeeper:       bankKeeper,
		FeegrantKeeper:   feegrantKeeper,
		GovKeeper:        govKeeper,
		SignModeHandler:  signModeHandler,
		SigGasConsumer:   ante.DefaultSigVerificationGasConsumer,
		IBCKeeper:        ibcKeeper,
		ParamsKeeper:     paramsKeeper,
		MaxTxGasWanted:   be.Config.MaxTxGasWanted,
		MaxTxSizeBytes:   uint64(be.Config.MaxTxSize),
		MinGasPriceCoins: be.Config.MinGasPrice,
	}
	
	return NewEnhancedAnteHandler(handlerOptions)
}

// GetOptimizedTendermintConfig returns Tendermint config with optimizations
func (be *BlockchainEnhancements) GetOptimizedTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()
	
	// Apply optimizations from our config
	cfg.Consensus.TimeoutCommit = be.Config.BlockTimeTarget
	cfg.Mempool.Size = be.Config.MempoolSize
	cfg.Mempool.CacheSize = be.Config.MempoolCacheSize
	cfg.Mempool.MaxTxsBytes = be.Config.MaxBlockSize
	cfg.P2P.MaxNumInboundPeers = be.Config.MaxPeers / 2
	cfg.P2P.MaxNumOutboundPeers = be.Config.MaxPeers / 2
	
	// Enable instrumentation
	if be.Config.MetricsEnabled {
		cfg.Instrumentation.Prometheus = true
		cfg.Instrumentation.PrometheusListenAddr = be.Config.PrometheusPort
	}
	
	return cfg
}

// ApplyPerformanceOptimizations applies all performance optimizations to the app
func (app *App) ApplyPerformanceOptimizations(enhancements *BlockchainEnhancements) error {
	// Store enhancements reference
	if app.enhancements == nil {
		app.enhancements = enhancements
	}
	
	// Set up caching for each module store
	modules := []string{"auth", "bank", "staking", "distribution", "gov", "mint"}
	for _, module := range modules {
		cache, err := enhancements.StateCache.GetOrCreateCache(
			module,
			enhancements.Config.StateCacheSize,
			enhancements.Config.StateCacheTTL,
		)
		if err != nil {
			return err
		}
		
		// Wrap the module's KVStore with caching
		// This would require modifying the module setup in app.go
		_ = cache // placeholder for now
	}
	
	return nil
}

// BeginBlockOptimizations runs optimizations at the beginning of each block
func (be *BlockchainEnhancements) BeginBlockOptimizations(ctx sdk.Context) {
	if be.PerformanceMonitor != nil {
		// Record block height
		be.PerformanceMonitor.RecordBlockMetrics(ctx, 0, 0)
		
		// Update mempool metrics
		if be.MemoryPool != nil {
			stats := be.MemoryPool.GetStats()
			be.PerformanceMonitor.RecordMempoolMetrics(stats.Size, stats.TotalBytes)
		}
	}
}

// EndBlockOptimizations runs optimizations at the end of each block
func (be *BlockchainEnhancements) EndBlockOptimizations(ctx sdk.Context, processingTime time.Duration) {
	if be.PerformanceMonitor != nil {
		// Record final block metrics
		blockSize := int64(len(ctx.TxBytes()))
		be.PerformanceMonitor.RecordBlockMetrics(ctx, processingTime, blockSize)
	}
}

// GetEnhancementStats returns statistics for all enhancements
func (be *BlockchainEnhancements) GetEnhancementStats() map[string]interface{} {
	stats := make(map[string]interface{})
	
	if be.StateCache != nil {
		stats["cache"] = be.StateCache.GetCacheStats()
	}
	
	if be.BatchProcessor != nil {
		stats["batch_processing"] = be.BatchProcessor.GetStats()
	}
	
	if be.MemoryPool != nil {
		stats["mempool"] = be.MemoryPool.GetStats()
	}
	
	return stats
}

// ValidateEnhancedConfig validates the enhanced configuration
func ValidateEnhancedConfig(config *EnhancedConfig) error {
	if config.MaxBlockSize <= 0 {
		return sdk.ErrInvalidRequest.Wrap("max block size must be positive")
	}
	
	if config.MaxGasPerBlock <= 0 {
		return sdk.ErrInvalidRequest.Wrap("max gas per block must be positive")
	}
	
	if config.BatchProcessingEnabled && config.MaxBatchSize <= 0 {
		return sdk.ErrInvalidRequest.Wrap("max batch size must be positive when batch processing is enabled")
	}
	
	if config.StateCacheSize <= 0 {
		return sdk.ErrInvalidRequest.Wrap("state cache size must be positive")
	}
	
	return nil
}

// Additional helper method to add to the main App struct
func (app *App) GetEnhancements() *BlockchainEnhancements {
	return app.enhancements
}

// Add this field to the main App struct in app.go:
// enhancements *BlockchainEnhancements