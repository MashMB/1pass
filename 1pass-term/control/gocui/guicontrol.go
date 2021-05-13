// Implementation of gocui GUI control.
//
// @author TSS

package gocui

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/facade"
)

type gocuiGuiControl struct {
	configFacade facade.ConfigFacade
	vaultFacade  facade.VaultFacade
}

func NewGocuiGuiControl(configFacade facade.ConfigFacade, vaultFacade facade.VaultFacade) *gocuiGuiControl {
	return &gocuiGuiControl{
		configFacade: configFacade,
		vaultFacade:  vaultFacade,
	}
}

func (ctrl *gocuiGuiControl) CountItems(category *domain.ItemCategory, trashed bool) int {
	return ctrl.vaultFacade.CountItems(category, trashed)
}

func (ctrl *gocuiGuiControl) GetItem(simple *domain.SimpleItem) *domain.Item {
	return ctrl.vaultFacade.GetItem(simple.Uid, simple.Trashed)
}

func (ctrl *gocuiGuiControl) GetItems(category *domain.ItemCategory, trashed bool) []*domain.SimpleItem {
	return ctrl.vaultFacade.GetItems(category, "", trashed)
}

func (ctrl *gocuiGuiControl) IsVaultUnlocked() bool {
	return ctrl.vaultFacade.IsUnlocked()
}

func (ctrl *gocuiGuiControl) LockVault() {
	ctrl.vaultFacade.Lock()
}

func (ctrl *gocuiGuiControl) Unlock(vault *domain.Vault, password string) error {
	return ctrl.vaultFacade.Unlock(vault, password)
}

func (ctrl *gocuiGuiControl) ValidateVault(vaultPath string) (*domain.Vault, error) {
	var vault *domain.Vault

	if vaultPath != "" {
		vault = domain.NewVault(vaultPath)
	} else {
		config := ctrl.configFacade.GetConfig()
		vault = domain.NewVault(config.Vault)
	}

	if err := ctrl.vaultFacade.Validate(vault); err != nil {
		return nil, err
	}

	return vault, nil
}
