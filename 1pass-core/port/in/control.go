// Definition of input controllers.
//
// @author TSS

package in

type CliControl interface {
	CheckForUpdate()

	Configure()

	GetCategories()

	GetItemDetails(vaultPath, uid string, trashed bool)

	GetItemOverview(vaultPath, uid string, trashed bool)

	GetItems(vaultPath, category string, trashed bool)
}
