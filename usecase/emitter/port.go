package usecase

import (
	"context"
	"sbit-emitter/domain/model"
)

type Usecase interface {
	UsecaseInput
	UsecaseOutput
}

// UsecaseInput is the interface that defines the emitter usecase input
type UsecaseInput interface {
	AddDeposit(ctx context.Context, deposit model.Deposit) (interface{}, error)
}

type UsecaseOutput interface {
	AddDepositResponse(ctx context.Context, response model.DepositResponse) (interface{}, error)
}
