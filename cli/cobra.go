// Definition of CLI commands for 1Pass.
//
// @author TSS

package cli

import (
	"fmt"
	"log"

	"github.com/mashmb/1pass/port/in"
	"github.com/spf13/cobra"
)

type cobraCli struct {
	cliControl in.CliControl
}

func NewCobraCli(cliControl in.CliControl) *cobraCli {
	return &cobraCli{
		cliControl: cliControl,
	}
}

func (cli *cobraCli) init() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "1pass",
		Short: "1Password Linux CLI explorer",
		Long: `Fast and Linux user friendly application used to explore 1Password OPVault format. Check your credentials 
efficiently in terminal.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World")
		},
	}

	return rootCmd
}

func (cli *cobraCli) Run() {
	rootCmd := cli.init()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
