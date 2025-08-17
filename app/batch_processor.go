package app

import (
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// BatchProcessor handles batch processing of transactions for improved throughput
type BatchProcessor struct {
	anteHandler    sdk.AnteHandler
	deliverHandler func(ctx sdk.Context, tx sdk.Tx) (*sdk.Result, error)
	maxBatchSize   int
	batchTimeout   time.Duration
	workers        int
	batchQueue     chan []sdk.Tx
	resultChan     chan *BatchResult
	mu             sync.RWMutex
	stats          BatchStats
}

// BatchResult contains the result of processing a batch of transactions
type BatchResult struct {
	Results   []*sdk.Result
	Errors    []error
	GasUsed   uint64
	Timestamp time.Time
}

// BatchStats tracks batch processing statistics
type BatchStats struct {
	TotalBatches     uint64
	TotalTxs         uint64
	AverageLatency   time.Duration
	ThroughputTPS    float64
	ErrorRate        float64
	LastUpdated      time.Time
}

// NewBatchProcessor creates a new batch processor
func NewBatchProcessor(
	anteHandler sdk.AnteHandler,
	deliverHandler func(ctx sdk.Context, tx sdk.Tx) (*sdk.Result, error),
	maxBatchSize int,
	batchTimeout time.Duration,
	workers int,
) *BatchProcessor {
	if maxBatchSize <= 0 {
		maxBatchSize = 100
	}
	if batchTimeout <= 0 {
		batchTimeout = 100 * time.Millisecond
	}
	if workers <= 0 {
		workers = 4
	}

	bp := &BatchProcessor{
		anteHandler:    anteHandler,
		deliverHandler: deliverHandler,
		maxBatchSize:   maxBatchSize,
		batchTimeout:   batchTimeout,
		workers:        workers,
		batchQueue:     make(chan []sdk.Tx, workers*2),
		resultChan:     make(chan *BatchResult, workers*2),
	}

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		go bp.worker()
	}

	return bp
}

// ProcessBatch processes a batch of transactions
func (bp *BatchProcessor) ProcessBatch(ctx sdk.Context, txs []sdk.Tx) *BatchResult {
	start := time.Now()
	results := make([]*sdk.Result, len(txs))
	errors := make([]error, len(txs))
	totalGasUsed := uint64(0)

	// Process transactions concurrently within the batch
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, bp.workers)

	for i, tx := range txs {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(idx int, transaction sdk.Tx) {
			defer wg.Done()
			defer func() { <-semaphore }()

			// Clone context for concurrent processing
			txCtx := ctx.WithTxBytes(ctx.TxBytes())

			// Run ante handler
			anteCtx, err := bp.anteHandler(txCtx, transaction, false)
			if err != nil {
				errors[idx] = err
				return
			}

			// Deliver transaction
			result, err := bp.deliverHandler(anteCtx, transaction)
			if err != nil {
				errors[idx] = err
				return
			}

			results[idx] = result
			totalGasUsed += result.GasUsed
		}(i, tx)
	}

	wg.Wait()

	// Update statistics
	bp.updateStats(len(txs), time.Since(start), errors)

	return &BatchResult{
		Results:   results,
		Errors:    errors,
		GasUsed:   totalGasUsed,
		Timestamp: time.Now(),
	}
}

// worker processes batches from the queue
func (bp *BatchProcessor) worker() {
	for batch := range bp.batchQueue {
		// Note: This would require a context to be passed
		// For now, we'll create a placeholder implementation
		result := &BatchResult{
			Results:   make([]*sdk.Result, len(batch)),
			Errors:    make([]error, len(batch)),
			Timestamp: time.Now(),
		}
		bp.resultChan <- result
	}
}

