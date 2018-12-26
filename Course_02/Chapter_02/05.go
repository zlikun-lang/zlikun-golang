// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md
package main

import "fmt"

// 方法，在Go中带接收者的函数即为方法
type student struct {
	name  string
	score float64
}

// 定义考试方法exam，接收者为student指针变量stu
func (stu *student) exam() {
	// 实际上指针是没有字段的，但Go智能的实现了这一点，访问了指针指向变量的字段
	stu.score = 67.5
}

func main() {
	john := student{"John", 0}

	john.exam()

	// {John 67.5}
	fmt.Println(john)
}
