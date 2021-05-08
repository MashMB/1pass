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

func (s *dfltUpdateService) CheckForUpdate(timeout int) (*domain.UpdateInfo, error) {
	info, err := s.updater.CheckForUpdate(int64(timeout))

	if err != nil {
		return nil, err
	}

	if info == nil || !info.Newer {
		return nil, domain.ErrNoUpdate
	}

	return info, nil
}

func (s *dfltUpdateService) Update(timeout int) error {
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
		if err := os.MkdirAll(domain.CacheDir, 0700); err != nil {
			return err
		}

		archive := filepath.Join(domain.CacheDir, domain.Archive)
		checksum := filepath.Join(domain.CacheDir, domain.Checksum)

		if err := s.updater.DownloadFile(archive, info.ArchiveUrl, int64(timeout)); err != nil {
			return err
		}

		if err := s.updater.DownloadFile(checksum, info.ChecksumUrl, int64(timeout)); err != nil {
			return err
		}

		if err := s.updater.ExtractArchive(archive, domain.CacheDir); err != nil {
			return err
		}

		binary := filepath.Join(domain.CacheDir, domain.AppName)

		if err := s.updater.ValidateChecksum(binary, checksum); err != nil {
			return err
		}

		if err := s.updater.ReplaceBinary(binary); err != nil {
			return err
		}

		if err := os.RemoveAll(domain.CacheDir); err != nil {
			return err
		}
	}

	return nil
}
