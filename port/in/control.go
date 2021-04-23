// Definition of input controllers.
//
// @author TSS

package in

type CliControl interface {
	GetItemOverview(vaultPath, password, uid string)

	GetItems(vaultPath, password string)
}
