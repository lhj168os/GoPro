package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func add(num1, num2 string) (result string, err error) {
	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}
	sNum1 := strings.IndexFunc(num1, f)
	sNum2 := strings.IndexFunc(num2, f)

	if sNum1 != -1 || sNum2 != -1 {
		err = errors.New("Input number error!")
		return
	}

	index1 := len(num1) - 1
	index2 := len(num2) - 1
	var carry int
	for index1 >= 0 && index2 >= 0 {
		c1 := num1[index1] - '0'
		c2 := num2[index2] - '0'
		sum := int(c1) + int(c2) + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := num1[index1] - '0'
		sum := int(c1) + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c2 := num2[index2] - '0'
		sum := int(c2) + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if carry == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	strs := strings.Split(string(str), "+")
	if len(strs) != 2 {
		fmt.Println("Input Number error!")
		return
	}
	num1 := strs[0]
	num2 := strs[1]
	result, err := add(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
