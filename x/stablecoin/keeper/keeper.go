package keeper

import (
	"encoding/binary"
	"errors"
	"fmt"

	"cosmossdk.io/log"
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	oracletypes "github.com/stateset/core/x/oracle/types"
	"github.com/stateset/core/x/stablecoin/types"
)

var (
	nextVaultIDKey   = []byte{0x03} // Keep nextVaultIDKey as it's not a param
	paramsKey        = []byte("params")
	reserveParamsKey = []byte("reserve_params")
)

// Keeper manages stablecoin state.
type Keeper struct {
	storeKey         storetypes.StoreKey
	cdc              codec.BinaryCodec
	authority        string
	bankKeeper       types.BankKeeper
	accountKeeper    types.AccountKeeper
	oracleKeeper     types.OracleKeeper
	complianceKeeper types.ComplianceKeeper
	hooks            types.StablecoinHooks
}

// NewKeeper instantiates a new keeper.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authority string,
	bank types.BankKeeper,
	account types.AccountKeeper,
	oracle types.OracleKeeper,
	compliance types.ComplianceKeeper,
) Keeper {
	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		authority:        authority,
		bankKeeper:       bank,
		accountKeeper:    account,
		oracleKeeper:     oracle,
		complianceKeeper: compliance,
	}
}

// SetHooks sets the stablecoin hooks
func (k *Keeper) SetHooks(sh types.StablecoinHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set stablecoin hooks twice")
	}
	k.hooks = sh
	return k
}

// GetAuthority returns the module authority address
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) ensureModuleAccount(ctx sdk.Context) error {
	macc := k.accountKeeper.GetModuleAddress(types.ModuleAccountName)
	if macc == nil {
		return errorsmod.Wrap(types.ErrModuleAccountNotFound, "stablecoin module account is not set")
	}
	return nil
}

// GetParams retrieves module params (vault params).
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(paramsKey)
	if len(bz) == 0 {
		return types.DefaultParams()
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams updates module params (vault params).
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(paramsKey, bz)
}

// GetReserveParams retrieves reserve parameters.
func (k Keeper) GetReserveParams(ctx sdk.Context) types.ReserveParams {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(reserveParamsKey)
	if len(bz) == 0 {
		return types.DefaultReserveParams()
	}
	var params types.ReserveParams
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetReserveParams updates reserve parameters.
func (k Keeper) SetReserveParams(ctx sdk.Context, params types.ReserveParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(reserveParamsKey, bz)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeReserveParamsUpdated,
			sdk.NewAttribute(types.AttributeKeyReserveRatio, fmt.Sprintf("%d", params.MinReserveRatioBps)),
		),
	)
	return nil
}

func (k Keeper) getNextVaultID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(nextVaultIDKey) {
		return 1
	}
	return binary.BigEndian.Uint64(store.Get(nextVaultIDKey))
}

func (k Keeper) setNextVaultID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(nextVaultIDKey, bz)
}

func (k Keeper) setVault(ctx sdk.Context, vault types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VaultKeyPrefix)
	store.Set(mustBz(vault.Id), types.ModuleCdc.MustMarshalJSON(&vault))
}

func (k Keeper) GetVault(ctx sdk.Context, id uint64) (types.Vault, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VaultKeyPrefix)
	bz := store.Get(mustBz(id))
	if len(bz) == 0 {
		return types.Vault{}, false
	}
	var vault types.Vault
	types.ModuleCdc.MustUnmarshalJSON(bz, &vault)
	return vault, true
}

func (k Keeper) IterateVaults(ctx sdk.Context, cb func(types.Vault) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VaultKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var vault types.Vault
		types.ModuleCdc.MustUnmarshalJSON(iter.Value(), &vault)
		if cb(vault) {
			break
		}
	}
}

func (k Keeper) removeVault(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VaultKeyPrefix)
	store.Delete(mustBz(id))
}

