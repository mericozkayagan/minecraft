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
	Short: "Create a new user",
	Long: `Create a new user with a policy attached to it`,
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
