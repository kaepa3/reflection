package main

import (
	"fmt"
	"reflect"
)

type S struct {
	F1 int
	F2 bool
	F3 string
}

func main() {
	v := S{
		F1: 1,
		F2: true,
		F3: "a",
	}
	rv := reflect.ValueOf(v)
	for i := 0; i < rv.Type().NumField(); i++ {
		f := rv.Field(i)
		fmt.Println(f.Type(), f.Interface())
	}

}
