// Default implementation of item service.
//
// @author TSS

package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func (s *dfltItemService) ClearMemory() {
	s.itemRepo.RemoveItems()
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
				if detailsJson["fields"] != nil {
					fieldsJson := detailsJson["fields"].([]interface{})
					fields := make([]*domain.ItemField, 0)

					for _, fieldJson := range fieldsJson {
						field := s.ParseItemField(false, fieldJson.(map[string]interface{}))

						if field != nil {
							fields = append(fields, field)
						}
					}

					if len(fields) != 0 {
						section := domain.NewItemSection("", fields)
						sections = append(sections, section)
					}
				}
			} else {
				sectionsJson := detailsJson["sections"].([]interface{})

				for _, sectionJson := range sectionsJson {
					section := s.ParseItemSection(sectionJson.(map[string]interface{}))

					if section != nil {
						sections = append(sections, section)
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

func (s *dfltItemService) GetItem(uid string, trashed bool) *domain.Item {
	return s.itemRepo.FindFirstByUidAndTrashed(uid, trashed)
}

func (s *dfltItemService) GetSimpleItems(category *domain.ItemCategory, trashed bool) []*domain.SimpleItem {
	items := make([]*domain.SimpleItem, 0)
	decodedItems := s.itemRepo.FindByCategoryAndTrashed(category, trashed)

	for _, decoded := range decodedItems {
		item := domain.NewSimpleItem(decoded.Category, decoded.Title, decoded.Uid)
		items = append(items, item)
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
				if dataType == domain.DataTypeEnum.Address {
					value = domain.DataTypeEnum.ParseValue(dataType, "", data["v"].(map[string]interface{}))
				} else {
					value = domain.DataTypeEnum.ParseValue(dataType, fmt.Sprint(data["v"]), nil)
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

	if data["fields"] != nil {
		fieldsData := data["fields"].([]interface{})

		for _, fieldData := range fieldsData {
			field := s.ParseItemField(true, fieldData.(map[string]interface{}))

			if field != nil {
				fields = append(fields, field)
			}
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
