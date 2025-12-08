package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/compliance/types"
)

// Keeper maintains the state and authority for compliance operations.
type Keeper struct {
	storeKey  storetypes.StoreKey
	authority string
}

// NewKeeper creates a new Keeper instance.
func NewKeeper(_ codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{storeKey: key, authority: authority}
}

// GetAuthority returns the authority address allowed to modify compliance data.
func (k Keeper) GetAuthority() string { return k.authority }

// SetAuthority updates the keeper authority (used during genesis or governance upgrades).
func (k *Keeper) SetAuthority(authority string) { k.authority = authority }

// SetProfile stores or updates a compliance profile.
func (k Keeper) SetProfile(ctx context.Context, profile types.Profile) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.ProfileKeyPrefix)
	bz := types.ModuleCdc.MustMarshalJSON(&profile)
	store.Set([]byte(profile.Address), bz)
}

// GetProfile retrieves a profile for the address.
func (k Keeper) GetProfile(ctx context.Context, addr sdk.AccAddress) (types.Profile, bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.ProfileKeyPrefix)
	bz := store.Get([]byte(addr.String()))
	if len(bz) == 0 {
		return types.Profile{}, false
	}
	var profile types.Profile
	types.ModuleCdc.MustUnmarshalJSON(bz, &profile)
	return profile, true
}

// RemoveProfile deletes a compliance profile.
func (k Keeper) RemoveProfile(ctx context.Context, addr sdk.AccAddress) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.ProfileKeyPrefix)
	store.Delete([]byte(addr.String()))
}

// AssertCompliant ensures an address is cleared for operations.
func (k Keeper) AssertCompliant(ctx context.Context, addr sdk.AccAddress) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	profile, found := k.GetProfile(ctx, addr)
	if !found {
		return types.ErrProfileNotFound
	}

	// Check if sanctioned
	if profile.Sanction {
		return types.ErrSanctionedAddress
	}

	// Check if blocked (status, jurisdiction)
	if profile.IsBlocked() {
		return errorsmod.Wrap(types.ErrComplianceBlocked, "profile is blocked from transacting")
	}

	// Check if expired
	if profile.IsExpired(sdkCtx.BlockTime()) {
		return errorsmod.Wrap(types.ErrProfileExpired, "compliance verification has expired")
	}

	return nil
}

// AssertCompliantForAmount ensures an address is cleared and within limits for a transaction.
func (k Keeper) AssertCompliantForAmount(ctx context.Context, addr sdk.AccAddress, amount sdk.Coin) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// First do basic compliance check
	if err := k.AssertCompliant(ctx, addr); err != nil {
		return err
	}

	profile, _ := k.GetProfile(ctx, addr)

	// Reset limits if needed (daily/monthly)
	profile = k.resetLimitsIfNeeded(sdkCtx, profile)

	// Check daily limit
	if !profile.DailyLimit.IsZero() && profile.DailyLimit.Denom == amount.Denom {
		newDaily := profile.DailyUsed.Amount.Add(amount.Amount)
		if newDaily.GT(profile.DailyLimit.Amount) {
			return errorsmod.Wrapf(types.ErrLimitExceeded,
				"daily limit exceeded: used %s + %s > limit %s",
				profile.DailyUsed.Amount, amount.Amount, profile.DailyLimit.Amount)
		}
	}

	// Check monthly limit
	if !profile.MonthlyLimit.IsZero() && profile.MonthlyLimit.Denom == amount.Denom {
		newMonthly := profile.MonthlyUsed.Amount.Add(amount.Amount)
		if newMonthly.GT(profile.MonthlyLimit.Amount) {
			return errorsmod.Wrapf(types.ErrLimitExceeded,
				"monthly limit exceeded: used %s + %s > limit %s",
				profile.MonthlyUsed.Amount, amount.Amount, profile.MonthlyLimit.Amount)
		}
	}

	// Check if enhanced due diligence is required for high-risk profiles
	if profile.RequiresEnhancedDueDiligence() && profile.KYCLevel != types.KYCEnhanced {
		return errorsmod.Wrap(types.ErrEnhancedDueDiligenceRequired,
			"enhanced KYC required for high-risk profile")
	}

	return nil
}

// RecordTransaction updates usage limits after a successful transaction.
func (k Keeper) RecordTransaction(ctx context.Context, addr sdk.AccAddress, amount sdk.Coin) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	profile, found := k.GetProfile(ctx, addr)
	if !found {
		return types.ErrProfileNotFound
	}

	// Reset limits if needed
	profile = k.resetLimitsIfNeeded(sdkCtx, profile)

	// Update daily usage
	if profile.DailyUsed.Denom == amount.Denom || profile.DailyUsed.IsZero() {
		profile.DailyUsed = sdk.NewCoin(amount.Denom, profile.DailyUsed.Amount.Add(amount.Amount))
	}

	// Update monthly usage
	if profile.MonthlyUsed.Denom == amount.Denom || profile.MonthlyUsed.IsZero() {
		profile.MonthlyUsed = sdk.NewCoin(amount.Denom, profile.MonthlyUsed.Amount.Add(amount.Amount))
	}

	k.SetProfile(ctx, profile)
	return nil
}

