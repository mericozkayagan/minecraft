package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/start_stop_instance"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the instance",
	Long:  `Stop the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
		start_stop_instance.StartStopInstance("stop")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
