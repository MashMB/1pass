// Definition of application use cases.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type VaultFacade interface {
	GetItemDetails(uid string, trashed bool) *domain.Item

	GetItemOverview(uid string, trashed bool) *domain.Item

	GetItems(category *domain.ItemCategory, trashed bool) []*domain.SimpleItem

	IsUnlocked() bool

	Lock()

	Unlock(path, password string) error
}
