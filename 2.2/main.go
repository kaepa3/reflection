package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(0).Kind() == reflect.Int)

	fmt.Println(reflect.TypeOf(uint(0)).Kind() == reflect.Uint)
	fmt.Println(reflect.TypeOf(float64(0)).Kind() == reflect.Float64)
	fmt.Println(reflect.TypeOf(false).Kind() == reflect.Bool)
	fmt.Println(reflect.TypeOf([1]int{}).Kind() == reflect.Array)
	fmt.Println(reflect.TypeOf([]int{}).Kind() == reflect.Slice)
	fmt.Println(reflect.TypeOf(map[int]bool{}).Kind() == reflect.Map)
	fmt.Println(reflect.TypeOf(make(chan int)).Kind() == reflect.Chan)
	fmt.Println(reflect.TypeOf(func() {}).Kind() == reflect.Func)
	type S struct{}
	fmt.Println(reflect.TypeOf(S{}).Kind() == reflect.Struct)
	fmt.Println(reflect.TypeOf(&S{}).Kind() == reflect.Ptr)
}
