package keeper

import (
	"github.com/stateset/core/x/agreement/types"
)

var _ types.QueryServer = Keeper{}
