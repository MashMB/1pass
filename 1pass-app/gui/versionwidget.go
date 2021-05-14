// Application version widget.
//
// @author TSS

package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/core/domain"
	"github.com/mashmb/1pass/1pass-core/port/in"
)

type verWidget struct {
	name       string
	title      string
	version    string
	guiControl in.GuiControl
}

func newVerWidget(version string, guiControl in.GuiControl) *verWidget {
	return &verWidget{
		name:       "verWidget",
		title:      "Version",
		version:    version,
		guiControl: guiControl,
	}
}

func (vw *verWidget) update(ui *gocui.Gui) {
	view, _ := ui.View(vw.name)
	info, err := vw.guiControl.CheckForUpdate()

	if err != nil {
		if err == domain.ErrNoUpdate {
			fmt.Fprint(view, fmt.Sprintf("(%v)", err))
		} else {
			fmt.Fprint(view, "(update check failed - run '1pass update' for more info)")
		}
	} else {
		fmt.Fprint(view, fmt.Sprintf("(new version is available: v%v - run '1pass update' to upgrade)", info.Version))
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
		go vw.update(ui)
	}

	return nil
}
