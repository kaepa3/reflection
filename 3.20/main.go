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
	sf := rt.Field(0)

	fmt.Println(sf.Name)
	fmt.Println(sf.Type)
	fmt.Println(sf.Tag)
	fmt.Println(sf.Index)
	fmt.Println(sf.Offset)

	fmt.Printf("%+v‘n", rt.Field(1))
}
