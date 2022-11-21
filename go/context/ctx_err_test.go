package context

import (
	"errors"
	"sync"
	"testing"
)

func BenchmarkCtxErr_ErrDefer(b *testing.B) {
	err := errors.New("error defer")
	ctx := ctxErr{
		lock: sync.Mutex{},
		err:  err,
	}
	for i := 0; i < b.N; i++ {
		res := ctx.ErrDefer()
		if err != res {
			b.Fatal(res)
		}
	}
}

func BenchmarkCtxErr_ErrNotDefer(b *testing.B) {
	err := errors.New("error not defer")
	ctx := ctxErr{
		lock: sync.Mutex{},
		err:  err,
	}
	for i := 0; i < b.N; i++ {
		res := ctx.ErrNotDefer()
		if err != res {
			b.Fatal(res)
		}
	}
}
