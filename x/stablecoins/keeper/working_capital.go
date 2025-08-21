package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/math"
	"github.com/stateset/core/x/stablecoins/types"
	"github.com/google/uuid"
)

// RequestWorkingCapital creates a new working capital request
func (k Keeper) RequestWorkingCapital(ctx sdk.Context, msg *types.MsgRequestWorkingCapital) (*types.WorkingCapitalRequest, error) {
	requestId := uuid.New().String()
	
	// Validate stablecoin exists
	_, found := k.GetStablecoin(ctx, msg.StablecoinDenom)
	if !found {
		return nil, types.ErrStablecoinNotFound
	}
	
	// Calculate risk score based on business profile
	riskScore := k.calculateRiskScore(ctx, msg.BusinessProfile)
	
	request := types.WorkingCapitalRequest{
		Id:                 requestId,
		Requester:          msg.Creator,
		BusinessProfile:    msg.BusinessProfile,
		RequestedAmount:    msg.Amount,
		StablecoinDenom:    msg.StablecoinDenom,
		Purpose:            msg.Purpose,
		OrderId:            msg.OrderId,
		ProposedCollateral: msg.Collateral,
		RequestedTerm:      msg.Term,
		Status:             types.LoanStatusPending,
		RiskScore:          riskScore,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	
	k.SetWorkingCapitalRequest(ctx, request)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"working_capital_requested",
			sdk.NewAttribute("request_id", requestId),
			sdk.NewAttribute("requester", msg.Creator),
			sdk.NewAttribute("amount", msg.Amount.String()),
		),
	})
	
	return &request, nil
}

// ApproveWorkingCapital approves a working capital request
func (k Keeper) ApproveWorkingCapital(ctx sdk.Context, msg *types.MsgApproveWorkingCapital) (*types.WorkingCapitalLoan, error) {
	request, found := k.GetWorkingCapitalRequest(ctx, msg.RequestId)
	if !found {
		return nil, types.ErrInvalidRequest
	}
	
	if request.Status != types.LoanStatusPending {
		return nil, types.ErrInvalidRequestStatus
	}
	
	loanId := uuid.New().String()
	now := time.Now()
	dueDate := now.Add(time.Duration(request.RequestedTerm) * 24 * time.Hour)
	
	loan := types.WorkingCapitalLoan{
		Id:               loanId,
		Borrower:         request.Requester,
		Amount:           msg.ApprovedAmount,
		StablecoinDenom:  request.StablecoinDenom,
		Purpose:          request.Purpose,
		OrderId:          request.OrderId,
		Status:           types.LoanStatusApproved,
		InterestRate:     msg.InterestRate,
		Term:             request.RequestedTerm,
		Collateral:       request.ProposedCollateral,
		CollateralRatio:  msg.CollateralRatio,
		DueDate:          dueDate,
		RepaidAmount:     sdk.NewCoins(),
		OutstandingAmount: msg.ApprovedAmount,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	
	k.SetWorkingCapitalLoan(ctx, loan)
	
	// Update request status
	request.Status = types.LoanStatusApproved
	request.ApprovedAmount = msg.ApprovedAmount
	request.ApprovedRate = msg.InterestRate
	request.ApprovedBy = msg.Authority
	approvedAt := time.Now()
	request.ApprovedAt = &approvedAt
	request.UpdatedAt = time.Now()
	k.SetWorkingCapitalRequest(ctx, request)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"working_capital_approved",
			sdk.NewAttribute("loan_id", loanId),
			sdk.NewAttribute("request_id", msg.RequestId),
			sdk.NewAttribute("borrower", request.Requester),
			sdk.NewAttribute("amount", msg.ApprovedAmount.String()),
		),
	})
	
	return &loan, nil
}

