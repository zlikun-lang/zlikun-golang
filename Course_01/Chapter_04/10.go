package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

// defer机制
func ReadFile(filename string) {
	start := time.Now()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	// defer 语句常用于处理成对的操作：打开-关闭、连接-断开、加锁-释放锁等
	// defer 语句会在程序退出前执行，类似于Java中的finally语句
	// 通过 defer 机制，可以保证无论函数有多复杂，资源都被正确释放
	// 通常 defer 语句跟在请求资源之后，如：打开文件后，使用defer语句关闭资源
	// defer 使用上的一个要注意的是，其执行时机是函数结束时，如果在循环中打开资源，则在循环内部不会释放
	defer f.Close()
	// 多个defer则按顺序执行
	defer log.Printf("文件已关闭")
	// 2018/11/29 10:48:18 elapsed (0s)
	defer log.Printf("elapsed (%s)", time.Since(start))

	log.Printf("读取文件：")
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf(string(data))
	}
}

func main() {
	ReadFile("Course_01/Chapter_04/10.go")
}
