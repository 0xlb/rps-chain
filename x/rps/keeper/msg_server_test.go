package keeper_test

import (
	"cosmossdk.io/collections"
	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func (s *KeeperTestSuite) TestMsgCreateGame() {
	var (
		_, _, playerA = testdata.KeyTestPubAddr()
		_, _, playerB = testdata.KeyTestPubAddr()
	)

	testCases := []struct {
		name        string
		req         *types.MsgCreateGame
		expectedErr bool
		expErrMsg   string
		expGame     types.Game
	}{
		{
			name: "invalid creator",
			req: &types.MsgCreateGame{
				Creator: "invalid_address",
				Oponent: playerB.String(),
				Rounds:  uint64(3),
			},
			expectedErr: true,
			expErrMsg:   types.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid oponent",
			req: &types.MsgCreateGame{
				Creator: playerA.String(),
				Oponent: "invalid_address",
				Rounds:  uint64(3),
			},
			expectedErr: true,
			expErrMsg:   types.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid oponent - is same as creator",
			req: &types.MsgCreateGame{
				Creator: playerA.String(),
				Oponent: playerA.String(),
				Rounds:  uint64(3),
			},
			expectedErr: true,
			expErrMsg:   types.ErrInvalidOponent.Error(),
		},
		{
			name: "invalid oponent - bech32 address but from other chain",
			req: &types.MsgCreateGame{
				Creator: playerA.String(),
				Oponent: "cosmos1tf20gjep7xwylml8hr8dl9t5wkcd7tfglw7h6k",
				Rounds:  uint64(3),
			},
			expectedErr: true,
			expErrMsg:   types.ErrInvalidAddress.Error(),
		},
		{
			name: "rounds > max rounds allowed",
			req: &types.MsgCreateGame{
				Creator: playerA.String(),
				Oponent: playerB.String(),
				Rounds:  uint64(4),
			},
			expectedErr: true,
			expErrMsg:   types.ErrRoundsOutOfBounds.Error(),
		},
		{
			name: "rounds = 0",
			req: &types.MsgCreateGame{
				Creator: playerA.String(),
				Oponent: playerB.String(),
				Rounds:  uint64(0),
			},
			expectedErr: true,
			expErrMsg:   types.ErrRoundsOutOfBounds.Error(),
		},
		{
			name: "create game successfully",
			req: &types.MsgCreateGame{
				Creator: playerA.String(),
				Oponent: playerB.String(),
				Rounds:  uint64(3),
			},
			expectedErr: false,
			expGame: types.Game{
				GameNumber:       1,
				PlayerA:          playerA.String(),
				PlayerB:          playerB.String(),
				Status:           rules.StatusWaiting,
				Rounds:           uint64(3),
				Score:            []uint64{0, 0},
				ExpirationHeight: uint64(types.DefaultTTL),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupTest()
			_, err := s.msgServer.CreateGame(s.ctx, tc.req)
			if tc.expectedErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				// game sequence should have increased
				gameID, err := s.keeper.GameNumber.Peek(s.ctx)
				s.Require().NoError(err)
				s.Require().Equal(gameID, uint64(1))
				// check that game was stored
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				s.Require().Equal(tc.expGame, game)
				// check that new game was added to the active games queue
				ok, err := s.keeper.ActiveGamesQueue.Has(s.ctx, collections.Join(tc.expGame.ExpirationHeight, gameID))
				s.Require().NoError(err)
				s.Require().True(ok)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgMakeMove() {
	const gameID uint64 = 1
	var (
		validHashedMove = types.CalculateHash(string(rules.Paper), "r4nd0m")
		_, _, playerA   = testdata.KeyTestPubAddr()
		_, _, playerB   = testdata.KeyTestPubAddr()
		_, _, otherAddr = testdata.KeyTestPubAddr()
	)

	testCases := []struct {
		name      string
		req       *types.MsgMakeMove
		malleate  func()
		expectErr bool
		expErrMsg string
		expGame   types.Game
	}{
		{
			name: "invalid game index",
			req: &types.MsgMakeMove{
				Player:    playerA.String(),
				GameIndex: uint64(2),
				Move:      validHashedMove,
			},
			expectErr: true,
			expErrMsg: "not found",
		},
		{
			name: "game ended",
			req: &types.MsgMakeMove{
				Player:    playerA.String(),
				GameIndex: gameID,
				Move:      validHashedMove,
			},
			malleate: func() {
				// update game to ended status
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				game.Status = rules.StatusCancelled
				err = s.keeper.Games.Set(s.ctx, gameID, game)
				s.Require().NoError(err)
			},
			expectErr: true,
			expErrMsg: types.ErrGameEnded.Error(),
		},
		{
			name: "invalid player - player not in the game",
			req: &types.MsgMakeMove{
				Player:    otherAddr.String(),
				GameIndex: gameID,
				Move:      validHashedMove,
			},
			expectErr: true,
			expErrMsg: types.ErrInvalidPlayer.Error(),
		},
		{
			name: "invalid move",
			req: &types.MsgMakeMove{
				Player:    playerA.String(),
				GameIndex: gameID,
				Move:      "invalid_move",
			},
			expectErr: true,
			expErrMsg: types.ErrInvalidCommitment.Error(),
		},
		{
			name: "make move successfully",
			req: &types.MsgMakeMove{
				Player:    playerA.String(),
				GameIndex: gameID,
				Move:      validHashedMove,
			},
			expectErr: false,
			expGame: types.Game{
				GameNumber:       1,
				PlayerA:          playerA.String(),
				PlayerB:          playerB.String(),
				PlayerAMoves:     []string{validHashedMove},
				Status:           rules.StatusInProgress,
				Rounds:           uint64(3),
				Score:            []uint64{0, 0},
				ExpirationHeight: uint64(types.DefaultTTL),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupTest()
			// setup - create a game
			_, err := s.msgServer.CreateGame(s.ctx, &types.MsgCreateGame{Creator: playerA.String(), Oponent: playerB.String(), Rounds: uint64(3)})
			s.Require().NoError(err)

			// call malleate function if defined
			// to make additional setups
			if tc.malleate != nil {
				tc.malleate()
			}

			// check game was created OK
			game, err := s.keeper.Games.Get(s.ctx, gameID)
			s.Require().NoError(err)

			// make move
			_, err = s.msgServer.MakeMove(s.ctx, tc.req)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
				// stored game should remain unchanged
				g, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				s.Require().Equal(game, g)
			} else {
				s.Require().NoError(err)
				// check game was updated successfully
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				s.Require().Equal(tc.expGame, game)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgRevealMove() {
	const gameID uint64 = 1
	var (
		salt            = "r4nd0m"
		revealedMove    = string(rules.Paper)
		validHashedMove = types.CalculateHash(revealedMove, salt)
		_, _, playerA   = testdata.KeyTestPubAddr()
		_, _, playerB   = testdata.KeyTestPubAddr()
		_, _, otherAddr = testdata.KeyTestPubAddr()
	)

	testCases := []struct {
		name      string
		req       *types.MsgRevealMove
		malleate  func()
		postCheck func()
		expectErr bool
		expErrMsg string
		expGame   types.Game
	}{
		{
			name: "invalid game index",
			req: &types.MsgRevealMove{
				Player:       playerA.String(),
				GameIndex:    uint64(2),
				RevealedMove: revealedMove,
				Salt:         salt,
			},
			expectErr: true,
			expErrMsg: "not found",
		},
		{
			name: "game ended",
			req: &types.MsgRevealMove{
				Player:       playerA.String(),
				GameIndex:    gameID,
				RevealedMove: revealedMove,
				Salt:         salt,
			},
			malleate: func() {
				// update game to ended status
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				game.Status = rules.StatusCancelled
				err = s.keeper.Games.Set(s.ctx, gameID, game)
				s.Require().NoError(err)
			},
			expectErr: true,
			expErrMsg: types.ErrGameEnded.Error(),
		},
		{
			name: "invalid player - player not in the game",
			req: &types.MsgRevealMove{
				Player:       otherAddr.String(),
				GameIndex:    gameID,
				RevealedMove: revealedMove,
				Salt:         salt,
			},
			expectErr: true,
			expErrMsg: types.ErrInvalidPlayer.Error(),
		},
		{
			name: "invalid move",
			req: &types.MsgRevealMove{
				Player:       playerA.String(),
				GameIndex:    gameID,
				RevealedMove: "invalid_move",
				Salt:         salt,
			},
			expectErr: true,
			expErrMsg: types.ErrInvalidMove.Error(),
		},
		{
			name: "reveal move successfully - round not revealed - score not updated",
			req: &types.MsgRevealMove{
				Player:       playerA.String(),
				GameIndex:    gameID,
				RevealedMove: revealedMove,
				Salt:         salt,
			},
			expectErr: false,
			expGame: types.Game{
				GameNumber:       1,
				PlayerA:          playerA.String(),
				PlayerB:          playerB.String(),
				PlayerAMoves:     []string{revealedMove},
				PlayerBMoves:     []string{validHashedMove},
				Status:           rules.StatusInProgress,
				Rounds:           uint64(3),
				Score:            []uint64{0, 0},
				ExpirationHeight: uint64(types.DefaultTTL),
			},
			postCheck: func() {
				// check that game still in active games queue
				ok, err := s.keeper.ActiveGamesQueue.Has(s.ctx, collections.Join(uint64(types.DefaultTTL), gameID))
				s.Require().NoError(err)
				s.Require().True(ok)
			},
		},
		{
			name: "reveal move successfully - round revealed - update score",
			req: &types.MsgRevealMove{
				Player:       playerA.String(),
				GameIndex:    gameID,
				RevealedMove: revealedMove,
				Salt:         salt,
			},
			malleate: func() {
				// update player B move to be revelead
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				game.PlayerBMoves = []string{string(rules.Rock)}
				err = s.keeper.Games.Set(s.ctx, gameID, game)
				s.Require().NoError(err)
			},
			expectErr: false,
			expGame: types.Game{
				GameNumber:       1,
				PlayerA:          playerA.String(),
				PlayerB:          playerB.String(),
				PlayerAMoves:     []string{revealedMove},
				PlayerBMoves:     []string{string(rules.Rock)},
				Status:           rules.StatusInProgress,
				Rounds:           uint64(3),
				Score:            []uint64{1, 0},
				ExpirationHeight: uint64(types.DefaultTTL),
			},
			postCheck: func() {
				// check that game still in active games queue
				ok, err := s.keeper.ActiveGamesQueue.Has(s.ctx, collections.Join(uint64(types.DefaultTTL), gameID))
				s.Require().NoError(err)
				s.Require().True(ok)
			},
		},
		{
			name: "reveal move successfully - round revealed - update score - game ends",
			req: &types.MsgRevealMove{
				Player:       playerA.String(),
				GameIndex:    gameID,
				RevealedMove: revealedMove,
				Salt:         salt,
			},
			malleate: func() {
				// update player B, player B moves and score to
				// end the game with this move
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				game.PlayerAMoves = []string{string(rules.Paper), validHashedMove}
				game.PlayerBMoves = []string{string(rules.Rock), string(rules.Rock)}
				game.Score = []uint64{1, 0}
				err = s.keeper.Games.Set(s.ctx, gameID, game)
				s.Require().NoError(err)
			},
			expectErr: false,
			expGame: types.Game{
				GameNumber:       1,
				PlayerA:          playerA.String(),
				PlayerB:          playerB.String(),
				PlayerAMoves:     []string{string(rules.Paper), string(rules.Paper)},
				PlayerBMoves:     []string{string(rules.Rock), string(rules.Rock)},
				Status:           rules.StatusPlayerAWins,
				Rounds:           uint64(3),
				Score:            []uint64{2, 0},
				ExpirationHeight: uint64(types.DefaultTTL),
			},
			postCheck: func() {
				// check that game is removed from active games queue
				ok, err := s.keeper.ActiveGamesQueue.Has(s.ctx, collections.Join(uint64(types.DefaultTTL), gameID))
				s.Require().NoError(err)
				s.Require().False(ok)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupTest()
			// setup - create a game with some commitments
			_, err := s.msgServer.CreateGame(s.ctx, &types.MsgCreateGame{Creator: playerA.String(), Oponent: playerB.String(), Rounds: uint64(3)})
			s.Require().NoError(err)

			game, err := s.keeper.Games.Get(s.ctx, gameID)
			s.Require().NoError(err)
			game.PlayerAMoves = []string{validHashedMove}
			game.PlayerBMoves = []string{validHashedMove}
			err = s.keeper.Games.Set(s.ctx, gameID, game)
			s.Require().NoError(err)

			// call malleate function if defined
			// to make additional setups
			if tc.malleate != nil {
				tc.malleate()
			}

			// get game after setup to use in check
			// after the reveal move call
			game, err = s.keeper.Games.Get(s.ctx, gameID)
			s.Require().NoError(err)

			// make move
			_, err = s.msgServer.RevealMove(s.ctx, tc.req)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
				// stored game should remain unchanged
				g, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				s.Require().Equal(game, g)
			} else {
				s.Require().NoError(err)
				// check game was updated successfully
				game, err := s.keeper.Games.Get(s.ctx, gameID)
				s.Require().NoError(err)
				s.Require().Equal(tc.expGame, game)
			}

			if tc.postCheck != nil {
				tc.postCheck()
			}
		})
	}
}

func (s *KeeperTestSuite) TestUpdateParams() {
	testCases := []struct {
		name      string
		req       *types.MsgUpdateParams
		expectErr bool
		expErrMsg string
		expParams types.Params
	}{
		{
			name: "set invalid authority",
			req: &types.MsgUpdateParams{
				Authority: "foo",
			},
			expectErr: true,
			expErrMsg: "invalid authority",
		},
		{
			name: "set invalid ttl",
			req: &types.MsgUpdateParams{
				Authority: s.keeper.GetAuthority(),
				Params: types.Params{
					Ttl: 0,
				},
			},
			expectErr: true,
			expErrMsg: types.ErrInvalidTTL.Error(),
		},
		{
			name: "update params successfully",
			req: &types.MsgUpdateParams{
				Authority: s.keeper.GetAuthority(),
				Params: types.Params{
					Ttl: 30,
				},
			},
			expectErr: false,
			expParams: types.Params{
				Ttl: 30,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupTest()
			_, err := s.msgServer.UpdateParams(s.ctx, tc.req)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				updatedParams, err := s.keeper.Params.Get(s.ctx)
				s.Require().NoError(err)
				s.Require().Equal(tc.expParams, updatedParams)
			}
		})
	}
}
