package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 带有缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		// 若channel缓冲区满，需要等有数据取出后再放入
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送元素=", i, "len(c)=", ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(3 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}
