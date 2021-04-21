// Implementation of file item repository.
//
// @author TSS

package file

import (
	"container/list"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mashmb/1pass/core/domain"
	"github.com/mashmb/1pass/core/domain/enum"
	"github.com/mashmb/1pass/port/out"
)

type fileItemRepo struct {
	items *list.List
}

func NewFileItemRepo(vaultPath string) *fileItemRepo {
	return &fileItemRepo{
		items: loadItems(vaultPath),
	}
}

func loadItems(vaultPath string) *list.List {
	items := list.New()
	itemsJson := loadItemsJson(vaultPath)

	for _, value := range itemsJson {
		v := value.(map[string]interface{})
		created := int64(v["created"].(float64))
		updated := int64(v["updated"].(float64))
		trashed := false

		if v["trashed"] != nil {
			trashed = v["trashed"].(bool)
		}

		item := domain.NewRawItem(v["category"].(string), v["d"].(string), v["hmac"].(string), v["k"].(string),
			v["o"].(string), created, updated, trashed)
		items.PushBack(item)
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

func (repo *fileItemRepo) FindByCategoryAndTrashed(category *enum.ItemCategory, trashed bool) *list.List {
	resultSet := list.New()

	for element := repo.items.Front(); element != nil; element = element.Next() {
		item := element.Value.(*domain.RawItem)
		cat, err := enum.ItemCategoryEnum.FromCode(item.Category)

		if err == nil && cat == category && item.Trashed == trashed {
			resultSet.PushBack(item)
		}
	}

	return resultSet
}
