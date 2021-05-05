package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F1 int
		f2 int
	}

	rv := reflect.ValueOf(S{})
	fmt.Println(rv.Field(0).CanInterface())
	fmt.Println(rv.Field(1).CanInterface())
	fmt.Println("-------------------------")
	i := 1
	fmt.Println(reflect.ValueOf(&i).Elem().Addr())
	fmt.Println(&i)
	fmt.Println("-------------------------")
	i = 0
	fmt.Println(reflect.ValueOf(i).CanAddr())
	fmt.Println(reflect.ValueOf(&i).Elem().CanAddr())
	a := [2]int{1, 2}
	fmt.Println(reflect.ValueOf(a).Index(0).CanAddr())
	fmt.Println(reflect.ValueOf(&a).Elem().Index(0).CanAddr())
	fmt.Println("-------------------------")
	{
		rv := reflect.ValueOf(func(a, b int) int {
			return a + b
		})
		results := rv.Call([]reflect.Value{
			reflect.ValueOf(1),
			reflect.ValueOf(2),
		})
		fmt.Println(results[0])
	}
	{
		rv := reflect.ValueOf(func(a, b int, in ...int) int {
			sum := a + b
			for _, n := range in {
				sum += n
			}
			return sum
		})
		results := rv.Call([]reflect.Value{
			reflect.ValueOf(1),
			reflect.ValueOf(2),
			reflect.ValueOf(3),
			reflect.ValueOf(4),
			reflect.ValueOf(5),
		})
		fmt.Println(results[0])
	}

	fmt.Println("-------------------------")
	{
		rv := reflect.ValueOf(func(in ...int) int {
			sum := 0
			for _, n := range in {
				sum += n
			}
			return sum
		})
		results := rv.CallSlice([]reflect.Value{
			reflect.ValueOf([]int{1, 2, 3, 4, 5}),
		})
		fmt.Println(results[0])

	}

}
