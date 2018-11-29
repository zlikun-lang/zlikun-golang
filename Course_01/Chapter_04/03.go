package main

import "fmt"

// 递归
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// 1 1 2 3 5
	fmt.Println(fib(5)) // 5
	// 1 1 2 3 5 8
	fmt.Println(fib(6)) // 8
	// 1 1 2 3 5 8 13
	fmt.Println(fib(7)) // 13
}
