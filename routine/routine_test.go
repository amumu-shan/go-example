package routine

import "testing"

func TestStart(t *testing.T) {
	Start()
}

func TestFor(t *testing.T) {
	done := false

	go func() {
		done = true
		println("set done true")
	}()

	for !done {
		println("not done !") // 并不内联执行
	}

	println("done !")
}
