package cmd

import (
	"fmt"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var kubernetesShowCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"get", "inspect"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Show kubernetes cluster",
	Long: `Show your current kubernetes cluster.
If you wish to use a custom format, the available fields are:

	* ID
	* Code
	* Name
	* AccountID
	* ImageID
	* VolumeID
	* ShortDescription
	* Description
	* DefaultUsername
	* CloudConfig

Example: civo kubernetes show ID/NAME -o custom -f "ID: Code (DefaultUsername)"`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			fmt.Printf("Unable to create a Civo API Client: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		kubernetesCluster, err := client.FindKubernetesCluster(args[0])
		if err != nil {
			fmt.Printf("Unable to search template: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()
		ow.StartLine()

		ow.AppendData("ID", kubernetesCluster.ID)
		ow.AppendData("Name", kubernetesCluster.Name)
		ow.AppendData("Nodes", strconv.Itoa(kubernetesCluster.NumTargetNode))
		ow.AppendData("Size", kubernetesCluster.TargetNodeSize)
		ow.AppendData("Status", kubernetesCluster.Status)
		ow.AppendData("Version", kubernetesCluster.Version)
		ow.AppendDataWithLabel("APIEndPoint", kubernetesCluster.APIEndPoint, "API Endpoint")
		ow.AppendDataWithLabel("DNSEntry", kubernetesCluster.DNSEntry, "DNS A record")

		if outputFormat == "json" || outputFormat == "custom" {
			//ow.AppendData("CloudConfig", template.CloudConfig)
			if outputFormat == "json" {
				ow.WriteSingleObjectJSON()
			} else {
				ow.WriteCustomOutput(outputFields)
			}
		} else {
			ow.WriteKeyValues()

			if len(kubernetesCluster.Instances) > 0 {
				fmt.Println()
				ow.WriteHeader("Nodes")
				owNode := utility.NewOutputWriter()

				for _, instance := range kubernetesCluster.Instances {
					owNode.StartLine()

					owNode.AppendData("Name", instance.Hostname)
					owNode.AppendData("IP", instance.PublicIP)
					owNode.AppendData("Status", instance.Status)
					owNode.AppendData("Size", instance.Size)
				}
				owNode.WriteTable()
			}

			if len(kubernetesCluster.InstalledApplications) > 0 {
				fmt.Println()
				ow.WriteHeader("Applications")
				owApp := utility.NewOutputWriter()

				for _, app := range kubernetesCluster.InstalledApplications {
					owApp.StartLine()

					owApp.AppendData("Name", app.Application)
					owApp.AppendData("Version", app.Version)
					owApp.AppendData("Installed", strconv.FormatBool(app.Installed))
					owApp.AppendData("Category", app.Category)
				}
				owApp.WriteTable()
			}

		}

	},
}
