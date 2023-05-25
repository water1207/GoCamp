package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.New(rand.NewSource(time.Now().UnixNano())) //设置随机数种子
	secretNum := rand.Intn(maxNum)                  //生成随机数

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Reading input err")
			continue
		}
		input = strings.TrimSuffix(input, "\n") //去除读入数据行中的换行符

		guess, err := strconv.Atoi(input) //将读入字符串转换为int
		if err != nil {
			fmt.Println("Invaild input.")
			continue
		}

		fmt.Println("Your guess is", guess)

		if guess > secretNum {
			fmt.Println("Your guess is bigger than secert num")
		} else if guess < secretNum {
			fmt.Println(("Your guess is smaller than secert num"))
		} else {
			fmt.Println("Correct, you win!")
			break
		}
	}
}
