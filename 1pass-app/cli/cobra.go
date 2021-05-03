// Definition of CLI commands for 1Pass.
//
// @author TSS

package cli

import (
	"fmt"

	"github.com/mashmb/1pass/1pass-core/port/in"
	"github.com/spf13/cobra"
)

type cobraCli struct {
	category   string
	trashed    bool
	version    string
	cliControl in.CliControl
}

func NewCobraCli(version string, cliControl in.CliControl) *cobraCli {
	return &cobraCli{
		version:    version,
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

	categoriesCmd := &cobra.Command{
		Use:   "categories",
		Short: "Get list of available item categories",
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.GetCategories()
		},
	}

	listCmd := &cobra.Command{
		Use:   "list [OPVault]",
		Short: "Get list of items stored in 1Passowrd OPVault format",
		Long: `Get list of items stored in 1Passowrd OPVault format. Items will be displayd in form of table. UID value is 
required to get item overview or details. If default OPVault was configured, [OPVault] argument is not required.`,
		Run: func(cmd *cobra.Command, args []string) {
			var vaultPath string

			if len(args) > 0 {
				vaultPath = args[0]
			}

			cli.cliControl.GetItems(vaultPath, cli.category, cli.trashed)
		},
	}

	listCmd.Flags().StringVarP(&cli.category, "category", "c", "", "filtering over item category")
	listCmd.Flags().BoolVarP(&cli.trashed, "trashed", "t", false, "work on trashed items")

	overviewCmd := &cobra.Command{
		Use:   "overview [UID] [OPVault]",
		Short: "Overview single item stored in 1Password OPVault format",
		Long: `Overview single item stored in 1Password OPVault format. Overview has no sensitive data like 
passwords. If default OPVault was configured, [OPVault] argument is not required.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var vaultPath string

			if len(args) > 1 {
				vaultPath = args[1]
			}

			cli.cliControl.GetItemOverview(vaultPath, args[0], cli.trashed)
		},
	}

	overviewCmd.Flags().BoolVarP(&cli.trashed, "trashed", "t", false, "search in trashed items")

	detailsCmd := &cobra.Command{
		Use:   "details [UID] [OPVault]",
		Short: "Details of single item stored in 1Password OPVault format",
		Long: `Details of single item stored in 1Password OPVault format. Details contains sensitive data 
like passwords. If default OPVault was configured, [OPVault] argument is not required.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var vaultPath string

			if len(args) > 1 {
				vaultPath = args[1]
			}

			cli.cliControl.GetItemDetails(vaultPath, args[0], cli.trashed)
		},
	}

	detailsCmd.Flags().BoolVarP(&cli.trashed, "trashed", "t", false, "search in trashed items")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Check application version",
		Run: func(cmd *cobra.Command, args []string) {
			ver := fmt.Sprintf("=== %v ===", cli.version)
			fmt.Println(ver)
		},
	}

	rootCmd.AddCommand(categoriesCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(overviewCmd)
	rootCmd.AddCommand(detailsCmd)
	rootCmd.AddCommand(versionCmd)

	return rootCmd
}

func (cli *cobraCli) Run() {
	rootCmd := cli.init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
