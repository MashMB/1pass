// Definition of application services.
//
// @author TSS

package service

type KeyService interface {
	CheckHmac(msg, key, desiredHmac []byte)

	DecodeData(key, initVector, data []byte) []byte

	DecodeKeys(key, derivedKey, derivedMac []byte) ([]byte, []byte)

	DecodeOpdata(cipherText, key, macKey []byte) []byte

	OverviewKeys(derivedKey, derivedMac []byte) ([]byte, []byte)
}
