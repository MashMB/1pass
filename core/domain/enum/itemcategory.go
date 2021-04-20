// Implementation of item category enumeration.
//
// @author TSS

package enum

import "github.com/mashmb/1pass/core/domain"

var (
	ItemCategoryEnum = newItemCategoryRegistry()
)

type itemCategory struct {
	Code string
	Name string
}

func newItemCategory(code, name string) *itemCategory {
	return &itemCategory{
		Code: code,
		Name: name,
	}
}

type itemCategoryRegistry struct {
	Login  *itemCategory
	Values []*itemCategory
}

func newItemCategoryRegistry() *itemCategoryRegistry {
	login := newItemCategory("001", "Login")

	return &itemCategoryRegistry{
		Login:  login,
		Values: []*itemCategory{login},
	}
}

func (r *itemCategoryRegistry) fromCode(code string) (*itemCategory, error) {
	var category *itemCategory

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

func (r *itemCategoryRegistry) fromName(name string) (*itemCategory, error) {
	var category *itemCategory

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
