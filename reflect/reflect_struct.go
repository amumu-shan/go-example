package _reflect

import (
	"fmt"
	"reflect"
)

type NotknowType struct {
	s1, s2, s3 string
}

func (n NotknowType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

type T struct {
	A int
	B string
}

var secret interface{} = NotknowType{"Ada", "Go", "Oberon"}

func SetStructValue() {
	val := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	fmt.Println(typ)
	kind := val.Kind()
	fmt.Println(kind)

	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("Field %d:%v\n", i, val.Field(i))
	}
	results := val.Method(0).Call(nil)
	fmt.Println(results)

	var inf = T{23, "skidoo"}

	elem := reflect.ValueOf(&inf).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fmt.Printf("%d:%s %s = %v \n", i, field.Type().Name(), field.Type(), field.Interface())
	}
	elem.Field(0).SetInt(88)
	elem.Field(1).SetString("B")
	fmt.Println(inf)
}
