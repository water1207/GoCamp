package main

import "fmt"

// 定义接口
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体类-Cat
type Cat struct {
	color string
}

// 实现接口
func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

// 具体类-Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) { // 接收指针
	animal.Sleep() // 多态
}

func main() {
	/*
		多态的现象，子类向上实现父类
		1. 有一个父类(有接口)
		2. 有子类(实现了父类的全部接口的方法)
		3. 父类类型的变量(指针)指向子类的具体数据变量
	*/
	var animal AnimalIF // 接口的数据类型， 父类指针
	animal = &Cat{"White"}

	animal.Sleep() // 调用了Cat的Sleep

	animal = &Dog{"Yellow"}

	animal.Sleep()

	// 多态的现象，封装一下
	cat := Cat{"Blue"}
	dog := Dog{"Black"}

	showAnimal(&cat)
	showAnimal(&dog)
}
