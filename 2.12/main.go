package main

import (
	"fmt"
	"reflect"
)

func main() {
	rv := reflect.ValueOf(map[int]string{
		10: "a",
		20: "b",
		30: "c",
	})
	iter := rv.MapRange()
	for iter.Next() {
		k, v := iter.Key(), iter.Value()
		fmt.Println(k, v)
	}
}
