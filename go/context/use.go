package context

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/dump"
	"log"
	"sync"
	"sync/atomic"
	"time"
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
	ctx = &valueCtx{
		Context: ctx,
		keys:    sync.Map{},
	}
	ctx1, ok := ctx.(*valueCtx)
	if ok {
		ctx1.SetValue("a", "a1")
		ctx1.SetValue("a", "a2")
	}

	fmt.Println(ctx.Value("a"))

	ctx = context.WithValue(ctx, "key", "a3")
	fmt.Println(ctx.Value("key"))
	dump.P(ctx)
}

func UseMap() {
	m := map[string]int{}
	for i := 0; i < 10000; i++ {
		m[fmt.Sprintf("%d", i)] = i
	}
	go func() {
		for k, v := range m {
			fmt.Println(k, v)
		}
	}()
	for i := 0; i < 10000; i++ {
		//m[fmt.Sprintf("%d", i)] = i
		delete(m, fmt.Sprintf("%d", i))
	}
	//time.Sleep(time.Second)
}

func UseCloseChan() {
	c := make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 1

	close(c)
	fmt.Println(<-c)
	c <- 1
}

func UseParentNil() {
	ctx := context.Background()
	//取消的上下文
	ctx, cancel := context.WithCancel(ctx)
	//自定义的上下文
	ctx = &valueCtx{
		ctx,
		sync.Map{},
	}
	//又一个取消的上下文
	ctx, _ = context.WithCancel(ctx)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		log.Println("cancel", ctx1.Err(), &ctx1)
	}(ctx)
	//dump.P(ctx)
	time.Sleep(time.Second * 1)
	cancel()
	time.Sleep(time.Second * 1)
}

func UseParentDone() {
	ctx := context.Background()
	//取消的上下文
	ctx, cancel := context.WithCancel(ctx)
	//自定义的上下文
	ctx1 := &ctxCancel{
		Context: ctx,
		mu:      sync.Mutex{},
		done:    nil,
		err:     nil,
	}
	//需要监听父级context是否退出了
	ctx1.Watch()
	//取消的上下文1
	ctx, _ = context.WithCancel(ctx1)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		log.Println("cancel1", ctx1.Err(), &ctx1)
	}(ctx)
	//取消的上下文2
	ctx, _ = context.WithCancel(ctx1)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		log.Println("cancel2", ctx1.Err(), &ctx1)
	}(ctx)

	time.Sleep(time.Second * 1)
	cancel()
	time.Sleep(time.Second * 1)
}

func UseCtxWrap() {
	ctx := context.TODO()
	ctx, _ = context.WithCancel(ctx)

	ctx1 := &ctxWrap{
		Context: ctx,
		done:    atomic.Value{},
	}
	ctx1.SetDone()

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	fmt.Println(c1, c2, c1 == c2)
	ctx, _ = context.WithCancel(ctx1)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		log.Println("cancel", ctx1.Err(), &ctx1)
	}(ctx)
	time.Sleep(time.Second * 1)
	ctx1.Close()
	time.Sleep(time.Second * 1)
}

func UseCancel() {
	ctx := context.TODO()
	//取消.父可以取消子
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		log.Println("f1", ctx1.Err(), &ctx1)
	}(ctx)
	ctx, _ = context.WithCancel(ctx)
	go func(ctx2 context.Context) {
		<-ctx2.Done()
		log.Println("f2", ctx2.Err(), &ctx2)
	}(ctx)
	//dump.P(ctx)
	time.Sleep(time.Second * 1)
	//取消父亲
	cancel()
	time.Sleep(time.Second * 1)

	//取消.子不可以取消父亲
	ctx = context.TODO()
	ctx, _ = context.WithCancel(ctx)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		log.Println("f1", ctx1.Err(), &ctx1)
	}(ctx)
	ctx, cancel2 := context.WithCancel(ctx)
	go func(ctx2 context.Context) {
		<-ctx2.Done()
		log.Println("f2", ctx2.Err(), &ctx2)
	}(ctx)
	//dump.P(ctx)
	time.Sleep(time.Second * 1)
	//取消子
	cancel2()
	time.Sleep(time.Second * 1)

}

func UseCtxClosure() {
	ctx := context.TODO()
	ctx, _ = context.WithCancel(ctx)
	go func() {
		<-ctx.Done()
		log.Println("f1", ctx.Err(), &ctx)
	}()
	ctx, cancel2 := context.WithCancel(ctx)
	go func() {
		<-ctx.Done()
		log.Println("f2", ctx.Err(), &ctx)
	}()
	time.Sleep(time.Second * 1)
	//取消子
	cancel2()
	time.Sleep(time.Second * 1)
}

func ContextParent() {
	ctx := context.TODO()
	ctx1, _ := context.WithCancel(ctx)
	go printName(ctx1, "ctx1")
	ctx2, cancel := context.WithCancel(ctx1)
	go printName(ctx2, "ctx2")
	ctx31, _ := context.WithCancel(ctx2)
	go printName(ctx31, "ctx31")
	ctx32, _ := context.WithCancel(ctx2)
	go printName(ctx32, "ctx32")
	time.Sleep(time.Second * 1)
	dump.P(ctx1)
	dump.P(ctx2)
	cancel()
	dump.P(ctx1)
	dump.P(ctx2)
	time.Sleep(time.Second * 1)
	/*dump.P(ctx1)
	dump.P(ctx2)
	dump.P(ctx3)*/
}

func ContextParentValue() {
	ctx := context.TODO()
	ctx1, cancel := context.WithCancel(ctx)
	go printName(ctx1, "ctx1")
	ctx2 := context.WithValue(ctx1, "k", "v")
	go printName(ctx2, "ctx2")
	ctx3 := context.WithValue(ctx2, "k", "v")
	go printName(ctx3, "ctx3")
	ctx4, _ := context.WithCancel(ctx3)
	go printName(ctx4, "ctx4")
	time.Sleep(time.Second * 1)
	dump.P(ctx2)
	cancel()
	time.Sleep(time.Second * 1)

}

func ContextMy() {
	ctx := &ctxCancel{}
	ctx1, _ := context.WithCancel(ctx)
	go printName(ctx1, "ctx1")
	time.Sleep(time.Second * 1)
	ctx.Close()
	time.Sleep(time.Second * 1)
}

func printName(ctx context.Context, name string) {
	<-ctx.Done()
	log.Println(name, ctx.Err())
}

func ContextFunc() {
	ctx := context.TODO()
	ctx, _ = context.WithCancel(ctx)
	log.Println(fmt.Sprintf("1 is %p", ctx))
	go func() {
		log.Println(fmt.Sprintf("1 ctx is %p", ctx))
		ok := ctx.Done()
		log.Println(fmt.Sprintf("1 ctx ok is %p", ok))
		<-ok
		log.Println(fmt.Sprintf("f1 %s is %p", ctx.Err(), ctx))
	}()
	ctx, cancel2 := context.WithCancel(ctx)
	log.Println(fmt.Sprintf("2 is %p", ctx))
	go func() {
		log.Println(fmt.Sprintf("2 ctx is %p", ctx))
		//<-ctx.Done()
		ok := ctx.Done()
		log.Println(fmt.Sprintf("2 ctx ok is %p", ok))
		<-ok
		log.Println(fmt.Sprintf("f2 %s is %p", ctx.Err(), ctx))
	}()
	time.Sleep(time.Second)
	cancel2()
	time.Sleep(time.Second * 10)
}
