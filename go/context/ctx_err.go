package context

import "sync"

type ctxErr struct {
	lock sync.Mutex
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
