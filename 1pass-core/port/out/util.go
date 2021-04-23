// Definition of application output utils.
//
// @author TSS

package out

import (
	"hash"
)

type CrytpoUtils interface {
	DeriveKey(password, salt []byte, iterations, keyLength int, h func() hash.Hash) []byte
}
