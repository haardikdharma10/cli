package cmd

import (
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var rolesListCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"all", "list"},
	Example: `civo roles ls`,
	Short:   "List all roles",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
		}

		roles, err := client.ListRoles()
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()
		for _, role := range roles {
			ow.StartLine()

			ow.AppendDataWithLabel("id", role.ID, "ID")
			ow.AppendDataWithLabel("name", role.Name, "Name")
			ow.AppendDataWithLabel("permissions", role.Permissions, "Permissions")
		}
		switch outputFormat {
		case "json":
			ow.WriteMultipleObjectsJSON(prettySet)
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			ow.WriteTable()
		}
	},
}
