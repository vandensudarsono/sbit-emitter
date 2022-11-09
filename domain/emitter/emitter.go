package emitdomain

import (
	"context"
	"sbit-emitter/domain/model"
)

type EmitDomain interface {
	EmitMessage(ctx context.Context, key string, message model.Deposit) error
}
