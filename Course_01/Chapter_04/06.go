package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func read() error {
	in := bufio.NewReader(os.Stdin)
	for {
		_, _, err := in.ReadRune()
		// 判断文件到达末尾（错误）
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}
		// ...use r…
	}
	return nil
}

func main() {

	err := read()
	if err != nil {
		log.Fatal(err)
	}

}
