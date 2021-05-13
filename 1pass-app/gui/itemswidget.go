// Widget for vault items.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type itemsWidget struct {
	name   string
	parent string
	title  string
	items  []*domain.SimpleItem
}

func newItemsWidget(parent string) *itemsWidget {
	return &itemsWidget{
		name:   "itemsWidget",
		title:  "Items",
		items:  make([]*domain.SimpleItem, 0),
		parent: parent,
	}
}

func (iw *itemsWidget) goBack(ui *gocui.Gui, view *gocui.View) error {
	if err := ui.DeleteView(iw.name); err != nil {
		return err
	}

	if _, err := ui.SetCurrentView(iw.parent); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) Keybindings(ui *gocui.Gui) error {
	if err := ui.SetKeybinding(iw.name, 'h', gocui.ModNone, iw.goBack); err != nil {
		return err
	}

	if err := ui.SetKeybinding(iw.name, gocui.KeyBackspace, gocui.ModNone, iw.goBack); err != nil {
		return err
	}

	return nil
}

func (iw *itemsWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(iw.name, 1, 1, int(0.5*float32(maxX-2)), maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = iw.title
		view.Highlight = true
		view.SelBgColor = gocui.ColorBlue

		if _, err := ui.SetCurrentView(iw.name); err != nil {
			return err
		}
	}

	return nil
}
