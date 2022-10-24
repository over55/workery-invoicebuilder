package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

    // Load up provider.
	uuidp := uuid.NewUUIDProvider()

	// Load up third-party logging library.
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Setup the application that handles the logic of generating PDFs.
	a, err := app.New(appConf, uuidp, logger)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("failed initializing app")
	}

	// Setup the RPC server to serve out the PDF generator.
	srv := rpc.NewServer(appConf, a, logger)

	// Run in the forground the RPC server. When the server gets termination
	// signal then the server will terminate.
	if err := srv.RunMainRuntimeLoop(); err != nil {
		log.Fatal().Err(err).Caller().Msg("failed running main runtime loop")
	}

}
