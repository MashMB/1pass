// Application entities.
//
// @author TSS

package domain

type Config struct {
	Timeout      int
	UpdateNotify bool
	Vault        string
}

type Item struct {
	Category *ItemCategory
	Created  int64
	Notes    string
	Title    string
	Trashed  bool
	Sections []*ItemSection
	Uid      string
	Url      string
	Updated  int64
}

type ItemField struct {
	Name  string
	Value string
}

type ItemSection struct {
	Title  string
	Fields []*ItemField
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
	Uid      string
	Updated  int64
}

type SimpleItem struct {
	Category *ItemCategory
	Title    string
	Uid      string
}

type UpdateInfo struct {
	ArchiveUrl  string
	ChecksumUrl string
	Newer       bool
	Version     string
}

type Vault struct {
	Path string
}

func NewConfig(timeout int, updateNotify bool, vault string) *Config {
	return &Config{
		Timeout:      timeout,
		UpdateNotify: updateNotify,
		Vault:        vault,
	}
}

func NewItem(uid, title, url, notes string, trashed bool, category *ItemCategory, sections []*ItemSection, created,
	updated int64) *Item {
	return &Item{
		Category: category,
		Created:  created,
		Notes:    notes,
		Title:    title,
		Trashed:  trashed,
		Sections: sections,
		Uid:      uid,
		Url:      url,
		Updated:  updated,
	}
}

func NewItemSection(title string, fields []*ItemField) *ItemSection {
	return &ItemSection{
		Title:  title,
		Fields: fields,
	}
}

func NewItemField(name, value string) *ItemField {
	return &ItemField{
		Name:  name,
		Value: value,
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

func NewRawItem(category, details, hmac, keys, overview, uid string, created, updated int64, trashed bool) *RawItem {
	return &RawItem{
		Category: category,
		Created:  created,
		Details:  details,
		Hmac:     hmac,
		Keys:     keys,
		Overview: overview,
		Trashed:  trashed,
		Uid:      uid,
		Updated:  updated,
	}
}

func NewSimpleItem(category *ItemCategory, title, uid string) *SimpleItem {
	return &SimpleItem{
		Category: category,
		Title:    title,
		Uid:      uid,
	}
}

func NewUpdateInfo(archiveUrl, checksumUrl, version string, newer bool) *UpdateInfo {
	return &UpdateInfo{
		ArchiveUrl:  archiveUrl,
		ChecksumUrl: checksumUrl,
		Newer:       newer,
		Version:     version,
	}
}

func NewVault(path string) *Vault {
	return &Vault{
		Path: path,
	}
}
