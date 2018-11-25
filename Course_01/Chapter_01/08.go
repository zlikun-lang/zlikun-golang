package main

import "fmt"

func main() {
	// 直接赋值一个数组
	medals := []string{"gold", "silver", "bronze"}
	// 单项赋值
	medals[0] = "gold"
	medals[1] = "silver"
	medals[2] = "bronze"

	fmt.Println(medals)
}
