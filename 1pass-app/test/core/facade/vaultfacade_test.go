// Unit tests for vault facade implementation.
//
// @author TSS

package facade

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	corefacade "github.com/mashmb/1pass/1pass-core/core/facade"
	"github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/util/crypto"
)

func setupVaultFacade() corefacade.VaultFacade {
	var configRepo out.ConfigRepo
	var cryptoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var configService service.ConfigService
	var itemService service.ItemService
	var keyService service.KeyService
	var vaultService service.VaultService

	configRepo = file.NewFileConfigRepo("../../../../assets/1pass.yml")
	cryptoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()

	configService = service.NewDfltConfigService(configRepo)
	keyService = service.NewDfltKeyService(cryptoUtils, profileRepo)
	itemService = service.NewDfltItemService(keyService, itemRepo)
	vaultService = service.NewDfltVaultService(itemRepo, profileRepo)

	return corefacade.NewDfltVaultFacade(configService, itemService, keyService, vaultService)
}

func TestGetItem(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	facade.Unlock("../../../../assets/onepassword_data", pass)
	expected := "Personal"
	uid := "0EDE2B13D7AC4E2C9105842682ACB187"
	trashed := false
	item := facade.GetItem(uid, trashed)

	if item.Title != expected {
		t.Errorf("GetItem() = %v; expected = %v", item.Title, expected)
	}

	expected = "A note to Trash"
	uid = "AE272805811C450586BA3EDEAEF8AE19"
	trashed = true
	item = facade.GetItem(uid, trashed)

	if item.Title != expected {
		t.Errorf("GetItem() = %v; expected = %v", item.Title, expected)
	}
}

func TestGetItems(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	expected := 10
	trashed := false
	first := "Bank of America"
	last := "YouTube"
	err := facade.Unlock("../../../../assets/onepassword_data", pass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}

	items := facade.GetItems(domain.ItemCategoryEnum.Login, trashed)

	if len(items) != expected {
		t.Errorf("[NOT-TRASHED-LENGTH] GetItems() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first {
		t.Errorf("[NOT-TRASHED-FIRST] GetItems() = %v; expected = %v", items[0].Title, first)
	}

	if items[len(items)-1].Title != last {
		t.Errorf("[NOT-TRASHED-LAST] GetItems() = %v; expected = %v", items[len(items)-1], last)
	}

	expected = 2
	trashed = true
	first = "A note to Trash"
	last = ""
	items = facade.GetItems(nil, trashed)

	if len(items) != expected {
		t.Errorf("[TRASHED-LENGTH] GetItems() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first {
		t.Errorf("[TRASHED-FIRST] GetItems() = %v; expected = %v", items[0].Title, first)
	}

	if items[len(items)-1].Title != last {
		t.Errorf("[TRASHED-LAST] GetItems() = %v; expected = %v", items[len(items)-1], last)
	}
}

func TestIsUnlocked(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	unlocked := facade.IsUnlocked()

	if unlocked == true {
		t.Errorf("IsUnlocked() = %v; expected %v", unlocked, false)
	}

	facade.Unlock("../../../../assets/onepassword_data", pass)
	unlocked = facade.IsUnlocked()

	if unlocked == false {
		t.Errorf("IsUnlocked() = %v; expected %v", unlocked, true)
	}
}

func TestLock(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	facade.Unlock("../../../../assets/onepassword_data", pass)

	if !facade.IsUnlocked() {
		t.Error("Unlock() should provide keys")
	}

	facade.Lock()

	if facade.IsUnlocked() {
		t.Error("Lock() should clear keys")
	}
}

func TestUnlock(t *testing.T) {
	facade := setupVaultFacade()
	goodPass := "freddy"
	badPass := ""
	err := facade.Unlock("../../../../assets/onepassword_data", badPass)

	if err == nil {
		t.Error("Unlock() should fail because of invalid password")
	}

	err = facade.Unlock("../../../../assets/onepassword_data", goodPass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}
}

func TestValidate(t *testing.T) {
	facade := setupVaultFacade()
	correctPath := "../../../../assets/onepassword_data"
	notVaultPath := "../../../../assets/empty"
	vault := domain.NewVault(correctPath)
	err := facade.Validate(vault)

	if err != nil {
		t.Error("Validate() should pass because of correct vault")
	}

	vault = domain.NewVault(notVaultPath)
	err = facade.Validate(vault)

	if err == nil {
		t.Error("Validate() should fail because of invalid vault")
	}
}
