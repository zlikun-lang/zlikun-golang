package main

import "fmt"

func main() {

	// 取模运算符号总是与被取模数符号一致
	fmt.Println(-5 % -3) // -2
	fmt.Println(5 % -3)  // -2

	// 除法运算符的行为依赖于操作数是否全为整数
	fmt.Println(5.0 / 4.0) // 1.25
	fmt.Println(5.0 / 4)   // 1.25
	fmt.Println(5 / 4)     // 1，全为整数时，向0方向截断为整数（即：舍弃小数）

	// 进制格式化
	fmt.Printf("十进制：%d\n", 192)  // 192
	fmt.Printf("八进制：%o\n", 192)  // 300
	fmt.Printf("十六进制：%x\n", 192) //  c0
	fmt.Printf("十六进制：%X\n", 192) //  C0

}
