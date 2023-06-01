package main

import "fmt"

// interface{} 是万能数据类型 类似JavaObject，但是包含了通用数据类型如int、string
func autoFunc(arg interface{}) {
	fmt.Println(" ==== autoFunc is called ==== ")
	fmt.Println(arg)

	// 可用类型断言，来判断底层数据类型, 返回类型与断言结果
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type")
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
	name string
}

func main() {
	book := Book{"Lee", "Golang"}

	autoFunc(book)
	autoFunc("hhhh")
}
