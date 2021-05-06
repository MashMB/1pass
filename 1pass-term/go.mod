// Dependencies for 1Pass terminal module.
//
// @author TSS

module github.com/mashmb/1pass/1pass-term

go 1.15

require (
	github.com/jedib0t/go-pretty/v6 v6.2.1
	github.com/mashmb/1pass/1pass-core v1.1.0
	golang.org/x/term v0.0.0-20210422114643-f5beecf764ed
)

replace github.com/mashmb/1pass/1pass-core => ../1pass-core
