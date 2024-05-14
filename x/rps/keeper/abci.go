package keeper

import (
	"context"
	"time"

	"cosmossdk.io/collections"
	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EndBlocker(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// fetch active games whose exportation period have ended (are equal or less than the current block)
	rng := collections.NewPrefixUntilPairRange[uint64, uint64](uint64(sdkCtx.BlockHeight()))
	if err := k.ActiveGamesQueue.Walk(ctx, rng, func(key collections.Pair[uint64, uint64]) (bool, error) {
		g, err := k.Games.Get(ctx, key.K2())
		if err != nil {
			return false, err
		}

		// check that status is InProgress or Waiting
		// then update the game status to cancelled
		if g.Status == rules.StatusWaiting || g.Status == rules.StatusInProgress {
			g.Status = rules.StatusCancelled
			if err := k.Games.Set(ctx, g.GameNumber, g); err != nil {
				return false, err
			}
		}

		// remove the game from the queue
		if err = k.ActiveGamesQueue.Remove(ctx, collections.Join(g.ExpirationHeight, g.GameNumber)); err != nil {
			return false, err
		}
		return false, nil
	}); err != nil {
		return err
	}
	return nil
}
