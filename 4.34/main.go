package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

type S struct {
	F1 int
}

func (s *S) M1() int {
	return s.F1
}
func (s *S) M2(a, b int) int {
	return a + b
}
func main() {
	rv := reflect.ValueOf(map[int]string{
		100: "a",
		200: "b",
		300: "c",
	})
	mapItr := rv.MapRange()
	for mapItr.Next() {
		k, v := mapItr.Key(), mapItr.Value()
		fmt.Println(k, v)
	}
	fmt.Println("-----------------")

	rv = reflect.ValueOf(&S{F1: 100})
	m1 := rv.Method(0)
	fmt.Println(m1.Kind())
	fmt.Println(m1.Type())
	result := m1.Call([]reflect.Value{})
	fmt.Println(result[0])

	fmt.Println("-----------------")
	fmt.Println(rv.NumMethod())

	fmt.Println("-----------------")
	rv = reflect.ValueOf(int8(0))
	fmt.Println(rv.OverflowInt(1))
	fmt.Println(rv.OverflowInt(math.MaxInt16))
	fmt.Println("-----------------")
	rv = reflect.ValueOf(float32(0))
	fmt.Println(rv.OverflowFloat(1))
	fmt.Println(rv.OverflowFloat(math.MaxFloat64))
	fmt.Println("-----------------")
	rv = reflect.ValueOf(complex64(0))
	fmt.Println(rv.OverflowComplex(1))
	fmt.Println(rv.OverflowComplex(complex(math.MaxFloat64, 0)))
	fmt.Println("-----------------")

	{
		i := 9999
		rv = reflect.ValueOf(&i)
		ptr := rv.Pointer()
		iptr := (*int)(unsafe.Pointer(ptr))
		fmt.Println(*iptr)
	}
	{
		s := []int{1, 2, 3}
		size := len(s)
		rv = reflect.ValueOf(s)
		ptr := rv.Pointer()
		var data []int

		sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
		sh.Data = ptr
		sh.Len = size
		sh.Cap = size
		fmt.Println(data)
	}
	{
		s := "hello"
		size := len(s)
		rv = reflect.ValueOf([]byte(s))
		ptr := rv.Pointer()
		var data string

		sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
		sh.Data = ptr
		sh.Len = size
		fmt.Println(data)
	}
}
