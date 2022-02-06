package cmd

import (
	"fmt"
	"os"

	"github.com/civo/civogo"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var newRolePermissions string
var accountToBeAssigned string

var rolePermission string
var defaultPermission *civogo.Permission

var rolesCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Example: "civo roles create NAME [flags]",
	Short:   "Create a new role",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		role, err := client.CreateRole(args[0], newRolePermissions, accountToBeAssigned)
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriterWithMap(map[string]string{"name": role.Name, "permissions": role.Permissions})

		switch outputFormat {
		case "json":
			ow.WriteSingleObjectJSON(prettySet)
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			fmt.Printf("Created a role called %s with role permissions %s\n", utility.Green(role.Name), utility.Green(role.Permissions))
		}
	},
}
