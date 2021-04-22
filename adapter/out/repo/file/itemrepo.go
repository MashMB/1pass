// Implementation of file item repository.
//
// @author TSS

package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mashmb/1pass/core/domain"
)

type fileItemRepo struct {
	items []*domain.RawItem
}

func NewFileItemRepo() *fileItemRepo {
	return &fileItemRepo{}
}

func (repo *fileItemRepo) FindByCategoryAndTrashed(category *domain.ItemCategory, trashed bool) []*domain.RawItem {
	resultSet := make([]*domain.RawItem, 0)

	for _, item := range repo.items {
		cat, err := domain.ItemCategoryEnum.FromCode(item.Category)

		if err == nil && cat == category && item.Trashed == trashed {
			resultSet = append(resultSet, item)
		}
	}

	return resultSet
}

func (repo *fileItemRepo) FindFirstByUidAndTrashed(uid string, trashed bool) *domain.RawItem {
	var item *domain.RawItem

	for _, rawItem := range repo.items {
		if rawItem.Uid == uid && rawItem.Trashed == trashed {
			item = rawItem
			break
		}
	}

	return item
}

func (repo *fileItemRepo) LoadItems(vault *domain.Vault) {
	if repo.items == nil {
		var itemsJson map[string]interface{}
		bandFiles, err := filepath.Glob(filepath.Join(vault.Path, domain.ProfileDir, domain.BandFilePattern))

		if err != nil {
			log.Fatalln(err)
		}

		itemsFile := "{"

		for _, bandFile := range bandFiles {
			file, err := ioutil.ReadFile(bandFile)

			if err != nil {
				log.Fatalln(err)
			}

			itemsFile = itemsFile + string(file[4:len(file)-3]) + ","
		}

		itemsFile = itemsFile[:len(itemsFile)-1] + "}"

		if err := json.Unmarshal([]byte(itemsFile), &itemsJson); err != nil {
			log.Fatalln(err)
		}

		items := make([]*domain.RawItem, 0)

		for uid, value := range itemsJson {
			v := value.(map[string]interface{})
			created := int64(v["created"].(float64))
			updated := int64(v["updated"].(float64))
			trashed := false

			if v["trashed"] != nil {
				trashed = v["trashed"].(bool)
			}

			item := domain.NewRawItem(v["category"].(string), v["d"].(string), v["hmac"].(string), v["k"].(string),
				v["o"].(string), uid, created, updated, trashed)
			items = append(items, item)
		}

		repo.items = items
	}
}
