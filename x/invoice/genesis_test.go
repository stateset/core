package invoice_test

import (
	"testing"

	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/invoice"
	"github.com/stateset/core/x/invoice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		InvoiceList: []types.Invoice{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		InvoiceCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InvoiceKeeper(t)
	invoice.InitGenesis(ctx, *k, genesisState)
	got := invoice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.InvoiceList, len(genesisState.InvoiceList))
	require.Subset(t, genesisState.InvoiceList, got.InvoiceList)
	require.Equal(t, genesisState.InvoiceCount, got.InvoiceCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
