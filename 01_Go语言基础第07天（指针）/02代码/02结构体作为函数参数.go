package main

import "fmt"

type Student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

//结构体变量作为函数参数
func test(stu Student) {
	stu.name = "野猪佩奇"
	// fmt.Println(stu)

}

func main() {
	stu := Student{101, "喜羊羊", 6, "男", 100, "羊村"}
	//值传递
	test(stu)

	fmt.Println(stu)
}
func test1(m map[int]Student) {
	//指针不能直接.成员
	//m[102].name = "威震天"//err

	stu := m[102]
	stu.name = "威震天"
	m[102] = stu
	//fmt.Println(stu)
	//fmt.Printf("%T\n", stu)
	//fmt.Println(m[102])
}
func main2222() {

	//将结构体作为map中的value
	m := make(map[int]Student)

	//map中的数据不建议排序操作
	m[101] = Student{101, "喜羊羊101", 6, "男", 101, "羊村11"}
	m[102] = Student{102, "喜羊羊102", 6, "男", 102, "羊村22"}

	//将map作为函数参数
	test1(m)
	fmt.Println(m)

}
