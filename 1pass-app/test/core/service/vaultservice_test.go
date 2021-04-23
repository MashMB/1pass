// Unit tests for default vault service.
//
// @author TSS

package service

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	coreservice "github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
)

func setupVaultService() coreservice.VaultService {
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()

	return coreservice.NewDfltVaultService(itemRepo, profileRepo)
}

func TestValidateVault(t *testing.T) {
	service := setupVaultService()
	correctPath := "../../../../assets/onepassword_data"
	noBandsPath := "../../../../assets/nobands"
	noProfilePath := "../../../../assets/noprofile"
	notVaultPath := "../../../../assets/empty"
	vault := domain.NewVault(correctPath)
	err := service.ValidateVault(vault)

	if err != nil {
		t.Error("ValidateVault() should pass because of correct path")
	}

	vault = domain.NewVault(noBandsPath)
	err = service.ValidateVault(vault)

	if err != nil {
		t.Error("ValidateVault() should pass because of correct path (no bands)")
	}

	vault = domain.NewVault(noProfilePath)
	err = service.ValidateVault(vault)

	if err == nil {
		t.Error("ValidateVault() should fail because of invalid path (no profile)")
	}

	vault = domain.NewVault(notVaultPath)
	err = service.ValidateVault(vault)

	if err == nil {
		t.Error("ValidateVault() should fail because of invalid path")
	}
}