// updateStats updates batch processing statistics
func (bp *BatchProcessor) updateStats(txCount int, latency time.Duration, errors []error) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	bp.stats.TotalBatches++
	bp.stats.TotalTxs += uint64(txCount)

	errorCount := 0
	for _, err := range errors {
		if err != nil {
			errorCount++
		}
	}

	// Update average latency (exponential moving average)
	if bp.stats.AverageLatency == 0 {
		bp.stats.AverageLatency = latency
	} else {
		bp.stats.AverageLatency = time.Duration(
			0.9*float64(bp.stats.AverageLatency) + 0.1*float64(latency),
		)
	}

	// Calculate throughput (transactions per second)
	if latency.Seconds() > 0 {
		currentTPS := float64(txCount) / latency.Seconds()
		if bp.stats.ThroughputTPS == 0 {
			bp.stats.ThroughputTPS = currentTPS
		} else {
			bp.stats.ThroughputTPS = 0.9*bp.stats.ThroughputTPS + 0.1*currentTPS
		}
	}

	// Calculate error rate
	if bp.stats.TotalTxs > 0 {
		bp.stats.ErrorRate = float64(errorCount) / float64(bp.stats.TotalTxs)
	}

	bp.stats.LastUpdated = time.Now()
}

// GetStats returns current batch processing statistics
func (bp *BatchProcessor) GetStats() BatchStats {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.stats
}

// TransactionBatcher collects transactions into batches
type TransactionBatcher struct {
	processor       *BatchProcessor
	currentBatch    []sdk.Tx
	maxBatchSize    int
	batchTimeout    time.Duration
	lastBatchTime   time.Time
	mu              sync.Mutex
	flushTimer      *time.Timer
	onBatchReady    func([]sdk.Tx)
}

// NewTransactionBatcher creates a new transaction batcher
func NewTransactionBatcher(
	processor *BatchProcessor,
	maxBatchSize int,
	batchTimeout time.Duration,
	onBatchReady func([]sdk.Tx),
) *TransactionBatcher {
	tb := &TransactionBatcher{
		processor:     processor,
		maxBatchSize:  maxBatchSize,
		batchTimeout:  batchTimeout,
		onBatchReady:  onBatchReady,
		lastBatchTime: time.Now(),
	}

	tb.resetFlushTimer()
	return tb
}

// AddTransaction adds a transaction to the current batch
func (tb *TransactionBatcher) AddTransaction(tx sdk.Tx) error {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.currentBatch = append(tb.currentBatch, tx)

	// Check if batch is full
	if len(tb.currentBatch) >= tb.maxBatchSize {
		tb.flushBatch()
	}

	return nil
}

// flushBatch sends the current batch for processing
func (tb *TransactionBatcher) flushBatch() {
	if len(tb.currentBatch) == 0 {
		return
	}

	if tb.onBatchReady != nil {
		tb.onBatchReady(tb.currentBatch)
	}

	tb.currentBatch = nil
	tb.lastBatchTime = time.Now()
	tb.resetFlushTimer()
}

// resetFlushTimer resets the batch timeout timer
func (tb *TransactionBatcher) resetFlushTimer() {
	if tb.flushTimer != nil {
		tb.flushTimer.Stop()
	}

	tb.flushTimer = time.AfterFunc(tb.batchTimeout, func() {
		tb.mu.Lock()
		defer tb.mu.Unlock()
		tb.flushBatch()
	})
}

// ForceFlush forces the current batch to be processed
func (tb *TransactionBatcher) ForceFlush() {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.flushBatch()
}

// MemoryPool represents an optimized transaction memory pool
type MemoryPool struct {
	pending     map[string]sdk.Tx
	queue       []sdk.Tx
	maxSize     int
	maxTxSize   int64
	minGasPrice sdk.DecCoins
	mu          sync.RWMutex
	stats       MempoolStats
}

// MempoolStats contains mempool statistics
type MempoolStats struct {
	Size        int
	TotalBytes  int64
	Pending     int
	Queued      int
	Rejected    uint64
	Accepted    uint64
	LastUpdated time.Time
}

