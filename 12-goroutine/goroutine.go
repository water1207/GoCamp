package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个go程 去执行newTask() 流程
	// go newTask()

	// 匿名函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 退出当前gorourine， return只会退出当前函数
			runtime.Goexit()

			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
