package main

import "fmt"

func main() {
	// 声明map
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用map前，先用make分配数据空间
	myMap1 = make(map[string]string, 10)
	/*
		myMap2 := make(map[string]string) // 可以不指定具体容量
		myMap3 := map[string]string{      // 定义并初始化
			"one":   "a",
			"two":   "b",
			"three": "c",
		}
	*/
	myMap1["one"] = "a"
	myMap1["two"] = "b"
	myMap1["three"] = "c"

	// 乱序的 map[one:a three:c two:b]
	fmt.Println(myMap1)
}
