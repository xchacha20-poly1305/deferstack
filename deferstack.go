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
	return d.length()
}

// Use it in internal. It doesn't has lock!
func (d *deferstack) length() int {
	return len(d.funcs)
}

func (d *deferstack) Run() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i := d.length() - 1; i >= 0; i-- {
		if d.funcs[i] == nil {
			continue
		}
		d.funcs[i]()
	}
}

// Clean() will remove all the funcs of your deferstack.
func (d *deferstack) Clean() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.funcs = []func(){}
}

// Remove final funcs.
func (d *deferstack) Remove(num int) error {
	if num < 0 {
		return errors.New("num less than zero")
	}
	if num > d.length() {
		return fmt.Errorf("the length is %d, but got %d", d.length(), num)
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.funcs = d.funcs[:num+1]
	return nil
}
