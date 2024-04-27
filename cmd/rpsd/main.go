package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/0xlb/rps-chain/app"
	"github.com/0xlb/rps-chain/app/params"
	"github.com/0xlb/rps-chain/cmd/rpsd/cmd"
)

func main() {
	params.SetAddressPrefixes()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
