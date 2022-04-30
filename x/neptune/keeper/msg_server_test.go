package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/JoseRFelix/neptune/testutil/keeper"
	"github.com/JoseRFelix/neptune/x/neptune/keeper"
	"github.com/JoseRFelix/neptune/x/neptune/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NeptuneKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
