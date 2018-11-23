package calc

import (
	"fmt"
)

func Add(a int, b int, c chan int) {
	c <- a + b
}

func init() {
	fmt.Println("add")
}
