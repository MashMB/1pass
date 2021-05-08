// 1Pass application entry point.
//
// @author TSS

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mashmb/1pass/1pass-app/cli"
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/facade"
	"github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/in"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/util/crypto"
	"github.com/mashmb/1pass/1pass-term/adapter/in/cobra"
	"github.com/mashmb/1pass/1pass-up/adapter/out/github"
)

func main() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln(err)
	}

	var configRepo out.ConfigRepo
	var cryptoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo
	var updater out.Updater

	var configService service.ConfigService
	var keyService service.KeyService
	var itemService service.ItemService
	var updateService service.UpdateService
	var vaultService service.VaultService

	var configFacade facade.ConfigFacade
	var updateFacade facade.UpdateFacade
	var vaultFacade facade.VaultFacade

	var cliControl in.CliControl

	configRepo = file.NewFileConfigRepo(filepath.Join(homeDir, ".config", "1pass"))
	cryptoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()
	updater = github.NewGithubUpdater()

	configService = service.NewDfltConfigService(configRepo)
	keyService = service.NewDfltKeyService(cryptoUtils, profileRepo)
	itemService = service.NewDfltItemService(keyService, itemRepo)
	updateService = service.NewDfltUpdateService(updater)
	vaultService = service.NewDfltVaultService(itemRepo, profileRepo)

	configFacade = facade.NewDfltConfigFacade(configService)
	updateFacade = facade.NewDfltUpdateFacade(configService, updateService)
	vaultFacade = facade.NewDfltVaultFacade(configService, itemService, keyService, vaultService)

	cliControl = cobra.NewCobraCliControl(configFacade, updateFacade, vaultFacade)

	cobraCli := cli.NewCobraCli(domain.Version, cliControl)
	cobraCli.Run()
}
