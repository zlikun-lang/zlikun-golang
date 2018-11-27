package main

import "fmt"

// slice（切片）
func main() {

	numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(numbers)
	// [1 2 3 4 5 6 7 8 9]
	fmt.Println(numbers[1:])
	// [0 1 2 3 4]
	fmt.Println(numbers[:5])
	// [2 3 4]
	fmt.Println(numbers[2:5])

	// 内置函数make创建一个指定长度的slice结构
	a := make([]string, 4)
	// [   ]
	fmt.Println(a)
	// 4
	fmt.Println(len(a))

}
