package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	type S1 struct {
		F1 string
		F2 int
		F3 bool
	}
	type S2 struct {
		F4 int
		F5 S1
	}
	rv := reflect.ValueOf(S2{
		F4: 200,
		F5: S1{
			F1: "bbb",
			F2: 300,
			F3: false,
		},
	})
	fmt.Println(rv.FieldByIndex([]int{0}))
	fmt.Println(rv.FieldByIndex([]int{1, 1}))
	fmt.Println(rv.FieldByIndex([]int{1, 1}))
	fmt.Println(rv.FieldByIndex([]int{1, 2}))

	fmt.Println("-------------")

	rv = reflect.ValueOf(S1{
		F1: "bbb",
		F2: 300,
		F3: false,
	})
	fmt.Println(rv.FieldByName("F1"))
	fmt.Println(rv.FieldByName("F2"))
	fmt.Println(rv.FieldByName("F3"))

	fmt.Println("-----------")
	fmt.Println(rv.FieldByNameFunc(func(s string) bool {
		return s == "F1"
	}))
	fmt.Println(rv.FieldByNameFunc(func(s string) bool {
		return s == "F2"
	}))
	fmt.Println(rv.FieldByNameFunc(func(s string) bool {
		return s == "F3"
	}))

	fmt.Println("-----------")
	rv = reflect.ValueOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(rv.Index(0))
	fmt.Println(rv.Index(3))

	v := struct {
		F1 interface{}
	}{
		F1: 9999,
	}
	rv = reflect.ValueOf(v).Field(0)

	data := rv.InterfaceData()
	ptr := unsafe.Pointer(data[1])
	value := (*int)(ptr)
	fmt.Println(*value)
}
