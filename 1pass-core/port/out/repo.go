// Definition of repositories.
//
// @author TSS

package out

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type ItemRepo interface {
	FindByCategoryAndTrashed(category *domain.ItemCategory, trashed bool) []*domain.RawItem

	FindFirstByUidAndTrashed(uid string, trashed bool) *domain.RawItem

	LoadItems(vault *domain.Vault) []*domain.RawItem
}

type ProfileRepo interface {
	GetIterations() int

	GetMasterKey() string

	GetOverviewKey() string

	GetSalt() string

	LoadProfile(vault *domain.Vault)
}
