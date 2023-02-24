/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/start_stop_instance"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the instance",
	Long: `Start the instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		
		start_stop_instance.StartStopInstance()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
