package main

import (
	"fmt"
	"go_dev/day1/goroute_example2/calc"
)

var a, b int

func main() {
	pipe := make(chan int, 1)
	go calc.Add(a, b, pipe)
	go calc.Sub(a, b, pipe)
	num := <-pipe
	fmt.Println(num)
	num = <-pipe
	fmt.Println(num)
}

func init() {
	a = 200
	b = 123
}
