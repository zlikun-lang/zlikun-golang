package main

import "fmt"

// 最大公约数
func gcd(x, y int) int {
	for y != 0 {
		// 元组赋值语句
		x, y = y, x%y
	}
	return x
}

// 斐波那契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {

	// 120, 36 -> 12
	fmt.Printf("%d, %d -> %d\n", 120, 36, gcd(120, 36))
	// 256, 64 -> 64
	fmt.Printf("%d, %d -> %d\n", 256, 64, gcd(256, 64))

	// fib(5) -> 5
	fmt.Printf("fib(%d) -> %d\n", 5, fib(5))
	// fib(7) -> 13
	fmt.Printf("fib(%d) -> %d\n", 7, fib(7))

}
