// Main 1Pass application widget.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
)

type onepassWidget struct {
	name  string
	title string
}

func newOnepassWidget() *onepassWidget {
	return &onepassWidget{
		title: "1Pass",
		name:  "1pass",
	}
}

func (ow *onepassWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(ow.name, 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = ow.title
	}

	return nil
}
