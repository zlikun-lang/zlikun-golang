package main

import "fmt"

func change(x int) {
	x = 120
	fmt.Println(x)
}

func change2(x *int) {
	*x = 64
	fmt.Println(*x)
}

func main() {
	x := 16

	// 16
	fmt.Println(x)

	// 120，副本传参
	change(x)

	// 16
	fmt.Println(x)

	// 64
	change2(&x)

	// 64，指针传参
	fmt.Println(x)

}
