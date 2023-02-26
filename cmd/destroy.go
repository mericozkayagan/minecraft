package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/destroy_instance"
	"github.com/mericozkayagan/minecraft/src/ec2/filter_by_tag"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the instance",
	Long:  `Destroy the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("destroy called")
		instanceId, _ := filter_by_tag.FilterByTag()
		destroy_instance.DestroyInstance(instanceId)
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
