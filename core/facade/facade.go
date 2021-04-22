// Definition of application use cases.
//
// @author TSS

package facade

import "github.com/mashmb/1pass/core/domain"

type VaultFacade interface {
	GetItemOverview(title string) []*domain.Item

	GetItems() []*domain.SimpleItem

	IsUnlocked() bool

	Lock()

	Unlock(password string) error
}
