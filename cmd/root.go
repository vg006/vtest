package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vtest",
	Short: "A Pen test tool",
	Long:  `vtest is a Pen test tool that helps you to create a new Pen test project with a predefined structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Welcome to vtest!")
		cmd.Println("Use 'vtest --help' for more information.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
