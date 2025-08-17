package ssusd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/cometbft/cometbft/abci/types"

	// "github.com/stateset/core/x/ssusd/client/cli"
	"github.com/stateset/core/x/ssusd/keeper"
	"github.com/stateset/core/x/ssusd/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
	cdc codec.Codec
}

func NewAppModuleBasic(cdc codec.Codec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	// return cli.GetTxCmd()
	return nil // CLI not implemented yet
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	// return cli.GetQueryCmd(types.StoreKey)
	return nil // CLI not implemented yet
}

type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	cdc.MustUnmarshalJSON(gs, &genState)
	InitGenesis(ctx, am.keeper, genState)
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(gs)
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	BeginBlocker(ctx, am.keeper)
}

func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.IterateLiquidationAuctions(ctx, func(auction types.LiquidationAuction) bool {
		if auction.Status == "active" && ctx.BlockTime().After(auction.EndTime) {
			auction.Status = "expired"
			k.SetLiquidationAuction(ctx, auction)
			
			pool := k.GetStabilityPool(ctx)
			if pool.TotalDeposits.Amount.GTE(auction.Debt.Amount) {
				pool.TotalDeposits = pool.TotalDeposits.Sub(auction.Debt)
				pool.TotalDebtAbsorbed = pool.TotalDebtAbsorbed.Add(auction.Debt)
				
				collateralPerUnit := sdk.NewDecFromInt(auction.Collateral.Amount).
					Quo(sdk.NewDecFromInt(auction.Debt.Amount))
				
				for i, provider := range pool.Providers {
					share := sdk.NewDecFromInt(provider.Deposit.Amount).
						Quo(sdk.NewDecFromInt(pool.TotalDeposits.Amount.Add(auction.Debt.Amount)))
					rewardAmount := share.Mul(sdk.NewDecFromInt(auction.Collateral.Amount)).TruncateInt()
					pool.Providers[i].RewardsEarned = pool.Providers[i].RewardsEarned.
						Add(sdk.NewCoin("stst", rewardAmount))
				}
				
				k.SetStabilityPool(ctx, pool)
			}
		}
		return false
	})
	
	k.IterateCollateralPositions(ctx, func(position types.CollateralPosition) bool {
		position.CollateralizationRatio = k.CalculateCollateralRatio(ctx, position.Collateral, position.Debt)
		position.IsLiquidatable = k.IsPositionLiquidatable(ctx, position)
		k.SetCollateralPosition(ctx, position)
		return false
	})
	
	k.UpdateSystemMetrics(ctx)
}