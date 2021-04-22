// Implementation of default vault service.
//
// @author TSS

package service

import (
	"os"
	"path/filepath"

	"github.com/mashmb/1pass/core/domain"
	"github.com/mashmb/1pass/port/out"
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
		return err
	}

	if !info.IsDir() {
		return domain.ErrInvalidVault
	}

	profilePath := filepath.Join(vault.Path, domain.ProfileDir, domain.ProfileFile)

	info, err = os.Stat(profilePath)

	if err != nil {
		return err
	}

	if info.IsDir() {
		return domain.ErrInvalidVault
	}

	s.profileRepo.LoadProfile(vault)
	s.itemRepo.LoadItems(vault)

	return nil
}
