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

	if len(logins) != 10 {
		t.Errorf("FindByCategoryAndTrashed() = %d; expected = %d", len(logins), expected)
	}
}
