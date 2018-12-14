package main

import (
	"fmt"
)

type emptyInterface interface {
}

func main() {
	a := 1003
	b := "jay"
	var c emptyInterface
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			c = a
		} else {
			c = b
		}
		switch c.(type) {
		case string:
			fmt.Println("string: ", c)
		case int:
			fmt.Println("int: ", c)
		}
	}
}
