package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate go gin starter kit",
	Long:  `generate go gin starter kit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		what := args[0]
		if what == "controller" {
			createController(args[1])
		}
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func createController(name string) {
	fmt.Println("createController: " + name)
}
