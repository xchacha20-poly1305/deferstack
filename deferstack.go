package deferstack

import (
	"errors"
	"fmt"
)

type deferstack struct {
	funcs []func()
}

func New(funcs ...func()) *deferstack {
	return &deferstack{
		funcs: funcs,
	}
}

func (d *deferstack) Funcs() []func() {
	return d.funcs
}

func (d *deferstack) Add(a ...func()) {
	d.funcs = append(d.funcs, a...)
}

func (d *deferstack) Length() int {
	return len(d.funcs)
}

func (d *deferstack) Run() {
	for i := d.Length() - 1; i >= 0; i-- {
		d.funcs[i]()
	}
}

func (d *deferstack) Clear() {
	clear(d.funcs)
}

// Remove the final funcs.
func (d *deferstack) Remove(num int) error {
	if num < 0 {
		return errors.New("num less than zero")
	}
	if num > d.Length() {
		return fmt.Errorf("the length is %d, but got %d", d.Length(), num)
	}
	d.funcs = d.funcs[:num]
	return nil
}
