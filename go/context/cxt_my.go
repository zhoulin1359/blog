package context

import (
	"errors"
	"sync"
	"time"
)

type ctxMy struct {
	mu   sync.Mutex
	done chan struct{}
	err  error
}

func (*ctxMy) Deadline() (deadline time.Time, ok bool) {
	return
}

func (m *ctxMy) Done() <-chan struct{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if nil == m.done {
		m.done = make(chan struct{})
	}
	return m.done
}

func (m *ctxMy) Close() {
	close(m.done)
	m.err = errors.New("my close")
}

func (m *ctxMy) Err() error {
	return m.err
}

func (*ctxMy) Value(key interface{}) interface{} {
	return nil
}
