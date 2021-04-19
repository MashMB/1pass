// Definition of application use cases.
//
// @author TSS

package facade

type VaultFacade interface {
	Unlock(password string) error
}
