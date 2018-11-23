package main

import (
	"bufio"
	"fmt"
	"os"
)

// 模拟shell中的uniq命令，找出相邻的重复行
func main() {
	// 声明一个名为counts的map变量
	counts := make(map[string]int)
	// 标准输入
	input := bufio.NewScanner(os.Stdin)
	// 每次回车将读入输入的行
	for input.Scan() {
		// 当输入exit时，结束输入
		if input.Text() == "exit" {
			break
		}
		// 对输入文本计数
		counts[input.Text()]++
	}
	// 遍历counts，找出计数大于1的即为重复行
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
