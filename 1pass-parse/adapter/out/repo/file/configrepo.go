// Implementation of file configuration repository.
//
// @author TSS

package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	"gopkg.in/yaml.v2"
)

type fileConfigRepo struct {
	config    map[string]interface{}
	configDir string
}

func NewFileConfigRepo(configDir string) *fileConfigRepo {
	return &fileConfigRepo{
		config:    loadConfigFile(configDir),
		configDir: configDir,
	}
}

func loadConfigFile(configDir string) map[string]interface{} {
	configFile := filepath.Join(configDir, domain.ConfigFile)

	if _, err := os.Stat(configFile); err != nil {
		return make(map[string]interface{})
	}

	file, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatalln(err)
	}

	var config map[string]interface{}

	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalln(err)
	}

	return config
}

func (repo *fileConfigRepo) GetDefaultVault() string {
	var vault string

	if repo.config["opvault"] != nil {
		vault = fmt.Sprint(repo.config["opvault"])
	}

	return vault
}

func (repo *fileConfigRepo) GetTimeout() int {
	timeout := 2

	if repo.config["timeout"] != nil {
		timeout = repo.config["timeout"].(int)
	}

	return timeout
}

func (repo *fileConfigRepo) GetUpdateNotification() bool {
	notification := true

	if repo.config["update-notification"] != nil {
		notification = repo.config["update-notification"].(bool)
	}

	return notification
}

func (repo *fileConfigRepo) GetUpdatePeriod() int {
	period := 1

	if repo.config["update-period"] != nil {
		period = repo.config["update-period"].(int)
	}

	return period
}

func (repo *fileConfigRepo) Save(config *domain.Config) {
	configFile := filepath.Join(repo.configDir, domain.ConfigFile)
	repo.config["opvault"] = config.Vault
	repo.config["timeout"] = config.Timeout
	repo.config["update-notification"] = config.UpdateNotify
	repo.config["update-period"] = config.UpdatePeriod
	file, err := yaml.Marshal(repo.config)

	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(repo.configDir); err != nil {
		os.MkdirAll(repo.configDir, 0700)
	}

	if err := ioutil.WriteFile(configFile, file, 0644); err != nil {
		log.Fatalln(err)
	}

	repo.config = loadConfigFile(repo.configDir)
}
