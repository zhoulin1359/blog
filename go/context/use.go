package context

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/dump"
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
