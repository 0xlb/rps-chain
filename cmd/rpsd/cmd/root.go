package cmd

import (
	"os"
	"time"

	dbm "github.com/cosmos/cosmos-db"
	cmtcfg "github.com/cometbft/cometbft/config"
	"github.com/spf13/cobra"

	clientv2keyring "cosmossdk.io/client/v2/autocli/keyring"
	"cosmossdk.io/core/address"
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	txmodule "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/0xlb/rpschain/app"
	"github.com/0xlb/rpschain/app/params"
)

// NewRootCmd creates a new root command for rpsd. It is called once in the
// main function.
func NewRootCmd() *cobra.Command {
	tempApp, err := app.NewRPSApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, simtestutil.NewAppOptionsWithFlagHome(tempDir()))
	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: tempApp.InterfaceRegistry(),
		Codec:             tempApp.AppCodec(),
		TxConfig:          tempApp.TxConfig(),
		Amino:             tempApp.LegacyAmino(),
	}
	if err != nil {
		panic(err)
	}

	clientCtx := ProvideClientContext(encodingConfig.Codec, encodingConfig.InterfaceRegistry, encodingConfig.TxConfig, encodingConfig.Amino)

	rootCmd := &cobra.Command{
		Use:   "rpsd",
		Short: "rpsd - the Rock, Paper & Scissors app chain",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			clientCtx = clientCtx.WithCmdContext(cmd.Context())
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			clientCtx, err = config.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}

			// This needs to go after ReadFromClientConfig, as that function
			// sets the RPC client needed for SIGN_MODE_TEXTUAL.
			enabledSignModes := append(tx.DefaultSignModes, signing.SignMode_SIGN_MODE_TEXTUAL) //nolint:gocritic // we know we aren't appending to the same slice
			txConfigOpts := tx.ConfigOptions{
				EnabledSignModes:           enabledSignModes,
				TextualCoinMetadataQueryFn: txmodule.NewGRPCCoinMetadataQueryFn(clientCtx),
			}
			txConfigWithTextual, err := tx.NewTxConfigWithOptions(
				codec.NewProtoCodec(encodingConfig.InterfaceRegistry),
				txConfigOpts,
			)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithTxConfig(txConfigWithTextual)
			if err := client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}
			clientCtx = clientCtx.WithTxConfig(txConfigWithTextual)

			// overwrite the minimum gas price from the app configuration
			srvCfg := serverconfig.DefaultConfig()
			srvCfg.MinGasPrices = "0rps"

			// overwrite the block timeout
			cmtCfg := cmtcfg.DefaultConfig()
			cmtCfg.Consensus.TimeoutCommit = 3 * time.Second
			cmtCfg.LogLevel = "*:error,p2p:info,state:info" // better default logging

			return server.InterceptConfigsPreRunHandler(cmd, serverconfig.DefaultConfigTemplate, srvCfg, cmtCfg)
		},
	}

	initRootCmd(rootCmd, clientCtx.TxConfig, tempApp.BasicModuleManager)

	opts := tempApp.AutoCliOpts()
	opts.ClientCtx = clientCtx
	if err := opts.EnhanceRootCommand(rootCmd); err != nil {
		panic(err)
	}

	return rootCmd
}

func ProvideClientContext(
	appCodec codec.Codec,
	interfaceRegistry codectypes.InterfaceRegistry,
	txConfig client.TxConfig,
	legacyAmino *codec.LegacyAmino,
) client.Context {
	clientCtx := client.Context{}.
		WithCodec(appCodec).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txConfig).
		WithLegacyAmino(legacyAmino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithViper("RPS") // env variable prefix

	// Read the config again to overwrite the default values with the values from the config file
	clientCtx, _ = config.ReadFromClientConfig(clientCtx)

	return clientCtx
}

func ProvideKeyring(clientCtx client.Context, addressCodec address.Codec) (clientv2keyring.Keyring, error) {
	kb, err := client.NewKeyringFromBackend(clientCtx, clientCtx.Keyring.Backend())
	if err != nil {
		return nil, err
	}

	return keyring.NewAutoCLIKeyring(kb)
}

func tempDir() string {
	dir, err := os.MkdirTemp("", "rps")
	if err != nil {
		dir = app.DefaultNodeHome
	}
	defer os.RemoveAll(dir)

	return dir
}