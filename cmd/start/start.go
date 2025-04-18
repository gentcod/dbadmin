package start

import (
	"fmt"
	"os"

	"github.com/gentcod/db-admin/internal/pkg/utils"

	"github.com/spf13/cobra"
)

var dbDriver string
var dbConnString string

var startCmd = &cobra.Command{
	Use:   "dbadmin",
	Short: "DbAdmin is a CLI database privilege manager",
	Long:  `A CLI for managing your database privilege or access right without the hassle of running SQL queries.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nWelcome to DbAdmin CLI. Use h or help for help.")
	},
}

func init() {
	startCmd.PersistentFlags().StringVarP(&dbDriver, "dbdriver", "d", "", "Database engine to use (e.g. mysql, postgres)")
	startCmd.PersistentFlags().StringVarP(&dbConnString, "dbstring", "D", "", "Database connection string")

	availableCmds := []*cobra.Command{
		versionCmd,
	}
	startCmd.AddCommand(availableCmds...)
}

func Execute() (string, string) {
	startCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		dbDriver, _ = cmd.Flags().GetString("dbdriver")
		dbConnString, _ = cmd.Flags().GetString("dbstring")

		if !cmd.Flags().Changed("dbdriver") || !cmd.Flags().Changed("dbstring") {
			return fmt.Errorf("both --dbdriver and --dbstring flags are required")
		}

		if !utils.ValidateDBDriver(dbDriver) {
			return fmt.Errorf("unsupported database driver: %s", dbDriver)
		}

		return nil
	}

	if err := startCmd.Execute(); err != nil {
		os.Exit(1)
	}

	return dbDriver, dbConnString
}
