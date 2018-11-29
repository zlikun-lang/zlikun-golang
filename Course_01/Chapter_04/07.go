package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {

	// 函数是第一类值，可以赋值给变量、从函数中返回等
	// 赋值
	t := add

	// 测试
	fmt.Println(t(1, 2))

}
