// Dependencies for 1Pass terminal module.
//
// @author TSS

module github.com/mashmb/1pass/1pass-term

go 1.15

require (
	github.com/mashmb/1pass/1pass-core v0.0.0
	golang.org/x/term v0.0.0-20210422114643-f5beecf764ed
)

replace github.com/mashmb/1pass/1pass-core => ../1pass-core
