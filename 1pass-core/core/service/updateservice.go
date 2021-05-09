// Implementation of default update service.
//
// @author TSS

package service

import (
	"os"
	"path/filepath"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/out"
)

type dfltUpdateService struct {
	updater out.Updater
}

func NewDfltUpdateService(updater out.Updater) *dfltUpdateService {
	return &dfltUpdateService{
		updater: updater,
	}
}

func (s *dfltUpdateService) CheckForUpdate(period, timeout int, force bool, configDir string) (*domain.UpdateInfo, error) {
	if !force {
		shouldCheck := s.updater.ShouldCheck(period, configDir)

		if !shouldCheck {
			return nil, domain.ErrNoUpdate
		}
	}

	info, err := s.updater.CheckForUpdate(int64(timeout))

	if err != nil {
		return nil, err
	}

	if info == nil || !info.Newer {
		s.updater.CheckTimestamp(configDir)

		return nil, domain.ErrNoUpdate
	}

	s.updater.CheckTimestamp(configDir)

	return info, nil
}

func (s *dfltUpdateService) Update(timeout int, stage func(int)) error {
	info, err := s.updater.CheckForUpdate(int64(timeout))

	if err != nil {
		return err
	}

	if info == nil {
		return domain.ErrNoUpdate
	}

	if !info.Newer {
		return domain.ErrNoUpdate
	} else {
		stage(1)

		if err := os.MkdirAll(domain.CacheDir, 0700); err != nil {
			return err
		}

		archive := filepath.Join(domain.CacheDir, domain.Archive)
		checksum := filepath.Join(domain.CacheDir, domain.Checksum)
		stage(2)

		if err := s.updater.DownloadFile(archive, info.ArchiveUrl, int64(timeout)); err != nil {
			return err
		}

		stage(3)

		if err := s.updater.DownloadFile(checksum, info.ChecksumUrl, int64(timeout)); err != nil {
			return err
		}

		stage(4)

		if err := s.updater.ExtractArchive(archive, domain.CacheDir); err != nil {
			return err
		}

		binary := filepath.Join(domain.CacheDir, domain.AppName)

		stage(5)

		if err := s.updater.ValidateChecksum(binary, checksum); err != nil {
			return err
		}

		stage(6)

		if err := s.updater.ReplaceBinary(binary); err != nil {
			return err
		}

		stage(7)

		if err := os.RemoveAll(domain.CacheDir); err != nil {
			return err
		}
	}

	return nil
}
