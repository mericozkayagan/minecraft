/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/iam/attach_policy"
	"github.com/mericozkayagan/minecraft/src/iam/create_policy"
	"github.com/mericozkayagan/minecraft/src/iam/create_user"
	"github.com/spf13/cobra"
)

// createuserCmd represents the createuser command
var createuserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "Create a new user",
	Long:  `Create a new user with a policy attached to it`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createuser called")

		create_user.CreateUser()
		create_policy.CreatePolicy()
		attach_policy.AttachPolicy()
	},
}

func init() {
	rootCmd.AddCommand(createuserCmd)
}
