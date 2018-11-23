package main

import (
	"fmt"
	"os"
)

// 01.go的改进版，使用range进行迭代
// $ go run Course_01/Chapter_00/02.go x y z
func main() {
	var s, sep = "", " "
	// 每次循环迭代，range产生一对值，索引及元素值，使用_接收表示忽略该项值
	// 这里对命令行参数对象作切片，因为与shell类似，命令行参数第一个元素指代脚本本身，在本例中为编译后的02.exe文件
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}
