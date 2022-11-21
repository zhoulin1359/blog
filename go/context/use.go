package context

import (
	"context"
	"fmt"
	"time"
)

func New() {
	context.TODO()
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			time.Sleep(time.Second)
			fmt.Println(1, n)
		}
	}()
	return ch
}
