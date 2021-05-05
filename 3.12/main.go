package main

import (
	"fmt"
	"reflect"
)

type I interface {
	Add(int, int) int
}
type S1 struct {
}
type S2 struct {
}

func (s S2) Add(a, b int) int {
	return a + b
}
func main() {
	rt := reflect.TypeOf((*I)(nil)).Elem()
	fmt.Println(reflect.TypeOf(S1{}).Implements(rt))
	fmt.Println(reflect.TypeOf(S2{}).Implements(rt))

}
