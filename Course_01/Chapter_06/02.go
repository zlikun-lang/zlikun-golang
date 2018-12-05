package main

import "fmt"

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	// convert int to ByteCounter
	*bc += ByteCounter(len(p))
	return len(p), nil
}

// 测试的两个函数参数部分都是使用空接口实现
func main() {
	// 返回一个字符串
	text := fmt.Sprintf("name = %s, age = %d", "John", 17)
	// 输出一个字符串
	fmt.Printf("message => %s\n", text)

	// 上述两个函数底层由 fmt.Fprintf() 实现
	var bc ByteCounter
	// ByteCounter 实现了 io.Write 接口，所以可以直接在 fmt.Fprintf() 函数里使用
	fmt.Fprintf(&bc, "Hello, %s", "John")
	fmt.Println(bc) // 11
}
