package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex" doc:"性别"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str)
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo)
		fmt.Println("doc:", tagDoc)
	}
}

func main() {
	var re resume
	findTag(re)
}
