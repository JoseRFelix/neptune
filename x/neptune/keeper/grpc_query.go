package keeper

import (
	"github.com/JoseRFelix/neptune/x/neptune/types"
)

var _ types.QueryServer = Keeper{}
