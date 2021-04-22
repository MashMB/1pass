// Unit tests for file implementation of item repository.
//
// @author TSS

package file

import (
	"testing"

	"github.com/mashmb/1pass/core/domain"
)

func setupFileItemRepo() *fileItemRepo {
	return NewFileItemRepo("../../../../assets/onepassword_data")
}

func TestFindByCategoryAndTrashed(t *testing.T) {
	repo := setupFileItemRepo()
	expected := 10
	logins := repo.FindByCategoryAndTrashed(domain.ItemCategoryEnum.Login, false)

	if len(logins) != expected {
		t.Errorf("FindByCategoryAndTrashed() = %d; expected = %d", len(logins), expected)
	}
}

func TestFindFirtByUidAndTrashed(t *testing.T) {
	repo := setupFileItemRepo()
	uid := "358B7411EB8B45CD9CE592ED16F3E9DE"
	trashed := false
	item := repo.FindFirstByUidAndTrashed(uid, trashed)

	if item.Uid != uid {
		t.Errorf("FindFirstByUidAndTrashed() = %v; expected = %v", item.Uid, uid)
	}

	if item.Trashed != trashed {
		t.Errorf("FindFirstByUidAndTrashed() = %v; expected = %v", item.Trashed, trashed)
	}
}
