package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Println("%v\n", this)
}

// 获取输入变量的字段与方法
func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	// 通过type 获取里面的字段
	// 1. 获取interface的reflect.Type, 通过Type得到NumField字段数量，进行遍历
	// 2. 得到每个field数据类型
	// 3. 通过field有一个Interface{}方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i) // 取出字段
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}

	// 通过type 获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func reflectNum(arg interface{}) {
	fmt.Println("Type: ", reflect.TypeOf(arg))
	fmt.Println("Value: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{1, "li4", 18}
	DoFiledAndMethod(user)
}
