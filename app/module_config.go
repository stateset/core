package app

import (
	"cosmossdk.io/x/evidence"
	evidencetypes "cosmossdk.io/x/evidence/types"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	transfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/spf13/cast"

	circuit "github.com/stateset/core/x/circuit"
	circuittypes "github.com/stateset/core/x/circuit/types"
	compliance "github.com/stateset/core/x/compliance"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	metrics "github.com/stateset/core/x/metrics"
	metricstypes "github.com/stateset/core/x/metrics/types"
	oracle "github.com/stateset/core/x/oracle"
	oracletypes "github.com/stateset/core/x/oracle/types"
	payments "github.com/stateset/core/x/payments"
	paymentstypes "github.com/stateset/core/x/payments/types"
	settlement "github.com/stateset/core/x/settlement"
	settlementtypes "github.com/stateset/core/x/settlement/types"
	stablecoin "github.com/stateset/core/x/stablecoin"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
	treasury "github.com/stateset/core/x/treasury"
	treasurytypes "github.com/stateset/core/x/treasury/types"
	zkpverify "github.com/stateset/core/x/zkpverify"
	zkpverifytypes "github.com/stateset/core/x/zkpverify/types"
)

// ModuleOrderConfig defines the order in which modules are processed
type ModuleOrderConfig struct {
	BeginBlockers []string
	EndBlockers   []string
	InitGenesis   []string
}

// GetDefaultModuleOrder returns the default module processing order
func GetDefaultModuleOrder() ModuleOrderConfig {
	return ModuleOrderConfig{
		BeginBlockers: []string{
			// Upgrade and capability modules first
			upgradetypes.ModuleName,
			capabilitytypes.ModuleName,
			consensusparamtypes.ModuleName,
			// Cosmos SDK core modules
			minttypes.ModuleName,
			distrtypes.ModuleName,
			slashingtypes.ModuleName,
			evidencetypes.ModuleName,
			stakingtypes.ModuleName,
			ibcexported.ModuleName,
			// Fee and circuit modules run early
			"feegrant",
			circuittypes.ModuleName,
			metricstypes.ModuleName,
			// Stateset business logic modules
			oracletypes.ModuleName,
			compliancetypes.ModuleName,
			treasurytypes.ModuleName,
			paymentstypes.ModuleName,
			stablecointypes.ModuleName,
			settlementtypes.ModuleName,
			zkpverifytypes.ModuleName,
		},
		EndBlockers: []string{
			crisistypes.ModuleName,
			govtypes.ModuleName,
			stakingtypes.ModuleName,
			consensusparamtypes.ModuleName,
			// Stateset modules
			oracletypes.ModuleName,
			compliancetypes.ModuleName,
			treasurytypes.ModuleName,
			paymentstypes.ModuleName,
			stablecointypes.ModuleName,
			settlementtypes.ModuleName,
			zkpverifytypes.ModuleName,
			circuittypes.ModuleName,
			metricstypes.ModuleName,
		},
		InitGenesis: []string{
			capabilitytypes.ModuleName,
			authtypes.ModuleName,
			banktypes.ModuleName,
			distrtypes.ModuleName,
			stakingtypes.ModuleName,
			slashingtypes.ModuleName,
			govtypes.ModuleName,
			minttypes.ModuleName,
			crisistypes.ModuleName,
			ibcexported.ModuleName,
			genutiltypes.ModuleName,
			evidencetypes.ModuleName,
			ibctransfertypes.ModuleName,
			// Stateset modules - circuit and metrics early
			circuittypes.ModuleName,
			metricstypes.ModuleName,
			oracletypes.ModuleName,
			compliancetypes.ModuleName,
			treasurytypes.ModuleName,
			paymentstypes.ModuleName,
			stablecointypes.ModuleName,
			settlementtypes.ModuleName,
			zkpverifytypes.ModuleName,
			consensusparamtypes.ModuleName,
		},
	}
}

// CreateModuleManager creates and configures the module manager
func (app *App) CreateModuleManager(
	appCodec codec.Codec,
	encodingConfig EncodingConfig,
	skipGenesisInvariants bool,
) *module.Manager {
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	mm := module.NewManager(
		// Cosmos SDK core modules
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName), app.interfaceRegistry),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper, app.AccountKeeper.AddressCodec()),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		// IBC modules
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		// Stateset custom modules
		oracle.NewAppModule(app.OracleKeeper),
		compliance.NewAppModule(app.ComplianceKeeper),
		treasury.NewAppModule(app.TreasuryKeeper),
		payments.NewAppModule(app.PaymentsKeeper),
		stablecoin.NewAppModule(app.StablecoinKeeper),
		settlement.NewAppModule(app.SettlementKeeper),
		circuit.NewAppModule(app.CircuitKeeper),
		metrics.NewAppModule(app.MetricsKeeper),
		zkpverify.NewAppModule(app.ZkpVerifyKeeper),
	)

	// Configure module order
	moduleOrder := GetDefaultModuleOrder()
	mm.SetOrderBeginBlockers(moduleOrder.BeginBlockers...)
	mm.SetOrderEndBlockers(moduleOrder.EndBlockers...)
	mm.SetOrderInitGenesis(moduleOrder.InitGenesis...)

	// Register invariants and services
	mm.RegisterInvariants(app.CrisisKeeper)
	mm.RegisterServices(module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()))

	return mm
}

// GetSkipGenesisInvariants returns the skipGenesisInvariants flag from app options
func GetSkipGenesisInvariants(appOpts interface {
	Get(string) interface{}
}) bool {
	return cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
}
