// Definition of application use cases.
//
// @author TSS

package facade

type VaultFacade interface {
	IsUnlocked() bool

	Lock()

	Unlock(password string) error
}
