// Unit tests for vault facade implementation.
//
// @author TSS

package facade

import (
	"testing"

	corefacade "github.com/mashmb/1pass/1pass-core/core/facade"
	"github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/util/crypto"
)

func setupVaultFacade() corefacade.VaultFacade {
	var cryptoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var itemService service.ItemService
	var keyService service.KeyService
	var vaultService service.VaultService

	cryptoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()

	keyService = service.NewDfltKeyService(cryptoUtils, profileRepo)
	itemService = service.NewDfltItemService(keyService, itemRepo)
	vaultService = service.NewDfltVaultService(itemRepo, profileRepo)

	return corefacade.NewDfltVaultFacade(itemService, keyService, vaultService)
}

func TestGetItemDetails(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	err := facade.Unlock("../../../../assets/onepassword_data", pass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}

	item := facade.GetItemDetails(uid)

	if item.Uid != uid {
		t.Errorf("GetItemDetails() = %v; expected = %v", item.Uid, uid)
	}
}

func TestGetItemOverview(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	err := facade.Unlock("../../../../assets/onepassword_data", pass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}

	item := facade.GetItemOverview(uid)

	if item.Uid != uid {
		t.Errorf("GetItemOverview() = %v; expected = %v", item.Uid, uid)
	}
}

func TestGetItems(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	expected := 10
	first := "Bank of America"
	last := "YouTube"
	err := facade.Unlock("../../../../assets/onepassword_data", pass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}

	items := facade.GetItems()

	if len(items) != expected {
		t.Errorf("GetItems() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first {
		t.Errorf("GetItems() = %v; expected = %v", items[0].Title, first)
	}

	if items[len(items)-1].Title != last {
		t.Errorf("GetItems() = %v; expected = %v", items[len(items)-1], last)
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
