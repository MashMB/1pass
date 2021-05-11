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
