// 1Pass application entry point.
//
// @author TSS

package main

import (
	"github.com/mashmb/1pass/cli"
)

func main() {
	cobraCli := cli.NewCobraCli()
	cobraCli.Run()
}
