package keeper

import (
	"github.com/stateset/core/x/proof/types"
)

var _ types.QueryServer = Keeper{}
