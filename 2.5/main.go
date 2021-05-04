package main

import (
	"fmt"
	"reflect"
)

func main() {
	rv1 := reflect.ValueOf(100)
	rv2 := reflect.ValueOf(100)
	rv3 := reflect.ValueOf(200)

	fmt.Println(rv1.Interface() == rv2.Interface())
	fmt.Println(rv1.Interface() == rv3.Interface())

	rv4 := reflect.ValueOf("aaa")
	rv5 := reflect.ValueOf("aaa")
	rv6 := reflect.ValueOf("bbb")
	fmt.Println(rv4.Interface() == rv5.Interface())
	fmt.Println(rv4.Interface() == rv6.Interface())

	rv7 := reflect.ValueOf([]int{1, 2, 3})
	rv8 := reflect.ValueOf([]int{1, 2, 3})
	rv9 := reflect.ValueOf([]int{4, 5, 6})
	fmt.Println(reflect.DeepEqual(rv7.Interface(), rv8.Interface()))
	fmt.Println(reflect.DeepEqual(rv7.Interface(), rv9.Interface()))
}
