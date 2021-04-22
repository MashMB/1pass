// Unit tests for vault facade implementation.
//
// @author TSS

package facade

import (
	"testing"

	"github.com/mashmb/1pass/adapter/out/repo/file"
	"github.com/mashmb/1pass/adapter/out/util/crypto"
	"github.com/mashmb/1pass/core/service"
	"github.com/mashmb/1pass/port/out"
)

func setupVaultFacade() *dfltVaultFacade {
	var cryptoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var itemService service.ItemService
	var keyService service.KeyService

	cryptoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo("../../assets/onepassword_data")
	profileRepo = file.NewFileProfileRepo("../../assets/onepassword_data")

	keyService = service.NewDfltKeyService(cryptoUtils, profileRepo)
	itemService = service.NewDfltItemService(keyService, itemRepo)

	return NewDfltVaultFacade(itemService, keyService)
}

func TestGetItemDetails(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	expected := 1
	title := "YouTube"
	err := facade.Unlock(pass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}

	items := facade.GetItemDetails(title)

	if len(items) != expected {
		t.Errorf("GetItemDetails() = %d; expected = %d", len(items), expected)
	}
}

func TestGetItemOverview(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	expected := 1
	title := "YouTube"
	err := facade.Unlock(pass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}

	items := facade.GetItemOverview(title)

	if len(items) != expected {
		t.Errorf("GetItemOverview() = %d; expected = %d", len(items), expected)
	}
}

func TestGetItems(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	expected := 10
	first := "Bank of America"
	last := "YouTube"
	err := facade.Unlock(pass)

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

	facade.Unlock(pass)
	unlocked = facade.IsUnlocked()

	if unlocked == false {
		t.Errorf("IsUnlocked() = %v; expected %v", unlocked, true)
	}
}

func TestLock(t *testing.T) {
	facade := setupVaultFacade()
	pass := "freddy"
	facade.Unlock(pass)

	if facade.keys == nil {
		t.Error("Unlock() should provide keys")
	}

	facade.Lock()

	if facade.keys != nil {
		t.Error("Lock() should clear keys")
	}
}

func TestUnlock(t *testing.T) {
	facade := setupVaultFacade()
	goodPass := "freddy"
	badPass := ""
	err := facade.Unlock(badPass)

	if err == nil {
		t.Error("Unlock() should fail because of invalid password")
	}

	err = facade.Unlock(goodPass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}
}
