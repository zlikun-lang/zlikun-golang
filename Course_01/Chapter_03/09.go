package main

import "fmt"

// 判断两个map是否相等，必须遍历map中的元素，依次比较
func equal(x, y map[string]int) bool {

	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}

	return true
}

func main() {

	// true
	fmt.Println(equal(map[string]int{"x": 1, "y": 2}, map[string]int{"x": 1, "y": 2}))
	// false
	fmt.Println(equal(map[string]int{"x": 1, "y": 2}, map[string]int{"x": 1, "y": 3}))
	// false
	fmt.Println(equal(map[string]int{"x": 1, "y": 2}, map[string]int{"x": 1}))

}
