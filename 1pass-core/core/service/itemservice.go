// Default implementation of item service.
//
// @author TSS

package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"sort"
	"strings"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/out"
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

func (s *dfltItemService) GetDetails(uid string, keys *domain.Keys) *domain.Item {
	var item *domain.Item
	rawItem := s.itemRepo.FindFirstByUidAndTrashed(uid, false)

	if rawItem != nil {
		cat, err := domain.ItemCategoryEnum.FromCode(rawItem.Category)

		if err == nil {
			detailsData, _ := base64.StdEncoding.DecodeString(rawItem.Details)
			itemKey, itemMac := s.keyService.ItemKeys(rawItem, keys)
			details, _ := s.keyService.DecodeOpdata(detailsData, itemKey, itemMac)
			var jsonBuffer bytes.Buffer
			json.Indent(&jsonBuffer, details, "", "  ")
			item = domain.NewItem(cat, string(jsonBuffer.Bytes()), rawItem.Uid, rawItem.Created, rawItem.Updated)
		}
	}

	return item
}

func (s *dfltItemService) GetOverview(uid string, keys *domain.Keys) *domain.Item {
	var item *domain.Item
	rawItem := s.itemRepo.FindFirstByUidAndTrashed(uid, false)

	if rawItem != nil {
		cat, err := domain.ItemCategoryEnum.FromCode(rawItem.Category)

		if err == nil {
			overviewData, _ := base64.StdEncoding.DecodeString(rawItem.Overview)
			overview, _ := s.keyService.DecodeOpdata(overviewData, keys.OverviewKey, keys.OverviewMac)
			var jsonBuffer bytes.Buffer
			json.Indent(&jsonBuffer, overview, "", "  ")
			item = domain.NewItem(cat, string(jsonBuffer.Bytes()), rawItem.Uid, rawItem.Created, rawItem.Updated)
		}
	}

	return item
}

func (s *dfltItemService) GetSimple(keys *domain.Keys, category *domain.ItemCategory) []*domain.SimpleItem {
	items := make([]*domain.SimpleItem, 0)
	rawItems := s.itemRepo.FindByCategoryAndTrashed(category, false)

	for _, rawItem := range rawItems {
		overviewData, _ := base64.StdEncoding.DecodeString(rawItem.Overview)
		overview, _ := s.keyService.DecodeOpdata(overviewData, keys.OverviewKey, keys.OverviewMac)
		var itemOverview map[string]interface{}
		json.Unmarshal(overview, &itemOverview)
		title := ""

		if itemOverview["title"] != nil {
			title = strings.TrimSpace(itemOverview["title"].(string))
		}

		cat, err := domain.ItemCategoryEnum.FromCode(rawItem.Category)

		if err == nil {
			item := domain.NewSimpleItem(cat, title, rawItem.Uid)
			items = append(items, item)
		}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Category.GetCode() != items[j].Category.GetCode() {
			return items[i].Category.GetCode() < items[j].Category.GetCode()
		}

		return items[i].Title < items[j].Title
	})

	return items
}
