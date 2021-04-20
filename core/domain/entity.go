// Application entities.
//
// @author TSS

package domain

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
