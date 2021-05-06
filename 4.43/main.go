package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	f4_43()
	fmt.Println("------------------")
	f4_44()
	fmt.Println("------------------")
	f4_45()
	fmt.Println("------------------")
	f4_46()
	fmt.Println("------------------")
	f4_47()
	fmt.Println("------------------")
	f4_48()
	fmt.Println("------------------")
	f4_49()
	fmt.Println("------------------")
	f4_50()
	fmt.Println("------------------")
	f4_51()
	fmt.Println("------------------")
	f4_52()
	fmt.Println("------------------")
	f4_53()
	fmt.Println("------------------")
	f4_54()
	fmt.Println("------------------")
	f4_55()
	fmt.Println("------------------")
	f4_56()
	fmt.Println("------------------")
	f4_57()
	fmt.Println("------------------")
	f4_58()
	fmt.Println("------------------")
	f4_59()
	fmt.Println("------------------")
	f4_60()
	fmt.Println("------------------")
	f4_61()
}
func f4_61() {
	i := 9999
	rv := reflect.ValueOf(&i).Elem()
	ptr := rv.UnsafeAddr()
	iptr := (*int)(unsafe.Pointer(ptr))
	fmt.Println(*iptr)
}
func f4_60() {
	c := make(chan int, 1)
	rv := reflect.ValueOf(c)

	ok := rv.TrySend(reflect.ValueOf(2))
	fmt.Println(ok)

	ok = rv.TrySend(reflect.ValueOf(3))
	fmt.Println(ok)

	ok = rv.TrySend(reflect.ValueOf(4))
	fmt.Println(ok)
	fmt.Println(<-c)

}
func f4_59() {
	c := make(chan int, 3)
	rv := reflect.ValueOf(c)
	c <- 1
	i, ok := rv.TryRecv()
	fmt.Println(i, ok)

	i, ok = rv.TryRecv()
	fmt.Println(i, ok)
}
func f4_58() {
	rv := reflect.ValueOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	s := rv.Slice3(0, 5, 5)
	fmt.Println(s, s.Len(), s.Cap())
	s = rv.Slice3(0, 5, 10)
	fmt.Println(s, s.Len(), s.Cap())
	s = rv.Slice3(5, 7, 9)
	fmt.Println(s, s.Len(), s.Cap())
}
func f4_57() {
	rv := reflect.ValueOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(rv.Slice(0, 5))
	fmt.Println(rv.Slice(5, 9))
}
func f4_56() {
	i := 100
	j := 200
	p := unsafe.Pointer(&i)
	q := unsafe.Pointer(&j)
	rv := reflect.ValueOf(&p).Elem()
	rv.SetPointer(q)

	fmt.Println(*(*int)(p))
}
func f4_55() {
	v := map[int]string{
		100: "1",
		200: "b",
	}
	rv := reflect.ValueOf(&v).Elem()

	fmt.Println(rv.Interface())
	rv.SetMapIndex(reflect.ValueOf(300), reflect.ValueOf("c"))
	fmt.Println(rv.Interface())
	rv.SetMapIndex(reflect.ValueOf(200), reflect.Value{})
	fmt.Println(rv.Interface())
}
func f4_54() {
	v := make([]int, 0, 10)
	rv := reflect.ValueOf(&v).Elem()
	fmt.Println(rv.Len(), rv.Interface())
	rv.SetLen(5)
	fmt.Println(rv.Len(), rv.Interface())
}
func f4_53() {
	v := "a"
	rv := reflect.ValueOf(&v).Elem()
	rv.SetString("hello")
	fmt.Println(v)
}
func f4_52() {
	v := uint(1)
	rv := reflect.ValueOf(&v).Elem()
	rv.SetUint(2)
	fmt.Println(v)
}
func f4_51() {
	v := 1
	rv := reflect.ValueOf(&v).Elem()
	rv.SetInt(3)
	fmt.Println(v)
}
func f4_50() {
	v := float64(1)
	rv := reflect.ValueOf(&v).Elem()
	rv.SetFloat(2)
	fmt.Println(v)
}
func f4_49() {
	v := complex128(1)
	rv := reflect.ValueOf(&v).Elem()
	rv.SetComplex(2)
	fmt.Println(v)
}
func f4_48() {
	v := make([]int, 0, 10)
	rv := reflect.ValueOf(&v).Elem()
	fmt.Println(rv.Cap())
	rv.SetCap(5)
	fmt.Println(rv.Cap())
}
func f4_47() {
	v := []byte("a")
	rv := reflect.ValueOf(&v).Elem()
	rv.SetBytes([]byte("b"))
	fmt.Println(v)
}
func f4_46() {
	v := true
	rv := reflect.ValueOf(&v).Elem()
	rv.SetBool(false)
	fmt.Println(v)
}
func f4_45() {
	i := 1
	rv := reflect.ValueOf(&i).Elem()
	rv.Set(reflect.ValueOf(2))
	fmt.Println(i)

	s := []int{1, 3, 5}
	rv = reflect.ValueOf(&s).Elem()
	rv.Set(reflect.ValueOf([]int{2, 4, 6}))
	fmt.Println(s)

	m := map[int]string{
		10: "aaa",
		20: "bbb",
	}
	rv = reflect.ValueOf(&m).Elem()
	rv.Set(reflect.ValueOf(map[int]string{20: "ccc", 30: "ddd"}))
	fmt.Println(m)

	type S1 struct {
		F1 int
	}
	s1 := S1{F1: 100}
	rv = reflect.ValueOf(&s1).Elem()
	rv.Set(reflect.ValueOf(S1{F1: 200}))
	fmt.Println(m)
}
func f4_43() {

	c := make(chan int, 3)
	rv := reflect.ValueOf(c)

	c <- 1
	l, ok := rv.Recv()
	fmt.Println(l, ok)
	close(c)

	l, ok = rv.Recv()
	fmt.Println(l, ok)
}
func f4_44() {
	c := make(chan int, 1)
	rv := reflect.ValueOf(c)
	rv.Send(reflect.ValueOf(1))
	fmt.Println(<-c)
}
