package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (p *Person) PrintInfo() {
	fmt.Print(" 姓名: ", p.name)
	fmt.Print(" 年龄: ", p.age)
	fmt.Print(" 性别: ", p.sex)
	fmt.Println()
}

func (p *Person) SetInfo(name string, age int, sex string) {
	p.name = name
	p.age = age
	p.sex = sex
}

func main() {
	var p = Person{
		"张三",
		18,
		"女",
	}

	p.PrintInfo()

	p.SetInfo("111", 12, "222")

	p.PrintInfo()
}
