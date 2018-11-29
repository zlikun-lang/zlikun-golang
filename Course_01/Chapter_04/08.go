package main

import "fmt"

// 匿名函数，函数handle接收一个匿名函数
func handle(fn func(int) int, x int) string {
	return fmt.Sprintf("%0x", fn(x))
}

func main() {
	// 3c，匿名函数将参数除以2，handle函数将其转换为16进制字符串
	fmt.Println(handle(func(x int) int { return x / 2 }, 120))
}
