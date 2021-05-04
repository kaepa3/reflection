package main

import (
	"fmt"
	"reflect"
)

func plus(v1, v2 interface{}) interface{} {
	rv1, rv2 := reflect.ValueOf(v1), reflect.ValueOf(v2)
	if isInt(rv1.Kind()) && isInt(rv2.Kind()) {
		return interface{}(rv1.Int() + rv2.Int())
	} else if rv1.Kind() == reflect.String && rv2.Kind() == reflect.String {
		return interface{}(rv1.String() + rv2.String())
	}

	return nil
}
func isInt(k reflect.Kind) bool {
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64

}
func main() {
	fmt.Println(plus(1, 20))
	fmt.Println(plus("1", "fsda"))
	fmt.Println(plus(false, 20))

}
