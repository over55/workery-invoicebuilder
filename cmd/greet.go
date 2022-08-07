package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/over55/workery-invoicebuilder/internal/config"
	"github.com/over55/workery-invoicebuilder/pkg/rpc"
)

func init() {
	rootCmd.AddCommand(greetCmd)
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Call `greet` RPC endpoint with sample data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Load up all the environment variables.
		appConf := config.AppConfig()

		// Connect to a running client.
		applicationAddress := fmt.Sprintf("%s:%s", appConf.Server.IP, appConf.Server.Port)
		client := rpc.NewClient(applicationAddress, 3, 15*time.Second)

		// Execute the remote call.
		message, err := client.Greet("Aki Kisaragi")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(message)
	},
}
