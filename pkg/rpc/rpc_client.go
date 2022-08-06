package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/over55/workery-invoicebuilder/internal/config"
)

type InvoiceBuilderService struct {
	Client        *rpc.Client
	RetryLimit    uint8
	retryCount    uint8
	DelayDuration time.Duration
	addr          string
}

func NewClient(appConfig *config.Conf) *InvoiceBuilderService {
	applicationAddress := fmt.Sprintf("%s:%s", appConfig.Server.IP, appConfig.Server.Port)

	client, err := rpc.DialHTTP("tcp", applicationAddress)
	if err != nil {
		log.Println("RPC CLIENT ERROR | InvoiceBuilderService | Dialing TCP Error:", err)
		return nil
	}

	log.Println("RPC CLIENT | Connected to RPC server.")

	return &InvoiceBuilderService{
		Client:        client,
		RetryLimit:    5,
		retryCount:    1,
		DelayDuration: 10,
		addr:          applicationAddress,
	}
}

// Function used to make RPC calls with retry functionality in case the RPC
// server has been shutdown and the connection was lost.
func (s *InvoiceBuilderService) call(serviceMethod string, args interface{}, reply interface{}) error {
	err := s.Client.Call(serviceMethod, args, reply)

	// Detect the `connection is shut down` error.
	if err == rpc.ErrShutdown {
		if s.retryCount < s.RetryLimit {
			s.retryCount += 1
			log.Println("RPC CLIENT ERROR | InvoiceBuilderService | Detected 'connection is shut down' | Retrying #", s.retryCount)

			// We need to apply an artifical delay in case we need to give time
			// for the server is starting up.
			time.Sleep(s.DelayDuration)

			// Attempt to re-connected and if successful then retry calling the
			// RPC endpoint, else return with error.
			client, err := rpc.DialHTTP("tcp", s.addr)
			if err != nil {
				log.Println("RPC CLIENT ERROR | InvoiceBuilderService | Detected 'connection is shut down' | Failed reconnecting | err:", err.Error())

				// Note: Use recursion to retry the call.
				return s.call(serviceMethod, args, reply)
			}

			log.Println("RPC CLIENT ERROR | InvoiceBuilderService | Detected 'connection is shut down' | Reconnected!")
			s.Client = client

			// Note: Use recursion to retry the call.
			return s.call(serviceMethod, args, reply)
		}
		log.Println("RPC CLIENT ERROR | InvoiceBuilderService | Detected 'connection is shut down' | Too many retries | err:", err.Error())
		return err
	}

	// If success then nil will be returned, else the specific error will be returned.
	return err
}
