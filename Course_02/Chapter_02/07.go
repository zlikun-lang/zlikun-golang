// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md
package main

import (
	"fmt"
	"runtime"
)

// 1. goroutine通过go关键字实现，本质上是一个普通函数
func say(msg string) {
	for i := 0; i < 5; i++ {
		// 表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched()
		fmt.Println(msg)
	}
}

func fn07A() {
	go say("***") // 启动一个协程
	say("###")    // 主线程执行
}

// 2. 使用channel进程协程间通信，channel使用chan关键字表示，管道传输数据使用 <- 符号，根据管道对象的位置确定是发送数据还是接收数据
// goroutine运行在相同的地址空间，因此访问共享内存必须做好同步，Channel相当于Unix中的双向管着，注意：必须使用make创建channel
// 管道发送数据和接收数据都是阻塞的，数据读取后管道会被阻塞，直到有数据接收，数据发出后会被阻塞，直到数据被读出（这不是一个意思么？！）
// 1) c chan int，表示一个int类型的管道
func fn07Sum(a []int, ch chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	// 2) 表示发送数据total到管道c
	ch <- total
}

func fn07B() {
	// 3) 使用make创建管道对象，第二个参数表示管道缓冲区大小（元素个数），为可选参数，当取值0时表示无缓冲阻塞读写
	ch := make(chan int)
	// 4) 使用协程进行并行计算
	go fn07Sum([]int{1, 1, 2, 3, 5, 8, 13}, ch)
	go fn07Sum([]int{21, 34, 55}, ch)
	// 5) 表示从管道c接收数据并赋值给变量（x, y）
	x, y := <-ch, <-ch
	// 110 33 143
	fmt.Println(x, y, x+y)
}

// 3. 通过range和close可以迭代使用管道
func fn07Fib(n int, ch chan int) {
	// 1) 函数结束时关闭管道
	defer close(ch)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		// 2) 循环向管道写入数据
		ch <- x
		x, y = y, x+y
	}
}

func fn07C() {
	// 3) 构造一个管道，缓冲容量为10
	ch := make(chan int, 10)
	// 4) 使用协程调用Fib函数，cap(ch)返回管道缓冲区大小，这里为10，表示只计算10个Fib数字
	go fn07Fib(cap(ch), ch)
	// 5) 使用range迭代管道（缓冲区）中数据，直接管道被显示关闭
	for i := range ch {
		fmt.Println(i)
	}
}

// 4. 通过 select 关键字，可以监听多个channel上的数据流动
func fn07Fib2(ch, quit chan int) {
	x, y := 1, 1
	for {
		// 1) 使用select选择多个管道，类似于switch的语法
		select {
		case ch <- x:
			// 2) ch <- x 接入数据管道中的数据
			x, y = y, x+y
		case <-quit:
			// 3) <-quit 返回true时函数结束，消费者显示传入0（在Go中0也表示false）用于结束生产者
			fmt.Println("Quit ...")
			return
		default:
			// 当case语句中的管道都没准备好时执行
		}
	}
}

func fn07D() {
	// 4) 声明数据管道和退出标识管道
	ch, quit := make(chan int), make(chan int)
	// 5) 启一个匿名函数协程，向Fib函数发送数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		// 6) 当发送10条数据后，发出退出信号（Fib函数需要正确处理该信息）
		quit <- 0
	}()
	// 7) 启动Fib函数
	fn07Fib2(ch, quit)
}

func main() {
	fn07A()
	fn07B()
	fn07C()
	fn07D()
}
