package cmd

import (
	"apikube/cgo"
	"github.com/spf13/cobra"
)

var exposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Exposes the deployment using a service",
	Run: func(cmd *cobra.Command, args []string) {
		cgo.CreateService()
	},
}

func init() {
	rootCmd.AddCommand(exposeCmd)
}
