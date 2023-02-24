/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/mericozkayagan/minecraft/src/iam/attach_policy"
	"github.com/mericozkayagan/minecraft/src/iam/create_policy"
	"github.com/mericozkayagan/minecraft/src/iam/create_user"
)

// createuserCmd represents the createuser command
var createuserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createuser called")
	},
}

func init() {
	rootCmd.AddCommand(createuserCmd)

	create_user.CreateUser(rootCmd.Flag("region").Value.String())
	create_policy.CreatePolicy(rootCmd.Flag("region").Value.String())
	attach_policy.AttachPolicy(rootCmd.Flag("region").Value.String())
}
