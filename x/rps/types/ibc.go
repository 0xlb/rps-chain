package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// valid tx types
const (
	CreateGame = "create"
	MakeMove   = "move"
	Reveal     = "reveal"
)

// IBC packet memo can trigger a transaction on the x/rps
// module. We expect the memo to be a json object with a
// field 'rps'. The value should have the fields of a valid messaga,
// e.g. MsgCreateGame, MsgMakeMove or MsgReveal
type PacketMemo struct {
	RPS Tx `json:"rps"`
}

type Tx struct {
	Type string      `json:"type"`
	Msg  interface{} `json:"msg"`
}

// Parse the memo if contains a valid message
func (pm PacketMemo) Parse() (interface{}, error) {
	bz, err := json.Marshal(pm.RPS.Msg)
	if err != nil {
		return nil, err
	}
	var msg interface{}
	switch pm.RPS.Type {
	case CreateGame:
		msg = &MsgCreateGame{}
	case MakeMove:
		msg = &MsgMakeMove{}
	case Reveal:
		msg = &MsgRevealMove{}
	default:
		return nil, fmt.Errorf("invalid tx type. Got %s", pm.RPS.Type)
	}

	if err := json.Unmarshal(bz, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// GetLocalAccount is a helper function to convert the sender account
// of the IBC packet to an address on the RPS chain
func GetLocalAccount(address string) (sdk.AccAddress, error) {
	bech32Prefix := strings.SplitN(address, "1", 2)[0]
	addressBz, err := sdk.GetFromBech32(address, bech32Prefix)
	if err != nil {
		return nil, err
	}
	if err := sdk.VerifyAddressFormat(addressBz); err != nil {
		return nil, err
	}
	return sdk.AccAddress(addressBz), nil
}
