// Implementation of default update facade.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/service"
)

type dfltUpdateFacade struct {
	configDir     string
	configService service.ConfigService
	updateService service.UpdateService
}

func NewDfltUpdateFacade(configDir string, configService service.ConfigService, updateService service.UpdateService) *dfltUpdateFacade {
	return &dfltUpdateFacade{
		configDir:     configDir,
		configService: configService,
		updateService: updateService,
	}
}

func (f *dfltUpdateFacade) CheckForUpdate() (*domain.UpdateInfo, error) {
	config := f.configService.GetConfig()

	return f.updateService.CheckForUpdate(config.UpdatePeriod, config.Timeout, f.configDir)
}

func (f *dfltUpdateFacade) Update() error {
	config := f.configService.GetConfig()

	return f.updateService.Update(config.Timeout)
}
