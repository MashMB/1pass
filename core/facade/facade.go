// Definition of application use cases.
//
// @author TSS

package facade

import "github.com/mashmb/1pass/core/domain"

type VaultFacade interface {
	GetItemDetails(uid string) *domain.Item

	GetItemOverview(uid string) *domain.Item

	GetItems() []*domain.SimpleItem

	IsUnlocked() bool

	Lock()

	Unlock(path, password string) error
}
