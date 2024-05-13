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

func (g Game) GetPlayerALastMove() string {
	movesCount := len(g.PlayerAMoves)
	if movesCount == 0 {
		return ""
	}
	return g.PlayerAMoves[movesCount-1]
}

func (g Game) GetPlayerBLastMove() string {
	movesCount := len(g.PlayerBMoves)
	if movesCount == 0 {
		return ""
	}
	return g.PlayerBMoves[movesCount-1]
}

func (g Game) GetPlayerType(playerAddr string) (player rules.Player, err error) {
	switch playerAddr {
	case g.PlayerA:
		player = rules.PlayerA
	case g.PlayerB:
		player = rules.PlayerB
	default:
		player = rules.InvalidPlayer
		err = ErrInvalidPlayer
	}
	return
}

func (g Game) GetPlayerLastMove(playerAddr string) (player rules.Player, move string, err error) {
	switch playerAddr {
	case g.PlayerA:
		player = rules.PlayerA
		move = g.GetPlayerALastMove()
	case g.PlayerB:
		player = rules.PlayerB
		move = g.GetPlayerBLastMove()
	default:
		player = rules.InvalidPlayer
		err = ErrInvalidPlayer
	}
	return
}

// IsRoundRevealed returns true if each player has
// the same amount of moves and these are revealed
func (g Game) IsRoundRevealed() bool {
	playerAMovesCount, playerBMovesCount := len(g.PlayerAMoves), len(g.PlayerBMoves)
	if playerAMovesCount != playerBMovesCount {
		return false
	}
	// empty list
	if playerAMovesCount == 0 {
		return true
	}

	return rules.IsValidMove(g.GetPlayerALastMove()) && rules.IsValidMove(g.GetPlayerBLastMove())
}

func (g *Game) AddWinToPlayerA() {
	g.Score[0]++
}

func (g *Game) AddWinToPlayerB() {
	g.Score[1]++
}


func (g *Game) AddPlayerMove(playerAddr, move string) error {
	// Player is in the game
	player, err := g.GetPlayerType(playerAddr)
	if err != nil {
		return err
	}

	if ok := isValidHash(move); !ok {
		return ErrInvalidCommitment
	}
	// To submit a commitment, need to check if previous
	// move was revealed
	var (
		prevMove        string
		oponentPrevMove string
	)
	// make a copy of players moves to avoid updating the game
	// object if any validation fails
	playerAMoves, playerBMoves := g.GetPlayerAMoves(), g.GetPlayerBMoves()

	switch playerAddr {
	case g.PlayerA:
		prevMove = g.GetPlayerALastMove()
		oponentPrevMove = g.GetPlayerBLastMove()
		playerAMoves = append(playerAMoves, move)
	case g.PlayerB:
		prevMove = g.GetPlayerBLastMove()
		oponentPrevMove = g.GetPlayerALastMove()
		playerBMoves = append(playerBMoves, move)
	default:
		return ErrInvalidPlayer
	}

	// check that the player's previous move was revealed
	// if no previous move, is an empty string
	ok := rules.IsValidMove(prevMove)
	if prevMove != "" && !ok {
		return errors.Wrapf(ErrRevealPreviousMove, "player with address %s has to reveal the move to finish the round", playerAddr)
	}

	playerAMovesCount, playerBMovesCount := len(playerAMoves), len(playerBMoves)
	// if previous move was revealed, but oponent's previous
	// move was not revealed, then cannot submit a new commitment.
	// Unless is evening out the moves (playerAmovesCount == playerBmovesCount)
	ok = rules.IsValidMove(oponentPrevMove)
	if oponentPrevMove != "" && !ok && playerAMovesCount != playerBMovesCount {
		return errors.Wrapf(ErrRevealPreviousMove, "oponent player has to reveal the move to finish the round")
	}

	// Can make the move - depends on:
	//  - rules: game status, rounds count, other player moves
	if ok := rules.CanMakeMove(player, playerAMovesCount, playerBMovesCount); !ok {
		return ErrPlayerCantMakeMove
	}

	// all validations passed
	// update the game object with the new move
	g.PlayerAMoves = playerAMoves
	g.PlayerBMoves = playerBMoves

	return nil
}


func (game *Game) RevealPlayerMove(playerAddr, revealedMove, salt string) error {
	if ok := rules.IsValidMove(revealedMove); !ok {
		return ErrInvalidMove
	}

	// Can only reveal if both players submitted the commitment
	playerAMovesCount, playerBMovesCount := len(game.PlayerAMoves), len(game.PlayerBMoves)
	if playerAMovesCount != playerBMovesCount {
		return ErrPlayerCantRevealMove
	}

	var commitment string
	switch playerAddr {
	case game.PlayerA:
		commitment = game.GetPlayerALastMove()
		game.PlayerAMoves[playerAMovesCount-1] = revealedMove
	case game.PlayerB:
		commitment = game.GetPlayerBLastMove()
		game.PlayerBMoves[playerBMovesCount-1] = revealedMove
	default:
		return ErrInvalidPlayer
	}

	// check that this move wasn't revealed previously
	if ok := isValidHash(commitment); !ok {
		return errors.Wrapf(ErrMoveAlreadyRevealed, "move: %s", commitment)
	}

	// calculate the hash and compare it with the commitment
	if ok := isMoveRevelead(commitment, revealedMove, salt); !ok {
		return ErrWrongMoveRevealed
	}

	return nil
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

func (game *Game) Ended() bool {
	return game.Status == rules.StatusDraw ||
		game.Status == rules.StatusPlayerAWins ||
		game.Status == rules.StatusPlayerBWins ||
		game.Status == rules.StatusCancelled
}

func getPlayerAddress(address string) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(address)
	return addr, errors.Wrapf(err, ErrInvalidAddress.Error(), address)
}
