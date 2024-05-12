package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/0xlb/rpschain/x/rps/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	k Keeper
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

func (qs queryServer) GetGame(ctx context.Context, req *types.QueryGetGameRequest) (*types.QueryGetGameResponse, error) {
	g, err := qs.k.Games.Get(ctx, req.Index)
	if err == nil {
		return &types.QueryGetGameResponse{Game: &g}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &types.QueryGetGameResponse{Game: nil}, nil
	}
	return nil, status.Error(codes.Internal, err.Error())
}

func (qs queryServer) GetParams(ctx context.Context, _ *types.QueryGetParamsRequest) (*types.QueryGetParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetParamsResponse{Param: &params}, nil
}
