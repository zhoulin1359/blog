package context

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/dump"
	"sync"
)

// UseValue 传值的context
func UseValue() {
	ctx := context.TODO()
	//设置值
	ctx = context.WithValue(ctx, "a", "a1")
	fmt.Println(ctx.Value("a"))
	ctx = context.WithValue(ctx, "a", "a2")
	fmt.Println(ctx.Value("a"))
	dump.P(ctx)
}

func UseCtxValue() {
	ctx := context.TODO()
	store := sync.Map{}
	store.Store("a", "a1")
	store.Store("a", "a2")
	ctx = &valueCtx{
		Context: ctx,
		keys:    store,
	}
	fmt.Println(ctx.Value("a"))

	ctx = context.WithValue(ctx, "key", "a3")
	fmt.Println(ctx.Value("key"))
	dump.P(ctx)
}
