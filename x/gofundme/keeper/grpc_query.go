package keeper

import (
	"github.com/cosmonaut/gofundme/x/gofundme/types"
)

var _ types.QueryServer = Keeper{}
