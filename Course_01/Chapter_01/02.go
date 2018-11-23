package main

import (
	"fmt"
	"math/rand"
)

// 简短变量声明
func main() {
	x := 1
	y := rand.Float64() * 3.0
	// x = 1, y = 1.813981
	fmt.Printf("x = %d, y = %f\n", x, y)

	// 简短声明语句，应至少有一个新的变量，否则编译错误，其中声明过的变量等价与简单的赋值语句
	x, z := 4, "good"
	fmt.Println(z)
}
