package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/0xlb/rpschain/x/rps/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// state management
	Schema     collections.Schema
	Params     collections.Item[types.Params]
	GameNumber collections.Sequence
	Games      collections.Map[uint64, types.Game]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		GameNumber:   collections.NewSequence(sb, types.GameNumberKey, "game_number"),
		Games: collections.NewMap(
			sb,
			types.GamesKey,
			"games",
			collections.Uint64Key,
			codec.CollValue[types.Game](cdc),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// NextGameNumber returns and increments the global game number counter.
func (k Keeper) NextGameNumber(ctx context.Context) uint64 {
	n, err := k.GameNumber.Next(ctx)
	if err != nil {
		panic(err)
	}
	// sequences starts in 0, but we want the first game
	// to have game number 1
	return n + 1
}
