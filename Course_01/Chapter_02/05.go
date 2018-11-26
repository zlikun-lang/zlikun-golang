package main

import (
	"fmt"
	"time"
)

// 声明一个常量
const pi = 3.14159

// 声明多个常量
const (
	e   = 2.71828182845904523536028747135266249775724709369995957496696763
	pi2 = 3.14159265358979323846264338327950288419716939937510582097494459
)

// 下面常量复制上面常量右边表达式
const (
	a = 1
	b
	c = 2
	d
)

// iota 常量生成器，实现类似Java枚举的功能
const (
	Sunday time.Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 复制右边表达式，这里用于计算容量单位
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main() {

	// 1 1 2 2
	fmt.Println(a, b, c, d)
	// Sunday Monday Tuesday Wednesday Thursday Friday Saturday
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	// Friday
	fmt.Println(Friday.String())
}
