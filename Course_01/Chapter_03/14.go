package main

import "fmt"

// 点结构体
type Point struct {
	X, Y int
}

// 圆形结构体
type Circle struct {
	Center Point // 复用Point结构体字段
	Radius int   // 增加半径字段
}

// 轮子结构体
type Wheel struct {
	Circle Circle // 复用圆形结构体字段
	Spokes int    // 增加辐条字段
}

func main() {

	// 在使用Wheel结构体时，其内部的Circle结构体是命名的，所以访问时需要使用其完整限定名
	var wheel Wheel

	wheel.Circle.Center.X = 1
	wheel.Circle.Center.Y = 2
	wheel.Circle.Radius = 3
	wheel.Spokes = 4

	// {{{1 2} 3} 4}
	fmt.Println(wheel)

}
