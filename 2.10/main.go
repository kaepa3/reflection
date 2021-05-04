package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	rv := reflect.Zero(reflect.TypeOf(1))
	fmt.Println(rv)
	type S struct {
		F1 string
		F2 bool
	}

	rv = reflect.Zero(reflect.TypeOf(time.Now()))
	fmt.Println(rv)

	rv = reflect.Zero(reflect.TypeOf(S{}))
	fmt.Println(rv)
}
