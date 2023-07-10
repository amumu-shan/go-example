package _struct

import (
	"fmt"
	"testing"
)

func TestCars_FindAll(t *testing.T) {

	bmw := &Car{Model: "轿车", Manufacturer: "BMW", BuildYear: 2020}
	ft := &Car{Model: "轿车", Manufacturer: "FT", BuildYear: 2021}
	byd := &Car{Model: "SUV", Manufacturer: "BYD", BuildYear: 2022}
	allCar := Cars([]*Car{bmw, ft, byd})

	allBMW := allCar.FindAll(func(car *Car) bool {
		return car.Manufacturer == "BMW" && car.BuildYear > 2010
	})
	fmt.Println(*(allBMW)[0])

	antes := allCar.Map(func(car *Car) Any {
		return car.Manufacturer == "BMW" && car.BuildYear > 2010
	})
	fmt.Println(antes)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "BYD"}
	appenderFunc, cars := MakeSortedAppender(manufacturers)
	allCar.Process(appenderFunc)
	fmt.Println("Map sortedCars:", cars)
	BMWCount := len(cars["BMW"])
	fmt.Println("we have ", BMWCount, " BMWs")
}
