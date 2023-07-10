package _struct

type Any interface {
}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

// 高阶函数，实际上就是把函数作为定义所需的方法的参数

// Process 函数，接收作用于每一辆car的f函数作参数
func (cs Cars) Process(f func(_ *Car)) {
	for _, c := range cs {
		f(c)
	}
}

func (cs Cars) FindAll(f func(_ *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})

	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