// DisburseWorkingCapital disburses funds from a capital pool to the borrower
func (k Keeper) DisburseWorkingCapital(ctx sdk.Context, msg *types.MsgDisburseWorkingCapital) error {
	loan, found := k.GetWorkingCapitalLoan(ctx, msg.LoanId)
	if !found {
		return types.ErrLoanNotFound
	}
	
	if loan.Status != types.LoanStatusApproved {
		return types.ErrLoanAlreadyDisbursed
	}
	
	pool, found := k.GetWorkingCapitalPool(ctx, msg.PoolId)
	if !found {
		return types.ErrPoolNotFound
	}
	
	// Check pool has sufficient funds
	if pool.AvailableFunds.IsAllLT(loan.Amount) {
		return types.ErrInsufficientPoolFunds
	}
	
	// Transfer funds from pool to borrower
	borrowerAddr, err := sdk.AccAddressFromBech32(loan.Borrower)
	if err != nil {
		return err
	}
	
	poolAddr := k.GetPoolAddress(ctx, msg.PoolId)
	err = k.bankKeeper.SendCoins(ctx, poolAddr, borrowerAddr, loan.Amount)
	if err != nil {
		return err
	}
	
	// Update loan status
	now := time.Now()
	loan.Status = types.LoanStatusDisbursed
	loan.DisbursedAt = &now
	loan.UpdatedAt = now
	k.SetWorkingCapitalLoan(ctx, loan)
	
	// Update pool balances
	pool.AvailableFunds = pool.AvailableFunds.Sub(loan.Amount...)
	pool.LentFunds = pool.LentFunds.Add(loan.Amount...)
	pool.UpdatedAt = time.Now()
	k.SetWorkingCapitalPool(ctx, pool)
	
	// Create repayment schedule
	k.createRepaymentSchedule(ctx, loan)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"working_capital_disbursed",
			sdk.NewAttribute("loan_id", msg.LoanId),
			sdk.NewAttribute("pool_id", msg.PoolId),
			sdk.NewAttribute("borrower", loan.Borrower),
			sdk.NewAttribute("amount", loan.Amount.String()),
		),
	})
	
	return nil
}

// RepayWorkingCapital processes a loan repayment
func (k Keeper) RepayWorkingCapital(ctx sdk.Context, msg *types.MsgRepayWorkingCapital) error {
	loan, found := k.GetWorkingCapitalLoan(ctx, msg.LoanId)
	if !found {
		return types.ErrLoanNotFound
	}
	
	if loan.Status != types.LoanStatusDisbursed && loan.Status != types.LoanStatusActive {
		return types.ErrInvalidRequestStatus
	}
	
	// Calculate interest
	interest := k.calculateInterest(ctx, loan)
	totalDue := loan.OutstandingAmount.Add(interest...)
	
	// Ensure repayment amount is valid
	if msg.Amount.IsAllGT(totalDue) {
		return types.ErrInvalidAmount
	}
	
	// Transfer repayment from borrower to pool
	borrowerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	
	poolAddr := k.GetPoolAddressByLoan(ctx, msg.LoanId)
	err = k.bankKeeper.SendCoins(ctx, borrowerAddr, poolAddr, msg.Amount)
	if err != nil {
		return err
	}
	
	// Update loan repayment status
	loan.RepaidAmount = loan.RepaidAmount.Add(msg.Amount...)
	loan.OutstandingAmount = loan.OutstandingAmount.Sub(msg.Amount...)
	
	if loan.OutstandingAmount.IsZero() {
		loan.Status = types.LoanStatusRepaid
	} else {
		loan.Status = types.LoanStatusActive
	}
	
	loan.UpdatedAt = time.Now()
	k.SetWorkingCapitalLoan(ctx, loan)
	
	// Update pool balances
	pool := k.GetPoolByLoan(ctx, msg.LoanId)
	pool.AvailableFunds = pool.AvailableFunds.Add(msg.Amount...)
	pool.LentFunds = pool.LentFunds.Sub(msg.Amount...)
	pool.UpdatedAt = time.Now()
	k.SetWorkingCapitalPool(ctx, pool)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"working_capital_repaid",
			sdk.NewAttribute("loan_id", msg.LoanId),
			sdk.NewAttribute("borrower", msg.Creator),
			sdk.NewAttribute("amount", msg.Amount.String()),
			sdk.NewAttribute("outstanding", loan.OutstandingAmount.String()),
		),
	})
	
	return nil
}

