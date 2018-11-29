package main

import "fmt"

// 可变参数，以 ... 表示（与Java一致）
// 与Java相同的是，可变参数只能是参数列表最后一个参数
func add(args ...int) int {
	var total int
	for _, val := range args {
		total += val
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4)) // 10

	// 如果传入参数是一个切片，可以在切片参数后加...传参
	fmt.Println(add([]int{1, 2, 3, 4}...)) // 10
}
