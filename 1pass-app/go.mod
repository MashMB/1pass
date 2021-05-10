// Dependencies for 1Pass application.
//
// @author TSS

module github.com/mashmb/1pass/1pass-app

go 1.15

require (
	github.com/jroimartin/gocui v0.4.0
	github.com/mashmb/1pass/1pass-core v1.2.0
	github.com/mashmb/1pass/1pass-parse v1.2.0
	github.com/mashmb/1pass/1pass-term v1.2.0
	github.com/mashmb/1pass/1pass-up v1.1.0
	github.com/nsf/termbox-go v1.1.1 // indirect
	github.com/spf13/cobra v1.1.3
)

replace (
	github.com/mashmb/1pass/1pass-core => ../1pass-core
	github.com/mashmb/1pass/1pass-parse => ../1pass-parse
	github.com/mashmb/1pass/1pass-term => ../1pass-term
	github.com/mashmb/1pass/1pass-up => ../1pass-up
)
