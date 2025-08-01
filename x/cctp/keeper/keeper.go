package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/stateset/core/x/cctp/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Owner management
func (k Keeper) GetOwner(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.OwnerKey)
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetOwner(ctx sdk.Context, owner string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.OwnerKey, []byte(owner))
}

func (k Keeper) GetPendingOwner(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PendingOwnerKey)
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetPendingOwner(ctx sdk.Context, pendingOwner string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PendingOwnerKey, []byte(pendingOwner))
}

func (k Keeper) DeletePendingOwner(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PendingOwnerKey)
}

// Attester Manager management
func (k Keeper) GetAttesterManager(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AttesterManagerKey)
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetAttesterManager(ctx sdk.Context, attesterManager string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AttesterManagerKey, []byte(attesterManager))
}

// Token Controller management
func (k Keeper) GetTokenController(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenControllerKey)
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetTokenController(ctx sdk.Context, tokenController string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.TokenControllerKey, []byte(tokenController))
}

// Pauser management
func (k Keeper) GetPauser(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PauserKey)
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k Keeper) SetPauser(ctx sdk.Context, pauser string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PauserKey, []byte(pauser))
}

// Signature Threshold management
func (k Keeper) GetSignatureThreshold(ctx sdk.Context) (uint32, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SignatureThresholdKey)
	if bz == nil {
		return 0, false
	}
	var threshold uint32
	k.cdc.MustUnmarshal(bz, &threshold)
	return threshold, true
}

func (k Keeper) SetSignatureThreshold(ctx sdk.Context, threshold uint32) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&threshold)
	store.Set(types.SignatureThresholdKey, bz)
}

// Nonce management
func (k Keeper) GetNextAvailableNonce(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextAvailableNonceKey)
	if bz == nil {
		return 0
	}
	var nonce uint64
	k.cdc.MustUnmarshal(bz, &nonce)
	return nonce
}

func (k Keeper) SetNextAvailableNonce(ctx sdk.Context, nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&nonce)
	store.Set(types.NextAvailableNonceKey, bz)
}

func (k Keeper) GetAndIncrementNextAvailableNonce(ctx sdk.Context) uint64 {
	nonce := k.GetNextAvailableNonce(ctx)
	k.SetNextAvailableNonce(ctx, nonce+1)
	return nonce
}

// Pause state management
func (k Keeper) GetBurningAndMintingPaused(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.BurningAndMintingPausedKey)
	if bz == nil {
		return false
	}
	var paused bool
	k.cdc.MustUnmarshal(bz, &paused)
	return paused
}

func (k Keeper) SetBurningAndMintingPaused(ctx sdk.Context, paused bool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&paused)
	store.Set(types.BurningAndMintingPausedKey, bz)
}

func (k Keeper) GetSendingAndReceivingMessagesPaused(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SendingAndReceivingPausedKey)
	if bz == nil {
		return false
	}
	var paused bool
	k.cdc.MustUnmarshal(bz, &paused)
	return paused
}

func (k Keeper) SetSendingAndReceivingMessagesPaused(ctx sdk.Context, paused bool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&paused)
	store.Set(types.SendingAndReceivingPausedKey, bz)
}

// Max message body size management
func (k Keeper) GetMaxMessageBodySize(ctx sdk.Context) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.MaxMessageBodySizeKey)
	if bz == nil {
		return 0, false
	}
	var size uint64
	k.cdc.MustUnmarshal(bz, &size)
	return size, true
}

func (k Keeper) SetMaxMessageBodySize(ctx sdk.Context, size uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&size)
	store.Set(types.MaxMessageBodySizeKey, bz)
}

// Attester management
func (k Keeper) GetAttester(ctx sdk.Context, attester string) (types.Attester, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AttesterKey(attester))
	if bz == nil {
		return types.Attester{}, false
	}
	var att types.Attester
	k.cdc.MustUnmarshal(bz, &att)
	return att, true
}

func (k Keeper) SetAttester(ctx sdk.Context, attester types.Attester) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&attester)
	store.Set(types.AttesterKey(attester.Attester), bz)
}

func (k Keeper) DeleteAttester(ctx sdk.Context, attester string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AttesterKey(attester))
}

func (k Keeper) GetAllAttesters(ctx sdk.Context) []types.Attester {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetAttesterKeyPrefix())
	defer iterator.Close()

	var attesters []types.Attester
	for ; iterator.Valid(); iterator.Next() {
		var attester types.Attester
		k.cdc.MustUnmarshal(iterator.Value(), &attester)
		attesters = append(attesters, attester)
	}
	return attesters
}

