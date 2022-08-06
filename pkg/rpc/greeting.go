package rpc

import (
	// "errors"

	"log"

	"github.com/over55/workery-invoicebuilder/pkg/dtos"
)

func (s *InvoiceBuilderService) Greet(name string) (string, error) {
	req := &dtos.GreetingRequestDTO{
		Name: name,
	}
	var reply dtos.GreetingResponseDTO
	err := s.call("RPC.Greet", req, &reply)
	if err != nil {
		log.Println("rpc_client | RPC.Greet | err", err)
		return "", err
	}
	return reply.Message, nil
}
