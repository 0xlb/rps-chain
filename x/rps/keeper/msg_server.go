package keeper

import (
	"context"

	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) CreateGame(ctx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	newGame := types.Game{
		GameNumber: ms.k.NextGameNumber(ctx),
		PlayerA:    msg.Creator,
		PlayerB:    msg.Oponent,
		Rounds:     msg.Rounds,
		Status:     rules.StatusWaiting,
		Score:      []uint64{0, 0},
	}

	if err := newGame.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Games.Set(ctx, newGame.GameNumber, newGame); err != nil {
		return nil, err
	}

	return &types.MsgCreateGameResponse{}, nil
}

func (ms msgServer) MakeMove(context.Context, *types.MsgMakeMove) (*types.MsgMakeMoveResponse, error) {
	return nil, nil
}

func (ms msgServer) UpdateParams(context.Context, *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	return nil, nil
}
