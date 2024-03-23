package deferstack

import (
	"sync"
)

type Deferstack interface {
	// Add used to add new functions.
	Add(...func())

	// Funcs show functions in stack.
	Funcs() []func()

	// Length get the length of stack.
	Length() int

	// Run starts running.
	Run()

	// Clean will remove all the functions.
	Clean()
}

type deferstack struct {
	funcs []func()

	mu sync.Mutex
}

var _ Deferstack = (*deferstack)(nil)

func New(funcs ...func()) Deferstack {
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

// Use it in internal. It doesn't have lock!
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

// Clean will remove all the funcs of deferstack.
func (d *deferstack) Clean() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.funcs = []func(){}
}
