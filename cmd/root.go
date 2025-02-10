package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	app "github.com/vg006/vtest/internal"
	asset "github.com/vg006/vtest/internal/assets"
)

func init() {
	rootCmd.Flags().IntP("port", "p", 8080, "Port to run the server on")
}

var rootCmd = &cobra.Command{
	Use:   "vtest",
	Short: "A Pen test tool",
	Long:  `vtest is a Pen test tool that helps you to create a new Pen test project with a predefined structure.`,

	Run: runRootCmd,
}

func runRootCmd(cli *cobra.Command, args []string) {
	cli.Println(asset.VtestLogo)
	cli.Println(asset.Text.Render("vgo is a Go project scaffolding tool.\n"))
	cli.Println(asset.Text.Render("Use 'vgo --help' for more information."))

	port := fmt.Sprintf(":%s", cli.Flag("port").Value.String())

	app := app.New(port)
	app.Exec()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(asset.Exit + asset.Msg.Render(err.Error()))
		os.Exit(1)
	}
}
