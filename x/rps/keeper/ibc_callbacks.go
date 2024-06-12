package keeper

import (
	"encoding/json"

	"github.com/0xlb/rpschain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

func (k *Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack ibcexported.Acknowledgement,
) ibcexported.Acknowledgement {
	var data transfertypes.FungibleTokenPacketData
	transfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &data)

	// check for memo field if it is not parsed successfully, no-op
	pm := types.PacketMemo{}
	if err := json.Unmarshal([]byte(data.Memo), &pm); err != nil {
		return ack
	}
	msg, err := pm.Parse()
	if err != nil {
		return ack
	}

	// get the local address for the packet sender
	// to use as the player that creates a game or makes a move
	playerAcc, err := types.GetLocalAccount(data.Sender)
	if err != nil {
		return ack
	}

	msgServer := NewMsgServerImpl(*k)

	switch pm.RPS.Type {
	case types.CreateGame:
		createMsg := msg.(*types.MsgCreateGame)
		createMsg.Creator = playerAcc.String()
		createMsg.Oponent = data.Receiver
		_, err = msgServer.CreateGame(ctx, createMsg)
	case types.MakeMove:
		moveMsg := msg.(*types.MsgMakeMove)
		moveMsg.Player = playerAcc.String()
		_, err = msgServer.MakeMove(ctx, moveMsg)
	case types.Reveal:
		revealMsg := msg.(*types.MsgRevealMove)
		revealMsg.Player = playerAcc.String()
		_, err = msgServer.RevealMove(ctx, revealMsg)
	default:
		return ack
	}

	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	return ack
}
