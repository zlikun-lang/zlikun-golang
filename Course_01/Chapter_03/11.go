package main

import (
	"fmt"
	"time"
)

// 定义结构体，注意结构体成员不能包含结构体本身
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// 定义结构体变量
var john Employee

func main() {

	john.ID = 1
	john.Name = "John"
	john.DoB = time.Now()

	// {1 John  2018-11-28 09:35:24.3268984 +0800 CST m=+0.007977101  0 0}
	fmt.Println(john)

	// 结构体中的变量可以使用指针来操作
	position := &john.Position
	*position = "Senior " + *position
	// {1 John  2018-11-28 09:41:49.1843753 +0800 CST m=+0.006981401 Senior  0 0}
	fmt.Println(john)
}
