package main

import (
	"fmt"
	"reflect"
)

func main() {
	rt := reflect.TypeOf([3]int{})
	fmt.Println(rt.Len())
}
