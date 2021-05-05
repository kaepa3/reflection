package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := 1
	fmt.Println(reflect.ValueOf(v).CanSet())
	fmt.Println(reflect.ValueOf(&v).Elem().CanSet())

	vv := struct {
		F1 int
		f2 int
	}{}
	rv := reflect.ValueOf(&vv).Elem()
	fmt.Println(rv.Field(0).CanSet())
	fmt.Println(rv.Field(1).CanSet())

	fmt.Println("---------------------------")
	rv = reflect.ValueOf([]int{1, 2, 3})
	fmt.Println(rv.Cap())
	rv = reflect.ValueOf([4]int{1, 2, 3, 4})
	fmt.Println(rv.Cap())
	rv = reflect.ValueOf(make(chan int, 5))
	fmt.Println(rv.Cap())
	fmt.Println("---------------------------")
	rv = reflect.ValueOf(make(chan int, 100))
	defer rv.Close()
	rv.Send(reflect.ValueOf(100))
	fmt.Println(rv.Recv())
	fmt.Println("---------------------------")
	rv1 := reflect.ValueOf(int8(1))
	rv2 := rv1.Convert(reflect.TypeOf(0))

	fmt.Println(rv1.Type(), rv2.Type())
	fmt.Println(rv1.Interface() == rv2.Interface())

	rv3 := reflect.ValueOf(1)
	fmt.Println(rv2.Interface() == rv3.Interface())
	fmt.Println("---------------------------")
	{
		v := 100
		rv := reflect.ValueOf(&v)
		fmt.Printf("%s, %#v\n", rv.Type(), rv.Elem())
	}
	{
		var v I
		rv := reflect.ValueOf(&v).Elem()
		fmt.Printf("%s, %#v\n", rv.Type(), rv.Elem())
	}
	type S1 struct {
		F1 string
		F2 int
		f3 bool
	}
	rv = reflect.ValueOf(S1{
		F1: "111",
		F2: 100,
		f3: true,
	})
	fmt.Println(rv.NumField())
}

type I interface {
	F1()
}
type S struct{}

func (s S) F1() {}
