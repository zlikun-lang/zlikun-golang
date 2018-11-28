package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := map[string]int{
		"bob":   12,
		"alice": 17,
		"kevin": 25,
	}

	// map本身是无序的，可以通过对键排序的方式，变相实现有序map
	// var names []string
	// 因为slice的长度是已知的，所以建议设定其容量（性能相对较好）
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	// 对slice排序（string）
	sort.Strings(names)

	// 遍历names，以有序方式获取map中的元素值
	for _, name := range names {
		fmt.Println(name, ages[name])
	}

}
