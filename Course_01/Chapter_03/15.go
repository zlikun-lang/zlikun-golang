package main

import "fmt"

// 点结构体
type Point struct {
	X, Y int
}

// 圆形结构体
type Circle struct {
	Point      // 复用Point结构体字段
	Radius int // 增加半径字段
}

// 轮子结构体
type Wheel struct {
	Circle     // 复用圆形结构体字段
	Spokes int // 增加辐条字段
}

func main() {

	// go提供了匿名成员数据类型（只声明数据类型，不指明名称）
	// Point类型嵌入到Circle类型中，Circle类型嵌入到Wheel类型中
	// 匿名成员包含了一个隐式名称，因此不能同时出现两个相同匿名成员（命名冲突）
	var wheel Wheel

	// 只是写法上的一种简化，本质仍是类型嵌套，这一点从输出结果可以看出
	wheel.X = 1
	wheel.Y = 2
	wheel.Radius = 3
	wheel.Spokes = 4

	// {{{1 2} 3} 4}
	fmt.Println(wheel)

	// 但该语法的副作用是不能直接使用字面值简单设置字段值，必须使用完整限定名
	wheel2 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 1,
				Y: 2,
			},
			Radius: 3,
		},
		Spokes: 4,
	}

	// {{{1 2} 3} 4}
	fmt.Println(wheel2)

	// 其等价写法
	// {{{1 2} 3} 4}
	fmt.Println(Wheel{Circle{Point{1, 2}, 3}, 4})
}
