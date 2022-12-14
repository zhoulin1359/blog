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

func (ctx *valueCtx) SetValue(key, value interface{}) {
	ctx.keys.Store(key, value)
}

func (ctx *valueCtx) Done() <-chan struct{} {
	return nil
}
