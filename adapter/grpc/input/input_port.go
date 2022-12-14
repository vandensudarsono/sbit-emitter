package input

import (
	"context"
	"sbit-emitter/domain/model"
	ucEmitter "sbit-emitter/usecase/emitter"

	pb "sbit-emitter/infrastructure/server/grpc/proto/emitter"

	"github.com/mitchellh/mapstructure"
)

type InputPort struct {
	usecaseInput ucEmitter.UsecaseInput
}

func NewEmitterInputPortService(usecaseInput ucEmitter.UsecaseInput) *InputPort {
	return &InputPort{usecaseInput: usecaseInput}
}

func (inport *InputPort) Deposit(ctx context.Context, in *pb.DepositRequest) (*pb.DepositResponse, error) {
	var out *pb.DepositResponse

	doc, err := inport.usecaseInput.AddDeposit(ctx, model.Deposit{
		WalletID: in.WalletId,
		Amount:   in.Amount,
	})

	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(doc, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
