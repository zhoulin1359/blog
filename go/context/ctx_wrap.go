package context

import (
	"context"
	"errors"
	"sync/atomic"
)

type ctxWrap struct {
	context.Context
	done atomic.Value
}

func (ctx *ctxWrap) SetDone() {
	ctx.done.Store(make(chan struct{}))
}

func (ctx *ctxWrap) Value(key interface{}) interface{} {
	return ctx.Context.Value(key)
}

func (ctx *ctxWrap) Done() <-chan struct{} {
	return ctx.done.Load().(chan struct{})
}

func (ctx *ctxWrap) Close() {
	c := ctx.done.Load().(chan struct{})
	close(c)
}

func (ctx *ctxWrap) Err() error {
	return errors.New("ctxWrap")
}
