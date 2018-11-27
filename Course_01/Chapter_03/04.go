package main

// 数组清零函数
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// 初始值默认为0，所以可以简化
func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
