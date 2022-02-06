package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var teamMembersCmd = &cobra.Command{
	Use:   "team-members",
	Short: "Manage team members inside your team",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	rootCmd.AddCommand(teamMembersCmd)
	teamMembersCmd.AddCommand(teamMembersListCmd)
	teamMembersCmd.AddCommand(teamMembersAddCmd)
	teamMembersCmd.AddCommand(teamMembersUpdateCmd)
	teamMembersCmd.AddCommand(teamMembersRemoveCmd)

	//teamMembersListCmd.Flags().StringVarP(&teamID, "team-id", "", "", "Team ID")

	// teamMembersAddCmd.Flags().StringVarP(&teamMemberTeamID, "team-id", "", "", "Team ID")
	// teamMembersAddCmd.Flags().StringVarP(&teamMemberUserID, "user-id", "", "", "User ID")
	// teamMembersAddCmd.Flags().StringVarP(&teamMemberPermissions, "permissions", "p", "", "Permissions")
	// teamMembersAddCmd.Flags().StringVarP(&teamMemberUserID, "roles", "r", "", "Roles")

	//teamMembersUpdateCmd.Flags().StringVarP(&teamID, "team-id", "", "", "Team ID")
	// teamMembersUpdateCmd.Flags().StringVarP(&teamMemberID, "id", "", "", "Team Member ID")
	// teamMembersUpdateCmd.Flags().StringVarP(&teamMemberPermissions, "permissions", "p", "", "Permissions")
	// teamMembersUpdateCmd.Flags().StringVarP(&teamMemberRoles, "roles", "r", "", "Roles")

	teamMembersRemoveCmd.Flags().StringVarP(&teamID2, "team-id", "", "", "Team ID")
	teamMembersRemoveCmd.Flags().StringVarP(&teamMemberRemoveID, "team-member-id", "", "", "Team Member ID")

}
