// Application version widget.
//
// @author TSS

package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type verWidget struct {
	name    string
	title   string
	version string
}

func newVerWidget(version string) *verWidget {
	return &verWidget{
		name:    "verWidget",
		title:   "Version",
		version: version,
	}
}

func (vw *verWidget) Layout(ui *gocui.Gui) error {
	maxX, _ := ui.Size()

	if view, err := ui.SetView(vw.name, 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = vw.title
		version := fmt.Sprintf("1Pass v%v ", vw.version)
		fmt.Fprint(view, version)
	}

	return nil
}
