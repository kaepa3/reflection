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

type I interface {
	M1(int, int) int
}

func main() {
	v := &S{}
	rt := reflect.TypeOf(v)

	{
		m1 := rt.Method(0)
		fmt.Printf("%+v\n", m1)
		ret := m1.Func.Call([]reflect.Value{
			reflect.ValueOf(v),
			reflect.ValueOf(1),
			reflect.ValueOf(2),
		})
		fmt.Println(ret[0])
	}
	{
		s := struct {
			Inter I
		}{
			Inter: &S{},
		}
		rt := reflect.TypeOf(s)
		ft := rt.Field(0).Type
		m1 := ft.Method(0)

		fmt.Printf("%+v\n", m1)
	}

}
