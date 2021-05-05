package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(0).Comparable())
	fmt.Println(reflect.TypeOf(make(chan int)).Comparable())

}
