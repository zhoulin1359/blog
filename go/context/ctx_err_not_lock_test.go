package context

import (
	"sync"
	"testing"
	"time"
)

func TestNotLockErr_Err(t *testing.T) {
	ctx := notLockErr{
		lock: sync.Mutex{},
		err:  nil,
		done: make(chan struct{}),
	}

	for i := 0; i < 1000; i++ {
		go func() {
			select {
			case <-ctx.Done():
				if nil == ctx.Err() {
					panic("not err")
				}
			}
		}()
	}
	time.Sleep(time.Second)
	ctx.Close()

	time.Sleep(time.Second * 10)

}

func TestNotLockErr_ErrLock(t *testing.T) {
	ctx := notLockErr{
		lock: sync.Mutex{},
		err:  nil,
		done: make(chan struct{}),
	}

	for i := 0; i < 1000; i++ {
		go func() {
			select {
			case <-ctx.Done():
				if nil == ctx.ErrLock() {
					panic("not err")
				}
			}
		}()
	}
	time.Sleep(time.Second)
	ctx.Close()

	time.Sleep(time.Second * 10)

}
