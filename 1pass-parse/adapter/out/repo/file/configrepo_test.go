// Unit tests for file configuration repository.
//
// @author TSS

package file

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

func setupFileConfigRepo() (*fileConfigRepo, *fileConfigRepo) {
	return NewFileConfigRepo("../../../../../assets"), NewFileConfigRepo("")
}

func TestGetDefaultVault(t *testing.T) {
	config, empty := setupFileConfigRepo()
	expected := "./assets/onepassword_data"
	vault := config.GetDefaultVault()

	if vault != expected {
		t.Errorf("GetDefaultVault() = %v; expected = %v", vault, expected)
	}

	vault = empty.GetDefaultVault()

	if vault != "" {
		t.Errorf("GetDefaultVault() = %v; expected = %v", vault, "")
	}
}

func TestGetUpdateNotification(t *testing.T) {
	config, empty := setupFileConfigRepo()
	expected := true
	notification := empty.GetUpdateNotification()

	if notification != expected {
		t.Errorf("GetUpdateNotification() = %v; expected = %v", notification, expected)
	}

	expected = false
	notification = config.GetUpdateNotification()

	if notification != expected {
		t.Errorf("GetUpdateNotification() = %v; expected = %v", notification, expected)
	}
}

func TestLoadConfigFile(t *testing.T) {
	config := loadConfigFile("")

	if len(config) != 0 {
		t.Error("loadConfigFile() should fail because config file do not exist")
	}

	config = loadConfigFile("../../../../../assets")

	if len(config) == 0 {
		t.Error("loadConfigFile() should pass because config file exist")
	}
}

func TestSave(t *testing.T) {
	config, _ := setupFileConfigRepo()
	expected := ""
	conf := domain.NewConfig(false, "")
	config.Save(conf)
	vault := config.GetDefaultVault()

	if vault != expected {
		t.Errorf("Save() = %v; expected = %v", vault, expected)
	}

	expected = "./assets/onepassword_data"
	conf.Vault = "./assets/onepassword_data"
	config.Save(conf)
	vault = config.GetDefaultVault()

	if vault != expected {
		t.Errorf("Save() = %v; expected = %v", vault, expected)
	}
}
