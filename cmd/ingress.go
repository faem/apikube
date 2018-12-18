package cmd

import (
	"apikube/cgo"
	"github.com/spf13/cobra"
)

var ingressCmd = &cobra.Command{
	Use:   "ingress",
	Short: "Creates an ingress to serve the api at linkendin.local",
	Run: func(cmd *cobra.Command, args []string) {
		cgo.CreateIngress()
	},
}

func init() {
	rootCmd.AddCommand(ingressCmd)
}
