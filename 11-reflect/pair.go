package main

import "fmt"

func main() {
	var a string
	a = "abcd"

	var allType interface{}
	allType = a

	val, _ := allType.(string)
	fmt.Println(val)
}
