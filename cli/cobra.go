// Definition of CLI commands for 1Pass.
//
// @author TSS

package cli

import (
	"fmt"

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
	}

	listCmd := &cobra.Command{
		Use:   "list [OPVault] [master_password]",
		Short: "Get list of items stored in 1Passowrd OPVault format (logins only)",
		Long: `Get list of items stored in 1Passowrd OPVault format (logins only). Items will be displayd in form of 
[<UID>] --- <title>. UID value is required for item overview and details identifiaction.`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.GetItems(args[0], args[1])
		},
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func (cli *cobraCli) Run() {
	rootCmd := cli.init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
