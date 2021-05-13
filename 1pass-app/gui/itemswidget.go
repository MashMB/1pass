// Widget for vault items.
//
// @author TSS

package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type itemsWidget struct {
	currIdx     int
	name        string
	parent      string
	title       string
	lockHandler func(ui *gocui.Gui, view *gocui.View) error
	items       []*domain.SimpleItem
}

func newItemsWidget(parent string, lockHandler func(ui *gocui.Gui, view *gocui.View) error) *itemsWidget {
	return &itemsWidget{
		currIdx:     -1,
		name:        "itemsWidget",
		title:       "Items",
		lockHandler: lockHandler,
		items:       make([]*domain.SimpleItem, 0),
		parent:      parent,
	}
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
	}

	return nil
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

func (iw *itemsWidget) update(ui *gocui.Gui) error {
	view, err := ui.View(iw.name)

	if err != nil {
		return err
	}

	iw.currIdx = -1
	iw.resetCursor(view)
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
	}

	return nil
}

func (iw *itemsWidget) Keybindings(ui *gocui.Gui) error {
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

		iw.update(ui)

		if _, err := ui.SetCurrentView(iw.name); err != nil {
			return err
		}
	}

	return nil
}
