package _struct

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	var lst List
	lst.Append(1)
	fmt.Printf("%v(len:%d)", lst, lst.Len())

	plst := new(List)
	plst.Append(2)
	fmt.Printf("----%v(len:%d)", plst, plst.Len())

	m := new(Mercedes)
	m.numberOfWheels()
	m.sayHiToMerkel()

	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()

	var i Integer = 5
	s := i.String()
	fmt.Println("Integer=", s)

	fmt.Println("--------")
	st := new(Stack)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	fmt.Printf("%v\n", st)
	fmt.Println("pop:", st.Pop())
	fmt.Println("pop:", st.Pop())
	fmt.Println("pop:", st.Pop())
	fmt.Println("pop:", st.Pop())
}
