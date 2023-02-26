package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/filter_by_tag"
	"github.com/mericozkayagan/minecraft/src/ec2/reboot_instance"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot the instance",
	Long:  `Reboot the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reboot called")

		instanceId, publicIp := filter_by_tag.FilterByTag()
		reboot_instance.RebootInstance(instanceId, publicIp)
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}
