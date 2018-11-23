package main

import (
	"fmt"
	"os"
	"strings"
)

// $ go run Course_01/Chapter_00/03.go x y z
func main() {
	// 比起使用循环迭代，使用strings包下的Join会简单很多
	fmt.Println(strings.Join(os.Args[1:], " "))
	// 如果只是为了调试，直接打印也是可以的
	fmt.Println(os.Args[1:])
}
