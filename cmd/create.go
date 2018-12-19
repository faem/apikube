package cmd

import (
	"apikube/cgo"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a deployment",
	Run: func(cmd *cobra.Command, args []string) {
		//cgo.CreateDeployment()
		cgo.CreateDeploymentKutil()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
