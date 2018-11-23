package main

import "fmt"

func main() {

	// 使用new创建一个指定类型的匿名变量，返回匿名变量的指针
	// 每次调用new(T)生成的指针是不相等的，这点等价于上例中的f函数
	p := new(int)
	// 指针p，指针p对应的值，指针本身的指针
	// 0xc00004a080 0 0xc000072018
	fmt.Println(p, *p, &p)

	// 修改针对对应变量的值
	*p = 2
	// 0xc00004a080 2 0xc000072018
	fmt.Println(p, *p, &p)

}
