package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/core/header"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/0xlb/rpschain/x/rps"
	"github.com/0xlb/rpschain/x/rps/keeper"
	"github.com/0xlb/rpschain/x/rps/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	keeper      keeper.Keeper
	queryClient types.QueryClient
	msgServer   types.MsgServer
	encCfg      moduletestutil.TestEncodingConfig
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.encCfg = moduletestutil.MakeTestEncodingConfig(rps.AppModule{})

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx.WithHeaderInfo(header.Info{})

	suite.keeper = keeper.NewKeeper(
		suite.encCfg.Codec,
		suite.encCfg.TxConfig.SigningContext().AddressCodec(),
		storeService,
		authtypes.NewModuleAddress("gov").String(),
	)

	suite.keeper.Params.Set(suite.ctx, types.DefaultParams())

	suite.msgServer = keeper.NewMsgServerImpl(suite.keeper)
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServerImpl(suite.keeper))
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
