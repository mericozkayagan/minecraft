package cmd

import (
	"fmt"

	"github.com/mericozkayagan/minecraft/src/ec2/create_instance"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Creates the ec2 instance and the necessary resources",
	Long:  `Creates the ec2 instance and the security group
	The security group has ports 22 and 25565`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Minecraft server...")
		create_instance.CreateInstance()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
