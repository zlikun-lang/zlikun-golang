// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.4.md
package main

import "fmt"

// 定义一个结构体
type person struct {
	name string
	age  int
}

type student struct {
	// 匿名字段，student结构体包含了person结构体的所有字段
	// 如果外层也定义了匿名结构体中相同字段，以外层为准（重载）
	person
	score float64
}

// 使用结构体，声明一个结构体变量
var John person

// 通过赋值声明一个结构体变量
var Kevin = person{"Kevin", 17}

func main() {
	fmt.Println(John)  // { 0}
	fmt.Println(Kevin) // {Kevin 17}

	// 访问结构体变量属性
	// The Kevin is 17 years old.
	fmt.Printf("The %s is %d years old.\n", Kevin.name, Kevin.age)

	// 给结构体变量赋值
	John.name = "John"
	fmt.Println(John) // {John 0}

	// 使用命名参数任意顺序赋值结构体变量
	John = person{age: 21, name: "John"}
	fmt.Println(John) // {John 21}

	// 与普通类型一样，通过new声明一个指针
	var kevinPointer *person = new(person)
	// { 0}
	fmt.Println(*kevinPointer)
	// 修改指针指向结构体字段值
	kevinPointer.name = "Kevin"
	// {Kevin 0}
	fmt.Println(*kevinPointer)

	// 嵌套结构体
	var ashe = student{person{"Ashe", 16}, 78.5}
	// {{Ashe 16} 78.5}
	fmt.Println(ashe)

	// 修改字段，因为定义时使用匿名字段，在使用时可以省略person前缀，而直接使用其定义的字段
	ashe.age = 17
	// {{Ashe 17} 78.5}
	fmt.Println(ashe)
	// 如果不省略，默认为结构体类型名称（如果是匿名定义，否则以显示定义的结构体字段名为准）
	ashe.person.age = 18
	// {{Ashe 18} 78.5}
	fmt.Println(ashe)
}
