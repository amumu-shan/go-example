package point

import (
	"fmt"
	"testing"
)

func TestReceiver(t *testing.T) {
	d1 := data{"one"}
	d1.print()
	//var in printer = data{"two"} //不可寻址的，通过interface引用的变量

	//m := map[string]data{
	//	"x": data{"three"},
	//}
	//m["x"].print()  //不可寻址的，map类型的元素

	//m["x"].name = "ff" //map的值是struct类型，则无法直接更新该struct的单个字段

	m2 := map[string]data{
		"x": {"Tome"},
	}
	//m2["x"].name = "Jerry" //错误，map中的元素是不可寻址的

	//方法1 局部变量
	r := m2["x"]
	r.name = "Jerry"
	m2["x"] = r
	fmt.Println(m2)

	//方法2 使用指向元素map的指针
	m3 := map[string]*data{
		"x": {"Tome"},
	}
	m3["x"].name = "Jerry"
	fmt.Println(m3["x"])
	//m3["y"].name = "Jerry" //必须是map中存在的key才能赋值
}