func (k Keeper) GetEnabledAttesters(ctx sdk.Context) []types.Attester {
	attesters := k.GetAllAttesters(ctx)
	var enabled []types.Attester
	for _, attester := range attesters {
		if attester.Status == types.AttesterStatus_ENABLED {
			enabled = append(enabled, attester)
		}
	}
	return enabled
}

// Remote Token Messenger management
func (k Keeper) GetRemoteTokenMessenger(ctx sdk.Context, domain uint32) (types.RemoteTokenMessenger, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RemoteTokenMessengerKey(domain))
	if bz == nil {
		return types.RemoteTokenMessenger{}, false
	}
	var rtm types.RemoteTokenMessenger
	k.cdc.MustUnmarshal(bz, &rtm)
	return rtm, true
}

func (k Keeper) SetRemoteTokenMessenger(ctx sdk.Context, rtm types.RemoteTokenMessenger) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rtm)
	store.Set(types.RemoteTokenMessengerKey(rtm.DomainId), bz)
}

func (k Keeper) DeleteRemoteTokenMessenger(ctx sdk.Context, domain uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RemoteTokenMessengerKey(domain))
}

func (k Keeper) GetAllRemoteTokenMessengers(ctx sdk.Context) []types.RemoteTokenMessenger {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetRemoteTokenMessengerKeyPrefix())
	defer iterator.Close()

	var messengers []types.RemoteTokenMessenger
	for ; iterator.Valid(); iterator.Next() {
		var messenger types.RemoteTokenMessenger
		k.cdc.MustUnmarshal(iterator.Value(), &messenger)
		messengers = append(messengers, messenger)
	}
	return messengers
}

// Token Pair management
func (k Keeper) GetTokenPair(ctx sdk.Context, remoteDomain uint32, remoteToken []byte) (types.TokenPair, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenPairKey(remoteDomain, remoteToken))
	if bz == nil {
		return types.TokenPair{}, false
	}
	var tp types.TokenPair
	k.cdc.MustUnmarshal(bz, &tp)
	return tp, true
}

func (k Keeper) SetTokenPair(ctx sdk.Context, tp types.TokenPair) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&tp)
	store.Set(types.TokenPairKey(tp.RemoteDomain, tp.RemoteToken), bz)
}

func (k Keeper) DeleteTokenPair(ctx sdk.Context, remoteDomain uint32, remoteToken []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.TokenPairKey(remoteDomain, remoteToken))
}

func (k Keeper) GetAllTokenPairs(ctx sdk.Context) []types.TokenPair {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetTokenPairKeyPrefix())
	defer iterator.Close()

	var pairs []types.TokenPair
	for ; iterator.Valid(); iterator.Next() {
		var pair types.TokenPair
		k.cdc.MustUnmarshal(iterator.Value(), &pair)
		pairs = append(pairs, pair)
	}
	return pairs
}

// Used Nonce management
func (k Keeper) IsNonceUsed(ctx sdk.Context, sourceDomain uint32, nonce uint64) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.UsedNonceKey(sourceDomain, nonce)
	return store.Has(key)
}

func (k Keeper) SetNonceUsed(ctx sdk.Context, sourceDomain uint32, nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	key := types.UsedNonceKey(sourceDomain, nonce)
	// We just need to mark it as used, the value doesn't matter
	store.Set(key, []byte{1})
}

func (k Keeper) GetUsedNonces(ctx sdk.Context, sourceDomain uint32) []uint64 {
	store := ctx.KVStore(k.storeKey)
	prefix := types.GetUsedNonceKeyPrefixForDomain(sourceDomain)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	var nonces []uint64
	for ; iterator.Valid(); iterator.Next() {
		_, nonce, err := types.ParseUsedNonceKey(iterator.Key())
		if err == nil {
			nonces = append(nonces, nonce)
		}
	}
	return nonces
}

// Per Message Burn Limit management
func (k Keeper) GetPerMessageBurnLimit(ctx sdk.Context, denom string) (sdk.Int, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PerMessageBurnLimitKey(denom))
	if bz == nil {
		return sdk.ZeroInt(), false
	}
	var limit sdk.Int
	k.cdc.MustUnmarshal(bz, &limit)
	return limit, true
}