// CreateCapitalPool creates a new capital pool for lending
func (k Keeper) CreateCapitalPool(ctx sdk.Context, msg *types.MsgCreateCapitalPool) (*types.WorkingCapitalPool, error) {
	poolId := uuid.New().String()
	
	// Validate stablecoin exists
	_, found := k.GetStablecoin(ctx, msg.StablecoinDenom)
	if !found {
		return nil, types.ErrStablecoinNotFound
	}
	
	// Transfer initial funds from creator to pool
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	
	poolAddr := k.GetPoolAddress(ctx, poolId)
	if !msg.InitialFunds.IsZero() {
		err = k.bankKeeper.SendCoins(ctx, creatorAddr, poolAddr, msg.InitialFunds)
		if err != nil {
			return nil, err
		}
	}
	
	pool := types.WorkingCapitalPool{
		PoolId:             poolId,
		Name:               msg.Name,
		StablecoinDenom:    msg.StablecoinDenom,
		TotalFunds:         msg.InitialFunds,
		AvailableFunds:     msg.InitialFunds,
		LentFunds:          sdk.NewCoins(),
		DefaultRate:        math.LegacyZeroDec(),
		AverageInterestRate: math.LegacyZeroDec(),
		MinLoanAmount:      msg.MinLoanAmount,
		MaxLoanAmount:      msg.MaxLoanAmount,
		MinCollateralRatio: msg.MinCollateralRatio,
		MaxTerm:            msg.MaxTerm,
		Status:             types.PoolStatusActive,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	
	k.SetWorkingCapitalPool(ctx, pool)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"capital_pool_created",
			sdk.NewAttribute("pool_id", poolId),
			sdk.NewAttribute("name", msg.Name),
			sdk.NewAttribute("initial_funds", msg.InitialFunds.String()),
		),
	})
	
	return &pool, nil
}

// FundCapitalPool adds funds to an existing capital pool
func (k Keeper) FundCapitalPool(ctx sdk.Context, msg *types.MsgFundCapitalPool) error {
	pool, found := k.GetWorkingCapitalPool(ctx, msg.PoolId)
	if !found {
		return types.ErrPoolNotFound
	}
	
	// Transfer funds from funder to pool
	funderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	
	poolAddr := k.GetPoolAddress(ctx, msg.PoolId)
	err = k.bankKeeper.SendCoins(ctx, funderAddr, poolAddr, msg.Amount)
	if err != nil {
		return err
	}
	
	// Update pool balances
	pool.TotalFunds = pool.TotalFunds.Add(msg.Amount...)
	pool.AvailableFunds = pool.AvailableFunds.Add(msg.Amount...)
	pool.UpdatedAt = time.Now()
	k.SetWorkingCapitalPool(ctx, pool)
	
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"capital_pool_funded",
			sdk.NewAttribute("pool_id", msg.PoolId),
			sdk.NewAttribute("funder", msg.Creator),
			sdk.NewAttribute("amount", msg.Amount.String()),
		),
	})
	
	return nil
}

// Helper functions

func (k Keeper) calculateRiskScore(ctx sdk.Context, profile *types.BusinessProfile) math.LegacyDec {
	// Simple risk scoring algorithm (can be enhanced)
	baseScore := math.LegacyNewDec(100)
	
	// Credit score factor
	creditFactor := math.LegacyNewDec(int64(profile.CreditScore)).Quo(math.LegacyNewDec(850))
	
	// Years in business factor
	yearsFactor := math.LegacyNewDec(int64(profile.YearsInBusiness)).Quo(math.LegacyNewDec(10))
	if yearsFactor.GT(math.LegacyOneDec()) {
		yearsFactor = math.LegacyOneDec()
	}
	
	// Calculate final score
	riskScore := baseScore.Mul(creditFactor).Mul(yearsFactor)
	
	return riskScore
}

func (k Keeper) calculateInterest(ctx sdk.Context, loan types.WorkingCapitalLoan) sdk.Coins {
	// Simple interest calculation
	principal := loan.Amount[0].Amount
	rate := loan.InterestRate
	term := math.LegacyNewDec(loan.Term)
	
	interest := math.LegacyNewDecFromInt(principal).Mul(rate).Mul(term).Quo(math.LegacyNewDec(365))
	
	return sdk.NewCoins(sdk.NewCoin(loan.StablecoinDenom, interest.TruncateInt()))
}

func (k Keeper) createRepaymentSchedule(ctx sdk.Context, loan types.WorkingCapitalLoan) {
	// Create simple repayment schedule (can be enhanced for installments)
	schedule := types.RepaymentSchedule{
		LoanId:            loan.Id,
		InstallmentNumber: 1,
		DueDate:           loan.DueDate,
		PrincipalAmount:   loan.Amount,
		InterestAmount:    k.calculateInterest(ctx, loan),
		TotalAmount:       loan.Amount.Add(k.calculateInterest(ctx, loan)...),
		Status:            types.RepaymentStatusPending,
	}
	
	k.SetRepaymentSchedule(ctx, schedule)
}

