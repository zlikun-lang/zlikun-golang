package main

import "flag"

// 使用命令行运行程序，观察输出
// $ go run Example/flag/01.go
func main() {
	var (
		host string
		port int
	)

	// flag 库用于输出 Usage 信息，类似于命令提示信息
	flag.StringVar(&host, "host", "", "主机名")
	flag.IntVar(&port, "port", 9091, "监听端口")
	flag.Parse()

	//   -host string
	//    	主机名
	//   -port int
	//    	监听端口 (default 9091)
	flag.Usage()
}
