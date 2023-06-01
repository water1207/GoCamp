package main

import "fmt"

type Hero struct {
	// 类属性首字母大写，表示该属性public，小写为private
	Name  string
	Ad    int
	Level int
}

func (hero *Hero) Show() {
	fmt.Printf("Hero: %v", hero)
}

func (hero *Hero) GetName() string {
	fmt.Println("Name = " + hero.Name)
	return hero.Name
}
func (hero *Hero) SetName(newName string) {
	hero.Name = newName
}
func main() {
	// 创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()
}
