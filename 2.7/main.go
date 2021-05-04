package main

import (
	"fmt"
	"reflect"
)

func main() {

	v := 1
	fmt.Println(reflect.ValueOf(v).CanSet())
	fmt.Println(reflect.ValueOf(&v).Elem().CanSet())

	type S struct {
		F1 int
		f2 bool
	}
	vv := S{F1: 10, f2: true}
	fmt.Println(reflect.ValueOf(&vv).Elem().CanSet())
	fmt.Println(reflect.ValueOf(&vv).Elem().FieldByName("F1").CanSet())
	fmt.Println(reflect.ValueOf(&vv).Elem().FieldByName("f2").CanSet())
}
