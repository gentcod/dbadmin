package add

import (
	"fmt"

	"github.com/gentcod/db-admin/internal/config"
	"github.com/gentcod/db-admin/internal/db"
	"github.com/gentcod/db-admin/internal/pkg/utils"
	"github.com/spf13/cobra"
)

var ptype string
var username string
var password string

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new privilege to a database user",
	Long: `The 'add' command creates a new user in the target PostgreSQL database with the appropriate access privileges.

		This command supports both manual and interactive modes for adding users:
		- You can supply the username and password via flags (--user and --password).
		- If the password is not provided, the CLI will prompt you to either generate a secure password or manually enter one.
		- You may optionally specify a privilege type using the --type flag to assign predefined roles or permission sets.
		`,
	Example: `
		add --user alice --password strongpassword123
			Creates a new database user 'alice' with the specified password.

		add --user bob --type readonly
			Prompts to either generate or manually input a password for user 'bob',
			and then assigns the 'readonly' access privileges after creation.

		Notes:
		- The --user flag is required.
		- If --password is not specified, the CLI will ask whether to autogenerate one.
		- If both --user and --password are provided, --type is ignored.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Cfg.CreateUserWithPassword(cmd.Context(), db.CreateUserWithPasswordParams{
			Username: username,
			Password: password,
		})
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		fmt.Printf("\nNew database user has been added successfully. %s with password: %s \n", username, password)
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("user") && !cmd.Flags().Changed("password") {
			if !cmd.Flags().Changed("type") {
				return cmd.MarkFlagRequired("type")
			}

			inputHandler := func(input string) (error, bool) {
				switch input {
				case "yes", "YES", "Yes", "Y", "y":
					password = utils.GenerateRandomPassword()
					fmt.Printf("\nAdding %s with password: %s... Make sure to write down generated password\n", username, password)
					return nil, false
				case "no", "NO", "No", "N", "n":
					return utils.RunInteractive(
						"\nEnter password: ",
						func(input string) (error, bool) {
							switch input {
							case "":
								return fmt.Errorf("password cannot be empty"), true
							default:
								if len(input) < 8 {
									return fmt.Errorf("minimum of 8 characters allowed"), true
								}
								password = input
								return nil, false
							}
						},
					), false
				default:
					return fmt.Errorf("invalid input: password is required"), true
				}
			}

			return utils.RunInteractive("\nDo you want you autogenerate password, Yes/No (Y/N)? ", inputHandler)
		}

		if cmd.Flags().Changed("user") && cmd.Flags().Changed("password") {
			if cmd.Flags().Changed("type") {
				fmt.Println("role type not needed in when user and password are parsed")
			}

			username, _ = cmd.Flags().GetString("user")
			password, _ = cmd.Flags().GetString("password")
			return nil
		}

		return nil
	},
}

func init() {
	AddCmd.Flags().StringVarP(&ptype, "type", "t", "", "Privilege type you'd like to assign (e.g. user, role) [required]")
	AddCmd.Flags().StringVarP(&username, "user", "u", "", "Username to add [required]")
	AddCmd.Flags().StringVarP(&password, "password", "p", "", "Password is the associated auth for a Username [required if type is user]")
}
