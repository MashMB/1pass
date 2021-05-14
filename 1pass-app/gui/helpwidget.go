// Contextual help widget.
//
// @author TSS

package gui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

const (
	mainHelp string = `Exit: CTRL+Q
Lock: CTRL+L`
)

type helpWidget struct {
	name  string
	title string
	help  string
}

func newHelpWidget() *helpWidget {
	return &helpWidget{
		name:  "help",
		title: "Help",
	}
}

func (hw *helpWidget) update(ui *gocui.Gui) error {
	view, err := ui.View(hw.name)

	if err != nil {
		return err
	}

	view.Clear()
	body := ""
	mainLines := strings.Split(mainHelp, "\n")

	for _, line := range mainLines {
		body = body + line + "\t\t\t\t"
	}

	if hw.help != "" {
		lines := strings.Split(hw.help, "\n")

		for idx, line := range lines {
			body = body + line

			if idx != len(lines)-1 {
				body = body + "\t\t\t\t"
			}
		}
	}

	fmt.Fprint(view, body)

	return nil
}

func (hw *helpWidget) Layout(ui *gocui.Gui) error {
	maxX, maxY := ui.Size()

	if view, err := ui.SetView(hw.name, 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		view.Title = hw.title

		hw.update(ui)
	}

	return nil
}
