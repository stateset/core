package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/x/ssusd/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) MintSSUSD(goCtx context.Context, msg *types.MsgMintSSUSD) (*types.MsgMintSSUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	
	if params.EmergencyShutdown {
		return nil, sdkerrors.Wrap(types.ErrEmergencyShutdown, "system is in emergency shutdown")
	}
	
	minterAddr, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return nil, err
	}
	
	if msg.Collateral.Denom != "stst" {
		return nil, sdkerrors.Wrap(types.ErrInvalidCollateral, "only STST is accepted as collateral")
	}
	
	collateralRatio := k.CalculateCollateralRatio(ctx, msg.Collateral, msg.MintAmount)
	if collateralRatio.LT(params.MinCollateralRatio) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientCollateral, 
			fmt.Sprintf("collateral ratio %s is below minimum %s", collateralRatio, params.MinCollateralRatio))
	}
	
	position, exists := k.GetCollateralPosition(ctx, msg.Minter)
	if exists {
		position.Collateral = position.Collateral.Add(msg.Collateral)
		position.Debt = position.Debt.Add(msg.MintAmount)
		position.CollateralizationRatio = k.CalculateCollateralRatio(ctx, position.Collateral, position.Debt)
	} else {
		position = types.CollateralPosition{
			Owner:                  msg.Minter,
			Collateral:             msg.Collateral,
			Debt:                   msg.MintAmount,
			CollateralizationRatio: collateralRatio,
			LastUpdate:             time.Now(),
			IsLiquidatable:         false,
		}
	}
	
	currentDebt := sdk.ZeroInt()
	k.IterateCollateralPositions(ctx, func(pos types.CollateralPosition) bool {
		currentDebt = currentDebt.Add(pos.Debt.Amount)
		return false
	})
	
	newTotalDebt := currentDebt.Add(msg.MintAmount.Amount)
	if newTotalDebt.GT(params.GlobalDebtCeiling) {
		return nil, sdkerrors.Wrap(types.ErrDebtCeilingExceeded, "global debt ceiling exceeded")
	}
	
	if position.Debt.Amount.GT(params.MaxDebtPerUser) {
		return nil, sdkerrors.Wrap(types.ErrDebtCeilingExceeded, "user debt ceiling exceeded")
	}
	
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, minterAddr, types.ModuleName, sdk.NewCoins(msg.Collateral))
	if err != nil {
		return nil, err
	}
	
	mintingFee := sdk.NewDecFromInt(msg.MintAmount.Amount).Mul(params.MintingFee).TruncateInt()
	netMintAmount := msg.MintAmount.Amount.Sub(mintingFee)
	
	ssUSDToMint := sdk.NewCoins(sdk.NewCoin("ssusd", netMintAmount))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, ssUSDToMint)
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minterAddr, ssUSDToMint)
	if err != nil {
		return nil, err
	}
	
	if mintingFee.IsPositive() {
		feeCoins := sdk.NewCoins(sdk.NewCoin("ssusd", mintingFee))
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, feeCoins)
		if err != nil {
			return nil, err
		}
		err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, feeCoins)
		if err != nil {
			return nil, err
		}
	}
	
	k.SetCollateralPosition(ctx, position)
	k.UpdateSystemMetrics(ctx)
	
	positionID := fmt.Sprintf("%s_%d", msg.Minter, k.GetNextPositionID(ctx))
	k.SetNextPositionID(ctx, k.GetNextPositionID(ctx)+1)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintSSUSD,
			sdk.NewAttribute(types.AttributeKeyMinter, msg.Minter),
			sdk.NewAttribute(types.AttributeKeyCollateral, msg.Collateral.String()),
			sdk.NewAttribute(types.AttributeKeyMinted, ssUSDToMint.String()),
			sdk.NewAttribute(types.AttributeKeyCollateralRatio, collateralRatio.String()),
		),
	})
	
	return &types.MsgMintSSUSDResponse{
		PositionId:       positionID,
		Minted:           sdk.NewCoin("ssusd", netMintAmount),
		CollateralRatio:  collateralRatio.String(),
	}, nil
}

