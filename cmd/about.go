package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "about go gin starter kit",
	Long:  `about go gin starter kit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go gin starter kit created by saeed heidari")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
