package main

import (
	"fmt"
)

func formatFileName(suffix string)func(string)string{
	return func(fileName string)string{
		return fileName+suffix
	}
}

func main() {
	f := formatFileName(".txt")
	fmt.Println(f("HelloWorld"))
}
