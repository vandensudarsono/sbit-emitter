package output

import (
	"context"
	"sbit-emitter/domain/model"
	pb "sbit-emitter/infrastructure/server/grpc/proto/emitter"
)

// OutputPort is an interface for output port
type OutputPort struct{}

// NewEmitterOutputPortService creates a new emitter output port service
func NewEmitterOutputPortService() *OutputPort {
	return &OutputPort{}
}

func (o *OutputPort) AddDepositResponse(ctx context.Context, response model.DepositResponse) (interface{}, error) {

	out := &pb.DepositResponse{
		Status: &pb.Status{
			Code:          response.Status.Code,
			MessageClient: response.Status.MessageClient,
		},
	}

	return out, nil
}
