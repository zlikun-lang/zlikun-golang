package main

import "fmt"

// 定义一个接口，定义说话、走路两个方法
// 在Go中接口表示一组方法，实际接口并不需要类似implements这样的关键字，而是只要实现了接口方法即实现了接口（鸭子谚语）
// 鸭子谚语：如果一只鸟长得像鸭子，叫声像鸭子，走路也像鸭子，那么它就是鸭子
// 空的interface，定义：interface{}，所有的类型都实现了该接口（没有方法）
type Human interface {
	// 说话，返回说话内容
	talk() string
	// 走路，返回走路速度
	walk() float64
}

type Student struct {
	name string
	age  int
}

// Student实现了talk方法
func (s *Student) talk() string {
	return fmt.Sprintf("%s 正在说话", s.name)
}

// Student实现了walk方法
func (s *Student) walk() float64 {
	speed := 1.5
	fmt.Printf("%s 正在以每秒%f米的速度走路\n", s.name, speed)
	return speed
}

// 测试方法，接收Human接口，执行接口方法
func test(human Human) {
	fmt.Println(human.walk())
	fmt.Println(human.talk())
}

func main() {
	// John 正在以每秒1.500000米的速度走路
	// 1.5
	// John 正在说话
	test(&Student{"John", 17})
}
