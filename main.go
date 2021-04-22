// 1Pass application entry point.
//
// @author TSS

package main

import (
	"github.com/mashmb/1pass/adapter/in/cobra"
	"github.com/mashmb/1pass/adapter/out/repo/file"
	"github.com/mashmb/1pass/adapter/out/util/crypto"
	"github.com/mashmb/1pass/cli"
	"github.com/mashmb/1pass/core/facade"
	"github.com/mashmb/1pass/core/service"
	"github.com/mashmb/1pass/port/in"
	"github.com/mashmb/1pass/port/out"
)

func main() {
	var cryptoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var keyService service.KeyService
	var itemService service.ItemService
	var vaultService service.VaultService

	var vaultFacade facade.VaultFacade

	var cliControl in.CliControl

	cryptoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()

	keyService = service.NewDfltKeyService(cryptoUtils, profileRepo)
	itemService = service.NewDfltItemService(keyService, itemRepo)
	vaultService = service.NewDfltVaultService(itemRepo, profileRepo)

	vaultFacade = facade.NewDfltVaultFacade(itemService, keyService, vaultService)

	cliControl = cobra.NewCobraCliControl(vaultFacade)

	cobraCli := cli.NewCobraCli(cliControl)
	cobraCli.Run()
}
