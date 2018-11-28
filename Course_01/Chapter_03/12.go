package main

import (
	"fmt"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Address string
}

// 结构体面值
func main() {

	// 以面值写法，要求按字段顺序且包含完整字段列表
	// 以名值写法，则无上述要求
	kevin := User{ID: 1, Name: "Kevin", Address: "NewYork"}

	// {1 Kevin 0 NewYork}
	fmt.Println(kevin)

}
