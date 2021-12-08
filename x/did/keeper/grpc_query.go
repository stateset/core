package keeper

import (
	"github.com/stateset/core/x/did/types"
)

var _ types.QueryServer = Keeper{}
