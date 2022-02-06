package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var teamMembersUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"edit", "modify"},
	Short:   "Update permissions for a team member inside your team.",
	Example: "civo team-members update TEAM_ID TEAM-MEMBER_ID PERMISSIONS ROLE_ID",
	Args:    cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		teamMemberTeamID := args[0]
		teamMemberID := args[1]
		teamMemberPermissions := args[2]
		teamMemberRoles := args[3]

		teamMembers, err := client.UpdateTeamMember(teamMemberTeamID, teamMemberID, teamMemberPermissions, teamMemberRoles)
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriterWithMap(map[string]string{"team_id": teamMembers.TeamID, "id": teamMembers.ID, "permissions": teamMembers.Permissions, "roles": teamMembers.Roles})

		switch outputFormat {
		case "json":
			ow.WriteSingleObjectJSON(prettySet)
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			fmt.Printf("Updated permissions of team-member with id %s to %s", utility.Green(string(teamMemberID)), utility.Green(teamMemberPermissions))
		}
	},
}
