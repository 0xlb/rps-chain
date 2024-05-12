package rules

// Player represents the possible players in the game
type Player int

const (
	InvalidPlayer Player = iota
	PlayerA
	PlayerB
)

// Choice represents a player's move in the game
type Choice string

const (
	Rock     Choice = "Rock"
	Paper    Choice = "Paper"
	Scissors Choice = "Scissors"
)

var validMoves = map[Choice]struct{}{
	Rock:     {},
	Paper:    {},
	Scissors: {},
}

// RoundResult represents the result of a round in the game
type RoundResult int

const (
	Win RoundResult = iota + 1
	Loss
	Draw
)

const (
	StatusWaiting     = "Waiting"
	StatusInProgress  = "In Progress"
	StatusPlayerAWins = "Player A Wins"
	StatusPlayerBWins = "Player B Wins"
	StatusDraw        = "Draw"
	StatusCancelled   = "Cancelled"
)

// IsValidStatus checks if the GameStatus is valid
func IsValidStatus(s string) bool {
	switch s {
	case StatusWaiting, StatusInProgress, StatusPlayerAWins, StatusPlayerBWins, StatusDraw, StatusCancelled:
		return true
	}
	return false
}

// DetermineRoundWinner determines if playerA is the winner of a round based on the players' choices
func DetermineRoundWinner(choicePlayerA, choicePlayerB Choice) RoundResult {
	switch {
	case choicePlayerA == choicePlayerB:
		return Draw
	case choicePlayerA == Rock && choicePlayerB == Scissors:
		return Win
	case choicePlayerA == Paper && choicePlayerB == Rock:
		return Win
	case choicePlayerA == Scissors && choicePlayerB == Paper:
		return Win
	default:
		return Loss
	}
}

// GameResultByMajority determines the result of the game based on the majority of rounds won
func GameResultByMajority(playerAWins, playerBWins, rounds uint64) string {
	halfRounds := rounds / 2
	if playerAWins > halfRounds {
		return StatusPlayerAWins
	} else if playerBWins > halfRounds {
		return StatusPlayerBWins
	}
	if playerAWins+playerBWins == rounds {
		return StatusDraw
	}
	return StatusInProgress
}

// IsValidMove validates the player's move
func IsValidMove(playerMoveStr string) bool {
	// Validate the move
	_, ok := validMoves[Choice(playerMoveStr)]
	return ok
}

// CanMakeMove checks if a player can make a move based on the current choices
func CanMakeMove(playerToMakeMove Player, playerAMovesCount, playerBMovesCount int) bool {
	if playerToMakeMove == InvalidPlayer {
		return false
	}
	movesDiff := absDiff(playerAMovesCount, playerBMovesCount)
	// Valid Cases:
	// - One of the players was one move behind and now they're even (diff is 0)
	// - They were even in moves and now one player made the next move
	return movesDiff <= 1
}

// absDiff is a helper function to get the absolute value of the difference
// between two integers
func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}