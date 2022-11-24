package context

import (
	"context"
	"errors"
	"sync"
)

type ctxCancel struct {
	context.Context
	mu   sync.Mutex
	done chan struct{}
	err  error
}

func (m *ctxCancel) Watch() {
	go func() {
		select {
		case <-m.Context.Done():
			//m.Done()
			m.Close()
		case <-m.Done():

		}
	}()
}

func (m *ctxCancel) Done() <-chan struct{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if nil == m.done {
		m.done = make(chan struct{})
	}
	return m.done
}

func (m *ctxCancel) Close() {
	close(m.done)
	m.err = errors.New("my close")
}

func (m *ctxCancel) Err() error {
	return m.err
}

func (*ctxCancel) Value(key interface{}) interface{} {
	return nil
}