// CreateVault creates a new vault with initial collateral and optional debt.
func (k Keeper) CreateVault(ctx sdk.Context, owner sdk.AccAddress, collateral sdk.Coin, debt sdk.Coin) (uint64, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	if err := k.ensureModuleAccount(ctx); err != nil {
		return 0, err
	}

	if !collateral.IsValid() || collateral.IsZero() {
		return 0, errorsmod.Wrap(types.ErrInvalidAmount, "collateral must be positive")
	}

	params := k.GetParams(ctx)
	if !params.VaultMintingEnabled {
		return 0, types.ErrVaultMintingDisabled
	}
	cp, ok := params.GetCollateralParam(collateral.Denom)
	if !ok || !cp.Active {
		return 0, types.ErrUnsupportedCollateral
	}

	if err := k.complianceKeeper.AssertCompliant(wrappedCtx, owner); err != nil {
		return 0, err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, owner, types.ModuleAccountName, sdk.NewCoins(collateral)); err != nil {
		return 0, err
	}

	vault := types.Vault{
		Owner:           owner.String(),
		Collateral:      collateral,
		CollateralDenom: collateral.Denom,
		Debt:            sdkmath.ZeroInt(),
		LastAccrued:     ctx.BlockHeight(),
	}

	nextID := k.getNextVaultID(ctx)
	vault.Id = nextID

	if !debt.IsZero() {
		if debt.Denom != types.StablecoinDenom {
			return 0, types.ErrInvalidAmount
		}
		if err := k.mintStablecoin(ctx, owner, &vault, cp, debt.Amount); err != nil {
			return 0, err
		}
	}

	k.setVault(ctx, vault)
	k.setNextVaultID(ctx, nextID+1)

	return vault.Id, nil
}

func (k Keeper) DepositCollateral(ctx sdk.Context, owner sdk.AccAddress, id uint64, collateral sdk.Coin) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	vault, found := k.GetVault(ctx, id)
	if !found {
		return types.ErrVaultNotFound
	}
	if vault.Owner != owner.String() {
		return types.ErrUnauthorized
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, owner, types.ModuleAccountName, sdk.NewCoins(collateral)); err != nil {
		return err
	}

	if vault.Collateral.Denom != collateral.Denom {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "collateral denom mismatch")
	}

	vault.Collateral = vault.Collateral.Add(collateral)
	k.setVault(ctx, vault)
	return nil
}

func (k Keeper) WithdrawCollateral(ctx sdk.Context, owner sdk.AccAddress, id uint64, collateral sdk.Coin) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	vault, found := k.GetVault(ctx, id)
	if !found {
		return types.ErrVaultNotFound
	}
	if vault.Owner != owner.String() {
		return types.ErrUnauthorized
	}
	if collateral.Denom != vault.Collateral.Denom {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "collateral denom mismatch")
	}
	if collateral.Amount.GT(vault.Collateral.Amount) {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "insufficient collateral")
	}

	params := k.GetParams(ctx)
	cp, ok := params.GetCollateralParam(vault.CollateralDenom)
	if !ok {
		return types.ErrUnsupportedCollateral
	}

	newCollateral := vault.Collateral.Sub(collateral)
	if err := k.assertCollateralization(ctx, newCollateral, vault.Debt, cp); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, owner, sdk.NewCoins(collateral)); err != nil {
		return err
	}

	vault.Collateral = newCollateral
	k.setVault(ctx, vault)
	return nil
}

func (k Keeper) MintStablecoin(ctx sdk.Context, owner sdk.AccAddress, id uint64, amount sdk.Coin) error {
	if amount.Denom != types.StablecoinDenom {
		return types.ErrInvalidAmount
	}

	params := k.GetParams(ctx)
	if !params.VaultMintingEnabled {
		return types.ErrVaultMintingDisabled
	}

	vault, found := k.GetVault(ctx, id)
	if !found {
		return types.ErrVaultNotFound
	}
	if vault.Owner != owner.String() {
		return types.ErrUnauthorized
	}

	cp, ok := params.GetCollateralParam(vault.CollateralDenom)
	if !ok {
		return types.ErrUnsupportedCollateral
	}

	if err := k.mintStablecoin(ctx, owner, &vault, cp, amount.Amount); err != nil {
		return err
	}

	k.setVault(ctx, vault)
	return nil
}

