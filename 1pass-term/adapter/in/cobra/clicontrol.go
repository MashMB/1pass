// Implementation of cobra CLI control.
//
// @author TSS

package cobra

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"syscall"

	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/facade"
	"golang.org/x/term"
)

type cobraCliControl struct {
	vaultFacade facade.VaultFacade
}

func NewCobraCliControl(vaultFacade facade.VaultFacade) *cobraCliControl {
	return &cobraCliControl{
		vaultFacade: vaultFacade,
	}
}

func (ctrl *cobraCliControl) GetCategories() {
	categories := domain.ItemCategoryEnum.GetValues()

	sort.Slice(categories, func(i, j int) bool {
		return categories[i].GetCode() < categories[j].GetCode()
	})

	for i, cat := range categories {
		msg := fmt.Sprintf("%d. %v", i+1, cat.GetName())
		fmt.Println(msg)
	}
}

func (ctrl *cobraCliControl) GetItemDetails(vaultPath, uid string) {
	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vaultPath, string(password))

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

func (ctrl *cobraCliControl) GetItemOverview(vaultPath, uid string) {
	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vaultPath, string(password))

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

func (ctrl *cobraCliControl) GetItems(vaultPath, category string, trashed bool) {
	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vaultPath, string(password))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var cat *domain.ItemCategory

	if category != "" {
		category = strings.TrimSpace(category)
		cat, err = domain.ItemCategoryEnum.FromName(category)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	items := ctrl.vaultFacade.GetItems(cat, trashed)

	for _, item := range items {
		row := fmt.Sprintf("[%v] (%v) --- %v", item.Uid, item.Category.GetName(), item.Title)
		fmt.Println(row)
	}
}
