// Default implementation of vault facade.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/service"
)

type dfltVaultFacade struct {
	keys          *domain.Keys
	configService service.ConfigService
	itemService   service.ItemService
	keyService    service.KeyService
	vaultService  service.VaultService
}

func NewDfltVaultFacade(configService service.ConfigService, itemService service.ItemService,
	keyService service.KeyService, vaultService service.VaultService) *dfltVaultFacade {
	return &dfltVaultFacade{
		configService: configService,
		itemService:   itemService,
		keyService:    keyService,
		vaultService:  vaultService,
	}
}

func (f *dfltVaultFacade) GetItem(uid string, trashed bool) *domain.Item {
	return f.itemService.GetItem(uid, trashed)
}

func (f *dfltVaultFacade) GetItems(category *domain.ItemCategory, trashed bool) []*domain.SimpleItem {
	return f.itemService.GetSimpleItems(category, trashed)
}

func (f *dfltVaultFacade) IsUnlocked() bool {
	unlocked := false

	if f.keys != nil {
		unlocked = true
	}

	return unlocked
}

func (f *dfltVaultFacade) Lock() {
	f.keys = nil
}

func (f *dfltVaultFacade) Unlock(path, password string) error {
	var vault *domain.Vault

	if path != "" {
		vault = domain.NewVault(path)
	} else {
		config := f.configService.GetConfig()
		vault = domain.NewVault(config.Vault)
	}

	err := f.vaultService.ValidateVault(vault)

	if err != nil {
		return err
	}

	derivedKey, derivedMac, err := f.keyService.DerivedKeys(password)

	if err != nil {
		return domain.ErrInvalidPass
	}

	masterKey, masterMac, err := f.keyService.MasterKeys(derivedKey, derivedMac)

	if err != nil {
		return domain.ErrInvalidPass
	}

	overviewKey, overviewMac, err := f.keyService.OverviewKeys(derivedKey, derivedMac)

	if err != nil {
		return domain.ErrInvalidPass
	}

	f.keys = domain.NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac)
	f.itemService.DecodeItems(vault, f.keys)

	return nil
}
