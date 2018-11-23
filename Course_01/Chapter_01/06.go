package main

// 声明普通变量
var x, p = 1, &x

// 声明数组
var arr [4]int

// 声明结构体
type Person struct {
	name string
	age  int32
}

var person Person

// 赋值
func main() {
	// 命名变量赋值
	x = 4
	// 通过指针间接赋值（这里取巧了，p指针刚好是x的指针）
	*p = 7
	// 数组元素赋值
	arr[0] = 120
	// 结构体赋值
	person.name = "john"

	v := 1
	v++
	v--
	// x = v++ // v++是语句而非表达式，所以不能用于赋值语句
}
