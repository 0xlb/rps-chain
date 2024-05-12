package types

import (
	"bytes"

	"cosmossdk.io/errors"
	"github.com/0xlb/rpschain/x/rps/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxRound = 3

func DefaultGames() (games []Game) {
	return
}

func (g Game) GetPlayerAAddress() (sdk.AccAddress, error) {
	return getPlayerAddress(g.PlayerA)
}

func (g Game) GetPlayerBAddress() (sdk.AccAddress, error) {
	return getPlayerAddress(g.PlayerB)
}

func (g Game) GetPlayerAScore() uint64 {
	return g.Score[0]
}

func (g Game) GetPlayerBScore() uint64 {
	return g.Score[1]
}

func (g *Game) AddWinToPlayerA() {
	g.Score[0]++
}

func (g *Game) AddWinToPlayerB() {
	g.Score[1]++
}

func (g Game) ValidateRounds() error {
	if g.Rounds <= MaxRound && g.Rounds > 0 {
		return nil
	}
	return ErrRoundsOutOfBounds
}

func (g Game) ValidateMovesCount() error {
	if len(g.PlayerAMoves) <= int(g.Rounds) && len(g.PlayerBMoves) <= int(g.Rounds) {
		return nil
	}
	return ErrInvalidMovesNumber
}

func (g Game) ValidateGameNumber() error {
	if g.GameNumber > 0 {
		return nil
	}
	return ErrInvalidGameNumber
}

func (g Game) ValidateStatus() error {
	if rules.IsValidStatus(g.Status) {
		return nil
	}
	return ErrInvalidGameStatus
}

func (g Game) ValidateScore() error {
	scLen := len(g.Score)
	if scLen != 2 {
		return ErrInvalidScore
	}
	if g.Score[0]+g.Score[1] > g.Rounds {
		return ErrInvalidScore
	}
	return nil
}

func (g Game) Validate() error {
	accA, err := g.GetPlayerAAddress()
	if err != nil {
		return err
	}
	accB, err := g.GetPlayerBAddress()
	if err != nil {
		return err
	}
	if bytes.Equal(accA, accB) {
		return ErrInvalidOponent
	}
	if err := g.ValidateGameNumber(); err != nil {
		return err
	}
	if err := g.ValidateStatus(); err != nil {
		return err
	}
	if err := g.ValidateRounds(); err != nil {
		return err
	}
	if err := g.ValidateMovesCount(); err != nil {
		return err
	}
	return g.ValidateScore()
}

func getPlayerAddress(address string) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(address)
	return addr, errors.Wrapf(err, ErrInvalidAddress.Error(), address)
}
