package other

import (
	"fmt"
)

func DeferResult() {
	var i = 1
	//对defer延迟执行的函数，它的参数会在声明时候就会求出具体值，而不是在执行时才求值
	defer fmt.Println("result:", func() int { return i * 2 }()) //2
	i++
}
