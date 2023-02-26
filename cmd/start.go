package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/start_stop_instance"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the instance",
	Long:  `Start the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		start_stop_instance.StartStopInstance("start")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
