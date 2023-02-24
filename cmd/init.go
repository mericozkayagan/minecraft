/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/mericozkayagan/minecraft/src/ec2/assoicate_eip"
	"github.com/mericozkayagan/minecraft/src/ec2/create_instance"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the ec2 instance",
	Long: `Creates the ec2 instance and associates the elastic ip`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	create_instance.CreateInstance(rootCmd.Flag("region").Value.String())
	assoicate_eip.AssociateEIP(rootCmd.Flag("region").Value.String())
}
