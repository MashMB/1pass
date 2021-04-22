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

func TestGetOverview(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	pass := "freddy"
	sucess := 1
	fail := 0
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, nil, nil, overviewKey, overviewMac)
	items := itemService.GetOverview("YouTube", keys)

	if len(items) != sucess {
		t.Errorf("[SUCESS] GetOverview() = %d; expected = %d", len(items), sucess)
	}

	items = itemService.GetOverview("", keys)

	if len(items) != fail {
		t.Errorf("[FAIL] GetOverview() = %d; expected = %d", len(items), fail)
	}
}

func TestGetSimple(t *testing.T) {
	itemService, keyService := setupItemAndKeyService()
	expected := 10
	first := "Bank of America"
	last := "YouTube"
	pass := "freddy"
	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, nil, nil, overviewKey, overviewMac)
	items := itemService.GetSimple(keys)

	if len(items) != expected {
		t.Errorf("GetSimple() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first {
		t.Errorf("GetSimple() = %v; expected = %v", items[0].Title, first)
	}

	if items[len(items)-1].Title != last {
		t.Errorf("GetSimple() = %v; expected = %v", items[len(items)-1], last)
	}
}
