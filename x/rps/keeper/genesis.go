package keeper

import (
	"context"

	"github.com/0xlb/rpschain/x/rps/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, g := range data.Games {
		if err := k.Games.Set(ctx, g.GameNumber, g); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var games []types.Game
	if err := k.Games.Walk(ctx, nil, func(index uint64, g types.Game) (stop bool, err error) {
		games = append(games, g)
		return
	}); err != nil {
		return nil, err
	}

	return &types.GenesisState{
		Params: params,
		Games:  games,
	}, nil
}