// NewMemoryPool creates a new optimized memory pool
func NewMemoryPool(maxSize int, maxTxSize int64, minGasPrice sdk.DecCoins) *MemoryPool {
	return &MemoryPool{
		pending:     make(map[string]sdk.Tx),
		queue:       make([]sdk.Tx, 0, maxSize),
		maxSize:     maxSize,
		maxTxSize:   maxTxSize,
		minGasPrice: minGasPrice,
	}
}

// AddTx adds a transaction to the mempool
func (mp *MemoryPool) AddTx(tx sdk.Tx, txBytes []byte) error {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	txHash := string(sdk.Tx(tx).Hash())

	// Check if transaction already exists
	if _, exists := mp.pending[txHash]; exists {
		return sdkerrors.Wrap(sdkerrors.ErrTxInMempoolCache, "transaction already in mempool")
	}

	// Check size limits
	if len(txBytes) > int(mp.maxTxSize) {
		mp.stats.Rejected++
		return sdkerrors.Wrapf(sdkerrors.ErrTxTooLarge, 
			"transaction size %d exceeds maximum %d", len(txBytes), mp.maxTxSize)
	}

	if len(mp.queue) >= mp.maxSize {
		mp.stats.Rejected++
		return sdkerrors.Wrap(sdkerrors.ErrMempoolIsFull, "mempool is full")
	}

	// Validate minimum gas price
	feeTx, ok := tx.(sdk.FeeTx)
	if ok && !mp.minGasPrice.IsZero() {
		feeCoins := feeTx.GetFee()
		gas := feeTx.GetGas()
		
		if gas > 0 {
			for _, minPrice := range mp.minGasPrice {
				requiredFee := minPrice.Amount.Mul(sdk.NewDec(int64(gas)))
				if !feeCoins.AmountOf(minPrice.Denom).GTE(requiredFee.TruncateInt()) {
					mp.stats.Rejected++
					return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee,
						"insufficient fee: required %s, got %s",
						sdk.NewCoin(minPrice.Denom, requiredFee.TruncateInt()),
						feeCoins.AmountOf(minPrice.Denom))
				}
			}
		}
	}

	// Add to mempool
	mp.pending[txHash] = tx
	mp.queue = append(mp.queue, tx)
	mp.stats.Size++
	mp.stats.TotalBytes += int64(len(txBytes))
	mp.stats.Accepted++
	mp.stats.LastUpdated = time.Now()

	return nil
}

// RemoveTx removes a transaction from the mempool
func (mp *MemoryPool) RemoveTx(tx sdk.Tx) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	txHash := string(sdk.Tx(tx).Hash())
	
	if _, exists := mp.pending[txHash]; exists {
		delete(mp.pending, txHash)
		
		// Remove from queue
		for i, queuedTx := range mp.queue {
			if string(sdk.Tx(queuedTx).Hash()) == txHash {
				mp.queue = append(mp.queue[:i], mp.queue[i+1:]...)
				break
			}
		}
		
		mp.stats.Size--
		mp.stats.LastUpdated = time.Now()
	}
}

// GetTxs returns a batch of transactions from the mempool
func (mp *MemoryPool) GetTxs(maxTxs int) []sdk.Tx {
	mp.mu.RLock()
	defer mp.mu.RUnlock()

	if maxTxs > len(mp.queue) {
		maxTxs = len(mp.queue)
	}

	return mp.queue[:maxTxs]
}

// GetStats returns mempool statistics
func (mp *MemoryPool) GetStats() MempoolStats {
	mp.mu.RLock()
	defer mp.mu.RUnlock()
	
	stats := mp.stats
	stats.Pending = len(mp.pending)
	stats.Queued = len(mp.queue)
	
	return stats
}

// Clear removes all transactions from the mempool
func (mp *MemoryPool) Clear() {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	mp.pending = make(map[string]sdk.Tx)
	mp.queue = mp.queue[:0]
	mp.stats.Size = 0
	mp.stats.TotalBytes = 0
	mp.stats.LastUpdated = time.Now()
}