// Implementation of item category enumeration.
//
// @author TSS

package enum

import "github.com/mashmb/1pass/core/domain"

var (
	ItemCategoryEnum = newItemCategoryRegistry()
)

type ItemCategory struct {
	Code string
	Name string
}

func newItemCategory(code, name string) *ItemCategory {
	return &ItemCategory{
		Code: code,
		Name: name,
	}
}

type itemCategoryRegistry struct {
	Login  *ItemCategory
	Values []*ItemCategory
}

func newItemCategoryRegistry() *itemCategoryRegistry {
	login := newItemCategory("001", "Login")

	return &itemCategoryRegistry{
		Login:  login,
		Values: []*ItemCategory{login},
	}
}

func (r *itemCategoryRegistry) FromCode(code string) (*ItemCategory, error) {
	var category *ItemCategory

	for _, value := range r.Values {
		if value.Code == code {
			category = value
			break
		}
	}

	if category == nil {
		return nil, domain.ErrUnknownItemCat
	}

	return category, nil
}

func (r *itemCategoryRegistry) FromName(name string) (*ItemCategory, error) {
	var category *ItemCategory

	for _, value := range r.Values {
		if value.Name == name {
			category = value
			break
		}
	}

	if category == nil {
		return nil, domain.ErrUnknownItemCat
	}

	return category, nil
}
