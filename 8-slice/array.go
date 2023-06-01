package main

import "fmt"

// printArray
func printArray2(myArray [10]int) {
	fmt.Println("This is printArray function")
	// 值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
func main() {
	// 固定长度数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	for i := 1; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	// for range
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	printArray2(myArray2)

	// 查看数组的数据类型
	fmt.Printf("types1 = %T\n", myArray1)
	fmt.Printf("types2 = %T\n", myArray2)
}
