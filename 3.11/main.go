package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(0).Kind() == reflect.Int)
	fmt.Println(reflect.TypeOf("").Kind() == reflect.String)

}
