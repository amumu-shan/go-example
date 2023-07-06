package block

import (
	"runtime"
)

var _ = runtime.GOMAXPROCS(4)
var a, b int

func u1() {
	a = 1
	b = 2
}
func u2() {
	a = 3
	b = 4
}
func p() {
	println(a)
	println(b)
}

func OrderOut() {
	in := make(chan func())
	out := make(chan func())

	go func(in <-chan func(), out chan<- func()) {
		for {
			select {
			case v := <-in:
				out <- v
			}
		}
	}(in, out)
	go func() {
		for {
			select {
			case v := <-out:
				v()
			}
		}
	}()

	in <- func() {
		u1()
	}
	in <- func() {
		u2()
	}
	in <- func() {
		p()
	}
}
