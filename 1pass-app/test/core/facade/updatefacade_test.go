// Unit tests for default update facade.
//
// @author TSS

package facade

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	corefacade "github.com/mashmb/1pass/1pass-core/core/facade"
	"github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-up/adapter/out/github"
)

func setupUpdateFacade() corefacade.UpdateFacade {
	var updater out.Updater

	var updateService service.UpdateService

	updater = github.NewGithubUpdater()

	updateService = service.NewDfltUpdateService(updater)

	return corefacade.NewDfltUpdateFacade(updateService)
}

func TestCheckForUpdate(t *testing.T) {
	facade := setupUpdateFacade()
	expected := domain.ErrNoUpdate
	_, err := facade.CheckForUpdate()

	if err != expected {
		t.Errorf("CheckForUpdate() = %v; expected = %v", err, expected)
	}
}

func TestUpdate(t *testing.T) {
	facade := setupUpdateFacade()
	expected := domain.ErrNoUpdate
	err := facade.Update()

	if err != expected {
		t.Errorf("Update() = %v; expected = %v", err, expected)
	}
}
