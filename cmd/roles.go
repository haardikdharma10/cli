package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rolesCmd = &cobra.Command{
	Use:     "roles",
	Aliases: []string{"role"},
	Short:   "\nManage roles in your team",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	rootCmd.AddCommand(rolesCmd)
	rolesCmd.AddCommand(rolesListCmd)
	rolesCmd.AddCommand(rolesCreateCmd)

	rolesCreateCmd.Flags().StringVarP(&newRolePermissions, "permissions", "p", "", "Permissions for the new role")
	rolesCreateCmd.Flags().StringVarP(&accountToBeAssigned, "account-id", "", "", "Account id for new role")

}
