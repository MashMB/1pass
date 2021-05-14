// Unit tests for file implementation of item repository.
//
// @author TSS

package file

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

func setupFileItemRepo() *fileItemRepo {
	vault := domain.NewVault("../../../assets/onepassword_data")
	repo := NewFileItemRepo()
	items := make([]*domain.Item, 0)
	rawItems := repo.LoadItems(vault)

	for _, rawItem := range rawItems {
		cat, err := domain.ItemCategoryEnum.FromCode(rawItem.Category)

		if err == nil {
			item := domain.NewItem(rawItem.Uid, "", "", "", rawItem.Trashed, cat, nil, rawItem.Created, rawItem.Updated)
			items = append(items, item)
		}
	}

	repo.StoreItems(items)

	return repo
}

func TestCountByCategoryAndTrashed(t *testing.T) {
	repo := setupFileItemRepo()
	expected := 27
	all := repo.CountByCategoryAndTrashed(nil, false)

	if all != expected {
		t.Errorf("CountByCategoryAndTrashed() = %d; expected %d", all, expected)
	}

	expected = 2
	trashed := repo.CountByCategoryAndTrashed(nil, true)

	if trashed != expected {
		t.Errorf("CountByCategoryAndTrashed() = %d; expected %d", trashed, expected)
	}

	expected = 10
	logins := repo.CountByCategoryAndTrashed(domain.ItemCategoryEnum.Login, false)

	if logins != expected {
		t.Errorf("CountByCategoryAndTrashed() = %d; expected %d", logins, expected)
	}
}

func TestFindByCategoryAndTitleAndTrashed(t *testing.T) {
	repo := setupFileItemRepo()
	expected := 27
	all := repo.FindByCategoryAndTitleAndTrashed(nil, "", false)

	if len(all) != expected {
		t.Errorf("[ALL] FindByCategoryAndTitleAndTrashed() = %d; expected = %d", len(all), expected)
	}

	expected = 10
	logins := repo.FindByCategoryAndTitleAndTrashed(domain.ItemCategoryEnum.Login, "", false)

	if len(logins) != expected {
		t.Errorf("[LOGINS] FindByCategoryAndTitleAndTrashed() = %d; expected = %d", len(logins), expected)
	}

	expected = 2
	cards := repo.FindByCategoryAndTitleAndTrashed(domain.ItemCategoryEnum.CreditCard, "", false)

	if len(cards) != expected {
		t.Errorf("[CARD] FindByCategoryAndTitleAndTrashed() = %d; expected = %d", len(cards), expected)
	}

	expected = 2
	trashed := repo.FindByCategoryAndTitleAndTrashed(nil, "", true)

	if len(trashed) != expected {
		t.Errorf("[TRASH] FindByCategoryAndTitleAndTrashed() = %d; expected = %d", len(trashed), expected)
	}
}

func TestFindFirtByUidAndTrashed(t *testing.T) {
	repo := setupFileItemRepo()
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	trashed := false
	item := repo.FindFirstByUidAndTrashed(uid, trashed)

	if item.Uid != uid {
		t.Errorf("[NOT-TRSHED] FindFirstByUidAndTrashed() = %v; expected = %v", item.Uid, uid)
	}

	if item.Trashed != trashed {
		t.Errorf("[NOT-TRASHED] FindFirstByUidAndTrashed() = %v; expected = %v", item.Trashed, trashed)
	}

	uid = "0C4F27910A64488BB339AED63565D148"
	trashed = true
	item = repo.FindFirstByUidAndTrashed(uid, trashed)

	if item.Uid != uid {
		t.Errorf("[TRSHED] FindFirstByUidAndTrashed() = %v; expected = %v", item.Uid, uid)
	}

	if item.Trashed != trashed {
		t.Errorf("[TRASHED] FindFirstByUidAndTrashed() = %v; expected = %v", item.Trashed, trashed)
	}
}

func TestLoadItems(t *testing.T) {
	repo := NewFileItemRepo()
	vault := domain.NewVault("../../../assets/onepassword_data")
	expected := 29
	items := repo.LoadItems(vault)

	if len(items) != expected {
		t.Errorf("LoadItems() = %d; expected = %d", len(items), expected)
	}
}
