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

const (
	onepassHelp string = `Navigate up/down: k/j
Open: ENTER`
)

type onepassWidget struct {
	currIdx     int
	name        string
	title       string
	errDialog   *errorDialog
	helpWidget  *helpWidget
	itemsWidget *itemsWidget
	passPrompt  *passwordPrompt
	categories  []*domain.ItemCategory
	vault       *domain.Vault
	guiControl  in.GuiControl
}

func newOnepassWidget(helpWidget *helpWidget, vault *domain.Vault, guiControl in.GuiControl) *onepassWidget {
	widget := &onepassWidget{
		currIdx:    -1,
		title:      "1Pass",
		name:       "1pass",
		errDialog:  newErrorDialog(),
		helpWidget: helpWidget,
		categories: make([]*domain.ItemCategory, 0),
		vault:      vault,
		guiControl: guiControl,
	}

	widget.passPrompt = newPasswordPrompt(widget.unlock)
	widget.itemsWidget = newItemsWidget(widget.name, helpWidget, widget.lock, widget.guiControl)

	return widget
}

func (ow *onepassWidget) cursorDown(ui *gocui.Gui, view *gocui.View) error {
	if view != nil && ow.currIdx != -1 && ow.currIdx < len(ow.categories)-1 {
		cx, cy := view.Cursor()

		if err := view.SetCursor(cx, cy+1); err != nil {
			ox, oy := view.Origin()

			if err := view.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		ow.currIdx++
	}

	return nil
}

func (ow *onepassWidget) cursorUp(ui *gocui.Gui, view *gocui.View) error {
	if view != nil && ow.currIdx > 0 {
		ox, oy := view.Origin()
		cx, cy := view.Cursor()

		if err := view.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := view.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}

		ow.currIdx--
	}

	return nil
}

func (ow *onepassWidget) lock(ui *gocui.Gui, view *gocui.View) error {
	ow.currIdx = -1
	ow.categories = make([]*domain.ItemCategory, 0)
	ow.guiControl.LockVault()

	if err := ow.resetCursor(view); err != nil {
		return err
	}

	view.Clear()

	if err := ow.update(ui); err != nil {
		return err
	}

	ow.helpWidget.help = onepassHelp

	if err := ow.helpWidget.update(ui); err != nil {
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

func (ow *onepassWidget) resetCursor(view *gocui.View) error {
	if view != nil {
		ox, oy := view.Origin()
		cx, _ := view.Cursor()

		if err := view.SetCursor(cx, 0); err != nil && oy > 0 {
			if err := view.SetOrigin(ox, 0); err != nil {
				return err
			}
		}
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

func (ow *onepassWidget) showItems(ui *gocui.Gui, view *gocui.View) error {
	if ow.currIdx != -1 {
		if err := ow.itemsWidget.Layout(ui); err != nil {
			return err
		}

		var items []*domain.SimpleItem

		if len(ow.categories) == 1 {
			items = ow.guiControl.GetItems(nil, true)
		} else {
			if ow.currIdx == len(ow.categories)-1 {
				items = ow.guiControl.GetItems(nil, true)
			} else if ow.currIdx == 0 {
				items = ow.guiControl.GetItems(nil, false)
			} else {
				items = ow.guiControl.GetItems(ow.categories[ow.currIdx], false)
			}
		}

		ow.itemsWidget.items = items

		if err := ow.itemsWidget.update(ui); err != nil {
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

		if len(ow.categories) > 0 {
			ow.currIdx = 0
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

	if err := ow.itemsWidget.Keybindings(ui); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyCtrlL, gocui.ModNone, ow.lock); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, 'j', gocui.ModNone, ow.cursorDown); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyArrowDown, gocui.ModNone, ow.cursorDown); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, 'k', gocui.ModNone, ow.cursorUp); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyArrowUp, gocui.ModNone, ow.cursorUp); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, 'l', gocui.ModNone, ow.showItems); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyArrowRight, gocui.ModNone, ow.showItems); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyEnter, gocui.ModNone, ow.showItems); err != nil {
		return err
	}

	return nil
}

func (ow *onepassWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(ow.name, 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		ui.Highlight = false

		view.Title = ow.title
		view.Highlight = true
		view.SelBgColor = gocui.ColorBlue

		if _, err := ui.SetCurrentView(ow.name); err != nil {
			return err
		}

		if err := ow.update(ui); err != nil {
			return err
		}

		ow.helpWidget.help = onepassHelp

		if err := ow.helpWidget.update(ui); err != nil {
			return err
		}
	}

	return nil
}
