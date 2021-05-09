// Definition of application use cases.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type ConfigFacade interface {
	IsConfigAvailable() bool

	GetConfig() *domain.Config

	SaveConfig(config *domain.Config)
}

type VaultFacade interface {
	GetItem(uid string, trashed bool) *domain.Item

	GetItems(category *domain.ItemCategory, title string, trashed bool) []*domain.SimpleItem

	IsUnlocked() bool

	Lock()

	Unlock(vault *domain.Vault, password string) error

	Validate(vault *domain.Vault) error
}

type UpdateFacade interface {
	CheckForUpdate(force bool) (*domain.UpdateInfo, error)

	Update(stage func(int)) error
}
