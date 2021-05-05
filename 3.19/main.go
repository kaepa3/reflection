package main

import (
	"fmt"
	"reflect"
)

func main() {

	rt := reflect.TypeOf(map[int64]float64{})
	fmt.Println(rt.Key())
	fmt.Println(rt.Elem())

	rt = reflect.TypeOf(map[string]bool{})
	fmt.Println(rt.Key())
	fmt.Println(rt.Elem())

}
