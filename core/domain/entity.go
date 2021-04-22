// Application entities.
//
// @author TSS

package domain

type Item struct {
	Category *ItemCategory
	Created  int64
	Details  string
	Updated  int64
}

type Keys struct {
	DerivedKey  []byte
	DerivedMac  []byte
	MasterKey   []byte
	MasterMac   []byte
	OverviewKey []byte
	OverviewMac []byte
}

type RawItem struct {
	Category string
	Created  int64
	Details  string
	Hmac     string
	Keys     string
	Overview string
	Trashed  bool
	Updated  int64
}

type SimpleItem struct {
	Title string
}

func NewItem(category *ItemCategory, details string, created, updated int64) *Item {
	return &Item{
		Category: category,
		Created:  created,
		Details:  details,
		Updated:  updated,
	}
}

func NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac []byte) *Keys {
	return &Keys{
		DerivedKey:  derivedKey,
		DerivedMac:  derivedMac,
		MasterKey:   masterKey,
		MasterMac:   masterMac,
		OverviewKey: overviewKey,
		OverviewMac: overviewMac,
	}
}

func NewRawItem(category, details, hmac, keys, overview string, created, updated int64, trashed bool) *RawItem {
	return &RawItem{
		Category: category,
		Created:  created,
		Details:  details,
		Hmac:     hmac,
		Keys:     keys,
		Overview: overview,
		Trashed:  trashed,
		Updated:  updated,
	}
}

func NewSimpleItem(title string) *SimpleItem {
	return &SimpleItem{
		Title: title,
	}
}
