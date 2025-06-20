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

func (c Car) start() {
	fmt.Sprintf("Car: %s", c.name)
}

func helloWorld(value Vehicle) {
	value.start()
}

func main() {
	var a Vehicle = Car{"volvo"} // initialization variable with interface data type
	b := Car{"Pajero"} // normal intialization for struct, the type of data is Car

	fmt.Println(reflect.TypeOf(a)) // ouput : Car
	fmt.Println(reflect.TypeOf(b)) // ouput : Car 

	helloWorld(a)
	fmt.Println(a)
}


// data type of a & b is Car even though a is declared as Vehicle, compiler recognizes a as Car