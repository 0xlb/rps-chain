package v2

import (
	"github.com/0xlb/rpschain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateGame(ctx sdk.Context, game types.Game) types.Game {
	currHeight := ctx.BlockHeight()
	game.ExpirationHeight = uint64(currHeight) + types.DefaultParams().Ttl
	return game
}
