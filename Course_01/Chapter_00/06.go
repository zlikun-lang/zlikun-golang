package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 使用net库获取HTTP网络信息
func main() {

	url := "https://baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("status = %s\n", resp.Status)
	fmt.Printf("headers = %s\n", resp.Header)

	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))

}
