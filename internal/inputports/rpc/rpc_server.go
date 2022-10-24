package rpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

    "github.com/rs/zerolog"

	"github.com/over55/workery-invoicebuilder/internal/app"
	"github.com/over55/workery-invoicebuilder/internal/config"
)

type RPC struct {
	Logger      zerolog.Logger
	App         app.InvoiceBuilder
	tcpListener *net.TCPListener
}

func NewServer(appConfig *config.Conf, a app.InvoiceBuilder, logger zerolog.Logger) *RPC {
	applicationAddress := fmt.Sprintf("%s:%s", appConfig.Server.IP, appConfig.Server.Port)
	logger.Info().Msgf("initializing address for %s", applicationAddress)

	tcpAddr, err := net.ResolveTCPAddr("tcp", applicationAddress)
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := &RPC{Logger: logger, App: a}

	// Attach the
	rpc.Register(rpcServer)
	rpc.HandleHTTP()

	logger.Info().Msg("rpc api was initialized")
	l, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		l.Close()
		logger.Fatal().Err(e).Caller().Msg("rpc api failed to initialize")
	}
	logger.Info().Msg("rpc api is listening now")
	rpcServer.tcpListener = l
	return rpcServer
}

func (rpcServer *RPC) RunMainRuntimeLoop() error {
	rpcServer.Logger.Info().Msg("rpc api is starting now...")

	// The following code will attach a background handler to run when the
	// application detects a shutdown signal.
	// Special thanks via https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.

		// Finish any RPC communication taking place at the moment before
		// shutting down the RPC server.
		rpcServer.tcpListener.Close()
	}()

	// Attach the following anonymous function to run on all cases (ex: panic,
	// termination signal, etc) so we can gracefully shutdown the service.
	defer func() {
		rpcServer.stopRuntimeLoop()
	}()

	// Safety net for 'too many open files' issue on legacy code.
	// Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	// Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	http.DefaultClient.Timeout = time.Minute * 10

	// DEVELOPER NOTES:
	// If you get "too many open files" then please read the following article
	// http://publib.boulder.ibm.com/httpserv/ihsdiag/too_many_open_files.html
	// so you can run in your console:
	// $ ulimit -H -n 4096
	// $ ulimit -n 4096

	rpcServer.Logger.Info().Msg("rpc api is ready and running")

	// Run the main loop blocking code.
	http.Serve(rpcServer.tcpListener, nil)

	return nil
}

func (rpcServer *RPC) stopRuntimeLoop() error {
	rpcServer.Logger.Info().Msg("starting graceful shutdown now...")
	rpcServer.tcpListener.Close()
	rpcServer.Logger.Info().Msg("graceful shutdown finished")
	return nil
}
