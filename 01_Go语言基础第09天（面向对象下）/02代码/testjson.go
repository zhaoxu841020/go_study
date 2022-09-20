package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     string
	Gender string
	Name   string
	Sno    string
}

func main() {
	var s1 = Student{
		ID:     "12",
		Gender: "男",
		Name:   "李四",
		Sno:    "s001",
	}
	// 结构体转换成Json（返回的是byte类型的切片）
	jsonByte, err := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

	var p2 Student
	err = json.Unmarshal(jsonByte, &p2)

	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)

	var p3 Student
	err = json.Unmarshal(jsonByte, &p3)

	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p4:%#v\n", p3)

	var str = `{"ID":"12","Gender":"男","Name":"李四","Sno":"s001"}`
	var s2 Student

	err2 := json.Unmarshal([]byte(str), &s2)

	if err2 != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err2)
		return
	}
	fmt.Printf("s2:%#v\n", s2)

	// Json字符串转换成结构体
	var str11 = `{"ID":"12","Gender":"男","Name":"李四","Sno":"s001"}`
	var s3 = Student{}
	// 第一个是需要传入byte类型的数据，第二参数需要传入转换的地址
	err3 := json.Unmarshal([]byte(str11), &s3)

	if err3 != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err3)
		return
	}
	fmt.Printf("s3:%#v\n", s3)
}
