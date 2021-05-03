// Definition of application use cases.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type ConfigFacade interface {
	GetConfig() *domain.Config
}

type VaultFacade interface {
	GetItem(uid string, trashed bool) *domain.Item

	GetItems(category *domain.ItemCategory, trashed bool) []*domain.SimpleItem

	IsUnlocked() bool

	Lock()

	Unlock(path, password string) error
}
