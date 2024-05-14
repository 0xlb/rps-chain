package app

import (
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	v2 "github.com/0xlb/rpschain/app/upgrades/v2"
)

func (app *RPSApp) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(
		v2.UpgradeName,
		v2.CreateUpgradeHandler(app.ModuleManager, app.Configurator()),
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	var storeUpgrades *storetypes.StoreUpgrades
	switch upgradeInfo.Name {
	case v2.UpgradeName:
		// no store upgrades
	default:
		// no-op
	}
	if storeUpgrades != nil {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}
}
