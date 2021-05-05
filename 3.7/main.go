package main

import (
	"fmt"
	"reflect"
)

type S struct {
}

func (s *S) M1(a, b int) int {
	return a + b
}
func (s *S) M2(in ...int) int {
	sum := 0
	for _, n := range in {
		sum += n
	}
	return sum
}
func main() {
	v := &S{}
	rt := reflect.TypeOf(v)
	fmt.Println(rt.NumMethod())

}
