// Implementation of item category enumeration.
//
// @author TSS

package domain

var (
	ItemCategoryEnum = newItemCategoryRegistry()
)

type ItemCategory struct {
	code string
	name string
}

type itemCategoryRegistry struct {
	Login  *ItemCategory
	values []*ItemCategory
}

func newItemCategory(code, name string) *ItemCategory {
	return &ItemCategory{
		code: code,
		name: name,
	}
}

func newItemCategoryRegistry() *itemCategoryRegistry {
	login := newItemCategory("001", "Login")

	return &itemCategoryRegistry{
		Login:  login,
		values: []*ItemCategory{login},
	}
}

func (ic *ItemCategory) GetCode() string {
	return ic.code
}

func (ic *ItemCategory) GetName() string {
	return ic.name
}

func (r *itemCategoryRegistry) GetValues() []*ItemCategory {
	return r.values
}

func (r *itemCategoryRegistry) FromCode(code string) (*ItemCategory, error) {
	var category *ItemCategory

	for _, value := range r.values {
		if value.code == code {
			category = value
			break
		}
	}

	if category == nil {
		return nil, ErrUnknownItemCat
	}

	return category, nil
}

func (r *itemCategoryRegistry) FromName(name string) (*ItemCategory, error) {
	var category *ItemCategory

	for _, value := range r.values {
		if value.name == name {
			category = value
			break
		}
	}

	if category == nil {
		return nil, ErrUnknownItemCat
	}

	return category, nil
}
