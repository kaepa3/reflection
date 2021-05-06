package main

import (
	"fmt"
	"reflect"
)

func main() {
	var rv reflect.Value
	fmt.Println(rv.IsValid())
	rv = reflect.ValueOf(0)
	fmt.Println(rv.IsValid())
	fmt.Println("------------------")

	fmt.Println(reflect.ValueOf(0).IsZero())
	fmt.Println(reflect.ValueOf(1).IsZero())
	fmt.Println(reflect.ValueOf("").IsZero())
	fmt.Println(reflect.ValueOf("aaa").IsZero())

	fmt.Println(reflect.ValueOf(true).IsZero())
	fmt.Println(reflect.ValueOf(false).IsZero())
	var v []int
	fmt.Println(reflect.ValueOf(v).IsZero())
	fmt.Println("------------------")

	rv = reflect.ValueOf([]int{1, 2, 3})
	fmt.Println(rv.Len())
	rv = reflect.ValueOf([4]int{1, 2, 3, 4})
	fmt.Println(rv.Len())
	rv = reflect.ValueOf(make(chan int, 5))
	fmt.Println(rv.Len())
	rv = reflect.ValueOf("hello")
	fmt.Println(rv.Len())

	fmt.Println("------------------")
	va := map[int]string{
		100: "a",
		200: "b",
	}
	rv = reflect.ValueOf(&va).Elem()
	fmt.Println(rv.MapIndex(reflect.ValueOf(100)))
	fmt.Println(rv.MapIndex(reflect.ValueOf(100)))

	fmt.Println("------------------")
	keys := rv.MapKeys()
	fmt.Println(len(keys))
	fmt.Println(keys[0])
	fmt.Println(keys[1])
}
