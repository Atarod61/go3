package main

import "fmt"

func main() {

	var name string
	var surname string
	var age int

	fmt.Println("start...")
	fmt.Scan(&name, &surname, &age)

	fmt.Println("your name is", name, surname, "and you are", age, "years old")
	fmt.Println("end...")

}
