package start

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DbAdmin",
	Long:  `All software has versions. This is DbAdmin's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DbAdmin Database Privilege Manager v0.9 -- HEAD")
	},
}
