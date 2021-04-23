// Definition of input controllers.
//
// @author TSS

package in

type CliControl interface {
	GetItems(vaultPath, password string)
}
