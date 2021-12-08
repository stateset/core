package keeper

import (
	"github.com/stateset/core/x/invoice/types"
)

var _ types.QueryServer = Keeper{}
