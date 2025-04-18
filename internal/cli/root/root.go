package root

import (
	"fmt"
	"os"

	"github.com/gentcod/db-admin/internal/cli/root/add"
	"github.com/spf13/cobra"
	// "github.com/spf13/pflag"
)

var RootCmd = &cobra.Command{
	Use: "",
	// SilenceUsage: true,
	Short: "DbAdmin is a CLI database privilege manager",
	Long:  `A CLI for managing your database privilege or access right without the hassle of running SQL queries.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to DbAdmin CLI. Use h or help for help.")
		fmt.Println()
	},
}

func init() {
	availableCmds := []*cobra.Command{
		add.AddCmd,
	}

	RootCmd.SetHelpCommand(HelpCmd)
	RootCmd.AddCommand(availableCmds...)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
