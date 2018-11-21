package main

import (
	"fmt"
	"time"
)

var pipe chan int

func test(num1 int, num2 int) {
	for i := num1; i < num2; i++ {
		pipe <- i
	}
}

func main() {
	pipe = make(chan int, 3 )
	go test(0, 10)
	go test(20, 30)
	go test(50, 60)
	go test(101, 122)
	for {
		num := <-pipe
		fmt.Println(num)
		if len(pipe) <= 0 {
			break
		}
		time.Sleep(time.Second)
	}
}
