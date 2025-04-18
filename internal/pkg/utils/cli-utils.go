package utils

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/gentcod/db-admin/internal/config"
	"github.com/spf13/cobra"
)

var dbDrivers = []string{"postgres", "mysql"}

// InputHandler specifies how the handler logic for an interactive CLI is managed.
type InputHandler func(string) (error, bool)

// IsValidSubcommand checks the validity of subcommand(s).
func IsValidSubcommand(available []*cobra.Command, sub string) bool {
	for _, s := range available {
		if sub == s.CalledAs() {
			return true
		}
	}
	return false
}

// ValidateDBDriver checks if the database driver choice is valid.
func ValidateDBDriver(driver string) bool {
	return slices.Contains(dbDrivers, driver)
}

// RunInteractive helps to run the CLI in an interactive manner
func RunInteractive(intro string, handler InputHandler) error {
	reader := bufio.NewReader(os.Stdin)
LOOP:
	for {
		fmt.Println(intro)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			continue LOOP
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}

		err, iter := handler(input)
		if iter {
			config.Cfg.Logger.Error(err)
			continue LOOP
		}

		return err
	}
}
