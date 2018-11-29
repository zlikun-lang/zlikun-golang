package main

import (
	"fmt"
	"math"
)

/* ----------------------------------------------------
func <function_name>(<parameter_list>) <return_type> {

}
---------------------------------------------------- */

// 根据勾股定理计算斜边边长
// 前面的参数x复用了后面y的类型
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func main() {
	fmt.Println(hypot(3, 4)) // 5

	fmt.Printf("%T\n", add)   // func(int, int) int
	fmt.Printf("%T\n", sub)   // func(int, int) int
	fmt.Printf("%T\n", first) // func(int, int) int
	fmt.Printf("%T\n", zero)  // func(int, int) int
}
