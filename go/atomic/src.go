package atomic

import (
	"sync"
	"sync/atomic"
)

func Add(num int) int32 {
	w := sync.WaitGroup{}
	w.Add(num)
	var res int32 = 0
	for i := 0; i < num; i++ {
		go func() {
			atomic.AddInt32(&res, 1)
			w.Done()
		}()
	}
	w.Wait()
	return res
}

func CompareAndSwap(num int) int32 {
	var addr, old, new int32 = 0, 0, 0
	addr = 1
	w := sync.WaitGroup{}
	w.Add(num)
	var res int32 = 0
	for i := 0; i < num; i++ {
		go func() {
			old = 1
			new = 2
			ok := atomic.CompareAndSwapInt32(&addr, old, new)
			if ok {
				atomic.AddInt32(&res, 1)
			}
			w.Done()
		}()
	}
	w.Wait()
	return res
}
