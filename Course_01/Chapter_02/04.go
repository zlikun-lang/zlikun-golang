package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "Hello World !"

	// 字符串长度
	fmt.Println(len(s))

	// 72 111，ASCII转换为整数（'h' and 'w'）
	fmt.Println(s[0], s[7])

	// 字符串切片
	// Hello
	fmt.Println(s[0:5])
	// Hello
	fmt.Println(s[:5])
	// World !
	fmt.Println(s[6:])
	// Hello World !
	fmt.Println(s[:])
	// goodbye World !
	fmt.Println("Goodbye" + s[5:])

	// 字符串与数字之间的转换
	x := 123
	y := fmt.Sprintf("%d", x)       // 格式化字符串
	fmt.Println(y, strconv.Itoa(x)) // 123 123

}
