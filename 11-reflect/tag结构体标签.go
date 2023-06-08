package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("info")
		fmt.Println("info: ", tag)
	}
}

func main() {
	var re resume

	findTag(re)
}
