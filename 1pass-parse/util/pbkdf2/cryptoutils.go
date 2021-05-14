// Implementation of PBKDF2 crypto utils.
//
// @author TSS

package pbkdf2

import (
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

type pbkdf2CrytpoUtils struct {
}

func NewPbkdf2CryptoUtils() *pbkdf2CrytpoUtils {
	return &pbkdf2CrytpoUtils{}
}

func (u *pbkdf2CrytpoUtils) DeriveKey(password, salt []byte, iterations, keyLength int, h func() hash.Hash) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLength, h)
}
