package zkpverify

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/zkpverify/client/cli"
	"github.com/stateset/core/x/zkpverify/keeper"
	"github.com/stateset/core/x/zkpverify/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.HasGenesis     = AppModule{}
	_ module.HasServices    = AppModule{}
)

// AppModuleBasic defines the basic application module for zkpverify
type AppModuleBasic struct{}

// Name returns the module's name
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the module's types on the legacy amino codec
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns the module's default genesis state
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis validates the module's genesis state
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var genesis types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genesis); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genesis.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	// gRPC gateway routes would be registered here
}

// GetTxCmd returns the root tx command for the module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// AppModule implements the AppModule interface for zkpverify
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// RegisterServices registers module services
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// Register the MsgServer and QueryServer implementations
	// Note: In production, you would use protobuf-generated registration functions
	// For now, we use the configurator directly
	_ = cfg.MsgServer()
	_ = cfg.QueryServer()
}

// InitGenesis initializes the module's genesis state
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesis types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesis)

	if err := am.keeper.InitGenesis(ctx, &genesis); err != nil {
		panic(fmt.Sprintf("failed to initialize %s genesis state: %v", types.ModuleName, err))
	}
}

// ExportGenesis exports the module's genesis state
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesis := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(genesis)
}

// ConsensusVersion returns the consensus version
func (am AppModule) ConsensusVersion() uint64 {
	return 1
}

// BeginBlock executes at the beginning of each block
func (am AppModule) BeginBlock(ctx context.Context) error {
	return nil
}

// EndBlock executes at the end of each block
func (am AppModule) EndBlock(ctx context.Context) error {
	return nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface
func (am AppModule) IsAppModule() {}

// RegisterMsgServer is a helper for registering the MsgServer
func RegisterMsgServer(s interface{}, srv types.MsgServer) {
	// Type assertion for gRPC server registration
	if grpcServer, ok := s.(interface {
		RegisterService(desc *interface{}, impl interface{})
	}); ok {
		_ = grpcServer // Would register service here
	}
}

// RegisterQueryServer is a helper for registering the QueryServer
func RegisterQueryServer(s interface{}, srv types.QueryServer) {
	// Type assertion for gRPC server registration
	if grpcServer, ok := s.(interface {
		RegisterService(desc *interface{}, impl interface{})
	}); ok {
		_ = grpcServer // Would register service here
	}
}
