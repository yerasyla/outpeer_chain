package keeper

import (
	"outpeer_chain/x/outpeerchain/types"
)

var _ types.QueryServer = Keeper{}
