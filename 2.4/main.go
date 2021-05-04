package main

import (
	"fmt"
	"reflect"
)

func main() {
	rt1 := reflect.TypeOf(100)
	rt2 := reflect.TypeOf(200)
	rt3 := reflect.TypeOf(true)
	fmt.Println(rt1 == rt1)
	fmt.Println(rt1 == rt2)
	fmt.Println(rt1 == rt3)

	rt4 := reflect.TypeOf([]int{})
	rt5 := reflect.TypeOf([]int{})
	rt6 := reflect.TypeOf([]string{})
	fmt.Println(rt4 == rt5)
	fmt.Println(rt4 == rt6)

	rt7 := reflect.TypeOf(map[int]bool{})
	rt8 := reflect.TypeOf(map[int]bool{})
	rt9 := reflect.TypeOf(map[bool]int{})
	fmt.Println(rt7 == rt8)
	fmt.Println(rt7 == rt9)

}
