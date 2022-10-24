package rpc

import (
	"github.com/over55/workery-invoicebuilder/pkg/dtos"
)

func (s *InvoiceBuilderService) Greet(name string) (string, error) {
	req := &dtos.GreetingRequestDTO{
		Name: name,
	}
	var reply dtos.GreetingResponseDTO
	err := s.call("RPC.Greet", req, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Msgf("failed greeting")
		return "", err
	}
	return reply.Message, nil
}