func (k msgServer) BurnSSUSD(goCtx context.Context, msg *types.MsgBurnSSUSD) (*types.MsgBurnSSUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	
	burnerAddr, err := sdk.AccAddressFromBech32(msg.Burner)
	if err != nil {
		return nil, err
	}
	
	position, exists := k.GetCollateralPosition(ctx, msg.Burner)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrPositionNotFound, "no collateral position found")
	}
	
	if msg.BurnAmount.Amount.GT(position.Debt.Amount) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAmount, "burn amount exceeds debt")
	}
	
	redemptionFee := sdk.NewDecFromInt(msg.BurnAmount.Amount).Mul(params.RedemptionFee).TruncateInt()
	totalBurnAmount := msg.BurnAmount.Amount.Add(redemptionFee)
	
	userBalance := k.bankKeeper.GetBalance(ctx, burnerAddr, "ssusd")
	if userBalance.Amount.LT(totalBurnAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient ssUSD balance")
	}
	
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, burnerAddr, types.ModuleName, 
		sdk.NewCoins(sdk.NewCoin("ssusd", totalBurnAmount)))
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ssusd", totalBurnAmount)))
	if err != nil {
		return nil, err
	}
	
	collateralToReturn := position.Collateral.Amount.Mul(msg.BurnAmount.Amount).Quo(position.Debt.Amount)
	position.Collateral.Amount = position.Collateral.Amount.Sub(collateralToReturn)
	position.Debt.Amount = position.Debt.Amount.Sub(msg.BurnAmount.Amount)
	
	if position.Debt.IsZero() {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, burnerAddr,
			sdk.NewCoins(position.Collateral))
		if err != nil {
			return nil, err
		}
		k.DeleteCollateralPosition(ctx, msg.Burner)
	} else {
		position.CollateralizationRatio = k.CalculateCollateralRatio(ctx, position.Collateral, position.Debt)
		position.LastUpdate = time.Now()
		k.SetCollateralPosition(ctx, position)
		
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, burnerAddr,
			sdk.NewCoins(sdk.NewCoin("stst", collateralToReturn)))
		if err != nil {
			return nil, err
		}
	}
	
	k.UpdateSystemMetrics(ctx)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnSSUSD,
			sdk.NewAttribute(types.AttributeKeyBurner, msg.Burner),
			sdk.NewAttribute(types.AttributeKeyBurned, msg.BurnAmount.String()),
			sdk.NewAttribute(types.AttributeKeyCollateralReturned, collateralToReturn.String()),
		),
	})
	
	return &types.MsgBurnSSUSDResponse{
		CollateralReturned: sdk.NewCoin("stst", collateralToReturn),
		RemainingDebt:      position.Debt,
	}, nil
}

func (k msgServer) CreateAgentWallet(goCtx context.Context, msg *types.MsgCreateAgentWallet) (*types.MsgCreateAgentWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	if _, exists := k.GetAgentWallet(ctx, msg.AgentId); exists {
		return nil, sdkerrors.Wrap(types.ErrAgentExists, "agent wallet already exists")
	}
	
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	
	if !msg.InitialDeposit.IsZero() {
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, 
			sdk.NewCoins(msg.InitialDeposit))
		if err != nil {
			return nil, err
		}
	}
	
	wallet := types.AgentWallet{
		AgentId:            msg.AgentId,
		Owner:              msg.Creator,
		Balances:           sdk.NewCoins(msg.InitialDeposit),
		AuthorizedSigners:  msg.AuthorizedSigners,
		TransactionCount:   0,
		ReputationScore:    sdk.NewDec(100),
		IsActive:           true,
		CreatedAt:          time.Now(),
	}
	
	k.SetAgentWallet(ctx, wallet)
	
	metrics := k.GetSystemMetrics(ctx)
	metrics.TotalAgents++
	k.SetSystemMetrics(ctx, metrics)
	
	walletAddr := sdk.AccAddress([]byte(msg.AgentId)).String()
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateAgent,
			sdk.NewAttribute(types.AttributeKeyAgentID, msg.AgentId),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyWalletAddress, walletAddr),
		),
	})
	
	return &types.MsgCreateAgentWalletResponse{
		WalletAddress: walletAddr,
		AgentId:       msg.AgentId,
	}, nil
}

