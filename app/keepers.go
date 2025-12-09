package app

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
	compliancekeeper "github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	metricskeeper "github.com/stateset/core/x/metrics/keeper"
	metricstypes "github.com/stateset/core/x/metrics/types"
	oraclekeeper "github.com/stateset/core/x/oracle/keeper"
	oracletypes "github.com/stateset/core/x/oracle/types"
	paymentskeeper "github.com/stateset/core/x/payments/keeper"
	paymentstypes "github.com/stateset/core/x/payments/types"
	settlementkeeper "github.com/stateset/core/x/settlement/keeper"
	settlementtypes "github.com/stateset/core/x/settlement/types"
	stablecoinkeeper "github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
	treasurykeeper "github.com/stateset/core/x/treasury/keeper"
	treasurytypes "github.com/stateset/core/x/treasury/types"
	zkpverifykeeper "github.com/stateset/core/x/zkpverify/keeper"
	zkpverifytypes "github.com/stateset/core/x/zkpverify/types"
)

// AppKeepers holds all the keepers used in the application
type AppKeepers struct {
	// Cosmos SDK keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             *govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper

	// IBC keepers
	IBCKeeper        *ibckeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper

	// Scoped keepers
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// Stateset custom keepers
	OracleKeeper     oraclekeeper.Keeper
	ComplianceKeeper compliancekeeper.Keeper
	TreasuryKeeper   treasurykeeper.Keeper
	PaymentsKeeper   paymentskeeper.Keeper
	StablecoinKeeper stablecoinkeeper.Keeper
	SettlementKeeper settlementkeeper.Keeper
	CircuitKeeper    circuitkeeper.Keeper
	MetricsKeeper    metricskeeper.Keeper
	ZkpVerifyKeeper  zkpverifykeeper.Keeper
}

// InitCosmosKeepers initializes all the Cosmos SDK standard keepers
func (app *App) InitCosmosKeepers(
	appCodec codec.Codec,
	cdc *codec.LegacyAmino,
	keys map[string]*storetypes.KVStoreKey,
	tkeys map[string]*storetypes.TransientStoreKey,
	memKeys map[string]*storetypes.MemoryStoreKey,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	logger log.Logger,
) {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	accAddrCodec := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	valAddrCodec := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix())
	consAddrCodec := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix())

	// Init ParamsKeeper
	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// Init ConsensusParamsKeeper
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[consensusparamtypes.StoreKey]),
		authority,
		runtime.EventService{},
	)

	// Init CapabilityKeeper
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// Init AccountKeeper
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		accAddrCodec,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authority,
	)

	// Init BankKeeper
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.AccountKeeper,
		app.ModuleAccountAddrs(),
		authority,
		logger,
	)

	// Init StakingKeeper
	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		authority,
		valAddrCodec,
		consAddrCodec,
	)

	// Init MintKeeper
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[minttypes.StoreKey]),
		app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authority,
	)

	// Init DistrKeeper
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[distrtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		authtypes.FeeCollectorName,
		authority,
	)

	// Init SlashingKeeper
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		cdc,
		runtime.NewKVStoreService(keys[slashingtypes.StoreKey]),
		app.StakingKeeper,
		authority,
	)

	// Init CrisisKeeper
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[crisistypes.StoreKey]),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authority,
		app.AccountKeeper.AddressCodec(),
	)

	// Init FeeGrantKeeper
	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[feegrant.StoreKey]),
		app.AccountKeeper,
	)

	// Init UpgradeKeeper
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		app.BaseApp,
		authority,
	)

	// Set staking hooks
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)
}

// InitIBCKeepers initializes all IBC-related keepers
func (app *App) InitIBCKeepers(
	appCodec codec.Codec,
	keys map[string]*storetypes.KVStoreKey,
	memKeys map[string]*storetypes.MemoryStoreKey,
) (capabilitykeeper.ScopedKeeper, capabilitykeeper.ScopedKeeper) {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// Scope IBC keepers
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	// Seal capability keeper
	app.CapabilityKeeper.Seal()

	// Init IBCKeeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		app.GetSubspace(ibcexported.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
		authority,
	)

	// Init TransferKeeper
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
		authority,
	)

	return scopedIBCKeeper, scopedTransferKeeper
}

// InitGovKeeper initializes the governance keeper
func (app *App) InitGovKeeper(
	appCodec codec.Codec,
	keys map[string]*storetypes.KVStoreKey,
) {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// Register legacy gov router
	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	// Init GovKeeper
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[govtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.MsgServiceRouter(),
		govtypes.DefaultConfig(),
		authority,
	)
	govKeeper.SetLegacyRouter(govRouter)
	app.GovKeeper = govKeeper.SetHooks(govtypes.NewMultiGovHooks())
}

// InitEvidenceKeeper initializes the evidence keeper
func (app *App) InitEvidenceKeeper(
	appCodec codec.Codec,
	keys map[string]*storetypes.KVStoreKey,
) {
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[evidencetypes.StoreKey]),
		app.StakingKeeper,
		app.SlashingKeeper,
		app.AccountKeeper.AddressCodec(),
		runtime.ProvideCometInfoService(),
	)
	app.EvidenceKeeper = *evidenceKeeper
}

// InitStatesetKeepers initializes all Stateset custom module keepers
func (app *App) InitStatesetKeepers(
	appCodec codec.Codec,
	keys map[string]*storetypes.KVStoreKey,
) {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// Init OracleKeeper
	app.OracleKeeper = oraclekeeper.NewKeeper(appCodec, keys[oracletypes.StoreKey], authority)

	// Init ComplianceKeeper
	app.ComplianceKeeper = compliancekeeper.NewKeeper(appCodec, keys[compliancetypes.StoreKey], authority)

	// Init TreasuryKeeper
	app.TreasuryKeeper = treasurykeeper.NewKeeper(appCodec, keys[treasurytypes.StoreKey], authority)

	// Init PaymentsKeeper
	app.PaymentsKeeper = paymentskeeper.NewKeeper(
		appCodec,
		keys[paymentstypes.StoreKey],
		app.BankKeeper,
		app.ComplianceKeeper,
		paymentstypes.ModuleAccountName,
	)

	// Init StablecoinKeeper
	app.StablecoinKeeper = stablecoinkeeper.NewKeeper(
		appCodec,
		keys[stablecointypes.StoreKey],
		app.BankKeeper,
		app.AccountKeeper,
		app.OracleKeeper,
		app.ComplianceKeeper,
	)

	// Init SettlementKeeper
	app.SettlementKeeper = settlementkeeper.NewKeeper(
		appCodec,
		keys[settlementtypes.StoreKey],
		app.BankKeeper,
		app.ComplianceKeeper,
		app.AccountKeeper,
		authority,
	)

	// Init CircuitKeeper for security controls
	app.CircuitKeeper = circuitkeeper.NewKeeper(
		appCodec,
		keys[circuittypes.StoreKey],
		authority,
	)

	// Init MetricsKeeper for monitoring
	app.MetricsKeeper = metricskeeper.NewKeeper(
		appCodec,
		keys[metricstypes.StoreKey],
	)

	// Init ZkpVerifyKeeper for STARK proof verification
	app.ZkpVerifyKeeper = zkpverifykeeper.NewKeeper(
		appCodec,
		keys[zkpverifytypes.StoreKey],
		authority,
	)
}
