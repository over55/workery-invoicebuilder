package rpc

import (
	// "errors"

	"log"

	"github.com/over55/workery-invoicebuilder/pkg/dtos"
)

func (s *InvoiceBuilderService) GeneratePDF(dto *dtos.WorkOrderInvoiceRequestDTO) (*dtos.WorkOrderInvoiceResponseDTO, error) {
	var reply dtos.WorkOrderInvoiceResponseDTO
	err := s.call("RPC.GeneratePDF", dto, &reply)
	if err != nil {
		log.Println("rpc_client | RPC.GeneratePDF | err", err)
		return nil, err
	}
	return &reply, nil
}
