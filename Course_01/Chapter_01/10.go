package main

import (
	"fmt"
	// 这里是自定义包的导入路径
	"zlikun-golang/Course_01/Chapter_01/pkg"
)

func main() {
	// 测试引入自定义包
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
}
