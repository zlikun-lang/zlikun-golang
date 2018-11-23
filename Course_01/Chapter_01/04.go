package main

import "fmt"

// 函数f返回一个int类型变量指针
func f() *int {
	v := 1
	return &v
}

// 每次返回的针对都不相同
func main() {
	fmt.Println(f()) // 0xc00004a080
	fmt.Println(f()) // 0xc00004a088
}
