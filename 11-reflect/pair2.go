package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type book struct {
}

func (this *book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// pair<type:book, value:book{}的地址>
	b := &book{}
	// pair<type: , value:>
	var r Reader
	// pair<type:book, value:book{}的地址>
	r = b // r被真正赋予为book类型
	r.ReadBook()

	// pair<type:book, value:book{}的地址>
	// 断言分两步: 得到r的动态类型的具体type(是book)，判断具体type(book)是否实现目标接口(WriteBook)
	w := r.(Writer) //断言还是成功的，因为w与r具体的type是一致的
	w.WriteBook()
}
