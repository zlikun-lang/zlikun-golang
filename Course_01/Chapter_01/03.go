package main

import "fmt"

func main() {
	x := 1  // 简短声明一个变量
	p := &x // &x表示x的指针
	// x = 1, p = 0xc00000a0b8
	fmt.Printf("x = %d, p = %v\n", x, p)
	// x = 1, *p = 1，*p表示获取指针p的值
	fmt.Printf("x = %d, *p = %d\n", x, *p)

	// 通过针对修改变量，原变量也会被修改（区别于副本传值）
	*p = 4
	// x = 4, *p = 4
	fmt.Printf("x = %d, *p = %d\n", x, *p)

}
