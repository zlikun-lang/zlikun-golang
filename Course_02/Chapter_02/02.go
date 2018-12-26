// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md
package main

import (
	"errors"
	"fmt"
)

// 定义变量
// 1. var 定义变量
var name string

// 2. 定义多个变量
var x, y, z int

// 3. 定义变量，并初始化
var m, n float64 = 3.14, 3.15

// 4. 类型推导变量（不显示指定变量类型，由值推导）
var a, b, c = 2, 3, 5

func fnA() {

	// 5. 简短声明（省略var关键字，只能用于函数内部）
	i, j, k := 8, 13, 21

	// 6. 丢弃变量值（_是一个特殊变量名，任何赋给它的值都被丢弃）
	_, p := -1, 1

	// 7. 在Go中已声明未使用的变量，在编译时会报错
	fmt.Printf("%v, %v, %v, %v\n", i, j, k, p) // 8, 13, 21, 1

}

// 8. 常量，编译阶段确定值，运行阶段不允许改变
const pi = 3.14

// 9. 内置基础类型

// 1) 布尔
var boolVar bool

// 2) 整型，虽然下面变量都属于整型，但仍不能相互赋值（编译报错）
var intVar int
var uintVar uint
var int8Var int8
var uint8Var uint8
var int16Var int16
var uint16Var uint16
var int32Var int32
var uint32Var uint32
var int64Var int64
var uint64Var uint64
var byteVar byte // byte是uint8的别称，两者等价
var runeVar rune // rune是int32的别称，两者等价

//var intVar2 = int32Var + runeVar	// YES，两者实际是同一类型
//var intVar3 = int32Var + int64Var	// No ，两者是不同类型

// 3) 字符串，字符串与Java中字符串一样，是不可变类型
var text string = "Hello guys !"

// 使用``声明多行字符串，并且会保留源格式（字符串排版，如换行、制表、空格等）
var multiText string = `林花谢了春红，
太匆匆，
无奈朝来寒雨晚来风`

// 4) 错误类型
var err = errors.New("遇到了一个错误")

func fnB() {
	if err != nil {
		fmt.Println(err)
	}
}

// 0. 分组
const (
	constA = 1
	constB = "UPPER"
	constC = 3.14
)

var (
	varA = 1
	varB = "UPPER"
	varC = 3.14
)

// 1. 枚举，使用关键字 iota 实现，其初始默认值为0，每遇到一次const关键字就重置为0，每赋值一次就自增一次
const (
	constD = iota // 0，遇到const被重置
	constE = iota // 1，自增
	constF        // 2，隐式的被赋值为iota，所以仍然会自增
)

const constG = iota // 0，遇到const被重置

const (
	constH, constI, constJ = iota, iota, iota // 0, 0, 0，这种声明常量方式，实际只遇到一次const关键字，所以被重置为0，没有自增
)

const (
	constK = iota // 0，每个const分组的第一个常量会重置为iota为0
	constL = 120  // 120，被定义为其它值，但iota仍会自增
	constM = iota // 2，上面一行使用了其它值，但iota仍然自增了
)

const (
	constN = 7    // 7
	constO = iota // 1，在第一行被隐式重置为0，这里自增一次
	constP        // 2，隐式赋值为iota，自增一次
)

func fnC() {
	// 1, UPPER, 3.14,
	// 0, 1, 2,
	// 0,
	// 0, 0, 0,
	// 0, 120, 2,
	// 7, 1, 2
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n",
		constA, constB, constC, constD, constE,
		constF, constG, constH, constI, constJ,
		constK, constL, constM, constN, constO, constP)
}

// 2. Go程序设计的一些规则
// 		1) 大写字母开头的变量是可导出的，可以在其它包中被读取，是公有变量。小写字母开头的是不可导出的，是私有变量。
// 		2) 同理，函数名也遵循该原则

// 3. array 是数组，与切片的区别在于需要明确指定长度
// 声明时指定长度
var arrA [4]int

// 通过值来推导长度
var arrB = [...]int{1, 1, 2, 3, 5}

// 声明了数组长度，其它未设置值的被填充默认值
var arrC = [7]int{1, 1, 2, 3, 5}

// 多维数组，声明一个二维数组，第一维中每个元素本身也是一个一维数组
var arrD = [2][4]int{[4]int{1, 2, 3, 4}, [...]int{5, 6, 7, 8}}

// 声明时，内部类型实际可以省略
var arrE = [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

// 4. slice 是切片，与数组区别在于不指定长度，其长度是可变的
var sliceA []int

// 与数组类似，但不指定长度，也不使用 ... 来自动计算长度（没有固定长度）
var sliceB = []int{1, 1, 2, 3, 5, 8}

// 5. map 是映射（字典），实际上可以理解为切片的一个变种，但其索引可以是任意类型
// 声明一个 <string, int> 字典
var mapA map[string]int

// 另一种声明方式
var mapB = make(map[string]int)

// 6. Go的一些内建操作函数
// 1) make，用于内建类型（map、slice、channel）的内存分配，make(T, args)其返回有初始值的T类型

// 2) new，用于各种类型的内存分配，new(T)分配了零值填充的T类型的内存空间，并返回其内存地址（指针：*T）

// 7. 关于零值，整型0、bool为false、string为""

// 测试
func main() {
	fnA()
	fnB()
	fnC()
}
