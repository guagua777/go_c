package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	ID   int
	Name string
}

func main4() {
	u := User{ID:1, Name:"test"}
	fmt.Println(unsafe.Sizeof(u)) // 大小>0，包含所有字段 + 内存对齐填充
	fmt.Println(u)


	u2 := new(User)
	// u 是 *User
	u2.ID = 3
	u2.Name = "Bob"

}
