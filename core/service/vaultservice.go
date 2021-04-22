// Implementation of default vault service.
//
// @author TSS

package service

import (
	"os"
	"path/filepath"

	"github.com/mashmb/1pass/core/domain"
)

type dfltVaultService struct {
}

func NewDfltVaultService() *dfltVaultService {
	return &dfltVaultService{}
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

	return nil
}
