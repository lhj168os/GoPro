package main

import (
	"fmt"
)

func factorial(num int) int {
	if num > 1 {
		return factorial(num-1) + num
	}
	return num
}

func main() {
	result := factorial(10)
	fmt.Println(result)
}
 