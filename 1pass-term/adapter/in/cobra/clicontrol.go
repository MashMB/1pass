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
	"time"

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

func (ctrl *cobraCliControl) GetItemDetails(vaultPath, uid string, trashed bool) {
	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vaultPath, string(password))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item := ctrl.vaultFacade.GetItem(uid, trashed)

	if item != nil {
		fmt.Println(item.Category.GetName())
		fmt.Println(fmt.Sprintf("[%v] --- %v", item.Uid, item.Title))
		updated := time.Unix(item.Updated, 0).Format("2006-01-02 15:04:05")
		created := time.Unix(item.Created, 0).Format("2006-01-02 15:04:05")
		fmt.Println(fmt.Sprintf("Updated: %v\tCreated: %v\tTrashed: %v", updated, created, item.Trashed))

		if item.Url != "" {
			fmt.Println(fmt.Sprintf("URL: %v", item.Url))
		}

		if item.Sections != nil {
			for _, section := range item.Sections {
				if section.Title != "" {
					fmt.Println(section.Title)
				}

				fmt.Println("------------------------------")

				if section.Fields != nil {
					for _, field := range section.Fields {
						fmt.Println(fmt.Sprintf("\t%v: %v", field.Name, field.Value))
					}
				}
			}
		}
	} else {
		msg := fmt.Sprintf("Item with UID %v do not exist", uid)
		fmt.Println(msg)
	}

}

func (ctrl *cobraCliControl) GetItemOverview(vaultPath, uid string, trashed bool) {
	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vaultPath, string(password))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item := ctrl.vaultFacade.GetItemOverview(uid, trashed)

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