// resetLimitsIfNeeded resets daily/monthly limits if enough time has passed.
func (k Keeper) resetLimitsIfNeeded(ctx sdk.Context, profile types.Profile) types.Profile {
	now := ctx.BlockTime()

	// Reset daily limit if more than 24 hours since last reset
	if now.Sub(profile.LastLimitReset) >= 24*time.Hour {
		if !profile.DailyUsed.IsZero() {
			profile.DailyUsed = sdk.NewCoin(profile.DailyUsed.Denom, sdkmath.ZeroInt())
		}
	}

	// Reset monthly limit if we're in a new month
	if profile.LastLimitReset.Month() != now.Month() || profile.LastLimitReset.Year() != now.Year() {
		if !profile.MonthlyUsed.IsZero() {
			profile.MonthlyUsed = sdk.NewCoin(profile.MonthlyUsed.Denom, sdkmath.ZeroInt())
		}
	}

	// Update last reset time if any reset occurred
	if now.Sub(profile.LastLimitReset) >= 24*time.Hour {
		profile.LastLimitReset = now
		k.SetProfile(sdk.WrapSDKContext(ctx), profile)
	}

	return profile
}

// SuspendProfile suspends a profile with audit logging.
func (k Keeper) SuspendProfile(ctx context.Context, addr sdk.AccAddress, actor, reason string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	profile, found := k.GetProfile(ctx, addr)
	if !found {
		return types.ErrProfileNotFound
	}

	oldStatus := profile.Status
	profile.Status = types.StatusSuspended
	profile.UpdatedBy = actor
	profile.UpdatedAt = sdkCtx.BlockTime()
	profile.AddAuditEntry("suspended", actor, reason, oldStatus, types.StatusSuspended)

	k.SetProfile(ctx, profile)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"compliance_profile_suspended",
			sdk.NewAttribute("address", addr.String()),
			sdk.NewAttribute("actor", actor),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// ReactivateProfile reactivates a suspended profile with audit logging.
func (k Keeper) ReactivateProfile(ctx context.Context, addr sdk.AccAddress, actor, reason string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	profile, found := k.GetProfile(ctx, addr)
	if !found {
		return types.ErrProfileNotFound
	}

	if profile.Status != types.StatusSuspended {
		return errorsmod.Wrap(types.ErrInvalidProfileStatus, "profile is not suspended")
	}

	oldStatus := profile.Status
	profile.Status = types.StatusActive
	profile.UpdatedBy = actor
	profile.UpdatedAt = sdkCtx.BlockTime()
	profile.AddAuditEntry("reactivated", actor, reason, oldStatus, types.StatusActive)

	k.SetProfile(ctx, profile)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"compliance_profile_reactivated",
			sdk.NewAttribute("address", addr.String()),
			sdk.NewAttribute("actor", actor),
			sdk.NewAttribute("reason", reason),
		),
	)

	return nil
}

// UpdateKYCLevel updates the KYC level of a profile with audit logging.
func (k Keeper) UpdateKYCLevel(ctx context.Context, addr sdk.AccAddress, newLevel types.KYCLevel, actor, reason string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	profile, found := k.GetProfile(ctx, addr)
	if !found {
		return types.ErrProfileNotFound
	}

	oldLevel := profile.KYCLevel
	profile.KYCLevel = newLevel
	profile.UpdatedBy = actor
	profile.UpdatedAt = sdkCtx.BlockTime()
	profile.AddAuditEntry("kyc_updated", actor, reason, types.ProfileStatus(oldLevel), types.ProfileStatus(newLevel))

	// If upgrading to standard or enhanced, set verification time and expiry
	if newLevel == types.KYCStandard || newLevel == types.KYCEnhanced {
		profile.VerifiedAt = sdkCtx.BlockTime()
		profile.ExpiresAt = sdkCtx.BlockTime().AddDate(1, 0, 0) // 1 year validity
	}

	k.SetProfile(ctx, profile)

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"compliance_kyc_updated",
			sdk.NewAttribute("address", addr.String()),
			sdk.NewAttribute("old_level", string(oldLevel)),
			sdk.NewAttribute("new_level", string(newLevel)),
			sdk.NewAttribute("actor", actor),
		),
	)

	return nil
}

// IterateProfiles iterates through stored profiles.
func (k Keeper) IterateProfiles(ctx context.Context, cb func(types.Profile) bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.ProfileKeyPrefix)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		types.ModuleCdc.MustUnmarshalJSON(iterator.Value(), &profile)
		if cb(profile) {
			break
		}
	}
}

// ExportGenesis exports module state.
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	state := types.DefaultGenesis()
	state.Authority = k.authority
	k.IterateProfiles(ctx, func(profile types.Profile) bool {
		state.Profiles = append(state.Profiles, profile)
		return false
	})
	return state
}

// InitGenesis initializes module state from genesis.
func (k Keeper) InitGenesis(ctx context.Context, state *types.GenesisState) {
	if state == nil {
		state = types.DefaultGenesis()
	}
	k.authority = state.Authority
	for _, profile := range state.Profiles {
		k.SetProfile(ctx, profile)
	}
}
