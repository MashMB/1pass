// Implementation of file configuration repository.
//
// @author TSS

package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type fileConfigRepo struct {
	config map[string]interface{}
}

func NewFileConfigRepo(configFile string) *fileConfigRepo {
	return &fileConfigRepo{
		config: loadConfigFile(configFile),
	}
}

func loadConfigFile(configFile string) map[string]interface{} {
	if _, err := os.Stat(configFile); err != nil {
		return nil
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

	if repo.config != nil && repo.config["opvault"] != nil {
		vault = fmt.Sprint(repo.config["opvault"])
	}

	return vault
}
