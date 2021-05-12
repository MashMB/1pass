// Main 1Pass application widget.
//
// @author TSS

package gui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/in"
)

type onepassWidget struct {
	name       string
	title      string
	errDialog  *errorDialog
	passPrompt *passwordPrompt
	categories []*domain.ItemCategory
	vault      *domain.Vault
	guiControl in.GuiControl
}

func newOnepassWidget(vault *domain.Vault, guiControl in.GuiControl) *onepassWidget {
	widget := &onepassWidget{
		title:      "1Pass",
		name:       "1pass",
		errDialog:  newErrorDialog(),
		categories: make([]*domain.ItemCategory, 0),
		vault:      vault,
		guiControl: guiControl,
	}

	widget.passPrompt = newPasswordPrompt(widget.unlock)

	return widget
}

func (ow *onepassWidget) lock(ui *gocui.Gui, view *gocui.View) error {
	ow.categories = make([]*domain.ItemCategory, 0)
	ow.guiControl.LockVault()
	view.Clear()

	if err := ow.update(ui); err != nil {
		return err
	}

	return nil
}

func (ow *onepassWidget) promptForPassword(ui *gocui.Gui) error {
	if err := ow.passPrompt.Layout(ui); err != nil {
		return err
	}

	return nil
}

func (ow *onepassWidget) unlock(ui *gocui.Gui, view *gocui.View) error {
	password := strings.TrimSpace(view.ViewBuffer())

	if err := ui.DeleteView(ow.passPrompt.name); err != nil {
		return err
	}

	if err := ow.guiControl.Unlock(ow.vault, password); err != nil {
		ow.errDialog.err = err

		if err := ow.errDialog.Layout(ui); err != nil {
			return err
		}
	} else {
		if _, err := ui.SetCurrentView(ow.name); err != nil {
			return err
		}

		if err := ow.update(ui); err != nil {
			return err
		}
	}

	return nil
}

func (ow *onepassWidget) update(ui *gocui.Gui) error {
	if ow.guiControl.IsVaultUnlocked() {
		maxX, _ := ui.Size()
		view, err := ui.View(ow.name)

		if err != nil {
			return err
		}

		allCount := ow.guiControl.CountItems(nil, false)

		if allCount != 0 {
			allPosition := text.AlignCenter.Apply(fmt.Sprintf("All (%d)", allCount), maxX-1)
			fmt.Fprint(view, allPosition)
			fmt.Fprint(view, "\n")
			ow.categories = append(ow.categories, nil)
			categories := domain.ItemCategoryEnum.GetValues()

			sort.Slice(categories, func(i, j int) bool {
				return categories[i].GetCode() < categories[j].GetCode()
			})

			for _, category := range categories {
				count := ow.guiControl.CountItems(category, false)

				if count != 0 {
					position := text.AlignCenter.Apply(fmt.Sprintf("%v (%d)", category.GetName(), count), maxX-1)
					fmt.Fprint(view, position)
					fmt.Fprint(view, "\n")
					ow.categories = append(ow.categories, category)
				}
			}
		}

		trashedCount := ow.guiControl.CountItems(nil, true)

		if trashedCount != 0 {
			trashedPosition := text.AlignCenter.Apply(fmt.Sprintf("Trashed (%d)", trashedCount), maxX-1)
			fmt.Fprint(view, trashedPosition)
			ow.categories = append(ow.categories, nil)
		}
	} else {
		if err := ow.promptForPassword(ui); err != nil {
			return err
		}
	}

	return nil
}

func (ow *onepassWidget) Keybindings(ui *gocui.Gui) error {
	if err := ow.errDialog.Keybindings(ui); err != nil {
		return err
	}

	if err := ow.passPrompt.Keybindings(ui); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyCtrlL, gocui.ModNone, ow.lock); err != nil {
		return err
	}

	return nil
}

func (ow *onepassWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(ow.name, 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = ow.title

		if _, err := ui.SetCurrentView(ow.name); err != nil {
			return err
		}

		if err := ow.update(ui); err != nil {
			return err
		}
	}

	return nil
}
