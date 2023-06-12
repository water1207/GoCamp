package main

import (
	"fmt"
)

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("Goroutine结束")

		fmt.Println("Goroutine ...")

		c <- 666 // 将666接入管道中
	}()

	// 管道会自动协调同步，若main执行更快，会让main阻塞，等待子进程发送管道数据
	// 无缓冲机制，如果子进程更快，main还不需要管道数据，会让子进程阻塞，等待main的需要
	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")

}
