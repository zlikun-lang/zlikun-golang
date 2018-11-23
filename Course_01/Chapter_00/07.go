package main

import (
	"fmt"
	"net/http"
)

// 实现一个简易的WEB服务器，http://127.0.0.1/
func main() {
	// 注册路由
	http.HandleFunc("/", handler)
	// 注册监听地址
	http.ListenAndServe("0.0.0.0:80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 将请求URL作为响应内容输出
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
