/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/reboot_instance"
	"github.com/spf13/cobra"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot the instance",
	Long:  `Reboot the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reboot called")

		reboot_instance.RebootInstance()
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}
