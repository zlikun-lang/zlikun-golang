package main

import "fmt"

type Point struct{ x, y float64 }

// 基于指针声明方法（避免大参数拷贝开销）
// 在使用时必须使用括号将指针变量括起来，以避免误解
func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func main() {
	// 定义一个变量
	p := Point{1, 2}
	// 使用变量指针调用方法
	(&p).ScaleBy(0.2)
	// 输出变量值
	// {0.2 0.4}
	fmt.Println(p)
	// 如果p只是一个变量，那么Go对此做了优化，可以直接用p来调用指针方法（实际上是编译器隐式的使用指针来调用）
	p.ScaleBy(0.5)
	// {0.1 0.2}
	fmt.Println(p)

}
