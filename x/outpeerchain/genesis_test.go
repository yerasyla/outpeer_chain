package outpeerchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "outpeer_chain/testutil/keeper"
	"outpeer_chain/testutil/nullify"
	"outpeer_chain/x/outpeerchain"
	"outpeer_chain/x/outpeerchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OutpeerchainKeeper(t)
	outpeerchain.InitGenesis(ctx, *k, genesisState)
	got := outpeerchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
