// Implementation of cobra CLI control.
//
// @author TSS

package cobra

import (
	"fmt"
	"os"

	"github.com/mashmb/1pass/core/facade"
)

type cobraCliControl struct {
	vaultFacade facade.VaultFacade
}

func NewCobraCliControl(vaultFacade facade.VaultFacade) *cobraCliControl {
	return &cobraCliControl{
		vaultFacade: vaultFacade,
	}
}

func (ctrl *cobraCliControl) GetItemDetails(vaultPath, password, uid string) {
	err := ctrl.vaultFacade.Unlock(vaultPath, password)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item := ctrl.vaultFacade.GetItemDetails(uid)

	if item != nil {
		fmt.Println(item.Details)
	} else {
		msg := fmt.Sprintf("Item with UID %v do not exist", uid)
		fmt.Println(msg)
	}

}

func (ctrl *cobraCliControl) GetItemOverview(vaultPath, password, uid string) {
	err := ctrl.vaultFacade.Unlock(vaultPath, password)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item := ctrl.vaultFacade.GetItemOverview(uid)

	if item != nil {
		fmt.Println(item.Details)
	} else {
		msg := fmt.Sprintf("Item with UID %v do not exist", uid)
		fmt.Println(msg)
	}
}

func (ctrl *cobraCliControl) GetItems(vaultPath, password string) {
	err := ctrl.vaultFacade.Unlock(vaultPath, password)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	items := ctrl.vaultFacade.GetItems()

	for _, item := range items {
		row := fmt.Sprintf("[%v] --- %v", item.Uid, item.Title)
		fmt.Println(row)
	}
}
