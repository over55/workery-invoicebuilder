package rpc

import (
	"github.com/over55/workery-invoicebuilder/pkg/dtos"
)

func (s *InvoiceBuilderService) GeneratePDF(dto *dtos.WorkOrderInvoiceRequestDTO) (*dtos.WorkOrderInvoiceResponseDTO, error) {
	var reply dtos.WorkOrderInvoiceResponseDTO
	err := s.call("RPC.GeneratePDF", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Msgf("failed generating pdf")
		return nil, err
	}
	return &reply, nil
}
