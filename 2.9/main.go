package main

import (
	"fmt"
	"reflect"
)

func main() {

	var v string
	fmt.Println(reflect.ValueOf(v).IsZero())
	v = ""
	fmt.Println(reflect.ValueOf(v).IsZero())
	v = "a"
	fmt.Println(reflect.ValueOf(v).IsZero())
}
