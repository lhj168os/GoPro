package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "Hello world, Hello you, hello me"
	fmt.Printf("str = \"%s\"\n", str)

	//比较两个字符串ASCII码的大小,1:>   0:=   -1:<
	cmp := strings.Compare(str, "Hello world")
	fmt.Println("strings.Compare(str, \"Hello world\") = ", cmp)

	//判断字符串是否包含另一个字符串
	result := strings.Contains(str, "world")
	fmt.Println("strings.Contains(str, \"world\") = ", result)

	//判断一个字符串是否包含另一个字符串中的任意一个字符
	result = strings.ContainsAny(str, "HwHZ")
	fmt.Println("strings.ContainsAny(str, \"HwHZ\") = ", result)

	//判断一个字符串中是否包含这个ASCII码对应的字符
	result = strings.ContainsRune(str, 33)
	fmt.Println("strings.ContainsRune(str, 33) = ", result)

	//计算一个字符串中出现另一个字符串的次数
	sum := strings.Count(str, "ll")
	fmt.Println("strings.Count(str, \"ll\") = ", sum)

	//忽略大小写两字符串是否相等
	result = strings.EqualFold(str, "hello world, hello you, hello me")
	fmt.Println("strings.EqualFold(str, \"hello world hello you hello me\") = ", result)

	//把一个字符串在有空白符的地方切开，分成多个字符串
	strs := strings.Fields(str)
	fmt.Println("strings.Fields(str) = ", strs)

	//对每个字符的ASCII码进行判断，将其传入函数f中，若返回的值为true则去掉，然后返回剩下的字符串的切片
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	strs = strings.FieldsFunc(str, f)
	fmt.Println("strings.FieldsFunc(str,f) = ", strs)

	//判断字符串是否以“Hell”开头
	result = strings.HasPrefix(str, "Hell")
	fmt.Println("strings.HasPrefix(str, \"Hell\") = ", result)

	//判断字符串是否以“me”结尾
	result = strings.HasSuffix(str, "me")
	fmt.Println("strings.HasSuffix(str, \"me\") = ", result)

	//查询一个字符串在str字符串中出现的位置，若str字符串中不含该字符串则返回-1
	idx := strings.Index(str, "lo")
	fmt.Println("strings.Index(str, \"lo\") = ", idx)

	//yggj字符串中任意一个字符在字符串str中首次出现的位置，都不出现返回-1
	idx = strings.IndexAny(str, "yggj")
	fmt.Println("strings.IndexAny(str, \"yggj\") = ", idx)

	//字符在字符串中首次出现的位置，不出现为-1
	idx = strings.IndexByte(str, 'y')
	fmt.Println("strings.IndexByte(str, 'y') = ", idx)

	
}
