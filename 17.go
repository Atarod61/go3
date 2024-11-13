package main

import "fmt"

func main() {

	fmt.Println("start...")
	var number int
	fmt.Scan(&number)
	if number > 0 {

		fmt.Printf("%d is positive", number)

	} else if number == 0 {

		fmt.Printf("%d is zero", number)

	} else {

		fmt.Printf("%d is negative", number)
	}

}
