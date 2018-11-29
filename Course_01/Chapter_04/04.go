package main

import "fmt"

// 多返回值
func reverse(x int, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(reverse(1, 2)) // 2 1
	fmt.Println(reverse(2, 1)) // 1 2
}
