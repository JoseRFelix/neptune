package keeper_test

import (
	"testing"

	testkeeper "github.com/JoseRFelix/neptune/testutil/keeper"
	"github.com/JoseRFelix/neptune/x/neptune/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NeptuneKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
