package types_test

import (
	"testing"

	"github.com/stateset/core/x/purchaseorder/types"
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
				PurchaseorderList: []types.Purchaseorder{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				PurchaseorderCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated purchaseorder",
			genState: &types.GenesisState{
				PurchaseorderList: []types.Purchaseorder{
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
			desc: "invalid purchaseorder count",
			genState: &types.GenesisState{
				PurchaseorderList: []types.Purchaseorder{
					{
						Id: 1,
					},
				},
				PurchaseorderCount: 0,
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
