package types

import "cosmossdk.io/errors"

var (
	ErrInvalidAddress     = errors.Register(ModuleName, 1, "invalid address")
	ErrRoundsOutOfBounds  = errors.Register(ModuleName, 2, "game rounds out of bounds")
	ErrInvalidMovesNumber = errors.Register(ModuleName, 3, "invalid player moves count")
	ErrInvalidGameNumber  = errors.Register(ModuleName, 4, "invalid game number. Should be greater than 0")
	ErrInvalidGameStatus  = errors.Register(ModuleName, 5, "invalid game status")
	ErrInvalidScore       = errors.Register(ModuleName, 6, "invalid score")
	ErrInvalidOponent     = errors.Register(ModuleName, 7, "invalid oponent address")
	ErrDuplicatedIndex    = errors.Register(ModuleName, 8, "duplicated game index")
	ErrInvalidMove        = errors.Register(ModuleName, 9, "invalid move")
	ErrGameEnded          = errors.Register(ModuleName, 10, "game has ended")
	ErrInvalidPlayer      = errors.Register(ModuleName, 11, "invalid player")
	ErrPlayerCantMakeMove = errors.Register(ModuleName, 12, "player cannot make move")
	ErrInvalidTTL         = errors.Register(ModuleName, 13, "ttl should be greater than 0")
)
