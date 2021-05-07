package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	f5_1()
	fmt.Println("------------------")
	f5_2()
	fmt.Println("------------------")
	f5_3()
	fmt.Println("------------------")
	f5_6()
	fmt.Println("------------------")
	f5_7()
	fmt.Println("------------------")
	f5_8()
	fmt.Println("------------------")
	f5_9()
	fmt.Println("------------------")
	f5_11()
	fmt.Println("------------------")
	f5_12()
	fmt.Println("------------------")
	f5_13()
	fmt.Println("------------------")
	f5_14()
	fmt.Println("------------------")
	f5_14_2()
	fmt.Println("------------------")
	f5_14_3()
	fmt.Println("------------------")
	f5_15()
	fmt.Println("------------------")
	f5_16()
	fmt.Println("------------------")
	f5_17()
	fmt.Println("------------------")
	f5_18()
	fmt.Println("------------------")
	f5_19()
	fmt.Println("------------------")
	f5_20()
	fmt.Println("------------------")
	f5_21()
	fmt.Println("------------------")
	f5_22()
	fmt.Println("------------------")
	f5_23()
	fmt.Println("------------------")
	f5_24()
	fmt.Println("------------------")
	f5_25()
	fmt.Println("------------------")
	f5_26()
	fmt.Println("------------------")
	f5_27()
	fmt.Println("------------------")
	f5_28()
	fmt.Println("------------------")
	f5_30()
	fmt.Println("------------------")
	f5_31()
	fmt.Println("------------------")
	f5_32()
	fmt.Println("------------------")
	f5_33()
	fmt.Println("------------------")
	f5_34()
}
func f5_34() {
	fmt.Println(reflect.Zero(reflect.TypeOf(1)))
	fmt.Println(reflect.Zero(reflect.TypeOf(true)))
	fmt.Println(reflect.Zero(reflect.TypeOf("a")))
}
func f5_33() {
	rv := reflect.ValueOf(0)
	fmt.Println(rv)

	rv = reflect.ValueOf(nil)
	fmt.Println(rv)
}
func f5_32() {
	v := []int{1, 2, 3}
	p := unsafe.Pointer(&v)
	vv := reflect.NewAt(reflect.ArrayOf(3, reflect.TypeOf(0)), p)
	fmt.Println(vv.Type())
	fmt.Println(vv)
}
func f5_31() {
	v := reflect.New(reflect.TypeOf(0))
	fmt.Println(v.Type())
	fmt.Println(v.Elem())

}
func f5_30() {
	rt := reflect.TypeOf([]int{})
	rv := reflect.MakeSlice(rt, 5, 5)
	fmt.Println(rv)
}
func f5_29() {

	rt := reflect.TypeOf(map[int]string{})
	rv := reflect.MakeMapWithSize(rt, 3)
	rv.SetMapIndex(reflect.ValueOf(300), reflect.ValueOf("ccc"))
	rv.SetMapIndex(reflect.ValueOf(400), reflect.ValueOf("ddd"))
	fmt.Println(rv)
}
func f5_28() {
	rt := reflect.TypeOf(map[int]string{})
	rv := reflect.MakeMap(rt)
	rv.SetMapIndex(reflect.ValueOf(100), reflect.ValueOf("aaa"))
	rv.SetMapIndex(reflect.ValueOf(200), reflect.ValueOf("bbb"))
	fmt.Println(rv)
}
func f5_27() {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}
	var fn func(int, int) (int, int)
	rv := reflect.ValueOf(&fn).Elem()
	v := reflect.MakeFunc(rv.Type(), swap)
	rv.Set(v)
	a, b := 1, 2
	a, b = fn(a, b)
	fmt.Printf("a:%d, b:%d\n", a, b)

}
func f5_26() {
	elemType := reflect.TypeOf(0)
	chanType := reflect.ChanOf(reflect.BothDir, elemType)
	rv := reflect.MakeChan(chanType, 5)
	v, ok := rv.Interface().(chan int)
	fmt.Println(v, ok)
}
func f5_25() {
	v := 100
	rv := reflect.ValueOf(&v)
	fmt.Println(reflect.Indirect(rv))

	var vv *int
	rv = reflect.ValueOf(vv)
	fmt.Println(reflect.Indirect(rv))

	vvv := "aaa"
	rv = reflect.ValueOf(vvv)
	fmt.Println(reflect.Indirect(rv))
}
func f5_24() {
	rv := reflect.ValueOf([]int{1, 2, 3})
	fmt.Println(rv.Interface())
	rv = reflect.AppendSlice(rv, reflect.ValueOf([]int{4, 5, 6}))
	fmt.Println(rv.Interface())

}
func f5_23() {
	rv := reflect.ValueOf([]int{1, 2, 3})
	fmt.Println(rv.Interface())
	rv = reflect.Append(rv, reflect.ValueOf(4), reflect.ValueOf(5))
	fmt.Println(rv.Interface())
}
func f5_22() {
	fmt.Println(reflect.TypeOf(0))
}
func f5_21() {
	rt := reflect.StructOf([]reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag:  `json:"name"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
			Tag:  `json:"age"`,
		},
	})
	v := reflect.New(rt).Elem()
	v.Field(0).SetString("aaa")
	v.Field(1).SetInt(100)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json: %s\n", w.Bytes())
}

func f5_20() {
	rt := reflect.SliceOf(reflect.TypeOf(0))
	fmt.Println(rt)
	rv := reflect.MakeSlice(rt, 5, 5)
	fmt.Println(rv)
}
func f5_19() {
	rt := reflect.PtrTo(reflect.TypeOf(0))
	fmt.Println(rt)

	rv := reflect.Zero(rt)
	fmt.Println(rv)
}
func f5_18() {

	rt := reflect.MapOf(reflect.TypeOf(0), reflect.TypeOf(""))
	fmt.Println(rt)

}
func f5_17() {
	{
		in := []reflect.Type{
			reflect.TypeOf(0),
			reflect.TypeOf(false),
		}
		out := []reflect.Type{reflect.TypeOf(false)}
		rt := reflect.FuncOf(in, out, false)
		fmt.Println(rt)
	}
	{
		in := []reflect.Type{
			reflect.TypeOf(0),
			reflect.TypeOf(reflect.TypeOf(0)),
		}
		out := []reflect.Type{reflect.TypeOf(false)}
		rt := reflect.FuncOf(in, out, false)
		fmt.Println(rt)
	}
}
func f5_16() {

	elem := reflect.TypeOf(0)
	rt := reflect.ChanOf(reflect.SendDir, elem)
	fmt.Println(rt)

	rt = reflect.ChanOf(reflect.RecvDir, elem)
	fmt.Println(rt)

	rt = reflect.ChanOf(reflect.BothDir, elem)
	fmt.Println(rt)
	rv := reflect.MakeChan(rt, 1)

	rv.Send(reflect.ValueOf(100))
	i, ok := rv.Recv()
	fmt.Println(i, ok)
	rv.Close()

}
func f5_15() {
	elem := reflect.TypeOf(0)
	rt := reflect.ArrayOf(5, elem)
	ary := reflect.Zero(rt)
	fmt.Println(ary)
}
func f5_14_3() {
	type S struct {
		F1 int `tag1:"aaa" `
		F2 int `tag1:"bbb" tag2:"ccc"`
	}
	rt := reflect.TypeOf(S{})
	f1 := rt.Field(0)
	f2 := rt.Field(1)

	t, ok := f1.Tag.Lookup("tag1")
	fmt.Printf("tag1:%s, ok:%t\n", t, ok)
	t, ok = f1.Tag.Lookup("tag2")
	fmt.Printf("tag2:%s, ok:%t\n", t, ok)

	t, ok = f2.Tag.Lookup("tag1")
	fmt.Printf("tag1:%s, ok:%t\n", t, ok)
	t, ok = f2.Tag.Lookup("tag2")
	fmt.Printf("tag2:%s, ok:%t\n", t, ok)
}
func f5_14_2() {
	type S struct {
		F1 int `tag1:"aaa" `
		F2 int `tag1:"bbb" tag2:"ccc"`
	}
	rt := reflect.TypeOf(S{})
	f1 := rt.Field(0)
	f2 := rt.Field(1)
	fmt.Printf("tag1 is %s\n", f1.Tag.Get("tag1"))
	fmt.Printf("tag2 is %s\n", f1.Tag.Get("tag2"))
	fmt.Printf("tag1 is %s\n", f2.Tag.Get("tag1"))
	fmt.Printf("tag2 is %s\n", f2.Tag.Get("tag2"))
}
func f5_14() {
	type S struct {
		F1 int `tag1:"aaa" tag2:"bbb"`
		F2 int `tag3:"k1=v1,k2=v2"`
	}
	rt := reflect.TypeOf(S{})
	f1 := rt.Field(0)
	f2 := rt.Field(1)
	fmt.Printf("tag1 is %s, tag2 is %s\n", f1.Tag.Get("tag1"), f1.Tag.Get("tag2"))
	text := f2.Tag.Get("tag3")
	values := strings.Split(text, ",")
	for _, v := range values {
		pair := strings.Split(v, "=")
		fmt.Printf("%+v\n", pair)

	}
}
func f5_13() {
	type S struct {
		A int64
		B bool `taga:"aaa"`
		c string
	}
	rt := reflect.TypeOf(S{})
	sf := rt.Field(0)
	fmt.Printf("%+v\n", sf)
	sf = rt.Field(1)
	fmt.Printf("%+v\n", sf)
	sf = rt.Field(2)
	fmt.Printf("%+v\n", sf)
}
func f5_12() {
	s := "hello"
	size := len(s)

	var data string
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&[]byte(s)[0]))
	sh.Len = size
	fmt.Println(data)

}
func f5_11() {
	s := []int{1, 2, 3}
	size := len(s)

	var data []int
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&s[0]))
	sh.Len = size
	sh.Cap = size
	fmt.Println(data)

}
func f5_9() {
	{
		cases := []reflect.SelectCase{
			{
				Dir: reflect.SelectSend,
			},
			{
				Dir: reflect.SelectDefault,
			},
		}
		index, recv, recvOk := reflect.Select(cases)
		fmt.Printf("index:%d,recv:%v, recvOk:%t\n", index, recv, recvOk)
	}
	{
		c := make(chan int, 3)
		c <- 100
		cases := []reflect.SelectCase{
			{
				Dir: reflect.SelectRecv,
			},
			{
				Dir: reflect.SelectDefault,
			},
		}
		index, recv, recvOk := reflect.Select(cases)
		fmt.Printf("index:%d,recv:%v, recvOk:%t\n", index, recv, recvOk)
	}
}

func f5_8() {
	{
		c := make(chan int, 3)
		cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(c),
				Send: reflect.ValueOf(100),
			},
		}
		index, recv, recvOk := reflect.Select(cases)
		fmt.Printf("index:%d,recv:%v, recvOk:%t\n", index, recv, recvOk)
	}
	{
		c := make(chan int, 3)
		c <- 100
		cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			},
		}
		index, recv, recvOk := reflect.Select(cases)
		fmt.Printf("index:%d,recv:%v, recvOk:%t\n", index, recv, recvOk)
	}
	{
		cases := []reflect.SelectCase{
			{
				Dir: reflect.SelectDefault,
			},
		}
		index, recv, recvOk := reflect.Select(cases)
		fmt.Printf("index:%d,recv:%v, recvOk:%t\n", index, recv, recvOk)
	}
}

type S struct {
	n int
}

func (s S) Add(n int) int {
	return s.n + n
}
func f5_7() {
	v := S{n: 100}
	rt := reflect.TypeOf(S{n: 100})
	m := rt.Method(0)

	fmt.Println(m.Name)
	fmt.Println(m.PkgPath)
	fmt.Println(m.Type)
	fmt.Println(m.Index)
	fmt.Println(m.Func.Call([]reflect.Value{reflect.ValueOf(v), reflect.ValueOf(200)})[0])
}
func f5_6() {
	rv := reflect.ValueOf(map[int]string{
		100: "a",
		200: "b",
		300: "c",
	})
	mapIter := rv.MapRange()
	for mapIter.Next() {
		k, v := mapIter.Key(), mapIter.Value()
		fmt.Println(k, v)
	}

}
func f5_3() {
	v := []int{1, 2, 3, 4, 5}
	f := reflect.Swapper(v)
	fmt.Println(v)

	f(0, 4)
	fmt.Println(v)
	f(1, 3)
	fmt.Println(v)
}
func f5_2() {
	{
		fmt.Println("# array")

		fmt.Println(reflect.DeepEqual([3]int{1, 2, 3}, [3]int{1, 2, 3}))
		fmt.Println(reflect.DeepEqual([3]int{1, 2, 3}, [5]int{1, 2, 3}))
		fmt.Println(reflect.DeepEqual([3]int{1, 2, 3}, [3]int64{1, 2, 3}))
	}
	{
		fmt.Println("# struct")
		type S1 struct {
			F1 int
			f2 bool
		}
		type S2 struct {
			F1 int
			f2 bool
		}
		fmt.Println(reflect.DeepEqual(S1{F1: 100, f2: true}, S1{F1: 100, f2: true}))
		fmt.Println(reflect.DeepEqual(S1{F1: 100, f2: true}, S1{F1: 100, f2: false}))
		fmt.Println(reflect.DeepEqual(S1{F1: 100, f2: true}, S2{F1: 100, f2: true}))
	}
}

func f5_1() {
	{
		v1 := []int{1, 2, 3, 4, 5}
		v2 := make([]int, len(v1), cap(v1))
		v3 := make([]int, 3, 5)
		v4 := make([]int, 10, 10)
		count := reflect.Copy(reflect.ValueOf(v2), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v2, count)
		count = reflect.Copy(reflect.ValueOf(v3), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v3, count)
		count = reflect.Copy(reflect.ValueOf(v4), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v4, count)
	}
	{
		v1 := []int{1, 2, 3, 4, 5}
		v2 := [5]int{}
		v3 := [3]int{}
		v4 := [10]int{}
		count := reflect.Copy(reflect.ValueOf(&v2).Elem(), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v2, count)
		count = reflect.Copy(reflect.ValueOf(&v3).Elem(), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v3, count)
		count = reflect.Copy(reflect.ValueOf(&v4).Elem(), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v4, count)
	}
	{

		v1 := "hello"
		v2 := make([]uint8, 5, 5)
		v3 := make([]uint8, 3, 3)
		v4 := make([]uint8, 10, 10)
		count := reflect.Copy(reflect.ValueOf(v2), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v2, count)
		count = reflect.Copy(reflect.ValueOf(v3), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v3, count)
		count = reflect.Copy(reflect.ValueOf(v4), reflect.ValueOf(v1))
		fmt.Printf("%+v, %+v, %+v\n", v1, v4, count)
	}
}
