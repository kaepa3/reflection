package main

import (
	"fmt"
	"reflect"
)

func isInt(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}

}
func main() {
	fmt.Println(isInt(1))
	fmt.Println(isInt("1"))
	fmt.Println(isInt(false))

}
