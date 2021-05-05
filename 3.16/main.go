package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(0).Bits())
	fmt.Println(reflect.TypeOf(make(chan int)).Bits())
}
