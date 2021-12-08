package types_test

import (
	"testing"

	"github.com/stateset/core/x/invoice/types"
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
				InvoiceList: []types.Invoice{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				InvoiceCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated invoice",
			genState: &types.GenesisState{
				InvoiceList: []types.Invoice{
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
			desc: "invalid invoice count",
			genState: &types.GenesisState{
				InvoiceList: []types.Invoice{
					{
						Id: 1,
					},
				},
				InvoiceCount: 0,
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
