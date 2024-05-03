package ccontext

import (
	"context"
	"github.com/OpenIMSDK/tools/mcontext"
)

func WithOperationID(ctx context.Context, operationID string) context.Context {
	return mcontext.SetOperationID(ctx, operationID)
}
