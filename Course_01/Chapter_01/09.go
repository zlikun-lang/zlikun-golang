package main

import "fmt"

// 类型通常在包一级声明，这里声明两温度单位，底层数据结构是float64
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 为 Celsius 类型声明一个 String() 方法
// 这里的 c 有点Java中的 this 的感觉，这个 String() 方法基本对应Java中的toString()方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {

	fmt.Println(CToF(AbsoluteZeroC)) // -459.66999999999996
	fmt.Println(CToF(FreezingC))     // 32
	fmt.Println(CToF(BoilingC))      // 212

	fmt.Println(FreezingC.String()) // 0°C
}
