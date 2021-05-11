// Main 1Pass application widget.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
)

type onepassWidget struct {
	name       string
	title      string
	passPrompt *passwordPrompt
}

func newOnepassWidget() *onepassWidget {
	widget := &onepassWidget{
		title: "1Pass",
		name:  "1pass",
	}

	widget.passPrompt = newPasswordPrompt(widget.unlock)

	return widget
}

func (ow *onepassWidget) promptForPassword(ui *gocui.Gui) error {
	if err := ow.passPrompt.Layout(ui); err != nil {
		return err
	}

	return nil
}

func (ow *onepassWidget) unlock(ui *gocui.Gui, view *gocui.View) error {
	if err := ui.DeleteView(ow.passPrompt.name); err != nil {
		return err
	}

	if _, err := ui.SetCurrentView(ow.name); err != nil {
		return err
	}

	ow.update(ui)

	return nil
}

func (ow *onepassWidget) update(ui *gocui.Gui) error {
	if err := ow.promptForPassword(ui); err != nil {
		return err
	}

	return nil
}

func (ow *onepassWidget) quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func (ow *onepassWidget) Keybindings(ui *gocui.Gui) error {
	if err := ow.passPrompt.Keybindings(ui); err != nil {
		return err
	}

	if err := ui.SetKeybinding(ow.name, gocui.KeyCtrlQ, gocui.ModNone, ow.quit); err != nil {
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

		ow.update(ui)
	}

	return nil
}
