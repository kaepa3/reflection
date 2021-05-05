package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S1 struct {
		F1 int `tag1:"aaa"`
		F2 int
	}
	type S2 struct {
		F3 bool
		F4 string
		F5 int
	}

	rt := reflect.TypeOf(S1{})
	fmt.Println(rt.NumField())

	rt = reflect.TypeOf(S2{})
	fmt.Println(rt.NumField())
}
