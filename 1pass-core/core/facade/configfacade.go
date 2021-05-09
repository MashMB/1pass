// Default implementation of configuration facade.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/service"
)

type dfltConfigFacade struct {
	configService service.ConfigService
}

func NewDfltConfigFacade(configService service.ConfigService) *dfltConfigFacade {
	return &dfltConfigFacade{
		configService: configService,
	}
}

func (f *dfltConfigFacade) IsConfigAvailable() bool {
	return f.configService.IsConfigAvailable()
}

func (f *dfltConfigFacade) GetConfig() *domain.Config {
	return f.configService.GetConfig()
}

func (f *dfltConfigFacade) SaveConfig(config *domain.Config) {
	f.configService.SaveConfig(config)
}
