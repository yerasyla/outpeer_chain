package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "outpeer_chain/testutil/keeper"
	"outpeer_chain/x/outpeerchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OutpeerchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
