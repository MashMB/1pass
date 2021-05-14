// GUI password prompt.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
)

type passwordPrompt struct {
	name    string
	title   string
	handler func(ui *gocui.Gui, view *gocui.View) error
}

func newPasswordPrompt(handler func(ui *gocui.Gui, view *gocui.View) error) *passwordPrompt {
	return &passwordPrompt{
		name:    "passPrompt",
		title:   "Password",
		handler: handler,
	}
}

func (pp *passwordPrompt) Keybindings(ui *gocui.Gui) error {
	if err := ui.SetKeybinding(pp.name, gocui.KeyEnter, gocui.ModNone, pp.handler); err != nil {
		return err
	}

	return nil
}

func (pp *passwordPrompt) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(pp.name, maxX/2-30, maxY/2-1, maxX/2+30, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		ui.Highlight = false

		view.Title = pp.title
		view.Editable = true
		view.Mask = '*'

		if _, err := ui.SetCurrentView(pp.name); err != nil {
			return err
		}
	}

	return nil
}
