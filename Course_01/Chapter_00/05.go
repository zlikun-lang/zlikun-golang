package main

import (
	"bufio"
	"fmt"
	"os"
)

// 读取文件
func main() {
	// 读取当前文件
	f, err := os.Open("Course_01/Chapter_00/05.go")
	if err != nil {
		// open ./hello.go: The system cannot find the file specified.
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	// 读取文件
	input := bufio.NewScanner(f)
	// 按行遍历
	for input.Scan() {
		fmt.Println(input.Text())
	}
	f.Close()
}
