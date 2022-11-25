package context

import "sync"

type ctxErr struct {
	lock sync.Mutex
	rw   sync.RWMutex
	err  error
}

func (c *ctxErr) ErrDefer() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.err
}

func (c *ctxErr) ErrNotDefer() error {
	c.lock.Lock()
	err := c.err
	c.lock.Unlock()
	return err
}

func (c *ctxErr) ErrNotDeferRW() error {
	c.rw.RLock()
	err := c.err
	c.rw.RUnlock()
	return err
}
