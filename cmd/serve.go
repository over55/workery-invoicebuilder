package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/over55/workery-invoicebuilder/internal/app"
	"github.com/over55/workery-invoicebuilder/internal/config"
	"github.com/over55/workery-invoicebuilder/internal/inputports/rpc"
	"github.com/over55/workery-invoicebuilder/pkg/uuid"
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

	uuidp := uuid.NewUUIDProvider()

	// Setup the application that handles the logic of generating PDFs.
	a, err := app.New(appConf, uuidp)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the RPC server to serve out the PDF generator.
	srv := rpc.NewServer(appConf, a)

	// Run in the forground the RPC server. When the server gets termination
	// signal then the server will terminate.
	if err := srv.RunMainRuntimeLoop(); err != nil {
		log.Fatal(err)
	}

}
