package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
)

func (s *KeeperTestSuite) TestGetGame() {
	var (
		_, _, playerA = testdata.KeyTestPubAddr()
		_, _, playerB = testdata.KeyTestPubAddr()
	)
	testCases := []struct {
		name      string
		req       *types.QueryGetGameRequest
		malleate  func()
		expectErr bool
		expErrMsg string
		expGame   *types.Game
	}{
		{
			name:      "non-existent game",
			req:       &types.QueryGetGameRequest{Index: 1},
			expectErr: false,
			expGame:   nil,
		},
		{
			name: "get game successfully",
			req:  &types.QueryGetGameRequest{Index: 1},
			malleate: func() {
				// create a game to query
				_, err := s.msgServer.CreateGame(s.ctx, &types.MsgCreateGame{Creator: playerA.String(), Oponent: playerB.String(), Rounds: uint64(3)})
				s.Require().NoError(err)
			},
			expectErr: false,
			expGame: &types.Game{
				GameNumber:       1,
				PlayerA:          playerA.String(),
				PlayerB:          playerB.String(),
				Status:           rules.StatusWaiting,
				Rounds:           uint64(3),
				Score:            []uint64{0, 0},
				ExpirationHeight: uint64(10),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupTest()
			// setup
			// call malleate function if defined
			if tc.malleate != nil {
				tc.malleate()
			}

			// make query
			res, err := s.queryClient.GetGame(s.ctx, tc.req)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.expGame, res.Game)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGetParams() {
	var (
		req           = &types.QueryGetParamsRequest{}
		defaultParams = types.DefaultParams()
	)

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
		expErrMsg string
		expParams *types.Params
	}{
		{
			name: "params not set",
			malleate: func() {
				s.keeper.Params.Remove(s.ctx)
			},
			expectErr: false,
			expParams: nil,
		},
		{
			name:      "get params successfully",
			expectErr: false,
			expParams: &defaultParams,
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupTest()
			// setup
			// call malleate function if defined
			if tc.malleate != nil {
				tc.malleate()
			}

			// make query
			res, err := s.queryClient.GetParams(s.ctx, req)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.expParams, res.Param)
			}
		})
	}
}