// Definition of application services.
//
// @author TSS

package service

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type ItemService interface {
	GetDetails(uid string, trashed bool, keys *domain.Keys) *domain.Item

	GetOverview(uid string, trashed bool, keys *domain.Keys) *domain.Item

	GetSimple(keys *domain.Keys, category *domain.ItemCategory, trashed bool) []*domain.SimpleItem

	ParseItemField(fromSection bool, data map[string]interface{}) *domain.ItemField

	ParseItemSection(data map[string]interface{}) *domain.ItemSection
}

type KeyService interface {
	CheckHmac(msg, key, desiredHmac []byte) error

	DecodeData(key, initVector, data []byte) ([]byte, error)

	DecodeKeys(key, derivedKey, derivedMac []byte) ([]byte, []byte, error)

	DecodeOpdata(cipherText, key, macKey []byte) ([]byte, error)

	DerivedKeys(password string) ([]byte, []byte, error)

	ItemKeys(item *domain.RawItem, keys *domain.Keys) ([]byte, []byte)

	MasterKeys(derivedKey, derivedMac []byte) ([]byte, []byte, error)

	OverviewKeys(derivedKey, derivedMac []byte) ([]byte, []byte, error)
}

type VaultService interface {
	ValidateVault(vault *domain.Vault) error
}
