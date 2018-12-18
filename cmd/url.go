package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Generate the url of the api",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet")
		/*command := exec.Command("minikube", "service","apiserver","--url")
		output, err := command.StdoutPipe()
		command.Start()
		if err!=nil{
			fmt.Println("Error getting URL!")
		}else{
			fmt.Println(output)
		}*/
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)
}
