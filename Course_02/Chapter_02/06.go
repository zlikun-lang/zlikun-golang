// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
package main

import (
	"fmt"
	"reflect"
)

// 在Go中interface表示一组方法的签名（大多数语言中也是这样，比如：Java），但Go的interface实现机制比较特别（这里主要指Java）
// 在Go中并不显示的定义一个实现类属性某个接口，而是只要这个类（结构体）包含了接口所定义的全部方法，即为该接口的实现者

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	salary float64
}

// 在Human中定义talk方法
func (h Human) talk() {
	fmt.Printf("[Human] %s 正在讲话 ...\n", h.name)
}

// Student重载了talk方法
func (stu Student) talk() {
	fmt.Printf("[Student] %s 正在谈论自己的学校 ...\n", stu.name)
}

// Employee重载了talk方法
func (emp Employee) talk() {
	fmt.Printf("[Employee] %s 正在谈论自己的工资 ...\n", emp.name)
}

// 定义一个Talk接口，上面三个结构体都实现了该接口
// 上面三个结构体都未显示定义与当前接口关系，但都实现了该接口全部方法，所以被认为实现了该接口
// [duck-typing]：当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子！
type Talker interface {
	talk()
}

// 定义一个函数，接收Talker接口
func hello(talker Talker) {
	talker.talk()
}

// !! 空接口因为不包含任何方法，所以可以认为所有类型都实现了该空接口

// 接口继承，隐式包含所有Talker接口的方法，并增加了自己的方法
type Joker interface {
	Talker // 继承
	joke() // 自定义方法
}

func main() {
	// 测试Human方法
	human := Human{"John", 17}
	// [Human] John 正在讲话 ...
	human.talk()

	// 测试Student方法
	student := Student{human, "银河学院"}
	// [Student] John 正在谈论自己的学校 ...
	student.talk()

	// 测试Employee方法
	employee := Employee{human, 3500.0}
	// [Employee] John 正在谈论自己的工资 ...
	employee.talk()

	// 通过接口调用方法
	var t1 Talker = student
	// [Student] John 正在谈论自己的学校 ...
	t1.talk()

	var t2 Talker = employee
	// [Employee] John 正在谈论自己的工资 ...
	t2.talk()

	// 调用hello函数
	// [Human] John 正在讲话 ...
	hello(&human)
	// [Student] John 正在谈论自己的学校 ...
	hello(&student)
	// [Employee] John 正在谈论自己的工资 ...
	hello(&employee)

	// 判断接口是否是某一类型变量，语法：interface_object.(T)
	value, ok := t1.(Student)
	// {{John 17} 银河学院} true
	fmt.Println(value, ok)

	// 反射
	// 反射对象类型
	t := reflect.TypeOf(t1)
	// 反向对象值
	v := reflect.ValueOf(t1)
	// main.Student {{John 17} 银河学院}
	fmt.Println(t, v)

}
