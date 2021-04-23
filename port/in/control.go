// Definition of input controllers.
//
// @author TSS

package in

type CliControl interface {
	GetItemDetails(vaultPath, password, uid string)

	GetItemOverview(vaultPath, password, uid string)

	GetItems(vaultPath string)
}
