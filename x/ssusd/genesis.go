package ssusd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/ssusd/keeper"
	"github.com/stateset/core/x/ssusd/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	
	for _, position := range genState.Positions {
		k.SetCollateralPosition(ctx, position)
	}
	
	for _, wallet := range genState.AgentWallets {
		k.SetAgentWallet(ctx, wallet)
	}
	
	for _, price := range genState.OraclePrices {
		k.SetOraclePrice(ctx, price)
	}
	
	for _, auction := range genState.ActiveAuctions {
		k.SetLiquidationAuction(ctx, auction)
	}
	
	k.SetStabilityPool(ctx, genState.StabilityPool)
	k.SetSystemMetrics(ctx, genState.SystemMetrics)
	k.SetNextPositionID(ctx, genState.NextPositionId)
	k.SetNextAuctionID(ctx, genState.NextAuctionId)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	
	k.IterateCollateralPositions(ctx, func(position types.CollateralPosition) bool {
		genesis.Positions = append(genesis.Positions, position)
		return false
	})
	
	k.IterateAgentWallets(ctx, func(wallet types.AgentWallet) bool {
		genesis.AgentWallets = append(genesis.AgentWallets, wallet)
		return false
	})
	
	k.IterateLiquidationAuctions(ctx, func(auction types.LiquidationAuction) bool {
		if auction.Status == "active" {
			genesis.ActiveAuctions = append(genesis.ActiveAuctions, auction)
		}
		return false
	})
	
	genesis.StabilityPool = k.GetStabilityPool(ctx)
	genesis.SystemMetrics = k.GetSystemMetrics(ctx)
	genesis.NextPositionId = k.GetNextPositionID(ctx)
	genesis.NextAuctionId = k.GetNextAuctionID(ctx)
	
	return genesis
}