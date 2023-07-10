package _struct

import (
	"fmt"
	"strconv"
)

type List []int

func (l *List) Len() int {
	return len(*l)
}
func (l *List) Append(val int) {
	*l = append(*l, val)
}

// Car1 -------------
type Car1 struct {
	wheelCount int
}

func (c *Car1) numberOfWheels() {
	fmt.Println("car numberOfWheels:", c.wheelCount)
}

type Mercedes struct {
	Car1
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("mercedes sayHiToMerkel")
}

// Base Magic -------------
type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type Integer int

func (i *Integer) String() string {
	return strconv.Itoa(int(*i))
}
func (p Integer) get() int {
	return int(p)
}
func f(i int) {

}

var v Integer
