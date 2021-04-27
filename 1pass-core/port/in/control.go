// Definition of input controllers.
//
// @author TSS

package in

import "github.com/mashmb/1pass/1pass-core/core/domain"

type CliControl interface {
	GetItemDetails(vaultPath, uid string)

	GetItemOverview(vaultPath, uid string)

	GetItems(vaultPath string, category *domain.ItemCategory)
}
