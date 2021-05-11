// Definition of input controllers.
//
// @author TSS

package in

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type CliControl interface {
	CheckForUpdate()

	Configure()

	FirstRun()

	GetCategories()

	GetItemDetails(vaultPath, uid string, trashed bool)

	GetItemOverview(vaultPath, uid string, trashed bool)

	GetItems(vaultPath, category, title string, trashed bool)

	Update()
}

type GuiControl interface {
	IsVaultUnlocked() bool

	ValidateVault(vaultPath string) (*domain.Vault, error)
}
