package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ x, y float64 }

// 通过嵌入结构体来扩展类型
type ColorPoint struct {
	Point
	Color color.RGBA
}

// 为Point增加一个方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// 测试，在ColorPoint中嵌入了Point结构体，所以隐式的在该结构体中包含了Point的方法
// 但ColorPoint与Point并非is-a关系，所以调用方法时，必须显示使用Point类型来调用
func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	cp := ColorPoint{Point{1, 2}, red}
	// ColorPoint可以使用Distance方法，但参数必须是Point类型
	// 2.8284271247461903
	fmt.Println(cp.Distance(Point{3, 4}))
	// 不能直接使用ColorPoint类型，需要获取其内嵌的Point类型
	target := ColorPoint{Point{3, 4}, blue}
	// 使用ColorPoint中内嵌的Point类型参数，本质是ColorPoint与Point的关系为has-a关系
	// 2.8284271247461903
	fmt.Println(cp.Distance(target.Point))
}
