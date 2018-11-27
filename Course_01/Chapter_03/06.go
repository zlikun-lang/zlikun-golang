package main

import "fmt"

func main() {
	// 使用 append() 内置函数向 slice 追加元素
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	// ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	fmt.Printf("%q\n", runes)

	// 上述仅用于演示，实际可以用内置转换逻辑完成
	runes = []rune("Hello, World")
	// ['H' 'e' 'l' 'l' 'o' ',' ' ' 'W' 'o' 'r' 'l' 'd']
	fmt.Printf("%q\n", runes)

}
