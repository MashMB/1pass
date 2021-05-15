// Unit tests for file configuration repository.
//
// @author TSS

package file

import (
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

func setupFileConfigRepo() (*fileConfigRepo, *fileConfigRepo) {
	return NewFileConfigRepo("../../../assets"), NewFileConfigRepo("")
}

func TestIsAvailable(t *testing.T) {
	config, empty := setupFileConfigRepo()
	expected := false
	exist := empty.IsAvailable()

	if exist != expected {
		t.Errorf("FileExist() = %v; expected = %v", exist, expected)
	}

	expected = true
	exist = config.IsAvailable()

	if exist != expected {
		t.Errorf("FileExist() = %v; expected = %v", exist, expected)
	}
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

func TestGetTimeout(t *testing.T) {
	config, empty := setupFileConfigRepo()
	expected := 1
	timeout := empty.GetTimeout()

	if timeout != expected {
		t.Errorf("GetTimeout() = %d; expected = %d", timeout, expected)
	}

	expected = 10
	timeout = config.GetTimeout()

	if timeout != expected {
		t.Errorf("GetTimeout() = %d; expected = %d", timeout, expected)
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

func TestGetUpdatePeriod(t *testing.T) {
	config, empty := setupFileConfigRepo()
	expected := 1
	period := empty.GetUpdatePeriod()

	if period != expected {
		t.Errorf("GetUpdatePeriod() = %d; expected = %d", period, expected)
	}

	expected = 7
	period = config.GetUpdatePeriod()

	if period != expected {
		t.Errorf("GetUpdatePeriod() = %d; expected = %d", period, expected)
	}
}

func TestLoadConfigFile(t *testing.T) {
	config := loadConfigFile("")

	if len(config) != 0 {
		t.Error("loadConfigFile() should fail because config file do not exist")
	}

	config = loadConfigFile("../../../assets")

	if len(config) == 0 {
		t.Error("loadConfigFile() should pass because config file exist")
	}
}

func TestSave(t *testing.T) {
	config, _ := setupFileConfigRepo()
	expected := ""
	conf := domain.NewConfig(10, 7, false, "")
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
