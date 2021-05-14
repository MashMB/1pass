// Widget for vault items.
//
// @author TSS

package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/in"
)

const (
	itemsHelp string = `Navigate up/down: k/j
Go back: Q
Reveal item: TAB`
)

type itemsWidget struct {
	currIdx       int
	name          string
	parent        string
	title         string
	lockHandler   func(ui *gocui.Gui, view *gocui.View) error
	detailsWidget *detailsWidget
	helpWidget    *helpWidget
	items         []*domain.SimpleItem
	guiControl    in.GuiControl
}

func newItemsWidget(parent string, helpWidget *helpWidget, lockHandler func(ui *gocui.Gui, view *gocui.View) error, guiControl in.GuiControl) *itemsWidget {
	widget := &itemsWidget{
		currIdx:     -1,
		name:        "itemsWidget",
		title:       "Items",
		lockHandler: lockHandler,
		helpWidget:  helpWidget,
		items:       make([]*domain.SimpleItem, 0),
		parent:      parent,
		guiControl:  guiControl,
	}

	widget.detailsWidget = newDetailsWidget(widget.name, widget.helpWidget, widget.lock)

	return widget
}

func (iw *itemsWidget) cursorDown(ui *gocui.Gui, view *gocui.View) error {
	if view != nil && iw.currIdx != -1 && iw.currIdx < len(iw.items)-1 {
		cx, cy := view.Cursor()

		if err := view.SetCursor(cx, cy+1); err != nil {
			ox, oy := view.Origin()

			if err := view.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		iw.currIdx++

		if err := iw.showOverview(ui); err != nil {
			return err
		}
	}

	return nil
}

func (iw *itemsWidget) cursorUp(ui *gocui.Gui, view *gocui.View) error {
	if view != nil && iw.currIdx > 0 {
		ox, oy := view.Origin()
		cx, cy := view.Cursor()

		if err := view.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := view.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}

		iw.currIdx--

		if err := iw.showOverview(ui); err != nil {
			return err
		}
	}

	return nil
}

func (iw *itemsWidget) goBack(ui *gocui.Gui, view *gocui.View) error {
	if err := ui.DeleteView(iw.detailsWidget.name); err != nil {
		return err
	}

	if err := ui.DeleteView(iw.name); err != nil {
		return err
	}

	if _, err := ui.SetCurrentView(iw.parent); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) lock(ui *gocui.Gui, view *gocui.View) error {
	if err := iw.goBack(ui, view); err != nil {
		return err
	}

	parentView, err := ui.View(iw.parent)

	if err != nil {
		return err
	}

	if err := iw.lockHandler(ui, parentView); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) resetCursor(view *gocui.View) error {
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

func (iw *itemsWidget) showOverview(ui *gocui.Gui) error {
	item := iw.guiControl.GetItem(iw.items[iw.currIdx])
	iw.detailsWidget.item = item

	if err := iw.detailsWidget.update(true, ui); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) toggleDetails(ui *gocui.Gui, view *gocui.View) error {
	item := iw.guiControl.GetItem(iw.items[iw.currIdx])
	iw.detailsWidget.item = item

	if err := iw.detailsWidget.update(false, ui); err != nil {
		return err
	}

	if _, err := ui.SetCurrentView(iw.detailsWidget.name); err != nil {
		return err
	}

	iw.helpWidget.help = detailsHelp

	if err := iw.helpWidget.update(ui); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) update(ui *gocui.Gui) error {
	view, err := ui.View(iw.name)

	if err != nil {
		return err
	}

	iw.currIdx = -1

	if err := iw.resetCursor(view); err != nil {
		return err
	}

	view.Clear()

	if len(iw.items) > 0 {
		for _, item := range iw.items {
			title := item.Title

			if title == "" {
				title = "<NO TITLE>"
			}

			position := fmt.Sprintf("%v\n", title)
			fmt.Fprint(view, position)
		}

		iw.currIdx = 0

		if err := iw.showOverview(ui); err != nil {
			return err
		}
	}

	return nil
}

func (iw *itemsWidget) Keybindings(ui *gocui.Gui) error {
	if err := iw.detailsWidget.Keybindings(ui); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, gocui.KeyCtrlL, gocui.ModNone, iw.lock); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, 'h', gocui.ModNone, iw.goBack); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, gocui.KeyArrowLeft, gocui.ModNone, iw.goBack); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, 'q', gocui.ModNone, iw.goBack); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, 'j', gocui.ModNone, iw.cursorDown); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, gocui.KeyArrowDown, gocui.ModNone, iw.cursorDown); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, 'k', gocui.ModNone, iw.cursorUp); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, gocui.KeyArrowUp, gocui.ModNone, iw.cursorUp); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, gocui.KeyTab, gocui.ModNone, iw.toggleDetails); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(iw.name, 1, 1, int(0.5*float32(maxX-2)), maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		ui.Highlight = true
		ui.SelFgColor = gocui.ColorBlue

		view.Title = iw.title
		view.Highlight = true
		view.SelFgColor = gocui.ColorDefault
		view.SelBgColor = gocui.ColorBlue

		iw.update(ui)

		if err := iw.detailsWidget.Layout(ui); err != nil {
			return err
		}

		if _, err := ui.SetCurrentView(iw.name); err != nil {
			return err
		}

		iw.helpWidget.help = itemsHelp

		if err := iw.helpWidget.update(ui); err != nil {
			return err
		}
	}

	return nil
}
