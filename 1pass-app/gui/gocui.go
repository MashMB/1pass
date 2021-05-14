// Main entry point for 1Pass terminal GUI.
//
// @author TSS

package gui

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/mashmb/1pass/1pass-core/port/in"
)

type GocuiGui struct {
	guiControl in.GuiControl
}

func NewGocuiGui(guiControl in.GuiControl) *GocuiGui {
	return &GocuiGui{
		guiControl: guiControl,
	}
}

func (gui *GocuiGui) quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func (gui *GocuiGui) Run(vaultPath string) {
	vault, err := gui.guiControl.ValidateVault(vaultPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ui, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Fatalln(err)
	}

	defer ui.Close()
	version := newVerWidget()
	help := newHelpWidget()
	onepass := newOnepassWidget(help, vault, gui.guiControl)
	ui.SetManager(version, help, onepass)

	if err := onepass.Keybindings(ui); err != nil {
		log.Fatalln(err)
	}

	if err := ui.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, gui.quit); err != nil {
		log.Fatalln(err)
	}

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}
