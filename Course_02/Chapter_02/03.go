// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.3.md
package main

import "fmt"

// 1. 分支语句
func fn02A(x int) {
	// if ... else if ... else ...
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equals 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// 在Go中，if语句内部甚至可以声明一个变量，其作用范围仅限if语句块本身（包含else部分）
	if y := x + 4; y > 10 {
		fmt.Println("x 加权后大于10")
	} else {
		fmt.Println("x 加权后不大于10")
	}

	// fmt.Println(y)	// 这里不能访问变量 y，超出了其作用范围

	// switch 语句
	switch x & 3 {
	case 0:
		fmt.Println("switch x & 3 case 0 .")
		// 在GO中，每个case语句隐式带了break语句，如果希望向下执行要显示执行下面语句
		fallthrough
	case 1:
		fmt.Println("switch x & 3 case 1 .")
	case 2:
		fmt.Println("switch x & 3 case 2 .")
	default:
		fmt.Println("no ..................")
	}
}

// 2. goto 跳转，谨慎使用（略）

// 3. 循环语句
func fn02B() {
	// for expression1; expression2; expression3 { $statements }
	// 求从1加到10计数之和
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i + 1
	}
	fmt.Printf("1 + ... + 10 = %v\n", sum)
}

func fn02C() {
	// for循环语句表达式1、表达式3都可以省略，进而把;也省略，此时其相当于Java中的while语句
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 配置 range 实现迭代
func fn02D(data map[string]int) {
	// 迭代map
	for _, v := range data {
		fmt.Println(v)
	}
}

func fn02E(s []int) {
	// 迭代slice
	for i, v := range s {
		fmt.Println(i, v)
	}
}

// 4. 函数，这里略过，后面专题强化学习
//	a. 值传参和指针传参
// 		1) *type，表示指针类型
// 		2) &value，表示一个值的指针，与*type对应
// 		3) *value，表示指针变量指向的值对象
// 	b. defer机制
// 		1) defer后的语句在函数退出时一定会执行，所以可以用于关闭资源，一般打开资源后紧接着就用defer关闭资源（函数退出前才会执行）
//		2) 同时存在多个defer，采用后进先出的方式（栈结构）执行（从下往上执行）
//		3) 需要注意的是defer语句只有在函数退出时才会执行，所以在循环体内打开资源，则defer语句并不能很好的使用

// 5. 异常机制，这里简单了解，后续专题学习
// panic 是一个内建函数，可以中断原有控制流程，进入一个panic状态，这个相当于Java中的异常，但并不完全一样
// recover 是一个内建函数，可以让进入panic的goroutine恢复过来，只在延迟函数中有用

// 6. 保留函数 main 和 init

// 7. 导入语句
// 	1) 导入时使用绝对路径（相对于$GOPATH/src目录）
//	2) 导入时使用相对路径，相对于当前目录（不推荐使用）
//	3) 点操作：import (. "fmt")，表示使用时省略包名，如：Println()，省略了fmt.部分
// 	4) 别名操作：import (f "fmt")，表示使用别名f代替fmt，如：f.Println()
// 	5) _操作，表示导入该包，但不直接使用包里的函数，但包时会调用其包含的init()函数（如果有的话）

func main() {
	fn02A(1)
	fn02A(10)
	fn02A(12)
	fn02B() // 1 + ... + 10 = 55
	fn02C() // 1024
	fn02D(map[string]int{"X": 1, "Y": 2, "Z": 3})
	fn02E([]int{1, 1, 2, 3, 5, 8})
}
