package config

import (
	"time"

	tmcfg "github.com/tendermint/tendermint/config"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// GetOptimizedConsensusParams returns optimized consensus parameters for better blockchain performance
func GetOptimizedConsensusParams() *tmproto.ConsensusParams {
	return &tmproto.ConsensusParams{
		Block: &tmproto.BlockParams{
			// Increase max block size to 10MB for higher throughput
			MaxBytes: 10485760, // 10MB
			// Increase max gas to allow more transactions per block
			MaxGas: 50000000, // 50M gas units
			// Time iota in milliseconds
			TimeIotaMs: 10,
		},
		Evidence: &tmproto.EvidenceParams{
			// Maximum age of evidence in blocks
			MaxAgeNumBlocks: 100000,
			// Maximum age of evidence in time
			MaxAgeDuration: 48 * time.Hour,
			// Maximum size of evidence in bytes
			MaxBytes: 1048576, // 1MB
		},
		Validator: &tmproto.ValidatorParams{
			PubKeyTypes: []string{
				tmproto.ABCIPubKeyTypeEd25519,
			},
		},
		Version: &tmproto.VersionParams{
			AppVersion: 0,
		},
	}
}

// GetOptimizedTendermintConfig returns optimized Tendermint configuration
func GetOptimizedTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()
	
	// Consensus optimizations
	cfg.Consensus.TimeoutPropose = 2000 * time.Millisecond
	cfg.Consensus.TimeoutProposeDelta = 500 * time.Millisecond
	cfg.Consensus.TimeoutPrevote = 1000 * time.Millisecond
	cfg.Consensus.TimeoutPrevoteDelta = 500 * time.Millisecond
	cfg.Consensus.TimeoutPrecommit = 1000 * time.Millisecond
	cfg.Consensus.TimeoutPrecommitDelta = 500 * time.Millisecond
	cfg.Consensus.TimeoutCommit = 1000 * time.Millisecond
	
	// Enable peer exchange for better network topology
	cfg.Consensus.PeerGossipSleepDuration = 50 * time.Millisecond
	cfg.Consensus.PeerQueryMaj23SleepDuration = 1000 * time.Millisecond
	
	// Mempool optimizations
	cfg.Mempool.Size = 10000
	cfg.Mempool.MaxTxsBytes = 1073741824 // 1GB
	cfg.Mempool.CacheSize = 100000
	cfg.Mempool.MaxTxBytes = 1048576 // 1MB per transaction
	
	// P2P optimizations
	cfg.P2P.MaxNumInboundPeers = 100
	cfg.P2P.MaxNumOutboundPeers = 50
	cfg.P2P.SendRate = 10240000 // 10 MB/s
	cfg.P2P.RecvRate = 10240000 // 10 MB/s
	cfg.P2P.FlushThrottleTimeout = 10 * time.Millisecond
	
	// RPC optimizations
	cfg.RPC.MaxOpenConnections = 5000
	cfg.RPC.MaxSubscriptionClients = 200
	cfg.RPC.MaxSubscriptionsPerClient = 10
	cfg.RPC.TimeoutBroadcastTxCommit = 10 * time.Second
	
	// Instrumentation
	cfg.Instrumentation.Prometheus = true
	cfg.Instrumentation.PrometheusListenAddr = ":26660"
	
	return cfg
}