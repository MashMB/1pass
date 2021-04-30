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

func (s *dfltItemService) DecodeDetails(encoded *domain.RawItem, keys *domain.Keys) map[string]interface{} {
	var detailsJson map[string]interface{}
	detailsData, _ := base64.StdEncoding.DecodeString(encoded.Details)
	itemKey, itemMac := s.keyService.ItemKeys(encoded, keys)
	details, _ := s.keyService.DecodeOpdata(detailsData, itemKey, itemMac)
	json.Unmarshal(details, &detailsJson)

	return detailsJson
}

func (s *dfltItemService) DecodeItems(vault *domain.Vault, keys *domain.Keys) {
	items := make([]*domain.Item, 0)
	encodedItems := s.itemRepo.LoadItems(vault)

	for _, encoded := range encodedItems {
		cat, err := domain.ItemCategoryEnum.FromCode(encoded.Category)

		if err == nil {
			overviewJson := s.DecodeOverview(encoded, keys)
			detailsJson := s.DecodeDetails(encoded, keys)
			sections := make([]*domain.ItemSection, 0)

			if detailsJson["sections"] == nil {
				fieldsJson := detailsJson["fields"].([]map[string]interface{})
				fields := make([]*domain.ItemField, 0)

				for _, fieldJson := range fieldsJson {
					field := s.ParseItemField(false, fieldJson)

					if field != nil {
						fields = append(fields, field)
					}
				}

				if len(fields) != 0 {
					section := domain.NewItemSection("", fields)
					sections = append(sections, section)
				}
			} else {
				sectionsJson := detailsJson["sections"].([]map[string]interface{})

				for _, sectionJson := range sectionsJson {
					section := s.ParseItemSection(sectionJson)

					if section != nil {
						sections = append(sections, s.ParseItemSection(sectionJson))
					}
				}
			}

			var title string
			var url string
			var notes string

			if overviewJson["title"] != nil {
				title = overviewJson["title"].(string)
			}

			if overviewJson["url"] != nil {
				url = overviewJson["url"].(string)
			}

			if detailsJson["notesPlain"] != nil {
				notes = detailsJson["notesPlain"].(string)
			}

			if len(sections) == 0 {
				sections = nil
			}

			item := domain.NewItem(encoded.Uid, title, url, notes, encoded.Trashed, cat, sections, encoded.Created, encoded.Updated)
			items = append(items, item)
		}
	}

	s.itemRepo.StoreItems(items)
}

func (s *dfltItemService) DecodeOverview(encoded *domain.RawItem, keys *domain.Keys) map[string]interface{} {
	var overviewJson map[string]interface{}
	overviewData, _ := base64.StdEncoding.DecodeString(encoded.Overview)
	overview, _ := s.keyService.DecodeOpdata(overviewData, keys.OverviewKey, keys.OverviewMac)
	json.Unmarshal(overview, &overviewJson)

	return overviewJson
}

func (s *dfltItemService) GetDetails(uid string, trashed bool, keys *domain.Keys) *domain.Item {
	var item *domain.Item
	rawItem := s.itemRepo.FindFirstByUidAndTrashed(uid, trashed)

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

func (s *dfltItemService) GetSimple(keys *domain.Keys, category *domain.ItemCategory, trashed bool) []*domain.SimpleItem {
	items := make([]*domain.SimpleItem, 0)
	rawItems := s.itemRepo.FindByCategoryAndTrashed(category, trashed)

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

func (s *dfltItemService) ParseItemField(fromSection bool, data map[string]interface{}) *domain.ItemField {
	var field *domain.ItemField
	var value string

	if !fromSection {
		if data["value"] != nil {
			value = data["value"].(string)
		}

		if value != "" {
			name := data["name"].(string)
			field = domain.NewItemField(strings.Title(name), value)
		}
	} else {
		if data["v"] != nil {
			dataType, err := domain.DataTypeEnum.FromName(data["k"].(string))

			if err != nil {
				value = data["v"].(string)
			} else {
				switch dataType {
				case domain.DataTypeEnum.Address:
					value = domain.DataTypeEnum.ParseValue(dataType, "", data["v"].(map[string]interface{}))

				case domain.DataTypeEnum.Date:
					unix := data["v"].(int64)
					value = domain.DataTypeEnum.ParseValue(dataType, string(unix), nil)

				default:
					value = domain.DataTypeEnum.ParseValue(dataType, data["v"].(string), nil)
				}
			}

			field = domain.NewItemField(strings.Title(data["t"].(string)), value)
		}
	}

	return field
}

func (s *dfltItemService) ParseItemSection(data map[string]interface{}) *domain.ItemSection {
	var title string
	fields := make([]*domain.ItemField, 0)
	fieldsData := data["fields"].([]map[string]interface{})

	for _, fieldData := range fieldsData {
		field := s.ParseItemField(true, fieldData)

		if field != nil {
			fields = append(fields, field)
		}
	}

	if len(fields) == 0 {
		fields = nil
	}

	if data["title"] != nil {
		title = strings.Title(data["title"].(string))
	}

	if fields == nil && title == "" {
		return nil
	}

	return domain.NewItemSection(strings.Title(title), fields)
}
