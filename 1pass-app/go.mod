// Dependencies for 1Pass application.
//
// @author TSS

module github.com/mashmb/1pass/1pass-app

go 1.15

require (
	github.com/mashmb/1pass/1pass-core v0.0.0
	github.com/mashmb/1pass/1pass-parse v0.0.0
	github.com/mashmb/1pass/1pass-term v0.0.0
	github.com/spf13/cobra v1.1.3
)

replace (
	github.com/mashmb/1pass/1pass-core => ../1pass-core
	github.com/mashmb/1pass/1pass-parse => ../1pass-parse
	github.com/mashmb/1pass/1pass-term => ../1pass-term
)
