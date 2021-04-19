// Default implementation of vault facade.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/core/domain"
	"github.com/mashmb/1pass/core/service"
)

type dfltVaultFacade struct {
	keys       *domain.Keys
	keyService service.KeyService
}

func NewDfltVaultFacade(keyService service.KeyService) *dfltVaultFacade {
	return &dfltVaultFacade{
		keyService: keyService,
	}
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

func (f *dfltVaultFacade) Unlock(password string) error {
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

	return nil
}
