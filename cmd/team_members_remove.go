package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	pluralize "github.com/alejandrojnm/go-pluralize"
	"github.com/civo/civogo"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var teamMemberRemoveID string

var teamID2 string

var teamMembersList []utility.ObjecteList
var teamMembersRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"delete", "rm"},
	Short:   "Remove a team-member from your team",
	Args:    cobra.MinimumNArgs(2),
	Example: "civo team-members remove TEAM_ID TEAM_MEMBER_ID",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		team, err := client.FindTeam(args[0])
		if err != nil {
			if errors.Is(err, civogo.ZeroMatchesError) {
				utility.Error("sorry there is no %s team in your account", utility.Red(args[0]))
				os.Exit(1)
			}
			if errors.Is(err, civogo.MultipleMatchesError) {
				utility.Error("sorry we found more than one team with that name in your account")
				os.Exit(1)
			}
		}

		if len(args) == 2 {
			teamMember, err := client.FindTeamMember(team.ID, args[1])
			if err != nil {
				if errors.Is(err, civogo.ZeroMatchesError) {
					utility.Error("sorry there is no %s team-member in your account", utility.Red(args[1]))
					os.Exit(1)
				}
				if errors.Is(err, civogo.MultipleMatchesError) {
					utility.Error("sorry we found more than one team-member in your account")
					os.Exit(1)
				}
			}
			teamMembersList = append(teamMembersList, utility.ObjecteList{ID: teamMember.ID, Name: teamMember.ID})
		} else {
			for _, v := range args[1:] {
				teamMember, err := client.FindTeamMember(team.ID, v)
				if err == nil {
					teamMembersList = append(teamMembersList, utility.ObjecteList{ID: teamMember.ID, Name: teamMember.ID})
				}
			}
		}

		teamMembersNameList := []string{}
		for _, v := range teamMembersList {
			teamMembersNameList = append(teamMembersNameList, v.Name)
		}

		if utility.UserConfirmedDeletion(fmt.Sprintf("team %s", pluralize.Pluralize(len(teamMembersList), "member")), defaultYes, strings.Join(teamMembersNameList, ", ")) {

			for _, v := range teamMembersList {
				_, err = client.RemoveTeamMember(team.ID, v.ID)
				if err != nil {
					utility.Error("error deleting the team-member: %s", err)
					os.Exit(1)
				}
			}

			ow := utility.NewOutputWriter()

			for _, v := range teamMembersList {
				ow.StartLine()
				ow.AppendDataWithLabel("id", v.ID, "ID")
				ow.AppendDataWithLabel("label", v.Name, "Label")
			}

			switch outputFormat {
			case "json":
				if len(teamMembersList) == 1 {
					ow.WriteSingleObjectJSON(prettySet)
				} else {
					ow.WriteMultipleObjectsJSON(prettySet)
				}
			case "custom":
				ow.WriteCustomOutput(outputFields)
			default:
				fmt.Printf("The team-member %s(%s) has been deleted\n", pluralize.Pluralize(len(teamMembersList), ""), strings.Join(teamMembersNameList, ", "))
			}
		} else {
			fmt.Println("Operation aborted.")
		}

	},
}
