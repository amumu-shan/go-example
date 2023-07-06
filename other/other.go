package other

import (
	"fmt"
	"go-example/routine"
	"strings"
	"time"
)

var a string

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func Start() {
	//f2()
	//x := min(1, 3, 2, 0)
	//fmt.Printf("The minimum is: %d\n", x)
	//slice := []int{7, 9, 3, 5, 1}
	//x = min(slice...)
	//fmt.Printf("The minimum in the slice is: %d", x)
	//a1()
	//func_closure()
	//fibonacci_closure()
	//var digitRegexp = regexp.MustCompile("[0-9]+")
	//fileBytes, _ := ioutil.ReadFile("a.text")
	//b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	//c := make([]byte, 0)
	//for _, bytes := range b {
	//	c = append(c, bytes...)
	//}
	//fmt.Println(c)
	//
	//m := matrix.NewMatrix()
	//fmt.Println(m)

	//y := new(Bar)
	//(*y).thingOne = "hello"
	//(*y).thingTwo = 1
	//x := make(Foo)
	//x["x"] = "goodbye"
	//x["y"] = "world"
	//u := new(Foo)
	//(*u)["x"] = "goodbye"
	//(*u)["y"] = "world"

	//mm := map[string]int{"one": 11, "two": 2, "three": 13, "four": 4}
	//for k, v := range mm {
	//	fmt.Println(k, v)
	//}
	//data := []int{1, 2, 3}
	//i := 0
	//i++
	//fmt.Println("data:", data[i])
	//
	//
	//
	//
	//char := "♥"
	//fmt.Println(len(char))
	//fmt.Println(utf8.RuneCountInString(char))
	//log.Fatal("fatal level log:log entry")
	//log.Println("Nomal level log")

	//var ch chan int // 未初始化，值为 nil
	//for i := 0; i < 3; i++ {
	//	go func(i int) {
	//		ch <- i
	//	}(i)
	//}
	//
	//fmt.Println("Result: ", <-ch)
	//time.Sleep(2 * time.Second)

	//routine.Start()
	//routine.UnCacheChan()
	routine.SingleChan()
}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func fibonacci_closure() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// println(f(), f(), f(), f(), f())
	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
}
func func_closure() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		fmt.Println("x:", x)
		return x
	}
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a2() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a2()
}
func a1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
func f2() {
	var x uint8 = 15
	var y uint8 = 4
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&^y)
	fmt.Printf("%08b\n", ^y)
	/*输出：
	a: 00001100
	b：00000100
	c: 8   二进制：00001000
	*/
	var str string = "hello,hello,hello"
	fmt.Println(strings.Replace(str, "hello", "ps", 2))

	fmt.Println("--------")

	t := time.Now()
	fmt.Println(t)
	s := t.Format("20060102")
	fmt.Println(s)

	i := 1
	switch i {
	case 1:
		fallthrough
	case 2:
		fmt.Println("2->", i)
	case 3:
		fmt.Println("3->", i)
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	for i = 1; i <= 10; i++ {
		fmt.Println(strings.Repeat("G", i))
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}
	s1 := ""
	for s1 != "aaaaa" {
		fmt.Println("Value of s:", s1)
		s1 = s1 + "a"
	}

	for i, s = 0, "G"; i < 6; i++ {
		fmt.Println(s)
		s = s + "G"
	}
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}
