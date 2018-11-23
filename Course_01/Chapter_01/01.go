package main

import "fmt"

// 声明一个变量：var $var_name $var_type = $var_value
var name string = "john"

// 变量类型可以省略，将根据初始化推导类型
var age = 17

// 变量类型与初始值可以省略其中一个，无初始值默认使用零值初始化
var score int32

func main() {
	// name = john, age = 17, score = 0
	fmt.Printf("name = %s, age = %d, score = %d\n", name, age, score)

	// 同时声明多个值（指定类型）
	var i, j, k int32
	// i + j + k = 0
	fmt.Printf("i + j + k = %d\n", i+j+k)

	// 同时声明多个值（指定初始值）
	var x, y, z = 'x', true, "hello"
	// x = x, y = true, z = hello
	fmt.Printf("x = %c, y = %t, z = %s\n", x, y, z)
}
