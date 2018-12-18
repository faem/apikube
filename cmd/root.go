package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apikube",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	/*Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},*/
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//ootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.apikube.yaml)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
