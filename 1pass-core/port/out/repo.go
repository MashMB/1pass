// Definition of repositories.
//
// @author TSS

package out

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type ConfigRepo interface {
	GetDefaultVault() string

	GetUpdateNotification() bool

	Save(config *domain.Config)
}

type ItemRepo interface {
	FindByCategoryAndTitleAndTrashed(category *domain.ItemCategory, title string, trashed bool) []*domain.Item

	FindFirstByUidAndTrashed(uid string, trashed bool) *domain.Item

	LoadItems(vault *domain.Vault) []*domain.RawItem

	RemoveItems()

	StoreItems(items []*domain.Item)
}

type ProfileRepo interface {
	GetIterations() int

	GetMasterKey() string

	GetOverviewKey() string

	GetSalt() string

	LoadProfile(vault *domain.Vault)
}
