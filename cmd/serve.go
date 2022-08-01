package cmd

import (
	"github.com/spf13/cobra"

	"github.com/over55/workery-invoicebuilder/internal/builder"
	"github.com/over55/workery-invoicebuilder/internal/config"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the JSON API over HTTP",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doRunServe()
	},
}

func doRunServe() {
	// Load up all the environment variables.
	appConf := config.AppConfig()

	builder.New(appConf.Server.PDFFilePath)

}
