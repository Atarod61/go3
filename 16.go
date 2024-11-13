package main

import (
	"fmt"
	"reflect"
)

func main() {
	my_float_number := 4.5
	fmt.Println("my float number:", my_float_number)
	my_int_number := int(my_float_number)
	fmt.Println("my int number:", my_int_number)
	fmt.Println(reflect.TypeOf(my_int_number))
}
