package main

import (
	"fmt"
	"reflect"
)

func main() {
	l, r := reflect.TypeOf(0), reflect.TypeOf(0)
	fmt.Println(l.AssignableTo(r))

	l, r = reflect.TypeOf(make(chan int)), reflect.TypeOf(0)
	fmt.Println(l.AssignableTo(r))
}
