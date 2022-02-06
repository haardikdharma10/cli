package cmd

import (
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var teamID string

var teamMembersListCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"all", "list"},
	Example: `civo team-members ls TEAM_ID`,
	Short:   "List all team members inside a team",
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
		}

		teamMembers, err := client.ListTeamMembers(args[0])
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()
		for _, teamMember := range teamMembers {
			ow.StartLine()

			ow.AppendDataWithLabel("team_id", teamMember.TeamID, "TeamID")
			ow.AppendDataWithLabel("user_id", teamMember.UserID, "UserID")
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
