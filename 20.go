package main

import "fmt"

func main() {

	MyNumbers := [5]int{18, 12, 15, 17, 20}

	for index, item := range MyNumbers {

		fmt.Printf("MyNumbers[%d] = %d\n", index, item)

	}

}
