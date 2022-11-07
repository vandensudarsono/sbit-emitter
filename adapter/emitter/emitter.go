package emitter

import (
	"context"
	"fmt"

	"github.com/lovoo/goka"
)

type EmitterServer struct {
	e *goka.Emitter
}

func NewEmitterServer(e *goka.Emitter) *EmitterServer {
	return &EmitterServer{e: e}
}

func (em *EmitterServer) EmitMessage(ctx context.Context, key string, message interface{}) error {
	// res, err := json.Marshal(message)
	// if err != nil {
	// 	return err
	// }
	data := fmt.Sprintf("%+v", message)
	_, err := em.e.Emit(key, data)

	return err
}