func (k msgServer) AgentTransfer(goCtx context.Context, msg *types.MsgAgentTransfer) (*types.MsgAgentTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	
	fromWallet, exists := k.GetAgentWallet(ctx, msg.FromAgent)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrAgentNotFound, "from agent not found")
	}
	
	toWallet, exists := k.GetAgentWallet(ctx, msg.ToAgent)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrAgentNotFound, "to agent not found")
	}
	
	if !fromWallet.IsActive || !toWallet.IsActive {
		return nil, sdkerrors.Wrap(types.ErrAgentInactive, "agent wallet is inactive")
	}
	
	if fromWallet.ReputationScore.LT(params.MinAgentReputation) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientReputation, "insufficient reputation score")
	}
	
	found := false
	for _, coin := range fromWallet.Balances {
		if coin.Denom == msg.Amount.Denom && coin.Amount.GTE(msg.Amount.Amount) {
			found = true
			break
		}
	}
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient balance in agent wallet")
	}
	
	baseFee := sdk.NewDecFromInt(msg.Amount.Amount).Mul(params.MintingFee)
	agentDiscount := baseFee.Mul(params.AgentFeeDiscount)
	finalFee := baseFee.Sub(agentDiscount).TruncateInt()
	
	netAmount := msg.Amount.Amount.Sub(finalFee)
	
	fromWallet.Balances = fromWallet.Balances.Sub(sdk.NewCoins(msg.Amount))
	toWallet.Balances = toWallet.Balances.Add(sdk.NewCoins(sdk.NewCoin(msg.Amount.Denom, netAmount)))
	
	fromWallet.TransactionCount++
	toWallet.TransactionCount++
	
	fromWallet.ReputationScore = fromWallet.ReputationScore.Add(sdk.NewDec(1))
	toWallet.ReputationScore = toWallet.ReputationScore.Add(sdk.NewDec(1))
	
	k.SetAgentWallet(ctx, fromWallet)
	k.SetAgentWallet(ctx, toWallet)
	
	if finalFee.IsPositive() && msg.Amount.Denom == "ssusd" {
		err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ssusd", finalFee)))
		if err != nil {
			return nil, err
		}
	}
	
	txID := fmt.Sprintf("%s_%s_%d", msg.FromAgent, msg.ToAgent, ctx.BlockHeight())
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAgentTransfer,
			sdk.NewAttribute(types.AttributeKeyFromAgent, msg.FromAgent),
			sdk.NewAttribute(types.AttributeKeyToAgent, msg.ToAgent),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyTransactionID, txID),
			sdk.NewAttribute(types.AttributeKeyMemo, msg.Memo),
		),
	})
	
	return &types.MsgAgentTransferResponse{
		TransactionId: txID,
		FeeCharged:    finalFee.String(),
		FromNonce:     fromWallet.TransactionCount,
		ToNonce:       toWallet.TransactionCount,
	}, nil
}

func (k msgServer) Liquidate(goCtx context.Context, msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	
	position, exists := k.GetCollateralPosition(ctx, msg.Debtor)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrPositionNotFound, "position not found")
	}
	
	if !k.IsPositionLiquidatable(ctx, position) {
		return nil, sdkerrors.Wrap(types.ErrNotLiquidatable, "position is not liquidatable")
	}
	
	auctionID := fmt.Sprintf("auction_%d", k.GetNextAuctionID(ctx))
	k.SetNextAuctionID(ctx, k.GetNextAuctionID(ctx)+1)
	
	startingPrice := k.GetSTSTPrice(ctx).Mul(sdk.NewDecWithPrec(120, 2))
	
	auction := types.LiquidationAuction{
		Id:              auctionID,
		Debtor:          msg.Debtor,
		Collateral:      position.Collateral,
		Debt:            position.Debt,
		StartingPrice:   startingPrice,
		CurrentPrice:    startingPrice,
		PriceDecayRate:  params.AuctionPriceDecay,
		StartTime:       time.Now(),
		EndTime:         time.Now().Add(time.Duration(params.AuctionDuration) * time.Second),
		Status:          "active",
	}
	
	k.SetLiquidationAuction(ctx, auction)
	k.DeleteCollateralPosition(ctx, msg.Debtor)
	
	metrics := k.GetSystemMetrics(ctx)
	metrics.TotalLiquidations++
	k.SetSystemMetrics(ctx, metrics)
	
	k.UpdateSystemMetrics(ctx)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeLiquidation,
			sdk.NewAttribute(types.AttributeKeyLiquidator, msg.Liquidator),
			sdk.NewAttribute(types.AttributeKeyDebtor, msg.Debtor),
			sdk.NewAttribute(types.AttributeKeyAuctionID, auctionID),
			sdk.NewAttribute(types.AttributeKeyCollateral, position.Collateral.String()),
			sdk.NewAttribute(types.AttributeKeyDebt, position.Debt.String()),
		),
	})
	
	return &types.MsgLiquidateResponse{
		AuctionId:           auctionID,
		CollateralAuctioned: position.Collateral,
		DebtRecovered:       position.Debt,
	}, nil
}

