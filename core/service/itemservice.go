// Default implementation of item service.
//
// @author TSS

package service

import (
	"container/list"
	"encoding/base64"
	"encoding/json"

	"github.com/mashmb/1pass/core/domain"
	"github.com/mashmb/1pass/core/domain/enum"
	"github.com/mashmb/1pass/port/out"
)

type dfltItemService struct {
	keyService KeyService
	itemRepo   out.ItemRepo
}

func NewDfltItemService(keyService KeyService, itemRepo out.ItemRepo) *dfltItemService {
	return &dfltItemService{
		keyService: keyService,
		itemRepo:   itemRepo,
	}
}

func (s *dfltItemService) GetSimple(keys *domain.Keys) *list.List {
	items := list.New()
	rawItems := s.itemRepo.FindByCategoryAndTrashed(enum.ItemCategoryEnum.Login, false)

	for element := rawItems.Front(); element != nil; element = element.Next() {
		rawItem := element.Value.(*domain.RawItem)
		overviewData, _ := base64.StdEncoding.DecodeString(rawItem.Overview)
		overview, _ := s.keyService.DecodeOpdata(overviewData, keys.OverviewKey, keys.OverviewMac)
		var itemData map[string]interface{}
		json.Unmarshal(overview, &itemData)
		title := ""

		if itemData["title"] != nil {
			title = itemData["title"].(string)
		}

		item := domain.NewSimpleItem(title)
		items.PushBack(item)
	}

	return items
}
