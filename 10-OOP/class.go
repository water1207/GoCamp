package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()....")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法

	level int
}

// 重定义父类方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	// 定义子类对象
	s := SuperMan{Human{"li4", "female"}, 99}

	//第二种定义方式
	var ss SuperMan
	ss.name = "wang5"
	ss.sex = "male"
	ss.level = 100

	s.Walk() // 父类方法
	s.Eat()  // 子类重写方法
	s.Fly()  // 子类的方法

}
