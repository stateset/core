package keeper

import (
	"encoding/binary"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/types"
)

var (
	paramsKey      = []byte{0x02}
	nextVaultIDKey = []byte{0x03}
)

// Keeper manages stablecoin state.
type Keeper struct {
	storeKey         storetypes.StoreKey
	bankKeeper       types.BankKeeper
	accountKeeper    types.AccountKeeper
	oracleKeeper     types.OracleKeeper
	complianceKeeper types.ComplianceKeeper
}

// NewKeeper instantiates a new keeper.
func NewKeeper(
	_ codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bank types.BankKeeper,
	account types.AccountKeeper,
	oracle types.OracleKeeper,
	compliance types.ComplianceKeeper,
) Keeper {
	return Keeper{
		storeKey:         storeKey,
		bankKeeper:       bank,
		accountKeeper:    account,
		oracleKeeper:     oracle,
		complianceKeeper: compliance,
	}
}

func (k Keeper) ensureModuleAccount(ctx sdk.Context) error {
	macc := k.accountKeeper.GetModuleAddress(types.ModuleAccountName)
	if macc == nil {
		return errorsmod.Wrap(types.ErrModuleAccountNotFound, "stablecoin module account is not set")
	}
	return nil
}

// GetParams retrieves module params.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(paramsKey) {
		return types.DefaultParams()
	}
	var params types.Params
	types.ModuleCdc.MustUnmarshalJSON(store.Get(paramsKey), &params)
	return params
}

// SetParams updates module params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	store.Set(paramsKey, types.ModuleCdc.MustMarshalJSON(&params))
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

	params := k.GetParams(ctx)
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
	vault, found := k.GetVault(ctx, id)
	if !found {
		return types.ErrVaultNotFound
	}
	if vault.Owner != owner.String() {
		return types.ErrUnauthorized
	}

	params := k.GetParams(ctx)
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
		return nil, errorsmod.Wrap(types.ErrUnderCollateralized, "vault still healthy")
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
	price, err := k.oracleKeeper.GetPriceDec(sdk.WrapSDKContext(ctx), collateral.Denom)
	if err != nil {
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
	k.setNextVaultID(ctx, state.NextVaultId)
	for _, vault := range state.Vaults {
		k.setVault(ctx, vault)
	}
}

// ExportGenesis exports module state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Params = k.GetParams(ctx)
	state.NextVaultId = k.getNextVaultID(ctx)
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