func (k msgServer) AddCollateral(goCtx context.Context, msg *types.MsgAddCollateral) (*types.MsgAddCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	position, exists := k.GetCollateralPosition(ctx, msg.Owner)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrPositionNotFound, "no collateral position found")
	}
	
	ownerAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddr, types.ModuleName, sdk.NewCoins(msg.Collateral))
	if err != nil {
		return nil, err
	}
	
	position.Collateral = position.Collateral.Add(msg.Collateral)
	position.CollateralizationRatio = k.CalculateCollateralRatio(ctx, position.Collateral, position.Debt)
	position.LastUpdate = time.Now()
	position.IsLiquidatable = k.IsPositionLiquidatable(ctx, position)
	
	k.SetCollateralPosition(ctx, position)
	k.UpdateSystemMetrics(ctx)
	
	return &types.MsgAddCollateralResponse{
		NewCollateralRatio: position.CollateralizationRatio.String(),
		TotalCollateral:    position.Collateral,
	}, nil
}

func (k msgServer) WithdrawCollateral(goCtx context.Context, msg *types.MsgWithdrawCollateral) (*types.MsgWithdrawCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	
	position, exists := k.GetCollateralPosition(ctx, msg.Owner)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrPositionNotFound, "no collateral position found")
	}
	
	if msg.Amount.Amount.GT(position.Collateral.Amount) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAmount, "withdrawal amount exceeds collateral")
	}
	
	newCollateral := position.Collateral.Sub(msg.Amount)
	newRatio := k.CalculateCollateralRatio(ctx, newCollateral, position.Debt)
	
	if newRatio.LT(params.MinCollateralRatio) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientCollateral, "withdrawal would violate minimum collateral ratio")
	}
	
	ownerAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddr, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}
	
	position.Collateral = newCollateral
	position.CollateralizationRatio = newRatio
	position.LastUpdate = time.Now()
	
	k.SetCollateralPosition(ctx, position)
	k.UpdateSystemMetrics(ctx)
	
	return &types.MsgWithdrawCollateralResponse{
		NewCollateralRatio:   newRatio.String(),
		RemainingCollateral:  position.Collateral,
	}, nil
}

func (k msgServer) UpdateOraclePrice(goCtx context.Context, msg *types.MsgUpdateOraclePrice) (*types.MsgUpdateOraclePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	
	isWhitelisted := false
	for _, oracle := range params.OracleWhitelist {
		if oracle == msg.Oracle {
			isWhitelisted = true
			break
		}
	}
	
	if !isWhitelisted {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "oracle not whitelisted")
	}
	
	oldPrice, exists := k.GetOraclePrice(ctx, msg.Asset)
	previousPrice := "0"
	if exists {
		previousPrice = oldPrice.Price.String()
	}
	
	newPrice := types.OraclePrice{
		Asset:     msg.Asset,
		Price:     msg.Price,
		Source:    msg.Oracle,
		Timestamp: time.Now(),
	}
	
	k.SetOraclePrice(ctx, newPrice)
	
	k.IterateCollateralPositions(ctx, func(position types.CollateralPosition) bool {
		position.CollateralizationRatio = k.CalculateCollateralRatio(ctx, position.Collateral, position.Debt)
		position.IsLiquidatable = k.IsPositionLiquidatable(ctx, position)
		k.SetCollateralPosition(ctx, position)
		return false
	})
	
	k.UpdateSystemMetrics(ctx)
	
	return &types.MsgUpdateOraclePriceResponse{
		Accepted:      true,
		PreviousPrice: previousPrice,
	}, nil
}

