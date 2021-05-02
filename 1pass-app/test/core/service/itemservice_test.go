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

func setupItemServiceAndKeys() (coreservice.ItemService, *domain.Keys) {
	pass := "freddy"
	vault := domain.NewVault("../../../../assets/onepassword_data")

	var crytpoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var keyService coreservice.KeyService

	crytpoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()
	profileRepo.LoadProfile(vault)

	keyService = coreservice.NewDfltKeyService(crytpoUtils, profileRepo)
	itemService := coreservice.NewDfltItemService(keyService, itemRepo)

	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	masterKey, masterMac, _ := keyService.MasterKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac)
	itemService.DecodeItems(vault, keys)

	return itemService, keys
}

func TestDecodeDetails(t *testing.T) {
	itemService, keys := setupItemServiceAndKeys()
	encodedDetails := "b3BkYXRhMDGFAwAAAAAAAJIAigy3ZztWl46Kx16K7KgQOG6mQpq9dv0LWtLF7vbFaK4wZU9bq9kv4FFt088kLAjAH2ToJMyF0QiUiQDxix56mahLDjee22iUbvVaza/QSK8SfHFEpyR1Ecg2MRsXvn2DrwUPNsIrJJ2X6kdZLN5duXZGuhqLDITxx7eBOF+J5UWyjIBGDJNs4q9kd/B+W30YBtolRhzaHNonaNAEwEKOYxBjOnEE1oO3TwVRYqq3IT/fqHpj2yVfTMKa8WtLy2g4rGGe+8NzkiXPMND2cYRo+8jwsCBY1Zxvqw0149k/Ly6cUam0nAlq7NuDfpoT9J5rCC1UdFNKjE88Cfoarxcl+Kr3ZbYFQA3POhVsgFusQX1YyKZdxZWlfyPWb/SvkiD1vQmM5DhOs5XLONnXDTKr1xbWL+zlJxruYSWRxM5qD9oexnv1U06FFOUVDGpFg9fzbWmbkS9KBSGsUIyeqmmNFMa+5WgJ2Q4olZXUE81WTsPi2FerncvHnGd95n5m5BW85icZwIH+0pPUUlYFljhruBVXa5+D/GMX1DAzpBHRaEmgDYMJrhsLgaArXJiDw9drHKP0gzqVM4Ma0TX9G1Cr/mMEW2DaVtGNVywiTNSMMoqazb7hxTgITiTttShLv9nBUiw94vDKhcigD05lVXAsbXPqnZUwwGz4yCkIkhB3dH6u/9UBDqtB5hXFz/3taMPu3dr2G61Aqe3EU40ihz5D7Bp147SoBcuCyPeiOGzdfGa5zuGwC/IWg2Ii8nJgfhBAD2Va8hsutTI7Yc4yU4Ufla4cX6d0NH/bU6ajZIb2oFw3ie5fzyk7pInlWUhR41z6CISwPRmC9SseTKQzZ3FsMqKZ5KcBNlYAu0v9IFawC+kjKrMwKl3W6NKUkAR8AgM4HCyskjx6cbs52Jac5J3UIaUwj3zjoV0dH+6fOEuu44Xr3sk9VBJ6zUHiSV8OgNV1XEFlrqal+XfP60Rr5RaeztFWT7y0q+CFoS3ZKzNItudi0y1zY0ZnNsUbll2RCYlHULXE3Idxy/gsQJ29Aj39DTfbxvqGB9u6PUxewfvHnLkNXp7cjl3wE4IsaVHCNsL/ZBNgqwcaVVos9Wx1BtIWWhkOnmt2nrRIFz9vVdry8nEW6G+/IIFe+1e34oQmBSkGD+4OCGxGq1/Yctu3yG4QlumKa7SDIATG7q0iX0IMfFE9ws2RMBa7YawjMkRItFw2gn4Egzz8IPyvh328JssIW9oo2K/O0if5+ITWKvHvDDbi3M3FjJKMxDGPvAr2GA=="
	encodedKeys := "3Yy1faUq8AlFc/zDAcePMDmbqw/Y6vAs4bjjW5Y7enTU/ww9XDh7HEpVFiffEI1ETzuBOF2mnj4pq5/Y2dOiIwFS0tUqLwTSrIwHx2bnIohKygGz/52SpsiAeo+AB5D7UEVCaQG+RENvlUcD99cZvw=="
	rawItem := domain.NewRawItem("", encodedDetails, "", encodedKeys, "", "", 0, 0, false)
	decoded := itemService.DecodeDetails(rawItem, keys)

	if len(decoded) == 0 {
		t.Error("DecodeDetails() should pass because of valid encoding")
	}
}

