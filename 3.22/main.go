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
	rt := reflect.TypeOf(S1{})
	f, ok := rt.FieldByName("F1")
	fmt.Printf("%s, %t\n", f.Name, ok)

	f, ok = rt.FieldByName("F2")
	fmt.Printf("%s, %t\n", f.Name, ok)

	f, ok = rt.FieldByName("F3")
	fmt.Printf("%s, %t\n", f.Name, ok)
}
