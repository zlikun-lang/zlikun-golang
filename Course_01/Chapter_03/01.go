package main

import "fmt"

// 数组，固定长度，对应不定长度结构为slice（切片）
var a = [3]int{1, 2}
var b = [3]int{1, 2, 3}
var c [3]int

func main() {
	// [1 2 0]
	fmt.Println(a)
	// 1
	fmt.Println(a[0])
	// 0
	fmt.Println(a[len(a)-1])

	// 迭代
	for i, v := range a {
		// i = 0, v = 1
		// i = 1, v = 2
		// i = 2, v = 0
		fmt.Printf("i = %d, v = %d\n", i, v)
	}

	// 如果方括号里不是数字而是...，表示数组长度是根据初始值数量定义的
	d := [...]int{1, 2, 3, 4}
	// [1 2 3 4]
	fmt.Println(d)

	// 数组长度必须是常量表达式，因为数组长度是在编译阶段确定
	e := [3]int{1, 2, 3}
	// [1 2 3]
	fmt.Println(e)
}
