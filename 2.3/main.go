package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func AnyToFloat(v interface{}) float64 {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	case reflect.Complex64, reflect.Complex128:
		return real(rv.Complex())
	case reflect.String:
		if f, err := strconv.ParseFloat(rv.String(), 64); err != nil {
			return f
		}
	}
	return 0
}
func main() {
	fmt.Println(AnyToFloat(100))
	fmt.Println(AnyToFloat(uint(200)))
	fmt.Println(AnyToFloat(float64(100)))
	fmt.Println(AnyToFloat(complex128(400)))
	fmt.Println(AnyToFloat("500"))
	fmt.Println(AnyToFloat(false))
	fmt.Println(AnyToFloat([]int{}))

}
