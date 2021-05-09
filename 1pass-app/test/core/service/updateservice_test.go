// Unit tests for default update service.
//
// @author TSS

package service

import (
	"net/http"
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	coreservice "github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-up/adapter/out/github"
)

func setupUpdateService() coreservice.UpdateService {
	var updater out.Updater

	updater = github.NewGithubUpdater()

	return coreservice.NewDfltUpdateService(updater)
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
		service := setupUpdateService()
		expected := domain.ErrNoUpdate
		timeout := 5
		_, err := service.CheckForUpdate(0, timeout, false, "../../../../assets")

		if err != expected {
			t.Errorf("CheckForUpdate() = %v; expected = %v", err, expected)
		}
	} else {
		t.Log("CheckForUpdate() no internet connection")
	}
}

func TestUpdate(t *testing.T) {
	if isOnline() {
		service := setupUpdateService()
		expected := domain.ErrNoUpdate
		timeout := 5
		err := service.Update(timeout)

		if err != expected {
			t.Errorf("Update() = %v; expected = %v", err, expected)
		}
	} else {
		t.Log("Update() no internet connection")
	}
}
