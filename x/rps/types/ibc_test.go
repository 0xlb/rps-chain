package types_test

import (
	"encoding/json"
	"testing"

	"github.com/0xlb/rpschain/x/rps/types"
	"github.com/stretchr/testify/require"
)

func TestParsePacketMemo(t *testing.T) {
	testCases := []struct {
		name       string
		memo       string
		expErr     bool
		parseCheck func(msg interface{})
	}{
		{
			name: "valid memo - Create game",
			memo: "{\"rps\": {\"type\": \"create\", \"msg\": {\"oponent\": \"rps189xq8ndsdlgkxtyfrgp263khwfzqqvn3aq6y2t\", \"rounds\": 3}}}",
			parseCheck: func(msg interface{}) {
				createMsg, ok := msg.(*types.MsgCreateGame)
				require.True(t, ok)
				require.Equal(t, createMsg.Oponent, "rps189xq8ndsdlgkxtyfrgp263khwfzqqvn3aq6y2t")
				require.Equal(t, createMsg.Rounds, uint64(3))
			},
		},
		{
			name: "valid memo (only rounds) - Create game",
			memo: "{\"rps\": {\"type\": \"create\", \"msg\": {\"rounds\": 3}}}",
			parseCheck: func(msg interface{}) {
				createMsg, ok := msg.(*types.MsgCreateGame)
				require.True(t, ok)
				require.Equal(t, createMsg.Oponent, "")
				require.Equal(t, createMsg.Creator, "")
				require.Equal(t, createMsg.Rounds, uint64(3))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pm := types.PacketMemo{}
			err := json.Unmarshal([]byte(tc.memo), &pm)
			require.NoError(t, err)
			msg, err := pm.Parse()

			if tc.expErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tc.parseCheck != nil {
				tc.parseCheck(msg)
			}
		})
	}
}

func TestGetLocalAddress(t *testing.T) {
	testCases := []struct {
		name            string
		address         string
		expErr          bool
		expLocalAddress string
	}{
		{
			name:            "valid address - Cosmos Hub",
			address:         "cosmos1paqywkh4vz3f8dd8fety8u2ymvnaef68xjx3ng",
			expLocalAddress: "rps1paqywkh4vz3f8dd8fety8u2ymvnaef68cc643g",
		},
		{
			name:            "valid address - Osmosis",
			address:         "osmo15e26mgge09qa7ajun99gq37hzvm9wgav6gztzg",
			expLocalAddress: "rps15e26mgge09qa7ajun99gq37hzvm9wgavvedlk6",
		},
		{
			name:    "invalid address - Osmosis",
			address: "osmo15e26mgge09qa7ajun99gq37hzvm9wgav6gztzz",
			expErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			addr, err := types.GetLocalAccount(tc.address)

			if tc.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expLocalAddress, addr.String())
		})
	}
}