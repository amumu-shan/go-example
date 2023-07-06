package point

import "fmt"

func nilInterface() {
	f := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil
		}
		return result
	}
	if res := f(-1); res != nil {
		fmt.Println("Good result: ", res)
		fmt.Printf("%T\n", res)
		fmt.Printf("%v\n", res)
	} else {
		fmt.Println("Bad result:", res)
	}
}
