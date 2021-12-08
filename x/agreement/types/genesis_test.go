package types_test

import (
	"testing"

	"github.com/stateset/core/x/agreement/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated sentAgreement",
			genState: &types.GenesisState{
				SentAgreementList: []types.SentAgreement{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid sentAgreement count",
			genState: &types.GenesisState{
				SentAgreementList: []types.SentAgreement{
					{
						Id: 1,
					},
				},
				SentAgreementCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated timedoutAgreement",
			genState: &types.GenesisState{
				TimedoutAgreementList: []types.TimedoutAgreement{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid timedoutAgreement count",
			genState: &types.GenesisState{
				TimedoutAgreementList: []types.TimedoutAgreement{
					{
						Id: 1,
					},
				},
				TimedoutAgreementCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated agreement",
			genState: &types.GenesisState{
				AgreementList: []types.Agreement{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid agreement count",
			genState: &types.GenesisState{
				AgreementList: []types.Agreement{
					{
						Id: 1,
					},
				},
				AgreementCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
