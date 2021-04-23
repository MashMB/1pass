// Definition of input controllers.
//
// @author TSS

package in

type CliControl interface {
	GetItemDetails(vaultPath, password, uid string)

	GetItemOverview(vaultPath, uid string)

	GetItems(vaultPath string)
}
