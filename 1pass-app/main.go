// 1Pass application entry point.
//
// @author TSS

package main

import (
	"github.com/mashmb/1pass/1pass-app/cli"
	"github.com/mashmb/1pass/1pass-core/core/facade"
	"github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/in"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/util/crypto"
	"github.com/mashmb/1pass/1pass-term/adapter/in/cobra"
)

const (
	Version string = "0.0.0"
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

	cobraCli := cli.NewCobraCli(Version, cliControl)
	cobraCli.Run()
}