func (k Keeper) mintStablecoin(ctx sdk.Context, owner sdk.AccAddress, vault *types.Vault, cp types.CollateralParam, amount sdkmath.Int) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	newDebt := vault.Debt.Add(amount)
	if newDebt.GT(cp.DebtLimit) {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "debt limit exceeded")
	}

	if err := k.assertCollateralization(ctx, vault.Collateral, newDebt, cp); err != nil {
		return err
	}

	mintCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, amount))
	if err := k.bankKeeper.MintCoins(wrappedCtx, types.ModuleAccountName, mintCoins); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, owner, mintCoins); err != nil {
		return err
	}

	vault.Debt = newDebt
	vault.LastAccrued = ctx.BlockHeight()
	return nil
}

func (k Keeper) RepayStablecoin(ctx sdk.Context, owner sdk.AccAddress, id uint64, amount sdk.Coin) error {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	if amount.Denom != types.StablecoinDenom {
		return types.ErrInvalidAmount
	}
	vault, found := k.GetVault(ctx, id)
	if !found {
		return types.ErrVaultNotFound
	}
	if vault.Owner != owner.String() {
		return types.ErrUnauthorized
	}

	repay := amount.Amount
	if repay.GT(vault.Debt) {
		repay = vault.Debt
	}

	coins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, repay))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, owner, types.ModuleAccountName, coins); err != nil {
		return err
	}
	if err := k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, coins); err != nil {
		return err
	}

	vault.Debt = vault.Debt.Sub(repay)
	k.setVault(ctx, vault)
	return nil
}

