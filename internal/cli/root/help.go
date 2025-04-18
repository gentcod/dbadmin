package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HelpCmd = &cobra.Command{
	Use:     "help",
	Aliases: []string{"h"},
	Hidden:  true,
	Short:   "Use this to get more information",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("\nDbAdmin help guide: ")
		fmt.Println()
		RootCmd.Help()
		return nil
	},
}
