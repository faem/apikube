package cmd

import (
	"apikube/cgo"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scales up the deployment to 5",
	Run: func(cmd *cobra.Command, args []string) {
		cgo.UpdateDeployment()
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
}
