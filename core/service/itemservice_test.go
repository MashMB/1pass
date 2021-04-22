// Unit tests for default item service.
//
// @author TSS

package service

import (
	"testing"

	"github.com/mashmb/1pass/adapter/out/repo/file"
	"github.com/mashmb/1pass/adapter/out/util/crypto"
	"github.com/mashmb/1pass/core/domain"
	"github.com/mashmb/1pass/port/out"
)

func setupItemAndKeyService() (*dfltItemService, KeyService) {
	vaultPath := "../../assets/onepassword_data"
	var crytpoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var keyService KeyService

	crytpoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo(vaultPath)
	profileRepo = file.NewFileProfileRepo(vaultPath)

	keyService = NewDfltKeyService(crytpoUtils, profileRepo)

	return NewDfltItemService(keyService, itemRepo), keyService
}

func TestGetDetails(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	pass := "freddy"
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	masterKey, masterMac, _ := keyService.MasterKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac)
	item := itemService.GetDetails(uid, keys)

	if item.Uid != uid {
		t.Errorf("[SUCESS] GetDetails() = %v; expected = %v", item.Uid, uid)
	}

	item = itemService.GetDetails("", keys)

	if item != nil {
		t.Errorf("[FAIL] GetDetails() = %v; expected = %v", item.Uid, nil)
	}
}

func TestGetOverview(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	pass := "freddy"
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, nil, nil, overviewKey, overviewMac)
	item := itemService.GetOverview(uid, keys)

	if item.Uid != uid {
		t.Errorf("[SUCESS] GetOverview() = %v; expected = %v", item.Uid, uid)
	}

	item = itemService.GetOverview("", keys)

	if item != nil {
		t.Errorf("[FAIL] GetOverview() = %v; expected = %v", item.Uid, nil)
	}
}

func TestGetSimple(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	expected := 10
	first, firstUid := "Bank of America", "EC0A40400ABB4B16926B7417E95C9669"
	last, lastUid := "YouTube", "358B7411EB8B45CD9CE592ED16F3E9DE"
	pass := "freddy"
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, nil, nil, overviewKey, overviewMac)
	items := itemService.GetSimple(keys)

	if len(items) != expected {
		t.Errorf("GetSimple() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first || items[0].Uid != firstUid {
		t.Errorf("[FIRST] GetSimple() = %v, %v; expected = %v, %v", items[0].Title, items[0].Uid, first, firstUid)
	}

	if items[len(items)-1].Title != last || items[len(items)-1].Uid != lastUid {
		t.Errorf("[LAST] GetSimple() = %v, %v; expected = %v, %v", items[len(items)-1].Title, items[len(items)-1].Uid, last, lastUid)
	}
}
