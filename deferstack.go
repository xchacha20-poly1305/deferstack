package deferstack

import (
	"errors"
	"fmt"
	"sync"
)

type deferstack struct {
	funcs []func()

	mu sync.Mutex
}

func New(funcs ...func()) *deferstack {
	return &deferstack{
		funcs: funcs,
	}
}

func (d *deferstack) Funcs() []func() {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.funcs
}

func (d *deferstack) Add(a ...func()) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.funcs = append(d.funcs, a...)
}

func (d *deferstack) Length() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.funcs)
}

func (d *deferstack) Run() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i := d.Length() - 1; i >= 0; i-- {
		d.funcs[i]()
	}
}

func (d *deferstack) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	clear(d.funcs)
}

// Remove `num` final funcs.
func (d *deferstack) Remove(num int) error {
	if num < 0 {
		return errors.New("num less than zero")
	}
	if num > d.Length() {
		return fmt.Errorf("the length is %d, but got %d", d.Length(), num)
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.funcs = d.funcs[:num+1]
	return nil
}
