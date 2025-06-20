package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	start()
}

type Car struct {
	name string
}

type Cir struct {
	name string
}

func (c Car) start() {
	fmt.Println("Car: ", c.name)
}

func brandCheck(value Vehicle) {
	value.start()
}

func returnCar() Vehicle {
	return Car{name: "Toyota"}
}

func main() {
	var a Vehicle = Car{"volvo"} // initialization variable with interface data type
	b := Car{"Pajero"}           // normal intialization for struct, the type of data is Car

	fmt.Println(reflect.TypeOf(a)) // ouput : Car
	fmt.Println(reflect.TypeOf(b)) // ouput : Car
	fmt.Println()

	fmt.Println("VAR a")
	a.start()      // Volvo
	brandCheck(a)  // volvo
	fmt.Println(a) // {volvo}
	fmt.Println()

	var c Vehicle = returnCar()

	fmt.Println("VAR c")
	c.start()                      // Toyota
	brandCheck(c)                  // Toyota
	fmt.Println(c)                 // {Toyota}
	fmt.Println(reflect.TypeOf(c)) // ouput : Car

}

// data type of a & b is Car even though a is declared as Vehicle, compiler recognizes a as Car

// data type of c is Car even though c is declared as Vehicle, compiler recognizes c as Car
