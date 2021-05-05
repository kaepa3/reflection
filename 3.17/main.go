package main

import (
	"fmt"
	"reflect"
)

func main() {
	rt := reflect.TypeOf([3]int{})
	fmt.Println(rt.Elem())

	rt = reflect.TypeOf([]string{})
	fmt.Println(rt.Elem())

	rt = reflect.TypeOf(map[int64]float64{})
	fmt.Println(rt.Elem())

	rt = reflect.TypeOf(make(chan int))
	fmt.Println(rt.Elem())
}
