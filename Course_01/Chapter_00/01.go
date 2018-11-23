package main

import (
	"fmt"
	"os"
)

// 以如下命令运行程序，输出命令行参数
// $ go run Course_01/Chapter_00/01.go x y z
func main() {
	// 声明s, sep两个string类型变量，没有显示初始化，则隐式使用零值初始化，数值为0，字符串为""
	var s, sep string
	// os.Args 用于获取命令行参数，len与python中同名函数一致，用于获取长度
	// i := 1，是一个短变量声明，等价：var i = 1
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