func (k Keeper) SetPerMessageBurnLimit(ctx sdk.Context, denom string, limit sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&limit)
	store.Set(types.PerMessageBurnLimitKey(denom), bz)
}

func (k Keeper) DeletePerMessageBurnLimit(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PerMessageBurnLimitKey(denom))
}

func (k Keeper) GetAllPerMessageBurnLimits(ctx sdk.Context) []types.PerMessageBurnLimit {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetPerMessageBurnLimitKeyPrefix())
	defer iterator.Close()

	var limits []types.PerMessageBurnLimit
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[len(types.GetPerMessageBurnLimitKeyPrefix()):])
		
		var limit sdk.Int
		k.cdc.MustUnmarshal(iterator.Value(), &limit)
		
		limits = append(limits, types.PerMessageBurnLimit{
			Denom: denom,
			Limit: limit,
		})
	}
	return limits
}

// Message tracking
func (k Keeper) SetSentMessage(ctx sdk.Context, domain uint32, nonce uint64, message []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.SentMessageKey(domain, nonce)
	store.Set(key, message)
}

func (k Keeper) GetSentMessage(ctx sdk.Context, domain uint32, nonce uint64) ([]byte, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.SentMessageKey(domain, nonce)
	bz := store.Get(key)
	if bz == nil {
		return nil, false
	}
	return bz, true
}

func (k Keeper) SetReceivedMessage(ctx sdk.Context, sourceDomain uint32, nonce uint64, message []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.ReceivedMessageKey(sourceDomain, nonce)
	store.Set(key, message)
}

func (k Keeper) GetReceivedMessage(ctx sdk.Context, sourceDomain uint32, nonce uint64) ([]byte, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ReceivedMessageKey(sourceDomain, nonce)
	bz := store.Get(key)
	if bz == nil {
		return nil, false
	}
	return bz, true
}

// Validation helpers
func (k Keeper) ValidateOwner(ctx sdk.Context, sender string) error {
	owner, found := k.GetOwner(ctx)
	if !found {
		return types.ErrOwnerNotSet
	}
	if sender != owner {
		return types.ErrNotOwner
	}
	return nil
}

func (k Keeper) ValidatePendingOwner(ctx sdk.Context, sender string) error {
	pendingOwner, found := k.GetPendingOwner(ctx)
	if !found {
		return types.ErrPendingOwnerNotSet
	}
	if sender != pendingOwner {
		return types.ErrNotPendingOwner
	}
	return nil
}

func (k Keeper) ValidateAttesterManager(ctx sdk.Context, sender string) error {
	attesterManager, found := k.GetAttesterManager(ctx)
	if !found {
		return types.ErrAttesterManagerNotSet
	}
	if sender != attesterManager {
		return types.ErrNotAttesterManager
	}
	return nil
}

func (k Keeper) ValidateTokenController(ctx sdk.Context, sender string) error {
	tokenController, found := k.GetTokenController(ctx)
	if !found {
		return types.ErrTokenControllerNotSet
	}
	if sender != tokenController {
		return types.ErrNotTokenController
	}
	return nil
}

func (k Keeper) ValidatePauser(ctx sdk.Context, sender string) error {
	pauser, found := k.GetPauser(ctx)
	if !found {
		return types.ErrPauserNotSet
	}
	if sender != pauser {
		return types.ErrNotPauser
	}
	return nil
}

// State validation
func (k Keeper) ValidateBurningAndMintingNotPaused(ctx sdk.Context) error {
	if k.GetBurningAndMintingPaused(ctx) {
		return types.ErrBurningAndMintingPaused
	}
	return nil
}

func (k Keeper) ValidateSendingAndReceivingNotPaused(ctx sdk.Context) error {
	if k.GetSendingAndReceivingMessagesPaused(ctx) {
		return types.ErrSendingAndReceivingPaused
	}
	return nil
}

func (k Keeper) ValidateBurnLimit(ctx sdk.Context, denom string, amount sdk.Int) error {
	limit, found := k.GetPerMessageBurnLimit(ctx, denom)
	if found && amount.GT(limit) {
		return types.ErrExceedsBurnLimit
	}
	return nil
}

func (k Keeper) ValidateMessageBodySize(ctx sdk.Context, messageBody []byte) error {
	maxSize, found := k.GetMaxMessageBodySize(ctx)
	if found && uint64(len(messageBody)) > maxSize {
		return types.ErrExceedsMaxMessageBodySize
	}
	return nil
}