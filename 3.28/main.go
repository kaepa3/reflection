package main

import (
	"fmt"
	"reflect"
)

func main() {
	rt := reflect.TypeOf(func(a int, b bool, c string) (float64, error) {
		return 0, nil
	})
	fmt.Println(rt.NumIn())
}
