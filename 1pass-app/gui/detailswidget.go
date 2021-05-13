// Widget use to display single item details.
//
// @author TSS

package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type detailsWidget struct {
	name  string
	title string
	item  *domain.Item
}

func newDetailsWidget() *detailsWidget {
	return &detailsWidget{
		name:  "detailsWidget",
		title: "Details",
	}
}

func (dw *detailsWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(dw.name, int(0.5*float32(maxX-2)+1), 1, maxX-2, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = dw.title
	}

	return nil
}
