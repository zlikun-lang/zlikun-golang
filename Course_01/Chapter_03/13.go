package main

import "fmt"

type Point struct{ x, y int }

func main() {
	m := Point{1, 2}
	n := Point{1, 2}

	// 结构体比较会比较内部所有字段
	fmt.Println(m == n)     // true
	fmt.Println(m.x == n.x) // true
	fmt.Println(m.y == n.y) // true
}
