// Implementation of file item repository.
//
// @author TSS

package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mashmb/1pass/port/out"
)

type fileItemRepo struct {
	itemsJson map[string]interface{}
}

func NewFileItemRepo(vaultPath string) *fileItemRepo {
	return &fileItemRepo{
		itemsJson: loadItemsJson(vaultPath),
	}
}

func loadItemsJson(vaultPath string) map[string]interface{} {
	var itemsJson map[string]interface{}
	bandFiles, err := filepath.Glob(filepath.Join(vaultPath, out.ProfileDir, out.BandFilePattern))

	if err != nil {
		log.Fatalln(err)
	}

	items := "{"

	for _, bandFile := range bandFiles {
		file, err := ioutil.ReadFile(bandFile)

		if err != nil {
			log.Fatalln(err)
		}

		items = items + string(file[4:len(file)-3]) + ","
	}

	items = items[:len(items)-1] + "}"

	if err := json.Unmarshal([]byte(items), &itemsJson); err != nil {
		log.Fatalln(err)
	}

	return itemsJson
}
