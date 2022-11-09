package emitter

import (
	"context"
	"sbit-emitter/domain/model"

	"github.com/lovoo/goka"
)

type EmitterServer struct {
	e *goka.Emitter
}

func NewEmitterServer(e *goka.Emitter) *EmitterServer {
	return &EmitterServer{e: e}
}

func (em *EmitterServer) EmitMessage(ctx context.Context, key string, message model.Deposit) error {

	_, err := em.e.Emit(key, &message)

	return err
}
