package main

import (
	"fmt"
	"reflect"
)

func main() {
	rt := reflect.TypeOf(0)
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(int8(0))
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(int16(0))
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(int32(0))
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(int64(0))
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(float64(0))
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(false)
	fmt.Println(rt.Align())

	rt = reflect.TypeOf("")
	fmt.Println(rt.Align())

	rt = reflect.TypeOf([10]int{})
	fmt.Println(rt.Align())

	rt = reflect.TypeOf([]int{})
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(map[int]int{})
	fmt.Println(rt.Align())

	rt = reflect.TypeOf(make(chan int, 2))
	fmt.Println(rt.Align())
}
