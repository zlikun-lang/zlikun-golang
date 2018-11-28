package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        // 类型后面部分表示JSON标签名
	Color  bool `json:"color,omitempty"` // omitempty表示成员为空或零值时，不生成JSON对象
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

type Data struct {
	Name string
	Age  int
}

func main() {
	// 转换为JSON（Encoder）
	// data, err := json.Marshal(movies)
	// 默认输出JSON没有格式化，难以阅读，可以通过后面两个参数进行格式化
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 反序列化
	data2 := Data{}
	err2 := json.Unmarshal([]byte("{\"name\":\"ashe\",\"age\":17}"), &data2)
	if err2 != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err2)
	} else {
		// {ashe 17}
		fmt.Println(data2)
	}
}
