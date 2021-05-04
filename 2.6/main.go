package main

import (
	"fmt"
	"reflect"
)

func main() {

	rv := reflect.ValueOf(100)
	fmt.Println(rv.Int())

	rv = reflect.ValueOf(false)
	fmt.Println(rv.Bool())

	i := rv.Interface()
	if b, ok := i.(bool); ok {
		fmt.Println(b)
	}
}
