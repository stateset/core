package keeper

import (
	"github.com/stateset/core/x/loan/types"
)

var _ types.QueryServer = Keeper{}
