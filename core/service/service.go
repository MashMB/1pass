// Definition of application services.
//
// @author TSS

package service

type KeyService interface {
	CheckHmac(msg, key, desiredHmac []byte)

	DecodeData(key, initVector, data []byte) []byte
}
