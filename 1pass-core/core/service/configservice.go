// Implementation of default configuration service.
//
// @author TSS

package service

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/out"
)

type dfltConfigService struct {
	configRepo out.ConfigRepo
}

func NewDfltConfigService(configRepo out.ConfigRepo) *dfltConfigService {
	return &dfltConfigService{
		configRepo: configRepo,
	}
}

func (s *dfltConfigService) GetConfig() *domain.Config {
	notification := s.configRepo.GetUpdateNotification()
	period := s.configRepo.GetUpdatePeriod()
	timeout := s.configRepo.GetTimeout()
	vault := s.configRepo.GetDefaultVault()

	return domain.NewConfig(timeout, period, notification, vault)
}

func (s *dfltConfigService) SaveConfig(config *domain.Config) {
	s.configRepo.Save(config)
}
