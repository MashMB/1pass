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

func (f *dfltUpdateFacade) CheckForUpdate(force bool) (*domain.UpdateInfo, error) {
	config := f.configService.GetConfig()

	return f.updateService.CheckForUpdate(config.UpdatePeriod, config.Timeout, force, f.configDir)
}

func (f *dfltUpdateFacade) Update(stage func(int)) error {
	config := f.configService.GetConfig()

	return f.updateService.Update(config.Timeout, stage)
}
