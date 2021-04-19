// Definition of application use cases.
//
// @author TSS

package facade

type VaultFacade interface {
	Lock()

	Unlock(password string) error
}
