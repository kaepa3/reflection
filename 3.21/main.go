package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S1 struct {
		F1 int `tag:"aaa"`
		F2 int
	}
	type S2 struct {
		F3 bool
		F4 S1
	}

	rt := reflect.TypeOf(S2{})

	f := rt.FieldByIndex([]int{0})
	fmt.Println(f.Name)

	f = rt.FieldByIndex([]int{1})
	fmt.Println(f.Name)

	f = rt.FieldByIndex([]int{1, 0})
	fmt.Println(f.Name)

	f = rt.FieldByIndex([]int{1, 1})
	fmt.Println(f.Name)
}
