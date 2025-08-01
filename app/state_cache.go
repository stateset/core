package app

import (
	"sync"
	"time"

	"cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	lru "github.com/hashicorp/golang-lru"
)

// StateCache provides a caching layer for frequently accessed state data
type StateCache struct {
	cache           *lru.ARCCache
	mu              sync.RWMutex
	hits            uint64
	misses          uint64
	evictions       uint64
	ttl             time.Duration
	expirationMap   map[string]time.Time
	expirationMutex sync.RWMutex
}

// CacheEntry represents a cached state entry
type CacheEntry struct {
	Value     []byte
	Timestamp time.Time
}

// NewStateCache creates a new state cache with specified size and TTL
func NewStateCache(size int, ttl time.Duration) (*StateCache, error) {
	cache, err := lru.NewARC(size)
	if err != nil {
		return nil, err
	}

	sc := &StateCache{
		cache:         cache,
		ttl:           ttl,
		expirationMap: make(map[string]time.Time),
	}

	// Start cleanup goroutine
	go sc.cleanupExpired()

	return sc, nil
}

// Get retrieves a value from the cache
func (sc *StateCache) Get(key []byte) ([]byte, bool) {
	sc.mu.RLock()
	defer sc.mu.RUnlock()

	keyStr := string(key)
	
	// Check expiration
	sc.expirationMutex.RLock()
	expTime, exists := sc.expirationMap[keyStr]
	sc.expirationMutex.RUnlock()
	
	if exists && time.Now().After(expTime) {
		sc.Delete(key)
		sc.misses++
		return nil, false
	}

	value, ok := sc.cache.Get(keyStr)
	if !ok {
		sc.misses++
		return nil, false
	}

	entry := value.(*CacheEntry)
	sc.hits++
	return entry.Value, true
}

// Set stores a value in the cache
func (sc *StateCache) Set(key []byte, value []byte) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	keyStr := string(key)
	entry := &CacheEntry{
		Value:     value,
		Timestamp: time.Now(),
	}

	evicted := sc.cache.Add(keyStr, entry)
	if evicted {
		sc.evictions++
	}

	// Set expiration
	if sc.ttl > 0 {
		sc.expirationMutex.Lock()
		sc.expirationMap[keyStr] = time.Now().Add(sc.ttl)
		sc.expirationMutex.Unlock()
	}
}

// Delete removes a value from the cache
func (sc *StateCache) Delete(key []byte) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	keyStr := string(key)
	sc.cache.Remove(keyStr)
	
	sc.expirationMutex.Lock()
	delete(sc.expirationMap, keyStr)
	sc.expirationMutex.Unlock()
}

// Clear removes all entries from the cache
func (sc *StateCache) Clear() {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	sc.cache.Purge()
	
	sc.expirationMutex.Lock()
	sc.expirationMap = make(map[string]time.Time)
	sc.expirationMutex.Unlock()
	
	sc.hits = 0
	sc.misses = 0
	sc.evictions = 0
}

// GetStats returns cache statistics
func (sc *StateCache) GetStats() CacheStats {
	sc.mu.RLock()
	defer sc.mu.RUnlock()

	total := sc.hits + sc.misses
	hitRate := float64(0)
	if total > 0 {
		hitRate = float64(sc.hits) / float64(total)
	}

	return CacheStats{
		Hits:      sc.hits,
		Misses:    sc.misses,
		Evictions: sc.evictions,
		Size:      sc.cache.Len(),
		HitRate:   hitRate,
	}
}

// cleanupExpired periodically removes expired entries
func (sc *StateCache) cleanupExpired() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		sc.removeExpiredEntries()
	}
}

// removeExpiredEntries removes all expired entries from the cache
func (sc *StateCache) removeExpiredEntries() {
	sc.expirationMutex.RLock()
	expiredKeys := []string{}
	now := time.Now()
	
	for key, expTime := range sc.expirationMap {
		if now.After(expTime) {
			expiredKeys = append(expiredKeys, key)
		}
	}
	sc.expirationMutex.RUnlock()

	// Remove expired entries
	for _, key := range expiredKeys {
		sc.Delete([]byte(key))
	}
}

// CacheStats contains cache performance statistics
type CacheStats struct {
	Hits      uint64
	Misses    uint64
	Evictions uint64
	Size      int
	HitRate   float64
}

// CachedStore wraps a KVStore with caching functionality
type CachedStore struct {
	parent types.KVStore
	cache  *StateCache
}

// NewCachedStore creates a new cached store
func NewCachedStore(parent types.KVStore, cache *StateCache) *CachedStore {
	return &CachedStore{
		parent: parent,
		cache:  cache,
	}
}

// Get implements KVStore
func (cs *CachedStore) Get(key []byte) []byte {
	// Try cache first
	if value, found := cs.cache.Get(key); found {
		return value
	}

	// Fall back to parent store
	value := cs.parent.Get(key)
	if value != nil {
		cs.cache.Set(key, value)
	}

	return value
}

// Has implements KVStore
func (cs *CachedStore) Has(key []byte) bool {
	// Check cache first
	if _, found := cs.cache.Get(key); found {
		return true
	}

	return cs.parent.Has(key)
}

// Set implements KVStore
func (cs *CachedStore) Set(key, value []byte) {
	cs.parent.Set(key, value)
	cs.cache.Set(key, value)
}

// Delete implements KVStore
func (cs *CachedStore) Delete(key []byte) {
	cs.parent.Delete(key)
	cs.cache.Delete(key)
}

// Iterator implements KVStore
func (cs *CachedStore) Iterator(start, end []byte) types.Iterator {
	return cs.parent.Iterator(start, end)
}

// ReverseIterator implements KVStore
func (cs *CachedStore) ReverseIterator(start, end []byte) types.Iterator {
	return cs.parent.ReverseIterator(start, end)
}

// CacheContext implements KVStore
func (cs *CachedStore) CacheContext() types.CacheWrap {
	return cs.parent.CacheContext()
}

// Write implements CacheWrap
func (cs *CachedStore) Write() {
	if cw, ok := cs.parent.(types.CacheWrap); ok {
		cw.Write()
	}
}

// GetStoreType implements Store
func (cs *CachedStore) GetStoreType() types.StoreType {
	return cs.parent.GetStoreType()
}

// CacheWrapper for module stores
type CacheWrapper struct {
	caches map[string]*StateCache
	mu     sync.RWMutex
}

// NewCacheWrapper creates a new cache wrapper
func NewCacheWrapper() *CacheWrapper {
	return &CacheWrapper{
		caches: make(map[string]*StateCache),
	}
}

// GetOrCreateCache gets or creates a cache for a specific module
func (cw *CacheWrapper) GetOrCreateCache(module string, size int, ttl time.Duration) (*StateCache, error) {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	if cache, exists := cw.caches[module]; exists {
		return cache, nil
	}

	cache, err := NewStateCache(size, ttl)
	if err != nil {
		return nil, err
	}

	cw.caches[module] = cache
	return cache, nil
}

// GetCacheStats returns statistics for all caches
func (cw *CacheWrapper) GetCacheStats() map[string]CacheStats {
	cw.mu.RLock()
	defer cw.mu.RUnlock()

	stats := make(map[string]CacheStats)
	for module, cache := range cw.caches {
		stats[module] = cache.GetStats()
	}

	return stats
}

// ClearAll clears all caches
func (cw *CacheWrapper) ClearAll() {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	for _, cache := range cw.caches {
		cache.Clear()
	}
}