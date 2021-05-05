package main

import (
	"fmt"
	"reflect"
)

func main() {
	rt := reflect.TypeOf(100)
	fmt.Println("#", rt.Name())

	rt = reflect.TypeOf("aaa")
	fmt.Println("#", rt.Name())

	rt = reflect.TypeOf([]int{})
	fmt.Println("#", rt.Name())
}
