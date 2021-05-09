// Definition of application services.
//
// @author TSS

package service

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type ConfigService interface {
	GetConfig() *domain.Config

	SaveConfig(config *domain.Config)
}

type ItemService interface {
	ClearMemory()

	DecodeDetails(encoded *domain.RawItem, keys *domain.Keys) map[string]interface{}

	DecodeItems(vault *domain.Vault, keys *domain.Keys)

	DecodeOverview(encoded *domain.RawItem, keys *domain.Keys) map[string]interface{}

	GetItem(uid string, trashed bool) *domain.Item

	GetSimpleItems(category *domain.ItemCategory, title string, trashed bool) []*domain.SimpleItem

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

type UpdateService interface {
	CheckForUpdate(period, timeout int, configDir string) (*domain.UpdateInfo, error)

	Update(timeout int) error
}

type VaultService interface {
	ValidateVault(vault *domain.Vault) error
}
