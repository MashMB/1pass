// GUI error dialog.
//
// @author TSS

package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type errorDialog struct {
	name         string
	title        string
	closeHandler func(ui *gocui.Gui, view *gocui.View) error
	err          error
}

func newErrorDialog(closeHandler func(ui *gocui.Gui, view *gocui.View) error) *errorDialog {
	return &errorDialog{
		name:         "errDialog",
		title:        "Error",
		closeHandler: closeHandler,
	}
}

func (ed *errorDialog) Keybindings(ui *gocui.Gui) error {
	if err := ui.SetKeybinding(ed.name, gocui.KeyEnter, gocui.ModNone, ed.closeHandler); err != nil {
		return err
	}

	return nil
}

func (ed *errorDialog) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(ed.name, maxX/2-30, maxY/2-1, maxX/2+30, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		ui.Highlight = true
		ui.SelFgColor = gocui.ColorRed

		view.Title = ed.title
		view.Highlight = true
		view.SelFgColor = gocui.ColorRed

		if _, err := ui.SetCurrentView(ed.name); err != nil {
			return err
		}

		fmt.Fprint(view, ed.err)
	}

	return nil
}
