package types_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
	"github.com/stretchr/testify/require"
)

func TestIsRoundRevealed(t *testing.T) {
	testCases := []struct {
		name         string
		playerAMoves []string
		playerBMoves []string
		expected     bool // Expected result
	}{
		{
			name:         "fail - uneven player moves count",
			playerAMoves: []string{string(rules.Rock), string(rules.Paper), string(rules.Scissors)},
			playerBMoves: []string{string(rules.Rock), string(rules.Paper)},
			expected:     false,
		},
		{
			name:         "success - valid moves",
			playerAMoves: []string{string(rules.Rock), string(rules.Paper), string(rules.Scissors)},
			playerBMoves: []string{string(rules.Rock), string(rules.Paper), string(rules.Paper)},
			expected:     true,
		},
		{
			name:         "fail - invalid moves",
			playerAMoves: []string{string(rules.Rock), string(rules.Paper), "invalid"},
			playerBMoves: []string{string(rules.Rock), string(rules.Paper), string(rules.Paper)},
			expected:     false,
		},
		{
			name:         "success - empty moves",
			playerAMoves: []string{},
			playerBMoves: []string{},
			expected:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game := types.Game{
				PlayerAMoves: tc.playerAMoves,
				PlayerBMoves: tc.playerBMoves,
			}
			require.Equal(t, tc.expected, game.IsRoundRevealed())

		})
	}
}

func TestAddPlayerMove(t *testing.T) {
	var (
		validHashedMove = types.CalculateHash(string(rules.Paper), "r4nd0m")
		_, _, playerA   = testdata.KeyTestPubAddr()
		_, _, playerB   = testdata.KeyTestPubAddr()
		_, _, otherAddr = testdata.KeyTestPubAddr()
	)
	testCases := []struct {
		name            string
		playerAddr      string   // Address of the player making the move
		move            string   // Move to be made
		playerAMoves    []string // Previous moves of the player A
		playerBMoves    []string // Previous moves of the player B
		expectedError   error    // Expected error
		expPlayerAMoves []string // Expected moves of the player A
		expPlayerBMoves []string // Expected moves of the player B
	}{
		{
			name:            "success - add valid move, no previous moves",
			playerAddr:      playerA.String(),
			move:            validHashedMove,
			playerAMoves:    nil,
			playerBMoves:    nil,
			expectedError:   nil,
			expPlayerAMoves: []string{validHashedMove},
			expPlayerBMoves: nil,
		},
		{
			name:            "success - add valid move with previous moves revealed",
			playerAddr:      playerA.String(),
			move:            validHashedMove,
			playerAMoves:    []string{string(rules.Paper)},
			playerBMoves:    []string{string(rules.Paper)},
			expectedError:   nil,
			expPlayerAMoves: []string{string(rules.Paper), validHashedMove},
			expPlayerBMoves: []string{string(rules.Paper)},
		},
		{
			name:          "fail - try to add valid move but oponent's previous move not revealed",
			playerAddr:    playerB.String(),
			move:          validHashedMove,
			playerAMoves:  []string{validHashedMove},
			playerBMoves:  []string{string(rules.Paper)},
			expectedError: types.ErrRevealPreviousMove,
		},
		{
			name:          "fail - try to add valid move but oponent didn't play yet",
			playerAddr:    playerB.String(),
			move:          validHashedMove,
			playerAMoves:  nil,
			playerBMoves:  []string{string(rules.Paper)},
			expectedError: types.ErrPlayerCantMakeMove,
		},
		{
			name:          "fail - cannot add revealed move. Should be hashed",
			playerAddr:    playerB.String(),
			move:          string(rules.Rock),
			playerAMoves:  nil,
			playerBMoves:  nil,
			expectedError: types.ErrInvalidCommitment,
		},
		{
			name:          "fail - try to add valid move but previous move not revealed",
			playerAddr:    playerA.String(),
			move:          validHashedMove,
			playerAMoves:  []string{validHashedMove},
			playerBMoves:  []string{validHashedMove},
			expectedError: types.ErrRevealPreviousMove,
		},
		{
			name:          "fail - player not in the game try to add valid move",
			playerAddr:    otherAddr.String(),
			move:          validHashedMove,
			playerAMoves:  nil,
			playerBMoves:  nil,
			expectedError: types.ErrInvalidPlayer,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// initialize a game with some moves
			g := types.Game{
				PlayerA:      playerA.String(),
				PlayerB:      playerB.String(),
				PlayerAMoves: tc.playerAMoves,
				PlayerBMoves: tc.playerBMoves,
			}

			err := g.AddPlayerMove(tc.playerAddr, tc.move)
			if tc.expectedError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError.Error())
				// expect moves didn't change
				require.Equal(t, tc.playerAMoves, g.PlayerAMoves)
				require.Equal(t, tc.playerBMoves, g.PlayerBMoves)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expPlayerAMoves, g.PlayerAMoves)
				require.Equal(t, tc.expPlayerBMoves, g.PlayerBMoves)
			}
		})
	}
}