func (k Keeper) LiquidateVault(ctx sdk.Context, liquidator sdk.AccAddress, id uint64) (sdk.Coins, error) {
	wrappedCtx := sdk.WrapSDKContext(ctx)
	vault, found := k.GetVault(ctx, id)
	if !found {
		return nil, types.ErrVaultNotFound
	}

	params := k.GetParams(ctx)
	cp, ok := params.GetCollateralParam(vault.CollateralDenom)
	if !ok {
		return nil, types.ErrUnsupportedCollateral
	}

	if err := k.assertCollateralization(ctx, vault.Collateral, vault.Debt, cp); err == nil {
		return nil, errorsmod.Wrap(types.ErrVaultHealthy, "vault still healthy")
	} else if !errors.Is(err, types.ErrUnderCollateralized) {
		// Fail safely if collateralization cannot be verified (e.g. missing/stale oracle price).
		return nil, err
	}

	// Liquidator must repay the outstanding stablecoin debt before collateral is released.
	if !vault.Debt.IsZero() {
		debtCoins := sdk.NewCoins(sdk.NewCoin(types.StablecoinDenom, vault.Debt))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(wrappedCtx, liquidator, types.ModuleAccountName, debtCoins); err != nil {
			return nil, err
		}
		if err := k.bankKeeper.BurnCoins(wrappedCtx, types.ModuleAccountName, debtCoins); err != nil {
			return nil, err
		}
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(wrappedCtx, types.ModuleAccountName, liquidator, sdk.NewCoins(vault.Collateral)); err != nil {
		return nil, err
	}

	k.removeVault(ctx, id)
	return sdk.NewCoins(vault.Collateral), nil
}

func (k Keeper) assertCollateralization(ctx sdk.Context, collateral sdk.Coin, debt sdkmath.Int, cp types.CollateralParam) error {
	if debt.IsZero() {
		return nil
	}
	price, err := k.oracleKeeper.GetPriceDecSafe(sdk.WrapSDKContext(ctx), collateral.Denom)
	if err != nil {
		if errors.Is(err, oracletypes.ErrPriceStale) {
			return types.ErrPriceStale
		}
		return types.ErrPriceNotFound
	}

	collateralValue := collateral.Amount.ToLegacyDec().Mul(price)
	required := sdkmath.LegacyNewDecFromInt(debt).Mul(cp.LiquidationRatio)

	if collateralValue.LT(required) {
		return types.ErrUnderCollateralized
	}
	return nil
}

// InitGenesis initializes module state from genesis.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.SetParams(ctx, state.Params)
	// Reserve-backed settings
	if err := k.SetReserveParams(ctx, state.ReserveParams); err != nil {
		panic(err)
	}
	k.SetReserve(ctx, state.Reserve)
	k.setNextDepositID(ctx, state.NextDepositId)
	k.setNextRedemptionID(ctx, state.NextRedemptionId)
	k.setNextAttestationID(ctx, state.NextAttestationId)

	for _, deposit := range state.ReserveDeposits {
		k.setReserveDeposit(ctx, deposit)
	}
	for _, redemption := range state.RedemptionRequests {
		k.setRedemptionRequest(ctx, redemption)
	}
	// Locked reserves are derived from pending redemption requests and are not part of genesis.
	locked := sdk.NewCoins()
	for _, redemption := range state.RedemptionRequests {
		if redemption.Status == types.RedeemStatusPending && redemption.OutputAmount.Denom != "" && redemption.OutputAmount.Amount.IsPositive() {
			locked = locked.Add(redemption.OutputAmount)
		}
	}
	k.setLockedReserves(ctx, locked)
	for _, stat := range state.DailyStats {
		k.SetDailyMintStats(ctx, stat)
	}
	for _, att := range state.Attestations {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OffChainAttestationKeyPrefix)
		store.Set(mustBz(att.Id), types.ModuleCdc.MustMarshalJSON(&att))
	}
	for _, addr := range state.ApprovedAttesters {
		k.SetApprovedAttester(ctx, addr, true)
	}

	k.setNextVaultID(ctx, state.NextVaultId)
	for _, vault := range state.Vaults {
		k.setVault(ctx, vault)
	}
}

// ExportGenesis exports module state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Params = k.GetParams(ctx)
	state.ReserveParams = k.GetReserveParams(ctx)
	state.Reserve = k.GetReserve(ctx)
	state.NextVaultId = k.getNextVaultID(ctx)

	state.NextDepositId = k.getNextDepositID(ctx)
	state.NextRedemptionId = k.getNextRedemptionID(ctx)
	state.NextAttestationId = k.getNextAttestationID(ctx)

	k.IterateReserveDeposits(ctx, func(deposit types.ReserveDeposit) bool {
		state.ReserveDeposits = append(state.ReserveDeposits, deposit)
		return false
	})
	k.IterateRedemptionRequests(ctx, func(request types.RedemptionRequest) bool {
		state.RedemptionRequests = append(state.RedemptionRequests, request)
		return false
	})
	k.IterateAttestations(ctx, func(att types.OffChainReserveAttestation) bool {
		state.Attestations = append(state.Attestations, att)
		return false
	})
	// Daily stats are keyed by date; export all present entries
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DailyMintStatsKeyPrefix)
	iter := store.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var stats types.DailyMintStats
		types.ModuleCdc.MustUnmarshalJSON(iter.Value(), &stats)
		state.DailyStats = append(state.DailyStats, stats)
	}
	// Approved attesters
	attStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ApprovedAttesterKeyPrefix)
	attIter := attStore.Iterator(nil, nil)
	defer attIter.Close()
	for ; attIter.Valid(); attIter.Next() {
		state.ApprovedAttesters = append(state.ApprovedAttesters, string(attIter.Key()))
	}

	k.IterateVaults(ctx, func(vault types.Vault) bool {
		state.Vaults = append(state.Vaults, vault)
		return false
	})
	return state
}

func mustBz(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