func TestDecodeItems(t *testing.T) {
	expected := 29
	pass := "freddy"
	vault := domain.NewVault("../../../../assets/onepassword_data")

	var crytpoUtils out.CrytpoUtils
	var itemRepo out.ItemRepo
	var profileRepo out.ProfileRepo

	var keyService coreservice.KeyService

	crytpoUtils = crypto.NewPbkdf2CryptoUtils()
	itemRepo = file.NewFileItemRepo()
	profileRepo = file.NewFileProfileRepo()
	profileRepo.LoadProfile(vault)

	keyService = coreservice.NewDfltKeyService(crytpoUtils, profileRepo)
	itemService := coreservice.NewDfltItemService(keyService, itemRepo)

	derivedKey, derivedMac, _ := keyService.DerivedKeys(pass)
	overviewKey, overviewMac, _ := keyService.OverviewKeys(derivedKey, derivedMac)
	masterKey, masterMac, _ := keyService.MasterKeys(derivedKey, derivedMac)
	keys := domain.NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac)
	itemService.DecodeItems(vault, keys)
	items := itemService.GetSimpleItems(nil, false)
	trashed := itemService.GetSimpleItems(nil, true)

	if expected != len(items)+len(trashed) {
		t.Errorf("DecodeItems() = %d; expected = %d", len(items)+len(trashed), expected)
	}
}

func TestDecodeOverview(t *testing.T) {
	itemService, keys := setupItemServiceAndKeys()
	encodedOverview := "b3BkYXRhMDGYAAAAAAAAAGRUHVTI1ig3dPmw3gdUasxYzglzq53+WXaerBgPS44zyc6fQEOjLHfD/qP/uRqwQvuW+PlRC9gKqFoTrptjy/ImutcydczWYgEp333LL7KMi7XEy5aJmxrITgHytmdguNn380ZmygliXTvWMZm7N4TkkOPWZ3FXRSGoJ77XEQDmexwJsVdoxUgATnYUwubBLaybVJkcVzGGeAeeFMJxo8Z/wbsOWV84yxVNRwW5kLyO73x+WMC0zAZG1utUy8qNPW30II9Rs5KRQcYa14Vvl48="
	rawItem := domain.NewRawItem("", "", "", "", encodedOverview, "", 0, 0, false)
	decoded := itemService.DecodeOverview(rawItem, keys)

	if len(decoded) == 0 {
		t.Error("DecodeOverview() should pass because of valid encoding")
	}
}

func TestGetItem(t *testing.T) {
	itemService, _ := setupItemServiceAndKeys()
	expected := "YouTube"
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	trashed := false
	item := itemService.GetItem(uid, trashed)

	if item.Title != expected {
		t.Errorf("GetItem() = %v; expected = %v", item.Title, expected)
	}

	expected = ""
	uid = "0C4F27910A64488BB339AED63565D148"
	trashed = true
	item = itemService.GetItem(uid, trashed)

	if item.Title != expected {
		t.Errorf("GetItem() = %v; expected = %v", item.Title, expected)
	}

	item = itemService.GetItem("", false)

	if item != nil {
		t.Error("GetItem() should return nil for not existing item")
	}
}

