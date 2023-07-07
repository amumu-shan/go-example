package _interface

import (
	"fmt"
	"testing"
)

type Any interface{}

func TestMethodSet(t *testing.T) {

	data := []int{1, 2, 3, 4, 5}
	var val Any
	switch t := val.(type) {
	case int:
		fmt.Println(t)
	}

	lst := List(data)
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	//CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID: Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
