package main

import (
	"fmt"
)

func main() {
	i, n, sum := 1, 10, 0
	for i <= n {
		sum += i
		i += 2
		fmt.Println(sum)
	}
}
