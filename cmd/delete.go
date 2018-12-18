package cmd

import (
	"apikube/cgo"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes deployment, service and ingress",
	Run: func(cmd *cobra.Command, args []string) {
		cgo.DeleteDeployment()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
