package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(reflect.ValueOf(1).Int())
	fmt.Println(reflect.ValueOf(int64(5)).Int())

	fmt.Println(reflect.ValueOf(uint32(3)).Uint())
	fmt.Println(reflect.ValueOf(uint(9)).Uint())

	fmt.Println(reflect.ValueOf(float32(2)).Float())
	fmt.Println(reflect.ValueOf(float64(4)).Float())

	fmt.Println(reflect.ValueOf(complex64(20)).Complex())
	fmt.Println(reflect.ValueOf(complex128(40)).Complex())

	fmt.Println(reflect.ValueOf("aaa").String())
	fmt.Println(reflect.ValueOf(1).String())
	fmt.Println(reflect.ValueOf([]int{}).String())

	fmt.Println(reflect.ValueOf(true).Bool())

	fmt.Println(reflect.ValueOf([]byte{1, 2, 3}).Bytes())

	fmt.Println(reflect.ValueOf(1).Interface())
	fmt.Println(reflect.ValueOf("aaa").Interface())
	fmt.Println(reflect.ValueOf([]byte{1, 2, 3}).Interface())
	fmt.Println(reflect.ValueOf(map[int]string{1: "a", 2: "b"}).Interface())

}
