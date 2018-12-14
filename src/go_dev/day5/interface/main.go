package main

import (
	"fmt"
)

type itf interface {
	printInfo()
	String() string
}

type people struct {
	name string
	age  int
	sex  int
}

type courseScore struct {
	ch int
	ma int
	en int
	pe int
}

type stu struct {
	schoolId string
	people
	courseScore
}

func (p stu) printInfo() {
	fmt.Println("ID: ", p.schoolId)
	fmt.Println("name: ", p.name)
	fmt.Println("age: ", p.age)
	fmt.Println("sex: ", p.sex)
	fmt.Println("score: ", p.ch, "(ch)", p.ma, "(ma)", p.en, "(en)", p.pe, "(pe)")
}

func (p stu) String() string {
	str := fmt.Sprintf("ID:%s  Name:%s  Age:%d  Sex:%d  CH:%d  MA:%d  EN:%d  PE:%d", p.schoolId, p.name, p.age, p.sex, p.ch, p.ma, p.en, p.pe)
	return str
}

func (s *stu) initInfo() {
	s.schoolId = "140201011034"
	s.name = "jay"
	s.age = 25
	s.sex = 1
	s.ch = 80
	s.ma = 90
	s.en = 78
	s.pe = 88
}
func main() {
	var s stu
	s.initInfo()
	fmt.Printf("【StudentInfo】  %s\n", s)
	var t itf
	t = &s
	t.printInfo()
}
