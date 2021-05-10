// File implementation of profile repository.
//
// @author TSS

package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type fileProfileRepo struct {
	profileJson map[string]interface{}
}

func NewFileProfileRepo() *fileProfileRepo {
	return &fileProfileRepo{}
}

func (repo *fileProfileRepo) GetIterations() int {
	return int(repo.profileJson["iterations"].(float64))
}

func (repo *fileProfileRepo) GetMasterKey() string {
	return repo.profileJson["masterKey"].(string)
}

func (repo *fileProfileRepo) GetOverviewKey() string {
	return repo.profileJson["overviewKey"].(string)
}

func (repo *fileProfileRepo) GetSalt() string {
	return repo.profileJson["salt"].(string)
}

func (repo *fileProfileRepo) LoadProfile(vault *domain.Vault) {
	if repo.profileJson == nil {
		file, err := ioutil.ReadFile(filepath.Join(vault.Path, domain.ProfileDir, domain.ProfileFile))

		if err != nil {
			log.Fatalln(err)
		}

		if err := json.Unmarshal(file[12:len(file)-1], &repo.profileJson); err != nil {
			log.Fatalln(err)
		}
	}
}
