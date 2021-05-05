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
	m1, ok := rt.MethodByName("M1")
	fmt.Printf("%t, %+v\n", ok, m1)

	m2, ok := rt.MethodByName("M2")
	fmt.Printf("%t, %+v\n", ok, m2)

}
