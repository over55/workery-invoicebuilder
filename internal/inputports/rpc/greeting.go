package rpc

import (
	"fmt"

	"github.com/over55/workery-invoicebuilder/pkg/dtos"
)

func (rpc *RPC) Greet(req *dtos.GreetingRequestDTO, res *dtos.GreetingResponseDTO) error {
	*res = dtos.GreetingResponseDTO{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}
	return nil
}
