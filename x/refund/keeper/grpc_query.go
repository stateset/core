package keeper

import (
	"github.com/stateset/core/x/refund/types"
)

var _ types.QueryServer = Keeper{}
