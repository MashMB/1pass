// Implementation of default update facade.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/service"
)

type dfltUpdateFacade struct {
	configService service.ConfigService
	updateService service.UpdateService
}

func NewDfltUpdateFacade(configService service.ConfigService, updateService service.UpdateService) *dfltUpdateFacade {
	return &dfltUpdateFacade{
		configService: configService,
		updateService: updateService,
	}
}

func (f *dfltUpdateFacade) CheckForUpdate() (*domain.UpdateInfo, error) {
	config := f.configService.GetConfig()

	return f.updateService.CheckForUpdate(config.Timeout)
}

func (f *dfltUpdateFacade) Update() error {
	config := f.configService.GetConfig()

	return f.updateService.Update(config.Timeout)
}
