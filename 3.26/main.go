package main

import (
	"fmt"
	"reflect"
)

func main() {

	rt := reflect.TypeOf(func(a int, b bool, c string) (float64, error) {
		return 0, nil
	})

	fmt.Println(rt.NumIn())
	fmt.Println(rt.In(0))
	fmt.Println(rt.In(1))
	fmt.Println(rt.In(2))

	fmt.Println(rt.NumOut())

	fmt.Println(rt.Out(0))
	fmt.Println(rt.Out(1))

	fmt.Println(rt.IsVariadic())

	rt = reflect.TypeOf(func(a ...int) (float64, error) {
		return 0, nil
	})
	fmt.Println(rt.IsVariadic())
}
