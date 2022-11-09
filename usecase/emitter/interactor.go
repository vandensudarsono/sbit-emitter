package usecase

import (
	"context"
	emitdomain "sbit-emitter/domain/emitter"
	"sbit-emitter/domain/model"
	logging "sbit-emitter/infrastructure/log"

	"github.com/spf13/viper"
)

type EmitterInterfactor struct {
	e   emitdomain.EmitDomain
	out UsecaseOutput
}

// NewEmitterInteractor is a constructor for emitterinteractor
func NewEmitterInteractor(emitter emitdomain.EmitDomain, out UsecaseOutput) *EmitterInterfactor {
	return &EmitterInterfactor{
		e:   emitter,
		out: out,
	}
}

func (e *EmitterInterfactor) AddDeposit(ctx context.Context, deposit model.Deposit) (interface{}, error) {
	var response model.DepositResponse
	//send data to goka emitter

	err := e.e.EmitMessage(ctx, viper.GetString("broker.topic"), deposit)
	if err != nil {
		logging.WithFields(logging.Fields{"component": "usecase", "action": "add deposit"}).
			Errorf("error emit message: %v", err)
		return nil, err
	}

	response.Status = model.Status{
		Code:          200,
		MessageClient: "success.",
	}

	return e.out.AddDepositResponse(ctx, response)
}
