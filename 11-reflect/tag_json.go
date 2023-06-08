package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"一路向北", 2000, 10, []string{"JayChou", "Edison"}}

	// 结构体编码 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal err", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码 json -> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal err", err)
	}
	fmt.Printf("%v\n", myMovie)
}
