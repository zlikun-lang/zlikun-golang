package main

import "fmt"

func main() {

	// 内置make函数可以创建一个map
	ages := make(map[string]int)

	// 也可以使用map字面值创建一个map
	ages = map[string]int{
		"alice":   17,
		"charlie": 21,
	}

	// map[alice:17 charlie:21]
	fmt.Println(ages)

	// 修改map元素
	ages["alice"] = 18
	// map[alice:18 charlie:21]
	fmt.Println(ages)

	// 第一句创建的map是一个空的map，其等价于
	ages = map[string]int{}

	// 添加一个元素
	ages["john"] = 12

	// 获取一个元素
	fmt.Println(ages["john"]) // 12

	// 删除一个元素
	delete(ages, "john")

	// 操作一个不存在的键也是可以的
	ages["kevin"] += 1
	fmt.Println(ages["kevin"]) // 1

	ages["bob"]++
	fmt.Println(ages["bob"]) // 1

	// map中的元素并非变量，所以无法通过&来获取其指针

	// 遍历map
	// kevin -> 1
	// bob -> 1
	for k, v := range ages {
		fmt.Printf("%s -> %d\n", k, v)
	}

	// 判断key不存在
	age, ok := ages["bob"]
	// 1 true
	fmt.Println(age, ok)
	// 所以判断语句可以写作
	if age, ok = ages["bob"]; ok {
		fmt.Println("键 bob 是存在的")
	}

}
