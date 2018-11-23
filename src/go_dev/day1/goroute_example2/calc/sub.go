package calc

import (
	"fmt"
)

func Sub(a int, b int, c chan int) {
	c <- a - b
}

func init() {
	fmt.Println("sub")
}

func init() {
	fmt.Println("sub2")
}
