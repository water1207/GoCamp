package main

import "fmt"

func main() {
	// 申明slice1是一个切片，并且初始化，默认值是1，2，3, 长度len为3
	slice1 := []int{1, 2, 3}

	// 申明slice2是一个切片，但是并没有给slice分配空间
	var slice2 []int
	slice2 = make([]int, 3) // 开辟3个空间， 默认值为0

	// 整合版
	slice3 := make([]int, 3)

	// %v 打印出详细数据
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice1 = %v\n", len(slice2), slice1)
	fmt.Printf("len = %d, slice1 = %v\n", len(slice3), slice1)

	//判断slice是否为空
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
