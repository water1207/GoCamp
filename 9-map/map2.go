package main

import "fmt"

// map也是引用传递，变量名cityMap就是一个指针
func printMap(cityMap map[string]string) {
	// 遍历
	for key, value := range cityMap {
		fmt.Printf("key = %v value = %s\n", key, value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "BJ"
	cityMap["Japan"] = "Tokyo"

	printMap(cityMap)

	// 删除
	fmt.Println(" ====删除后==== ")
	delete(cityMap, "China")
	printMap(cityMap)

	// 验证引用传递
	fmt.Println(" ====修改后 引用传递=== ")
	changeValue(cityMap)
	printMap(cityMap)

	cityMap2 := cityMap
	printMap(cityMap2)
	cityMap2["Japan"] = "Hangzhou"
	fmt.Println(" ==== 浅拷贝1 ====")
	printMap(cityMap)
	fmt.Println(" ==== 浅拷贝2 ====")
	printMap(cityMap2)

	/*	map不可以copy
		cityMap3 := make(map[string]string)
		copy(cityMap3, cityMap)
	*/

}
