// Widget for vault items.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type itemsWidget struct {
	name  string
	title string
	items []*domain.SimpleItem
}

func newItemsWidget() *itemsWidget {
	return &itemsWidget{
		name:  "itemsWidget",
		title: "Items",
		items: make([]*domain.SimpleItem, 0),
	}
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
