package main

import (
	"fmt"
	"reflect"
)

func main() {

	rt := reflect.TypeOf(make(chan int))
	fmt.Println(rt.ChanDir())

	rt = reflect.TypeOf(make(chan<- int))
	fmt.Println(rt.ChanDir())

	rt = reflect.TypeOf(make(<-chan int))
	fmt.Println(rt.ChanDir())
}
