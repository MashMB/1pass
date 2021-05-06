// Unit tests for default update service.
//
// @author TSS

package service

import (
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

func TestCheckForUpdate(t *testing.T) {
	service := setupUpdateService()
	expected := domain.ErrNoUpdate
	_, err := service.CheckForUpdate()

	if err != expected {
		t.Errorf("CheckForUpdate() = %v; expected = %v", err, expected)
	}
}

func TestUpdate(t *testing.T) {
	service := setupUpdateService()
	expected := domain.ErrNoUpdate
	err := service.Update()

	if err != expected {
		t.Errorf("Update() = %v; expected = %v", err, expected)
	}
}
