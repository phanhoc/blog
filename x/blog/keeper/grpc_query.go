package keeper

import (
	"github.com/phanhoc/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
