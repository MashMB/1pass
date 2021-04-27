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
	BankAccount     *ItemCategory
	CreditCard      *ItemCategory
	Database        *ItemCategory
	DriverLicense   *ItemCategory
	Email           *ItemCategory
	Identity        *ItemCategory
	Login           *ItemCategory
	Membership      *ItemCategory
	OutdoorLicense  *ItemCategory
	Passport        *ItemCategory
	Password        *ItemCategory
	Rewards         *ItemCategory
	Router          *ItemCategory
	SecureNote      *ItemCategory
	Server          *ItemCategory
	SoftwareLicense *ItemCategory
	Ssn             *ItemCategory
	Tombstone       *ItemCategory
	values          []*ItemCategory
}

func newItemCategory(code, name string) *ItemCategory {
	return &ItemCategory{
		code: code,
		name: name,
	}
}

func newItemCategoryRegistry() *itemCategoryRegistry {
	bankAccount := newItemCategory("101", "Bank Account")
	creditCard := newItemCategory("002", "Credit Card")
	database := newItemCategory("102", "Database")
	driverLicense := newItemCategory("103", "Driver License")
	email := newItemCategory("111", "Email")
	identity := newItemCategory("004", "Identity")
	login := newItemCategory("001", "Login")
	membership := newItemCategory("105", "Membership")
	outdoorLicense := newItemCategory("104", "Outdoor License")
	passport := newItemCategory("106", "Passport")
	password := newItemCategory("005", "Password")
	rewards := newItemCategory("107", "Rewards")
	router := newItemCategory("109", "Router")
	secureNote := newItemCategory("003", "Secure Note")
	server := newItemCategory("110", "Server")
	softwareLicense := newItemCategory("100", "Software License")
	ssn := newItemCategory("108", "SSN")
	tombstone := newItemCategory("099", "Tombstone")

	return &itemCategoryRegistry{
		BankAccount:     bankAccount,
		CreditCard:      creditCard,
		Database:        database,
		DriverLicense:   driverLicense,
		Email:           email,
		Identity:        identity,
		Login:           login,
		Membership:      membership,
		OutdoorLicense:  outdoorLicense,
		Passport:        passport,
		Password:        password,
		Rewards:         rewards,
		Router:          router,
		SecureNote:      secureNote,
		Server:          server,
		SoftwareLicense: softwareLicense,
		Ssn:             ssn,
		Tombstone:       tombstone,
		values: []*ItemCategory{bankAccount, creditCard, database, driverLicense, email, identity, login,
			membership, outdoorLicense, passport, password, rewards, router, secureNote, server, softwareLicense,
			ssn, tombstone},
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
