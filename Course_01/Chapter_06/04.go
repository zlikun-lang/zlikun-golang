package main

import (
	"fmt"
	"strconv"
)

type IntSet struct {
	numbers []int
}

func (i *IntSet) String() string {
	buf := ""
	for _, number := range i.numbers {
		buf += strconv.Itoa(number) + ","
	}
	return buf[:len(buf)-1]
}

func main() {

	i := IntSet{numbers: []int{1, 2, 3, 4}}
	// 1,2,3,4
	fmt.Println((&i).String())
}
