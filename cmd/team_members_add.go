package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var teamMembersAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"new", "create"},
	Short:   "Add a new team member",
	Example: "civo team-members add TEAM_ID USER_ID PERMISSIONS ROLE_ID",
	Args:    cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		teamMemberTeamID := args[0]
		teamMemberUserID := args[1]
		teamMemberPermissions := args[2]
		teamMemberRoles := args[3]

		teamMembers, err := client.AddTeamMember(teamMemberTeamID, teamMemberUserID, teamMemberPermissions, teamMemberRoles)
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()
		for _, teamMember := range teamMembers {
			ow.StartLine()

			ow.AppendDataWithLabel("team_id", teamMember.TeamID, "TeamID")
			ow.AppendDataWithLabel("user_id", teamMember.UserID, "UserID")
			ow.AppendDataWithLabel("permissions", teamMember.Permissions, "Permissions")
			ow.AppendDataWithLabel("roles", teamMember.Roles, "Roles")
		}

		switch outputFormat {
		case "json":
			ow.WriteSingleObjectJSON(prettySet)
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			fmt.Printf("Added team member with User ID %s, permissions %s and roles %s", utility.Green(string(teamMemberUserID)), utility.Green(teamMemberPermissions), utility.Green(teamMemberRoles))
		}
	},
}
