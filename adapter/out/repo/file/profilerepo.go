// File implementation of profile repository.
//
// @author TSS

package file

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mashmb/1pass/port/out"
)

const (
	profileFile string = "profile.js"
)

type fileProfileRepo struct {
	profileJson map[string]interface{}
}

func NewFileProfileRepo(vaultPath string) *fileProfileRepo {
	return &fileProfileRepo{
		profileJson: loadProfileJson(vaultPath),
	}
}

func loadProfileJson(vaultPath string) map[string]interface{} {
	var profileJson map[string]interface{}
	file, err := ioutil.ReadFile(vaultPath + "/" + out.ProfileDir + "/" + profileFile)

	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(file[12:len(file)-1], &profileJson); err != nil {
		log.Fatalln(err)
	}

	return profileJson
}

func (repo *fileProfileRepo) getIterations() int {
	return int(repo.profileJson["iterations"].(float64))
}

func (repo *fileProfileRepo) GetMasterKey() string {
	return repo.profileJson["masterKey"].(string)
}

func (repo *fileProfileRepo) GetOverviewKey() string {
	return repo.profileJson["overviewKey"].(string)
}

func (repo *fileProfileRepo) getSalt() string {
	return repo.profileJson["salt"].(string)
}
