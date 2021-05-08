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
	name       string
	trashed    bool
	vault      string
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
efficiently in terminal. Run '1pass --help' for more informations.`,
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.CheckForUpdate()
			fmt.Println(fmt.Sprintf("%v\n%v", cmd.Short, cmd.Long))
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update 1pass application. Run '1pass update --help' for more detailed informations.",
		Long: `Update 1pass application. It is recommended to update application with root permissions. 

Whole update process:
1. Check if there is new release on GitHub.
2. Download newer release to temporary directory (with checksums).
3. Extract downloaded archive.
4. Compare checksums.
5. Replace running binary.
6. Clean cache (temporary files and directories).`,
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.Update()
		},
	}

	configureCmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure 1pass application",
		Long: `Configure 1pass application. Configuration process is interactive - answer the questions. Available settings: 
1. Configure default OPVault path.
2. Configure update notifications.`,
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.Configure()
		},
	}

	categoriesCmd := &cobra.Command{
		Use:   "categories",
		Short: "Get list of available item categories",
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.GetCategories()
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Get list of items stored in 1Passowrd OPVault format",
		Long: `Get list of items stored in 1Passowrd OPVault format. Items will be displayd in form of table. UID value is 
required to get item overview or details. If default OPVault is not configured,  [-v, --vault] flag is needed. Run 
'1pass list --help' for more info.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.GetItems(cli.vault, cli.category, cli.name, cli.trashed)
		},
	}

	listCmd.Flags().StringVarP(&cli.vault, "vault", "v", "", "OPVault path")
	listCmd.Flags().StringVarP(&cli.category, "category", "c", "", "filtering over item category")
	listCmd.Flags().StringVarP(&cli.name, "name", "n", "", "filtering over item name (title)")
	listCmd.Flags().BoolVarP(&cli.trashed, "trashed", "t", false, "work on trashed items")

	overviewCmd := &cobra.Command{
		Use:   "overview [UID]",
		Short: "Overview single item stored in 1Password OPVault format",
		Long: `Overview single item stored in 1Password OPVault format. Overview has no sensitive data like 
passwords. If default OPVault is not configured,  [-v, --vault] flag is needed. Run '1pass overview --help' for more info.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.GetItemOverview(cli.vault, args[0], cli.trashed)
		},
	}

	overviewCmd.Flags().StringVarP(&cli.vault, "vault", "v", "", "OPVault path")
	overviewCmd.Flags().BoolVarP(&cli.trashed, "trashed", "t", false, "search in trashed items")

	detailsCmd := &cobra.Command{
		Use:   "details [UID]",
		Short: "Details of single item stored in 1Password OPVault format",
		Long: `Details of single item stored in 1Password OPVault format. Details contains sensitive data 
like passwords. If default OPVault is not configured,  [-v, --vault] flag is needed. Run '1pass details --help' for more info.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.cliControl.GetItemDetails(cli.vault, args[0], cli.trashed)
		},
	}

	detailsCmd.Flags().StringVarP(&cli.vault, "vault", "v", "", "OPVault path")
	detailsCmd.Flags().BoolVarP(&cli.trashed, "trashed", "t", false, "search in trashed items")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Check application version",
		Run: func(cmd *cobra.Command, args []string) {
			ver := fmt.Sprintf("=== %v ===", cli.version)
			fmt.Println(ver)
		},
	}

	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(configureCmd)
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
