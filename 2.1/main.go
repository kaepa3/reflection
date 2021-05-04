package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {

	rt := reflect.TypeOf(100)
	fmt.Println(rt.Name())
	fmt.Println(rt.String())

	type T1 struct {
		F1 int
		F2 string
	}

	rt = reflect.TypeOf(T1{})
	fmt.Println(rt.Name())
	fmt.Println(rt.String())

	rt = reflect.TypeOf(rand.New(rand.NewSource(99)))
	fmt.Println(rt.Name())
	fmt.Println(rt.String())

}
