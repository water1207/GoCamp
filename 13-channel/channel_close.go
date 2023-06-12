package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close可以关闭channel,确定不用向channel发送数据了，才需要显示数据
		close(c)
	}()

	for {
		// ok如果为true表示channel没有关闭, 如果false表示channel已关闭, 若关闭则退出死循环
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	/*	可以使用range不断迭代操作channel
		for data := range c {
			fmt.Println(data)
		}*/

	// main结束
	fmt.Println("Main Finished..")
}
