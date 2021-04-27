// Unit tests for default item service.
//
// @author TSS

package service

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	coreservice "github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/util/crypto"
)

func setupItemAndKeyService() (coreservice.ItemService, coreservice.KeyService) {
	vault := domain.NewVault("../../../../assets/onepassword_data")
	var crytpoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var keyService coreservice.KeyService

	crytpoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	itemRepo.LoadItems(vault)
	profileRepo = file.NewFileProfileRepo()
	profileRepo.LoadProfile(vault)

	keyService = coreservice.NewDfltKeyService(crytpoUtils, profileRepo)

	return coreservice.NewDfltItemService(keyService, itemRepo), keyService
}

func TestGetDetails(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	pass := "freddy"
	trashed := false
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	masterKey, masterMac, _ := keyService.MasterKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac)
	item := itemService.GetDetails(uid, trashed, keys)

	if item.Uid != uid {
		t.Errorf("[NOT-TRASHED-SUCESS] GetDetails() = %v; expected = %v", item.Uid, uid)
	}

	item = itemService.GetDetails("", trashed, keys)

	if item != nil {
		t.Errorf("[FAIL] GetDetails() = %v; expected = %v", item.Uid, nil)
	}

	trashed = true
	uid = "AE272805811C450586BA3EDEAEF8AE19"
	item = itemService.GetDetails(uid, trashed, keys)

	if item.Uid != uid {
		t.Errorf("[TRASHED-SUCESS] GetDetails() = %v; expected = %v", item.Uid, uid)
	}
}

func TestGetOverview(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	pass := "freddy"
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	trashed := false
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, nil, nil, overviewKey, overviewMac)
	item := itemService.GetOverview(uid, trashed, keys)

	if item.Uid != uid {
		t.Errorf("[NOT-TRASHED-SUCESS] GetOverview() = %v; expected = %v", item.Uid, uid)
	}

	item = itemService.GetOverview("", trashed, keys)

	if item != nil {
		t.Errorf("GetOverview() = %v; expected = %v", item.Uid, nil)
	}

	uid = "AE272805811C450586BA3EDEAEF8AE19"
	trashed = true
	item = itemService.GetOverview(uid, trashed, keys)

	if item.Uid != uid {
		t.Errorf("[TRASHED-SUCESS] GetOverview() = %v; expected = %v", item.Uid, uid)
	}
}

func TestGetSimple(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	expected := 27
	trashed := false
	first, firstUid := "Bank of America", "EC0A40400ABB4B16926B7417E95C9669"
	last, lastUid := "Email Account", "FD2EADB43C4F4FC7BEB35A1692DDFDEA"
	pass := "freddy"
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, nil, nil, overviewKey, overviewMac)
	items := itemService.GetSimple(keys, nil, trashed)

	if len(items) != expected {
		t.Errorf("[NOT-TRASHED] GetSimple() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first || items[0].Uid != firstUid {
		t.Errorf("[FIRST-NOT-TRSHED] GetSimple() = %v, %v; expected = %v, %v", items[0].Title, items[0].Uid, first, firstUid)
	}

	if items[len(items)-1].Title != last || items[len(items)-1].Uid != lastUid {
		t.Errorf("[LAST-NOT-TRASHED] GetSimple() = %v, %v; expected = %v, %v", items[len(items)-1].Title, items[len(items)-1].Uid, last, lastUid)
	}

	trashed = true
	items = itemService.GetSimple(keys, nil, trashed)
	expected = 2
	first, firstUid = "A note to Trash", "AE272805811C450586BA3EDEAEF8AE19"
	last, lastUid = "", "0C4F27910A64488BB339AED63565D148"

	if len(items) != expected {
		t.Errorf("[TRASHED] GetSimple() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first || items[0].Uid != firstUid {
		t.Errorf("[FIRST-TRSHED] GetSimple() = %v, %v; expected = %v, %v", items[0].Title, items[0].Uid, first, firstUid)
	}

	if items[len(items)-1].Title != last || items[len(items)-1].Uid != lastUid {
		t.Errorf("[LAST-TRASHED] GetSimple() = %v, %v; expected = %v, %v", items[len(items)-1].Title, items[len(items)-1].Uid, last, lastUid)
	}
}
