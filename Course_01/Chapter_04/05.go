package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 错误处理
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		// 2018/11/29 10:02:36 Http Get Error => Get http2://httpbin.org/ip: unsupported protocol scheme "http2"
		log.Fatal("Http Get Error => ", err)
	} else {
		// 200 OK
		fmt.Println(resp.Status)
	}
}
