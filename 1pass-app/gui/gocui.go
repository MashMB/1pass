// Main entry point for 1Pass terminal GUI.
//
// @author TSS

package gui

import (
	"log"

	"github.com/jroimartin/gocui"
)

type GocuiGui struct{}

func NewGocuiGui() *GocuiGui {
	return &GocuiGui{}
}

func (gui *GocuiGui) Run() {
	ui, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Fatalln(err)
	}

	defer ui.Close()
	onepass := newOnepassWidget()
	ui.SetManager(onepass)

	if err := onepass.Keybindings(ui); err != nil {
		log.Fatalln(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}
