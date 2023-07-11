package panic

import (
	"fmt"
)

func DoRecover() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	panic("not good")
}

// DoRecover1 错误的用法
func DoRecover1() {
	defer func() {
		doRecover()
	}()
	panic("not good")
}
func doRecover() {
	fmt.Println("recovered:", recover())
}
