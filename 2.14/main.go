package main

import (
	"fmt"
	"reflect"
)

type S struct {
	F1 int    `tag1:"a"`
	F2 bool   `tag1:"b" tag2:"x"`
	F3 string `tag3:"c"`
}

func main() {
	rt := reflect.TypeOf(S{})
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fmt.Println("#", f.Name)
		text := f.Tag.Get("tag1")
		fmt.Println("tag1:", text)
		text, ok := f.Tag.Lookup("tag2")
		fmt.Println("tag2", ok, text)
	}

}