func TestRevealPlayerMove(t *testing.T) {
	var (
		salt            = "r4nd0m"
		validHashedMove = types.CalculateHash(string(rules.Paper), salt)
		_, _, playerA   = testdata.KeyTestPubAddr()
		_, _, playerB   = testdata.KeyTestPubAddr()
		_, _, otherAddr = testdata.KeyTestPubAddr()
	)
	testCases := []struct {
		name            string
		playerAddr      string
		revealedMove    string
		salt            string
		playerAMoves    []string // Previous moves of the player A
		playerBMoves    []string // Previous moves of the player B
		expectedError   error
		expPlayerAMoves []string // Expected moves of the player A
		expPlayerBMoves []string // Expected moves of the player B
	}{
		{
			name:            "success - valid move reveal by Player A",
			playerAddr:      playerA.String(),
			revealedMove:    string(rules.Paper),
			salt:            salt,
			playerAMoves:    []string{validHashedMove},
			playerBMoves:    []string{validHashedMove},
			expectedError:   nil,
			expPlayerAMoves: []string{string(rules.Paper)},
			expPlayerBMoves: []string{validHashedMove},
		},
		{
			name:            "success - valid move reveal by Player B after oponent revealed",
			playerAddr:      playerB.String(),
			revealedMove:    string(rules.Paper),
			salt:            salt,
			playerAMoves:    []string{string(rules.Paper)},
			playerBMoves:    []string{validHashedMove},
			expectedError:   nil,
			expPlayerAMoves: []string{string(rules.Paper)},
			expPlayerBMoves: []string{string(rules.Paper)},
		},
		{
			name:          "fail - player not in the game",
			playerAddr:    otherAddr.String(),
			revealedMove:  string(rules.Paper),
			salt:          salt,
			playerAMoves:  []string{validHashedMove},
			playerBMoves:  []string{validHashedMove},
			expectedError: types.ErrInvalidPlayer,
		},
		{
			name:          "fail - invalid move reveal",
			playerAddr:    playerA.String(),
			revealedMove:  "InvalidMove",
			salt:          salt,
			playerAMoves:  []string{validHashedMove},
			playerBMoves:  []string{validHashedMove},
			expectedError: types.ErrInvalidMove,
		},
		{
			name:          "fail - previous move was revealed",
			playerAddr:    playerA.String(),
			revealedMove:  string(rules.Paper),
			salt:          salt,
			playerAMoves:  []string{string(rules.Paper)},
			playerBMoves:  []string{validHashedMove},
			expectedError: types.ErrMoveAlreadyRevealed,
		},
		{
			name:          "fail - oponent did not submit commitment yet",
			playerAddr:    playerB.String(),
			revealedMove:  string(rules.Paper),
			salt:          salt,
			playerAMoves:  nil,
			playerBMoves:  []string{validHashedMove},
			expectedError: types.ErrPlayerCantRevealMove,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// initialize a game with some moves
			g := types.Game{
				PlayerA:      playerA.String(),
				PlayerB:      playerB.String(),
				PlayerAMoves: tc.playerAMoves,
				PlayerBMoves: tc.playerBMoves,
			}

			err := g.RevealPlayerMove(tc.playerAddr, tc.revealedMove, tc.salt)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError.Error())
				// expect moves didn't change
				require.Equal(t, tc.playerAMoves, g.PlayerAMoves)
				require.Equal(t, tc.playerBMoves, g.PlayerBMoves)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expPlayerAMoves, g.PlayerAMoves)
				require.Equal(t, tc.expPlayerBMoves, g.PlayerBMoves)
			}
		})
	}
}
