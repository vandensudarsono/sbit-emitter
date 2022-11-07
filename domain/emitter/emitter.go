package emitdomain

import "context"

type EmitDomain interface {
	EmitMessage(ctx context.Context, key string, message interface{}) error
}
