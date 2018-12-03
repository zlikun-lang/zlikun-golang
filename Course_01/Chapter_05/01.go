package main

import (
	"fmt"
	"math"
)

type Point struct{ x, y float64 }

// 定义方法与函数不同的时，方法名前面放了一个变量，表示为该变量类型添加方法
// 附加参数p叫做方法接收器，类似其它语言中的this或self等，在Go里可以任意指定该参数名，一般用类型首字母表示
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// 测试
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	// 5
	fmt.Println(p.Distance(q))
}