func (k Keeper) GetPoolAddress(ctx sdk.Context, poolId string) sdk.AccAddress {
	return sdk.AccAddress([]byte(fmt.Sprintf("pool_%s", poolId)))
}

func (k Keeper) GetPoolAddressByLoan(ctx sdk.Context, loanId string) sdk.AccAddress {
	// Logic to find pool associated with loan
	// This is simplified - in production you'd store the pool-loan mapping
	return sdk.AccAddress([]byte("default_pool"))
}

func (k Keeper) GetPoolByLoan(ctx sdk.Context, loanId string) types.WorkingCapitalPool {
	// Logic to find pool associated with loan
	// This is simplified - in production you'd store the pool-loan mapping
	pool, _ := k.GetWorkingCapitalPool(ctx, "default_pool_id")
	return pool
}

// CRUD operations for WorkingCapitalRequest

func (k Keeper) SetWorkingCapitalRequest(ctx sdk.Context, request types.WorkingCapitalRequest) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&request)
	store.Set(append(types.WorkingCapitalRequestKeyPrefix, types.WorkingCapitalRequestKey(request.Id)...), b)
}

func (k Keeper) GetWorkingCapitalRequest(ctx sdk.Context, requestId string) (types.WorkingCapitalRequest, bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(append(types.WorkingCapitalRequestKeyPrefix, types.WorkingCapitalRequestKey(requestId)...))
	if b == nil {
		return types.WorkingCapitalRequest{}, false
	}
	
	var request types.WorkingCapitalRequest
	k.cdc.MustUnmarshal(b, &request)
	return request, true
}

// CRUD operations for WorkingCapitalLoan

func (k Keeper) SetWorkingCapitalLoan(ctx sdk.Context, loan types.WorkingCapitalLoan) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&loan)
	store.Set(append(types.WorkingCapitalLoanKeyPrefix, types.WorkingCapitalLoanKey(loan.Id)...), b)
}

func (k Keeper) GetWorkingCapitalLoan(ctx sdk.Context, loanId string) (types.WorkingCapitalLoan, bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(append(types.WorkingCapitalLoanKeyPrefix, types.WorkingCapitalLoanKey(loanId)...))
	if b == nil {
		return types.WorkingCapitalLoan{}, false
	}
	
	var loan types.WorkingCapitalLoan
	k.cdc.MustUnmarshal(b, &loan)
	return loan, true
}

// CRUD operations for WorkingCapitalPool

func (k Keeper) SetWorkingCapitalPool(ctx sdk.Context, pool types.WorkingCapitalPool) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&pool)
	store.Set(append(types.WorkingCapitalPoolKeyPrefix, types.WorkingCapitalPoolKey(pool.PoolId)...), b)
}

func (k Keeper) GetWorkingCapitalPool(ctx sdk.Context, poolId string) (types.WorkingCapitalPool, bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(append(types.WorkingCapitalPoolKeyPrefix, types.WorkingCapitalPoolKey(poolId)...))
	if b == nil {
		return types.WorkingCapitalPool{}, false
	}
	
	var pool types.WorkingCapitalPool
	k.cdc.MustUnmarshal(b, &pool)
	return pool, true
}

// CRUD operations for RepaymentSchedule

func (k Keeper) SetRepaymentSchedule(ctx sdk.Context, schedule types.RepaymentSchedule) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&schedule)
	store.Set(append(types.RepaymentScheduleKeyPrefix, types.RepaymentScheduleKey(schedule.LoanId, schedule.InstallmentNumber)...), b)
}

func (k Keeper) GetRepaymentSchedule(ctx sdk.Context, loanId string, installmentNumber int) (types.RepaymentSchedule, bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(append(types.RepaymentScheduleKeyPrefix, types.RepaymentScheduleKey(loanId, installmentNumber)...))
	if b == nil {
		return types.RepaymentSchedule{}, false
	}
	
	var schedule types.RepaymentSchedule
	k.cdc.MustUnmarshal(b, &schedule)
	return schedule, true
}