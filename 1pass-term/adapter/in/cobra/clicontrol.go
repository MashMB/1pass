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

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/core/facade"
	"golang.org/x/term"
)

type cobraCliControl struct {
	configFacade facade.ConfigFacade
	updateFacade facade.UpdateFacade
	vaultFacade  facade.VaultFacade
}

func NewCobraCliControl(configFacade facade.ConfigFacade, updateFacade facade.UpdateFacade,
	vaultFacade facade.VaultFacade) *cobraCliControl {
	return &cobraCliControl{
		configFacade: configFacade,
		updateFacade: updateFacade,
		vaultFacade:  vaultFacade,
	}
}

func (ctrl *cobraCliControl) CheckForUpdate() {
	config := ctrl.configFacade.GetConfig()

	if config.UpdateNotify {
		info, err := ctrl.updateFacade.CheckForUpdate()

		if err == nil {
			msg := fmt.Sprintf("New version of 1pass is available (current: %v, available: %v). Run 'sudo 1pass update' to upgrade.\n",
				domain.Version, info.Version)
			fmt.Println(msg)
		}
	}
}

func (ctrl *cobraCliControl) Configure() {
	ctrl.CheckForUpdate()
	var vault string
	var notify string
	config := ctrl.configFacade.GetConfig()

	fmt.Println("Configuring 1pass:")
	fmt.Print(fmt.Sprintf("  1. Default OPVault path (%v): ", config.Vault))
	fmt.Scanln(&vault)
	config.Vault = strings.TrimSpace(vault)

	fmt.Print(fmt.Sprintf("  2. Update notifications? (%v) [y - for yes/n - for no]: ",
		domain.LogicValEnum.FromValue(config.UpdateNotify).GetName()))
	fmt.Scanln(&notify)
	notifyVal, err := domain.LogicValEnum.FromName(notify)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.UpdateNotify = notifyVal.GetValue()

	ctrl.configFacade.SaveConfig(config)
}

func (ctrl *cobraCliControl) GetCategories() {
	ctrl.CheckForUpdate()
	t := table.NewWriter()
	t.SetStyle(table.StyleDouble)
	t.AppendHeader(table.Row{"lp.", "category"})
	categories := domain.ItemCategoryEnum.GetValues()

	sort.Slice(categories, func(i, j int) bool {
		return categories[i].GetCode() < categories[j].GetCode()
	})

	for i, cat := range categories {
		t.AppendRow(table.Row{i + 1, cat.GetName()})
	}

	fmt.Println(t.Render())
}

func (ctrl *cobraCliControl) GetItemDetails(vaultPath, uid string, trashed bool) {
	ctrl.CheckForUpdate()
	var vault *domain.Vault

	if vaultPath != "" {
		vault = domain.NewVault(vaultPath)
	}

	err := ctrl.vaultFacade.Validate(vault)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vault, string(password))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item := ctrl.vaultFacade.GetItem(uid, trashed)

	if item != nil {
		fmt.Println()
		fmt.Println(item.Category.GetName())
		fmt.Println("------------------------------")
		fmt.Println(item.Title)
		fmt.Println()
		updated := time.Unix(item.Updated, 0).Format("2006-01-02 15:04:05")
		created := time.Unix(item.Created, 0).Format("2006-01-02 15:04:05")
		fmt.Println(fmt.Sprintf("Updated: %v\nCreated: %v\nTrashed: %v", updated, created, item.Trashed))

		if item.Url != "" {
			fmt.Println(fmt.Sprintf("URL: %v", item.Url))
		}

		if item.Sections != nil {
			for _, section := range item.Sections {
				fmt.Println()

				if section.Title != "" {
					fmt.Println(section.Title)
				}

				fmt.Println("------------------------------")

				if section.Fields != nil {
					for _, field := range section.Fields {
						fmt.Println(fmt.Sprintf("%v: %v", field.Name, field.Value))
					}
				}
			}

			fmt.Println()
		}

		if item.Notes != "" {
			fmt.Println("Notes")
			fmt.Println("------------------------------")
			fmt.Println(item.Notes)
			fmt.Println()
		}
	} else {
		msg := fmt.Sprintf("Item with UID %v do not exist", uid)
		fmt.Println(msg)
	}

}

func (ctrl *cobraCliControl) GetItemOverview(vaultPath, uid string, trashed bool) {
	ctrl.CheckForUpdate()
	var vault *domain.Vault

	if vaultPath != "" {
		vault = domain.NewVault(vaultPath)
	}

	err := ctrl.vaultFacade.Validate(vault)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vault, string(password))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item := ctrl.vaultFacade.GetItem(uid, trashed)

	if item != nil {
		fmt.Println()
		fmt.Println(item.Category.GetName())
		fmt.Println("------------------------------")
		fmt.Println(item.Title)
		fmt.Println()
		updated := time.Unix(item.Updated, 0).Format("2006-01-02 15:04:05")
		created := time.Unix(item.Created, 0).Format("2006-01-02 15:04:05")
		fmt.Println(fmt.Sprintf("Updated: %v\nCreated: %v\nTrashed: %v", updated, created, item.Trashed))

		if item.Url != "" {
			fmt.Println(fmt.Sprintf("URL: %v", item.Url))
		}

		if item.Sections != nil {
			for _, section := range item.Sections {
				fmt.Println()

				if section.Title != "" {
					fmt.Println(section.Title)
				}

				fmt.Println("------------------------------")

				if section.Fields != nil {
					for _, field := range section.Fields {
						fmt.Println(fmt.Sprintf("%v: %v", field.Name, "**********"))
					}
				}
			}

			fmt.Println()
		}

		if item.Notes != "" {
			fmt.Println("Notes")
			fmt.Println("------------------------------")
			fmt.Println("**********")
			fmt.Println()
		}
	} else {
		msg := fmt.Sprintf("Item with UID %v do not exist", uid)
		fmt.Println(msg)
	}
}

func (ctrl *cobraCliControl) GetItems(vaultPath, category, title string, trashed bool) {
	ctrl.CheckForUpdate()
	var vault *domain.Vault

	if vaultPath != "" {
		vault = domain.NewVault(vaultPath)
	}

	err := ctrl.vaultFacade.Validate(vault)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Password:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	err = ctrl.vaultFacade.Unlock(vault, string(password))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleDouble)
	t.AppendHeader(table.Row{"lp.", "uid", "category", "title"})

	var cat *domain.ItemCategory

	if category != "" {
		category = strings.TrimSpace(category)
		cat, err = domain.ItemCategoryEnum.FromName(category)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	title = strings.TrimSpace(title)
	items := ctrl.vaultFacade.GetItems(cat, title, trashed)

	if len(items) > 0 {
		for i, item := range items {
			t.AppendRow(table.Row{i + 1, item.Uid, item.Category.GetName(), item.Title})
		}

		fmt.Println(t.Render())
	} else {
		msg := fmt.Sprintf("No results for search (category: %v, name: %v, trashed: %v)", category, title, trashed)
		fmt.Println(msg)
	}
}

func (ctrl *cobraCliControl) Update() {
	fmt.Println("Checking for 1pass application updates...")

	info, err := ctrl.updateFacade.CheckForUpdate()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	msg := fmt.Sprintf("Updating 1pass application from version %v to %v...", domain.Version, info.Version)
	fmt.Println(msg)

	if err := ctrl.updateFacade.Update(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("1pass application updated to version %v", info.Version))
}
