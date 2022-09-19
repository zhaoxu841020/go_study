package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
	gorm.Model
	Name string
}

func main() {

	var testUser = new(User)

	testUser.Name = "zhaoxu"

	fmt.Println(testUser.Name)

	// db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// // 自动迁移
	// db.AutoMigrate(&UserInfo{})

	// u3 := UserInfo{11, "22", "dd", "dd"}
	// u1 := UserInfo{1, "七米", "男", "篮球"}
	// u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// // 创建记录
	// db.Create(&u1)
	// db.Create(&u2)
	// db.Create(&u3)
	// // 查询
	// var u = new(UserInfo)
	// db.First(u)
	// fmt.Printf("%#v\n", u)

	// var uu UserInfo
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)

	// // 更新
	// db.Model(&u).Update("hobby", "双色球")
	// // 删除
	// db.Delete(&u)
}
