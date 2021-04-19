// Definition of application services.
//
// @author TSS

package service

type KeyService interface {
	CheckHmac(msg, key, desiredHmac []byte) error

	DecodeData(key, initVector, data []byte) ([]byte, error)

	DecodeKeys(key, derivedKey, derivedMac []byte) ([]byte, []byte, error)

	DecodeOpdata(cipherText, key, macKey []byte) ([]byte, error)

	DerivedKeys(password string) ([]byte, []byte, error)

	MasterKeys(derivedKey, derivedMac []byte) ([]byte, []byte, error)

	OverviewKeys(derivedKey, derivedMac []byte) ([]byte, []byte, error)
}
