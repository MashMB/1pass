// Unit tests for default configuration service.
//
// @author TSS

package service

import (
	"testing"

	coreservice "github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
)

func setupConfigService() coreservice.ConfigService {
	var configRepo out.ConfigRepo

	configRepo = file.NewFileConfigRepo("../../../../assets")

	return coreservice.NewDfltConfigService(configRepo)
}

func TestIsConfigAvailable(t *testing.T) {
	service := setupConfigService()
	expected := true
	available := service.IsConfigAvailable()

	if available != expected {
		t.Errorf("IsConfigAvailable() = %v; expected = %v", available, expected)
	}
}

func TestGetConfig(t *testing.T) {
	service := setupConfigService()
	config := service.GetConfig()

	if config == nil {
		t.Error("GetConfig() should pass because of valid config")
	}
}

func TestSaveConfig(t *testing.T) {
	service := setupConfigService()
	expected := ""
	config := service.GetConfig()
	config.Vault = ""
	service.SaveConfig(config)
	config = service.GetConfig()

	if config.Vault != expected {
		t.Errorf("SaveConfig() = %v; expected = %v", config.Vault, expected)
	}

	expected = "./assets/onepassword_data"
	config.Vault = "./assets/onepassword_data"
	service.SaveConfig(config)
	config = service.GetConfig()

	if config.Vault != expected {
		t.Errorf("SaveConfig() = %v; expected = %v", config.Vault, expected)
	}
}
