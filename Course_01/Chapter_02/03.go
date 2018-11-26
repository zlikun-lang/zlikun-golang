package main

import "fmt"

func main() {
	var x complex64 = complex(1, 2)
	var y complex128 = complex(1, 2)

	// (-3+4i)，不同类型不能直接运算，所以这里强制转换
	fmt.Println((complex128(x)) * y)

	var z complex128 = complex(1, 2)

	// -3，实数部分
	fmt.Println(real(y * z))
	//  4，虚数部分
	fmt.Println(imag(y * z))

	// 任何浮点数或10进制整数后跟一个i，表示一个实数部为0的复数
	fmt.Println(3i)         // (0+3i)
	fmt.Println(3.14i)      // (0+3.14i)
	fmt.Println(3i * 3.14i) // (-9.42+0i)
}
