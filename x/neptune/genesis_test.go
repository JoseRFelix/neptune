package neptune_test

import (
	"testing"

	keepertest "github.com/JoseRFelix/neptune/testutil/keeper"
	"github.com/JoseRFelix/neptune/testutil/nullify"
	"github.com/JoseRFelix/neptune/x/neptune"
	"github.com/JoseRFelix/neptune/x/neptune/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NeptuneKeeper(t)
	neptune.InitGenesis(ctx, *k, genesisState)
	got := neptune.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
