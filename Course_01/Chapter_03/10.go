package main

import "fmt"

// go 中并未提供 set 数据结构类型，但 map 的键本身相当于 set 结构，因此可以模拟 set
func main() {

	seen := make(map[string]bool)

	// 添加重复元素，后者将覆盖前者
	seen["x"] = true
	seen["x"] = true
	seen["y"] = true

	// 查询长度
	fmt.Println(len(seen)) // 2

}
