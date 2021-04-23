// Implementation of default vault service.
//
// @author TSS

package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/out"
)

type dfltVaultService struct {
	itemRepo    out.ItemRepo
	profileRepo out.ProfileRepo
}

func NewDfltVaultService(itemRepo out.ItemRepo, profileRepo out.ProfileRepo) *dfltVaultService {
	return &dfltVaultService{
		itemRepo:    itemRepo,
		profileRepo: profileRepo,
	}
}

func (s *dfltVaultService) ValidateVault(vault *domain.Vault) error {
	info, err := os.Stat(vault.Path)

	if err != nil {
		errMsg := fmt.Sprintf(domain.ErrInvalidVault.Error(), vault.Path+" is not a valid path")

		return errors.New(errMsg)
	}

	if !info.IsDir() {
		errMsg := fmt.Sprintf(domain.ErrInvalidVault.Error(), vault.Path+" is not a directory")

		return errors.New(errMsg)
	}

	profilePath := filepath.Join(vault.Path, domain.ProfileDir, domain.ProfileFile)

	info, err = os.Stat(profilePath)

	if err != nil {
		errMsg := fmt.Sprintf(domain.ErrInvalidVault.Error(), profilePath+" do not exist")

		return errors.New(errMsg)
	}

	if info.IsDir() {
		errMsg := fmt.Sprintf(domain.ErrInvalidVault.Error(), profilePath+" is not a file")

		return errors.New(errMsg)
	}

	s.profileRepo.LoadProfile(vault)
	s.itemRepo.LoadItems(vault)

	return nil
}