func (k msgServer) DepositStability(goCtx context.Context, msg *types.MsgDepositStability) (*types.MsgDepositStabilityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	depositorAddr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositorAddr, types.ModuleName, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}
	
	pool := k.GetStabilityPool(ctx)
	
	var providerFound bool
	for i, provider := range pool.Providers {
		if provider.Address == msg.Depositor {
			pool.Providers[i].Deposit = pool.Providers[i].Deposit.Add(msg.Amount)
			providerFound = true
			break
		}
	}
	
	if !providerFound {
		pool.Providers = append(pool.Providers, types.StabilityProvider{
			Address:       msg.Depositor,
			Deposit:       msg.Amount,
			RewardsEarned: sdk.NewCoin("stst", sdk.ZeroInt()),
			DepositTime:   time.Now(),
		})
	}
	
	pool.TotalDeposits = pool.TotalDeposits.Add(msg.Amount)
	k.SetStabilityPool(ctx, pool)
	
	poolShare := sdk.NewDecFromInt(msg.Amount.Amount).Quo(sdk.NewDecFromInt(pool.TotalDeposits.Amount))
	
	return &types.MsgDepositStabilityResponse{
		TotalDeposit: msg.Amount,
		PoolShare:    poolShare.String(),
	}, nil
}

func (k msgServer) WithdrawStability(goCtx context.Context, msg *types.MsgWithdrawStability) (*types.MsgWithdrawStabilityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	pool := k.GetStabilityPool(ctx)
	
	var provider *types.StabilityProvider
	var providerIndex int
	for i, p := range pool.Providers {
		if p.Address == msg.Withdrawer {
			provider = &pool.Providers[i]
			providerIndex = i
			break
		}
	}
	
	if provider == nil {
		return nil, sdkerrors.Wrap(types.ErrProviderNotFound, "stability provider not found")
	}
	
	if msg.Amount.Amount.GT(provider.Deposit.Amount) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAmount, "withdrawal amount exceeds deposit")
	}
	
	withdrawerAddr, err := sdk.AccAddressFromBech32(msg.Withdrawer)
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawerAddr, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}
	
	rewards := provider.RewardsEarned
	if rewards.IsPositive() {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawerAddr, sdk.NewCoins(rewards))
		if err != nil {
			return nil, err
		}
		provider.RewardsEarned = sdk.NewCoin("stst", sdk.ZeroInt())
	}
	
	provider.Deposit = provider.Deposit.Sub(msg.Amount)
	pool.TotalDeposits = pool.TotalDeposits.Sub(msg.Amount)
	
	if provider.Deposit.IsZero() {
		pool.Providers = append(pool.Providers[:providerIndex], pool.Providers[providerIndex+1:]...)
	}
	
	k.SetStabilityPool(ctx, pool)
	
	return &types.MsgWithdrawStabilityResponse{
		Withdrawn: msg.Amount,
		Rewards:   rewards,
	}, nil
}

func (k msgServer) BidAuction(goCtx context.Context, msg *types.MsgBidAuction) (*types.MsgBidAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	auction, exists := k.GetLiquidationAuction(ctx, msg.AuctionId)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrAuctionNotFound, "auction not found")
	}
	
	if auction.Status != "active" {
		return nil, sdkerrors.Wrap(types.ErrAuctionInactive, "auction is not active")
	}
	
	if time.Now().After(auction.EndTime) {
		auction.Status = "expired"
		k.SetLiquidationAuction(ctx, auction)
		return nil, sdkerrors.Wrap(types.ErrAuctionExpired, "auction has expired")
	}
	
	secondsElapsed := time.Since(auction.StartTime).Seconds()
	decayFactor := sdk.NewDec(1).Sub(auction.PriceDecayRate.Mul(sdk.NewDec(int64(secondsElapsed))))
	currentPrice := auction.StartingPrice.Mul(decayFactor)
	
	requiredPayment := sdk.NewDecFromInt(auction.Collateral.Amount).Mul(currentPrice).TruncateInt()
	
	if msg.BidAmount.Amount.LT(requiredPayment) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientBid, "bid amount is insufficient")
	}
	
	bidderAddr, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAddr, types.ModuleName, 
		sdk.NewCoins(sdk.NewCoin("ssusd", requiredPayment)))
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ssusd", requiredPayment)))
	if err != nil {
		return nil, err
	}
	
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidderAddr, sdk.NewCoins(auction.Collateral))
	if err != nil {
		return nil, err
	}
	
	refund := msg.BidAmount.Amount.Sub(requiredPayment)
	
	auction.Status = "completed"
	k.SetLiquidationAuction(ctx, auction)
	
	return &types.MsgBidAuctionResponse{
		Success:        true,
		CollateralWon:  auction.Collateral,
		Refund:         sdk.NewCoin("ssusd", refund),
	}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	k.SetParams(ctx, msg.Params)
	
	return &types.MsgUpdateParamsResponse{}, nil
}