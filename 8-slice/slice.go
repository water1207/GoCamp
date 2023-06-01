package main

import "fmt"

func printArray(myArray []int) {
	// _ 表示匿名变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	// 区分于普通数组，这里是引用传递，切片的变量名即为指针
	myArray[0] = 99
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组， 切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray(myArray)

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
