package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result, "is", reflect.TypeOf(result))
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt, "is", reflect.TypeOf(resultInt))
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary, "is", reflect.TypeOf(binary))

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt, "is", reflect.TypeOf(stringInt))
}
