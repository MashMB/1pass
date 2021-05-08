// Unit tests for default update facade.
//
// @author TSS

package facade

import (
	"net/http"
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

func isOnline() bool {
	_, err := http.Get("http://google.com")

	if err != nil {
		return false
	}

	return true
}

func TestCheckForUpdate(t *testing.T) {
	if isOnline() {
		facade := setupUpdateFacade()
		expected := domain.ErrNoUpdate
		_, err := facade.CheckForUpdate()

		if err != expected {
			t.Errorf("CheckForUpdate() = %v; expected = %v", err, expected)
		}
	} else {
		t.Log("CheckForUpdate() no internet connection")
	}
}

func TestUpdate(t *testing.T) {
	if isOnline() {
		facade := setupUpdateFacade()
		expected := domain.ErrNoUpdate
		err := facade.Update()

		if err != expected {
			t.Errorf("Update() = %v; expected = %v", err, expected)
		}
	} else {
		t.Log("Update() no internet connection")
	}
}
