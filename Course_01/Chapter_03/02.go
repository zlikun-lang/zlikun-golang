package main

import "fmt"

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	// 数组字面值，初始化顺序无关紧要，未用到的索引可以省略，使用零值初始化
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	// [$ € ￡ ￥]
	fmt.Println(symbol)
	// $
	fmt.Println(symbol[0])
	// €
	fmt.Println(symbol[1])
	// ￡
	fmt.Println(symbol[2])
	// ￥
	fmt.Println(symbol[3])

	// 这里使用数组字面值，最大索引为4（第5个元素），其它元素以零值填充
	r := [...]int{4: -1}
	// [0 0 0 0 -1]
	fmt.Println(r)
}
