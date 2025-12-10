package oracle

import (
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/oracle/client/cli"
	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

var _ module.AppModule = AppModule{}
var _ module.AppModuleBasic = AppModuleBasic{}

// AppModuleBasic implements the basic methods for the oracle module.
type AppModuleBasic struct{}

// Name returns the module name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the legacy Amino types.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module interface types.
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns default module genesis state.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	bz, _ := json.Marshal(types.DefaultGenesis())
	return bz
}

// ValidateGenesis ensures the genesis data is valid.
func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	if len(bz) == 0 {
		return nil
	}
	var state types.GenesisState
	if err := json.Unmarshal(bz, &state); err != nil {
		return err
	}
	return state.Validate()
}

// RegisterGRPCGatewayRoutes is a no-op for now.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

// GetTxCmd wires oracle tx subcommands.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return cli.NewTxCmd() }

// GetQueryCmd wires oracle query subcommands.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return cli.NewQueryCmd() }

// AppModule implements the full module interface.
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		keeper: k,
	}
}

// ConsensusVersion returns the module consensus version.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// IsAppModule indicates this module satisfies the Cosmos SDK AppModule interface.
func (AppModule) IsAppModule() {}

// IsOnePerModuleType is part of the depinject.OnePerModuleType marker interface.
func (AppModule) IsOnePerModuleType() {}

// RegisterInvariants registers invariants (none for now).
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// RegisterServices registers the module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
}

// InitGenesis initializes module state.
func (am AppModule) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var state types.GenesisState
	if len(data) == 0 {
		state = *types.DefaultGenesis()
	} else {
		if err := json.Unmarshal(data, &state); err != nil {
			panic(fmt.Sprintf("failed to unmarshal %s genesis state: %v", types.ModuleName, err))
		}
	}
	InitGenesis(ctx, am.keeper, &state)
	return nil
}

// ExportGenesis exports the module state.
func (am AppModule) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) json.RawMessage {
	state := ExportGenesis(ctx, am.keeper)
	bz, _ := json.Marshal(state)
	return bz
}

// EndBlock performs end-block processing for the oracle module.
// It marks stale prices and slashes providers who haven't updated.
func (am AppModule) EndBlock(ctx sdk.Context) []abci.ValidatorUpdate {
	am.keeper.ProcessStalePrices(ctx)
	return nil
}
