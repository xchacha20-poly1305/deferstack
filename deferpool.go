package deferpool

import (
	"errors"
	"fmt"
)

type Deferpool struct {
	funcs []func()
}

func New(funcs ...func()) *Deferpool {
	return &Deferpool{
		funcs: funcs,
	}
}

func (d *Deferpool) Funcs() []func() {
	return d.funcs
}

func (d *Deferpool) Add(a ...func()) {
	d.funcs = append(d.funcs, a...)
}

func (d *Deferpool) Length() int {
	return len(d.funcs)
}

func (d *Deferpool) Run() {
	for i := d.Length() - 1; i >= 0; i-- {
		d.funcs[i]()
	}
}

func (d *Deferpool) Clear() {
	clear(d.funcs)
}

// Remove the final funcs.
func (d *Deferpool) Remove(num int) error {
	if num < 0 {
		return errors.New("num less than zero")
	}
	if num > d.Length() {
		return fmt.Errorf("the length is %d, but got %d", d.Length(), num)
	}
	d.funcs = d.funcs[:num]
	return nil
}
