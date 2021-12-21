package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultStatesetInstanceCost is initially set the same as in wasmd
	DefaultStatesetInstanceCost uint64 = 60_000
	// DefaultStatesetCompileCost set to a large number for testing
	DefaultStatesetCompileCost uint64 = 100
)

// StatesetGasRegisterConfig is defaults plus a custom compile amount
func StatesetGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultStatesetInstanceCost
	gasConfig.CompileCost = DefaultStatesetCompileCost

	return gasConfig
}

func NewStatesetWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(StatesetGasRegisterConfig())
}
