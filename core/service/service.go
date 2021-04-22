// Definition of application services.
//
// @author TSS

package service

import (
	"github.com/mashmb/1pass/core/domain"
)

type ItemService interface {
	GetOverview(title string, keys *domain.Keys) []*domain.Item

	GetSimple(keys *domain.Keys) []*domain.SimpleItem
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
