// Definition of application output utils.
//
// @author TSS

import (
	"hash"
)

type CrytpoUtils interface {
	DeriveKey(password, salt []byte, iterations, keyLength int, h func() hash.Hash) []byte
}
