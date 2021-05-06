package main

import (
	"fmt"
	"reflect"
)

func main() {
	{
		var v []int
		fmt.Println(reflect.ValueOf(v).IsNil())
		v = make([]int, 0, 0)
		fmt.Println(reflect.ValueOf(v).IsNil())
	}
	{
		var v *int
		fmt.Println(reflect.ValueOf(v).IsNil())
		vv := 1
		v = &vv
		fmt.Println(reflect.ValueOf(v).IsNil())
	}
	{
		var v chan int
		fmt.Println(reflect.ValueOf(v).IsNil())
		v = make(chan int, 1)
		fmt.Println(reflect.ValueOf(v).IsNil())
	}
	{
		var v func()
		fmt.Println(reflect.ValueOf(v).IsNil())
		v = func() {}
		fmt.Println(reflect.ValueOf(v).IsNil())
	}
	{
		var v interface{}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		fmt.Println(reflect.ValueOf(v).IsNil())
	}
	fmt.Println("hello")

}
