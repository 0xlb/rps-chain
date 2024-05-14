package keeper

import (
	"cosmossdk.io/collections"
	v2 "github.com/0xlb/rpschain/x/rps/migrations/v2"
	"github.com/0xlb/rpschain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place state migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns Migrator instance for the state migration.
func NewMigrator(k Keeper) Migrator {
	return Migrator{
		keeper: k,
	}
}

// Migrate1to2 migrates the module state from version 1 to version 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	// migrate games - add the expiration height for existing games
	if err := m.keeper.Games.Walk(ctx, nil, func(id uint64, game types.Game) (stop bool, err error) {
		mg := v2.MigrateGame(ctx, game)
		if err := m.keeper.Games.Set(ctx, id, mg); err != nil {
			return true, err
		}
		// add these to the active games list
		if err := m.keeper.ActiveGamesQueue.Set(ctx, collections.Join(mg.ExpirationHeight, id)); err != nil {
			return true, err
		}
		return false, nil
	}); err != nil {
		return err
	}

	return m.keeper.Params.Set(ctx, types.DefaultParams())
}
