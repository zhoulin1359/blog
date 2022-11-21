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
	v, _ := ctx.keys.Load(key)
	return v
}
