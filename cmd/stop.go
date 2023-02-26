package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/filter_by_tag"
	"github.com/mericozkayagan/minecraft/src/ec2/start_stop_instance"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the instance",
	Long:  `Stop the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")

		fmt.Println()
		instanceId, publicIp := filter_by_tag.FilterByTag()
		start_stop_instance.StartStopInstance("stop", instanceId, publicIp)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
