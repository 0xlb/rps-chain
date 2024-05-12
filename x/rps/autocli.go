package rps

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	rpsv1 "github.com/0xlb/rpschain/api/rps/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetGame",
					Use:       "game [index]",
					Short:     "Get the current game with the provided game index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
					},
				},
				{
					RpcMethod: "GetParams",
					Use:       "params",
					Short:     "Get the current module parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateGame",
					Use:       "create oponent rounds",
					Short:     "Creates a new Rock, Paper & Scissors game for the message sender and the chose oponent",
					Long:      "Creates a new Rock, Paper & Scissors game for the message sender and the chose oponent. Input parameters are the oponent address and the rounds number for the game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "oponent"},
						{ProtoField: "rounds"},
					},
				},
				{
					RpcMethod: "MakeMove",
					Use:       "make-move game_index move",
					Short:     "Submits a new move for a specific Rock, Paper & Scissors game",
					Long:      "Submits a new move for a specific Rock, Paper & Scissors game. Valid move options are 'Rock', 'Paper' or 'Scissors'",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "game_index"},
						{ProtoField: "move"},
					},
				},
			},
		},
	}
}
