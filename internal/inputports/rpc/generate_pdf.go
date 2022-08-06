package rpc

import (
	"github.com/over55/workery-invoicebuilder/pkg/dtos"
)

func (rpc *RPC) GeneratePDF(req *dtos.WorkOrderInvoiceRequestDTO, res *dtos.WorkOrderInvoiceResponseDTO) error {

	pdf, err := rpc.App.GeneratePDF(req)
	if err != nil {
		return err
	}

	*res = *pdf
	return nil
}
