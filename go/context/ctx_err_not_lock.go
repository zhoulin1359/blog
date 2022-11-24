package context

import (
	"errors"
	"sync"
)

type notLockErr struct {
	lock sync.Mutex
	err  error
	done chan struct{}
}

func (n *notLockErr) Done() <-chan struct{} {
	return n.done
}

func (n *notLockErr) Close() {
	n.lock.Lock()
	//注意顺序
	close(n.done)
	n.err = errors.New("err")
	n.lock.Unlock()
}

func (n *notLockErr) Err() error {
	return n.err
}

func (n *notLockErr) ErrLock() error {
	n.lock.Lock()
	err := n.err
	n.lock.Unlock()
	return err
}
