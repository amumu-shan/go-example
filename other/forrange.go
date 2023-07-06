package other

import "time"

func ForRange() {
	//data := []field{{"one"}, {"two"}, {"three"}}

	//错误方式
	//for _, v := range data {
	//	go v.print()
	//}
	//
	//正确方式1
	//for _, v := range data {
	//	go func(f field) {
	//		f.print()
	//	}(v)
	//}
	//正确方式2
	//for _, v := range data {
	//	v := v
	//	go v.print()
	//}
	//正确方式3
	data2 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
