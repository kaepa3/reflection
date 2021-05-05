package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

type S1 struct{}
type S2 S1
type S3 = S1

func main() {

	rt := reflect.TypeOf(0)
	fmt.Println(rt.PkgPath())

	rt = reflect.TypeOf("")
	fmt.Println(rt.PkgPath())

	rt = reflect.TypeOf(S1{})
	fmt.Println(rt.PkgPath())

	rt = reflect.TypeOf(S2{})
	fmt.Println(rt.PkgPath())

	rt = reflect.TypeOf(S3{})
	fmt.Println(rt.PkgPath())

	rt = reflect.TypeOf(base64.Encoding{})
	fmt.Println(rt.PkgPath())
}
