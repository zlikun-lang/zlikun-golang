package main

import "fmt"

// panic类似于Java中的异常，但其使用场景有所不同
// panic会引起程序崩溃，因此主要用于严重错误，所以程序中应尽可能使用错误机制，而非panic
// 通常来说panic表示一个严重错误，不应对该异常做任何处理，但如果需要从panic中恢复，可以使用recover函数
func div(x, y float64) float64 {
	if y == 0 {
		panic("y is 0")
	}
	return x / y
}

func main() {
	fmt.Println(div(1, 4)) // 0.25
	fmt.Println(div(1, 0)) // panic: y is 0
	/* --------------------------------------------------------------
	panic: y is 0

	goroutine 1 [running]:
	main.div(...)
		E:/go/src/zlikun-golang/Course_01/Chapter_04/11.go:10
	main.main()
		E:/go/src/zlikun-golang/Course_01/Chapter_04/11.go:17 +0x9e
	-------------------------------------------------------------- */
}
