// Unit tests for default configuration facade.
//
// @author TSS

package facade

import (
	"testing"

	corefacade "github.com/mashmb/1pass/1pass-core/core/facade"
	"github.com/mashmb/1pass/1pass-core/core/service"
	"github.com/mashmb/1pass/1pass-core/port/out"
	"github.com/mashmb/1pass/1pass-parse/adapter/out/repo/file"
)

func setupConfigFacade() corefacade.ConfigFacade {
	var configRepo out.ConfigRepo

	var configService service.ConfigService

	configRepo = file.NewFileConfigRepo("../../../../assets")

	configService = service.NewDfltConfigService(configRepo)

	return corefacade.NewDfltConfigFacade(configService)
}

func TestGetConfig(t *testing.T) {
	facade := setupConfigFacade()
	expected := "./assets/onepassword_data"
	config := facade.GetConfig()

	if config.Vault != expected {
		t.Errorf("GetConfig() = %v; expected = %v", config.Vault, expected)
	}
}

func TestSaveConfig(t *testing.T) {
	facade := setupConfigFacade()
	expected := ""
	config := facade.GetConfig()
	config.Vault = expected
	facade.SaveConfig(config)
	config = facade.GetConfig()

	if config.Vault != expected {
		t.Errorf("SaveConfig() = %v; expected = %v", config.Vault, expected)
	}

	expected = "./assets/onepassword_data"
	config.Vault = expected
	facade.SaveConfig(config)
	config = facade.GetConfig()

	if config.Vault != expected {
		t.Errorf("SaveConfig() = %v; expected = %v", config.Vault, expected)
	}
}
