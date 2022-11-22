package context

import (
	"context"
	"sync"
)

type valueCtx struct {
	context.Context
	keys sync.Map
}

func (ctx *valueCtx) Value(key interface{}) interface{} {
	v, ok := ctx.keys.Load(key)
	if ok {
		return v
	}
	return ctx.Context.Value(key)
}
