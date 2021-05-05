// Implementation of default update facade.
//
// @author TSS

package facade

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/service"
)

type dfltUpdateFacade struct {
	updateService service.UpdateService
}

func NewDfltUpdateFacade(updateService service.UpdateService) *dfltUpdateFacade {
	return &dfltUpdateFacade{
		updateService: updateService,
	}
}

func (f *dfltUpdateFacade) CheckforUpdate() (*domain.UpdateInfo, error) {
	return f.updateService.CheckForUpdate()
}