func TestGetSimpleItems(t *testing.T) {
	itemService, _ := setupItemServiceAndKeys()
	expected := 27
	trashed := false
	first, firstUid := "Bank of America", "EC0A40400ABB4B16926B7417E95C9669"
	last, lastUid := "Email Account", "FD2EADB43C4F4FC7BEB35A1692DDFDEA"
	items := itemService.GetSimpleItems(nil, trashed)

	if len(items) != expected {
		t.Errorf("[NOT-TRASHED] GetSimpleItems() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first || items[0].Uid != firstUid {
		t.Errorf("[FIRST-NOT-TRSHED] GetSimpleItems() = %v, %v; expected = %v, %v", items[0].Title, items[0].Uid, first, firstUid)
	}

	if items[len(items)-1].Title != last || items[len(items)-1].Uid != lastUid {
		t.Errorf("[LAST-NOT-TRASHED] GetSimpleItems() = %v, %v; expected = %v, %v", items[len(items)-1].Title, items[len(items)-1].Uid, last, lastUid)
	}

	trashed = true
	items = itemService.GetSimpleItems(nil, trashed)
	expected = 2
	first, firstUid = "A note to Trash", "AE272805811C450586BA3EDEAEF8AE19"
	last, lastUid = "", "0C4F27910A64488BB339AED63565D148"

	if len(items) != expected {
		t.Errorf("[TRASHED] GetSimpleItems() = %d; expected = %d", len(items), expected)
	}

	if items[0].Title != first || items[0].Uid != firstUid {
		t.Errorf("[FIRST-TRSHED] GetSimpleItems() = %v, %v; expected = %v, %v", items[0].Title, items[0].Uid, first, firstUid)
	}

	if items[len(items)-1].Title != last || items[len(items)-1].Uid != lastUid {
		t.Errorf("[LAST-TRASHED] GetSimpleItems() = %v, %v; expected = %v, %v", items[len(items)-1].Title, items[len(items)-1].Uid, last, lastUid)
	}
}

func TestParseItemField(t *testing.T) {
	itemService, _ := setupItemServiceAndKeys()
	fieldJson := make(map[string]interface{})
	fieldJson["name"] = "username"
	field := itemService.ParseItemField(false, fieldJson)

	if field != nil {
		t.Errorf("ParseItemField() = %v; expected = %v", field, nil)
	}

	expected := "Username"
	fieldJson["value"] = "test@test.org"
	field = itemService.ParseItemField(false, fieldJson)

	if field.Name != expected {
		t.Errorf("ParseItemField() = %v; expected = %v", field.Name, expected)
	}

	sectionFieldJson := make(map[string]interface{})
	sectionFieldJson["t"] = "password"
	sectionFieldJson["k"] = "string"
	field = itemService.ParseItemField(true, sectionFieldJson)

	if field != nil {
		t.Errorf("ParseItemField() = %v; expected = %v", field, nil)
	}

	expected = "1234"
	sectionFieldJson["v"] = "1234"
	field = itemService.ParseItemField(true, sectionFieldJson)

	if field.Value != expected {
		t.Errorf("ParseItemField = %v; expected = %v", field.Value, expected)
	}
}

func TestParseItemSection(t *testing.T) {
	itemService, _ := setupItemServiceAndKeys()
	sectionJson := make(map[string]interface{})
	section := itemService.ParseItemSection(sectionJson)

	if section != nil {
		t.Errorf("ParseItemSection() = %v; expected %v", section, nil)
	}

	expected := "Details"
	sectionJson["title"] = "details"
	section = itemService.ParseItemSection(sectionJson)

	if section.Title != expected {
		t.Errorf("ParseItemSection() = %v; expected %v", section.Title, expected)
	}

	expectedLen := 1
	fieldsJson := make([]interface{}, 0)
	fieldJson := make(map[string]interface{})
	fieldJson["t"] = "password"
	fieldJson["k"] = "string"
	fieldJson["v"] = "1234"
	fieldsJson = append(fieldsJson, fieldJson)
	sectionJson["fields"] = fieldsJson
	section = itemService.ParseItemSection(sectionJson)

	if len(section.Fields) != expectedLen {
		t.Errorf("ParseItemSection() = %d; expected %d", len(section.Fields), expectedLen)
	}
}
