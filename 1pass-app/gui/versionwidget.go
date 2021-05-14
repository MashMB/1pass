// Application version widget.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
)

type verWidget struct {
	name  string
	title string
}

func newVerWidget() *verWidget {
	return &verWidget{
		name:  "verWidget",
		title: "Version",
	}
}

func (vw *verWidget) Layout(ui *gocui.Gui) error {
	maxX, _ := ui.Size()

	if view, err := ui.SetView(vw.name, 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = vw.title
	}

	return nil
}
