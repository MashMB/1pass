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
	"github.com/mashmb/1pass/port/out"
)

type fileItemRepo struct {
	items []*domain.RawItem
}

func NewFileItemRepo(vaultPath string) *fileItemRepo {
	return &fileItemRepo{
		items: loadItems(vaultPath),
	}
}

func loadItems(vaultPath string) []*domain.RawItem {
	items := make([]*domain.RawItem, 0)
	itemsJson := loadItemsJson(vaultPath)

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

	return items
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
