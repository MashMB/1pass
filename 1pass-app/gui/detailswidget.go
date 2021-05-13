// Widget use to display single item details.
//
// @author TSS

package gui

import (
	"fmt"
	"time"

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

func (dw *detailsWidget) update(overview bool, ui *gocui.Gui) error {
	view, err := ui.View(dw.name)

	if err != nil {
		return err
	}

	view.Clear()

	if dw.item != nil {
		fmt.Fprint(view, fmt.Sprintf("%v\n", dw.item.Category.GetName()))
		fmt.Fprint(view, "------------------------------\n")
		fmt.Fprint(view, fmt.Sprintf("%v\n\n", dw.item.Title))
		updated := time.Unix(dw.item.Updated, 0).Format("2006-01-02 15:04:05")
		created := time.Unix(dw.item.Created, 0).Format("2006-01-02 15:04:05")
		fmt.Fprint(view, fmt.Sprintf("Updated: %v\nCreated: %v\nTrashed: %v\n", updated, created, dw.item.Trashed))
	}

	return nil
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
