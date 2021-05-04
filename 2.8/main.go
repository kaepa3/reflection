package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	rv := reflect.ValueOf(&i).Elem()
	rv.SetInt(2)
	fmt.Println(i)

	b := false
	rv = reflect.ValueOf(&b).Elem()
	rv.SetBool(true)
	fmt.Println(b)

	s := []int{10, 11, 12}
	rv = reflect.ValueOf(&s).Elem()
	rv.Set(reflect.ValueOf([]int{20, 21, 22}))
	fmt.Println(s)

}
