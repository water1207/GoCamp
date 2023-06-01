package main

import "fmt"

// 切片的扩容机制： append长度增加后超过容量则将容量增加2倍
func main() {
	var numbers = make([]int, 3, 5) // 长度3， 容量5的切片 类似vector， 若缺省cap则默认与len相同

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素1， len = 4, [0,0,0,1], cap = 5
	numbers = append(numbers, 1)
	// [0,0,0,1,2]
	numbers = append(numbers, 2)

	// 再追加，超出预先的cap，自动开辟一倍cap， 当前cap = 10, len = 6, [0,0,0,1,2,3]
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 切片截取， [0, 2)
	numbers2 := numbers[0:3]
	fmt.Printf("cut = %v\n", numbers2)

	//numbers1、numbers2底层指向同一个数组属于浅拷贝，同时引用传递，修改2也会影响1，
	numbers2[0] = 100
	fmt.Println(" ==== 引用传递 =====")

	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("numbers2 = %v", numbers2)

	// copy输入深拷贝，为其numbers3开辟新内存空间
	numbers3 := make([]int, 3)
	copy(numbers3, numbers2)
	numbers3[0] = 999
	fmt.Println(" ==== 深拷贝 不影响原有值=====")

	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("numbers2 = %v", numbers2)

}
