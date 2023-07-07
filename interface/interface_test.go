package _interface

import (
	"fmt"
	"testing"
)

func TestSquare_Area(t *testing.T) {
	sq1 := new(Square)
	sq1.side = 5
	sub := sq1.Sub()
	fmt.Println("sub", sub)
	//父类指向子类，可以调用子类的方法（继承有的方法）
	var areaInf Shaper
	areaInf = sq1
	fmt.Println(areaInf.Area())
}
