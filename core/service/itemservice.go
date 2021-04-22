// Default implementation of item service.
//
// @author TSS

package service

import (
	"encoding/base64"
	"encoding/json"
	"sort"
	"strings"

	"github.com/mashmb/1pass/core/domain"
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

func (s *dfltItemService) GetOverview(title string, keys *domain.Keys) []*domain.Item {
	title = strings.TrimSpace(strings.ToLower(title))
	items := make([]*domain.Item, 0)
	rawItems := s.itemRepo.FindByCategoryAndTrashed(domain.ItemCategoryEnum.Login, false)

	for _, rawItem := range rawItems {
		overviewData, _ := base64.StdEncoding.DecodeString(rawItem.Overview)
		overview, _ := s.keyService.DecodeOpdata(overviewData, keys.OverviewKey, keys.OverviewMac)
		var itemData map[string]interface{}
		json.Unmarshal(overview, &itemData)
		itemTitle := ""

		if itemData["title"] != nil {
			itemTitle = strings.TrimSpace(strings.ToLower(itemData["title"].(string)))
		}

		if itemTitle == title {
			cat, err := domain.ItemCategoryEnum.FromCode(rawItem.Category)

			if err == nil {
				overview, _ = json.MarshalIndent(itemData, "", "  ")
				item := domain.NewItem(cat, string(overview), rawItem.Created, rawItem.Updated)
				items = append(items, item)
			}
		}
	}

	return items
}

func (s *dfltItemService) GetSimple(keys *domain.Keys) []*domain.SimpleItem {
	items := make([]*domain.SimpleItem, 0)
	rawItems := s.itemRepo.FindByCategoryAndTrashed(domain.ItemCategoryEnum.Login, false)

	for _, rawItem := range rawItems {
		overviewData, _ := base64.StdEncoding.DecodeString(rawItem.Overview)
		overview, _ := s.keyService.DecodeOpdata(overviewData, keys.OverviewKey, keys.OverviewMac)
		var itemData map[string]interface{}
		json.Unmarshal(overview, &itemData)
		title := ""

		if itemData["title"] != nil {
			title = strings.TrimSpace(itemData["title"].(string))
		}

		item := domain.NewSimpleItem(title)
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Title < items[j].Title
	})

	return items
}
