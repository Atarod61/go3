package main

import (
	"fmt"
)

func main() {

	i := 8

	fmt.Println("Start Loop...")
	for i <= 8 {

		if i == 5 {
			break
		}
		fmt.Println(i)
		i += 1
	}
	fmt.Println("End Loop...")
}
