package types

import "cosmossdk.io/collections"

const ModuleName = "rps"

var (
	ParamsKey     = collections.NewPrefix(0)
	GamesKey      = collections.NewPrefix(1)
	GameNumberKey = collections.NewPrefix(2)
)
