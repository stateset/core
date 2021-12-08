package agreement_test

import (
	"testing"

	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/agreement"
	"github.com/stateset/core/x/agreement/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		SentAgreementList: []types.SentAgreement{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentAgreementCount: 2,
		TimedoutAgreementList: []types.TimedoutAgreement{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimedoutAgreementCount: 2,
		AgreementList: []types.Agreement{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AgreementCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AgreementKeeper(t)
	agreement.InitGenesis(ctx, *k, genesisState)
	got := agreement.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.SentAgreementList, len(genesisState.SentAgreementList))
	require.Subset(t, genesisState.SentAgreementList, got.SentAgreementList)
	require.Equal(t, genesisState.SentAgreementCount, got.SentAgreementCount)
	require.Len(t, got.TimedoutAgreementList, len(genesisState.TimedoutAgreementList))
	require.Subset(t, genesisState.TimedoutAgreementList, got.TimedoutAgreementList)
	require.Equal(t, genesisState.TimedoutAgreementCount, got.TimedoutAgreementCount)
	require.Len(t, got.AgreementList, len(genesisState.AgreementList))
	require.Subset(t, genesisState.AgreementList, got.AgreementList)
	require.Equal(t, genesisState.AgreementCount, got.AgreementCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
