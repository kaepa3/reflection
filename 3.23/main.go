package main

import (
	"fmt"
	"reflect"
)

func main() {

	type S1 struct {
		F1 int `tag1:"aaa"`
		F2 int
	}
	rt := reflect.TypeOf(S1{})
	f, ok := rt.FieldByNameFunc(func(s string) bool {
		return s == "F1"
	})
	fmt.Printf("%s, %t\n", f.Name, ok)

	f, ok = rt.FieldByNameFunc(func(s string) bool {
		return s == "F2"
	})
	fmt.Printf("%s, %t\n", f.Name, ok)

	f, ok = rt.FieldByNameFunc(func(s string) bool {
		return s == "F3"
	})
	fmt.Printf("%s, %t\n", f.Name, ok)
}
