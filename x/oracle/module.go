package oracle

import (
	"context"
	"encoding/json"
	"fmt"
	
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	
	abci "github.com/cometbft/cometbft/abci/types"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	
	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the oracle module
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the oracle module's name
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterCodec registers the oracle module's types for the given codec
func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (b AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns default genesis state as raw bytes for the oracle module
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the oracle module
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return data.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the oracle module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the oracle module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns no root query command for the oracle module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// AppModule implements an application module for the oracle module
type AppModule struct {
	AppModuleBasic
	
	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Name returns the oracle module's name
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers the oracle module invariants
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	// Register invariants
}

// RegisterServices registers module services
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// InitGenesis performs genesis initialization for the oracle module
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	cdc.MustUnmarshalJSON(gs, &genState)
	InitGenesis(ctx, am.keeper, genState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the oracle module
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule/ConsensusVersion
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock returns the begin blocker for the oracle module
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	BeginBlocker(ctx, am.keeper)
}

// EndBlock returns the end blocker for the oracle module
func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlocker(ctx, am.keeper)
}

// BeginBlocker handles block beginning logic for oracle module
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// Check for stale prices and emit events
	params := k.GetParams(ctx)
	
	// Get all aggregated prices
	allAssets := k.GetAllAggregatedAssets(ctx)
	
	for _, asset := range allAssets {
		aggregated, found := k.GetAggregatedPrice(ctx, asset)
		if !found {
			continue
		}
		
		// Check if price is stale
		if ctx.BlockTime().Sub(aggregated.LastUpdate) > params.MaxPriceAge {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"stale_price_detected",
					sdk.NewAttribute("asset", asset),
					sdk.NewAttribute("last_update", aggregated.LastUpdate.String()),
					sdk.NewAttribute("age", ctx.BlockTime().Sub(aggregated.LastUpdate).String()),
				),
			)
		}
		
		// Check for emergency conditions
		if aggregated.StandardDev.GT(params.EmergencyThreshold) {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"emergency_threshold_exceeded",
					sdk.NewAttribute("asset", asset),
					sdk.NewAttribute("standard_dev", aggregated.StandardDev.String()),
					sdk.NewAttribute("threshold", params.EmergencyThreshold.String()),
				),
			)
		}
	}
}

// EndBlocker handles block ending logic for oracle module
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	// Clean up expired price feeds
	allFeeds := k.GetAllPriceFeeds(ctx)
	currentTime := ctx.BlockTime()
	
	for _, feed := range allFeeds {
		if feed.IsExpired(currentTime) {
			k.RemovePriceFeed(ctx, feed.FeedID)
		}
	}
	
	// Update provider performance metrics
	providers := k.GetActiveOracleProviders(ctx)
	params := k.GetParams(ctx)
	
	for _, provider := range providers {
		// Check if provider hasn't submitted recently
		if currentTime.Sub(provider.LastUpdate) > params.UpdateInterval*2 {
			// Reduce reputation for inactivity
			k.UpdateProviderReputation(ctx, provider.Address, sdk.NewDecWithPrec(-5, 3)) // -0.005
			
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"provider_inactive",
					sdk.NewAttribute("provider", provider.Address),
					sdk.NewAttribute("last_update", provider.LastUpdate.String()),
				),
			)
		}
	}
	
	return []abci.ValidatorUpdate{}
}