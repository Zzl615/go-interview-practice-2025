// lock.go
package lock

import "sync"

type RW interface {
	Write()
	Read()
}

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	_ = l.count
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	l.mu.RUnlock()
}
