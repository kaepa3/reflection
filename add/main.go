package main

import (
	"fmt"
	"reflect"
)

func main() {
	A_1()
	fmt.Println("----------------------")
	A_2()
	fmt.Println("----------------------")
	A_3()
	fmt.Println("----------------------")
	A_4()
	fmt.Println("----------------------")
	A_5()
	fmt.Println("----------------------")
	A_6()
	fmt.Println("----------------------")
	A_7()
	fmt.Println("----------------------")
	A_8()
	fmt.Println("----------------------")
	A_9()
	fmt.Println("----------------------")
	A_10()
	fmt.Println("----------------------")
	A_11()
	fmt.Println("----------------------")
	A_12()
	fmt.Println("----------------------")
	A_13()
	fmt.Println("----------------------")
	A_14()
}
func A_14() {
	v := make(chan int, 3)
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	fmt.Println(rt.String())

	fmt.Println(rt.Elem().String())
	fmt.Println(rt.ChanDir())
	switch rt.Kind() {
	case reflect.Chan:
		fmt.Println(rt, "is chan")
	}
	fmt.Println(rv.Interface().(chan int))

	rv.Send(reflect.ValueOf(100))
	fmt.Println(<-v)
	v <- 200
	ret, ok := rv.Recv()
	fmt.Println(ret, ok)

	v = make(chan int, 1)
	rv = reflect.ValueOf(v)

	ok = rv.TrySend(reflect.ValueOf(500))
	fmt.Println(ok)

	ok = rv.TrySend(reflect.ValueOf(230))
	fmt.Println(ok)
	fmt.Println(<-v)
	v <- 555
	ret, ok = rv.TryRecv()
	fmt.Println(ret, ok)

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c2 <- 200
	index, recv, ok := reflect.Select([]reflect.SelectCase{
		reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c1),
		},
		reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c2),
		},
	})
	fmt.Printf("%d, %v ,%t\n", index, recv, ok)
}
func A_13() {
	var v interface{}
	rv := reflect.ValueOf(&v).Elem()
	fmt.Println(rv.Type().String())

	switch rv.Kind() {
	case reflect.Interface:
		fmt.Println(rv, "is interface")
	}
	rv.Set(reflect.ValueOf(3))
	fmt.Println(v)

	fmt.Println(rv.Interface())
	el := rv.Elem()
	fmt.Println(el.Type(), el)
	fmt.Println(reflect.TypeOf(0).Implements(rv.Type()))

}
func A_12() {
	v := 1
	rv := reflect.ValueOf(&v)
	switch rv.Kind() {
	case reflect.Ptr:
		fmt.Println(rv, "is Pointer")
	}
	fmt.Println(rv.Interface().(*int))
	fmt.Println(rv.Elem().Int())

	fmt.Println(reflect.Indirect(rv).Int())
	fmt.Println(reflect.Indirect(reflect.Indirect(rv)).Int())

	rv.Elem().Set(reflect.ValueOf(3))
	fmt.Println(v)
}
func A_11() {
	v := func(l, r int) int {
		return l + r
	}
	rv := reflect.ValueOf(v)
	rt := rv.Type()
	fmt.Println(rt.String())

	fmt.Println(rt.NumIn())
	fmt.Println(rt.NumOut())
	fmt.Println(rt.IsVariadic())
	fmt.Println(rt.In(0).String(), rt.In(1).String())
	fmt.Println(rt.Out(0).String())

	switch rv.Kind() {
	case reflect.Func:
		fmt.Println(rv, "is Func")
	}
	fmt.Println(rv.Interface().(func(int, int) int))

	rv = reflect.ValueOf(&v).Elem()
	rv.Set(reflect.ValueOf(func(r, l int) int {
		return r - l
	}))
	fmt.Println(v)

	ret := rv.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf(10)})
	fmt.Println(len(ret), ret[0].Int())

	rv = reflect.ValueOf(func(in ...int) int {
		sum := 0
		for _, n := range in {
			sum += n
		}
		return sum
	})
	args := []reflect.Value{
		reflect.ValueOf([]int{1, 2, 3, 4, 5}),
	}
	ret = rv.CallSlice(args)
	fmt.Println(len(ret), ret[0].Int())

}

type S1 struct {
	F1 int `t1:"aaa" t2:"bbb"`
	f2 int `t1:"ccc"`
}

func (s *S1) Func1(l, r int) int {
	return l + r
}

type S2 struct {
	F3 string
	F4 S1
}

func A_10() {
	v1 := S1{F1: 100, f2: 200}
	v2 := S2{
		F3: "x",
		F4: S1{
			F1: 10,
			f2: 20,
		},
	}
	rt1 := reflect.TypeOf(v1)
	rt2 := reflect.TypeOf(v2)

	fmt.Println(rt1.String())

	switch rt1.Kind() {
	case reflect.Struct:
		fmt.Println(rt1, "is struct")
	}
	fmt.Println(rt1.Field(0).Name)

	f, ok := rt1.FieldByName("F1")
	fmt.Println(f.Name, ok)

	f = rt2.FieldByIndex([]int{1, 0})
	fmt.Println(f.Name)

	f = rt1.Field(0)
	fmt.Println(f.Tag.Get("t1"))
	fmt.Println(f.Tag.Get("t2"))
	fmt.Println(f.Tag.Get("t3"))

	f = rt1.Field(1)
	t, ok := f.Tag.Lookup("t1")
	fmt.Println(t, ok)
	t, ok = f.Tag.Lookup("t2")
	fmt.Println(t, ok)

	rv := reflect.ValueOf(v1)
	fmt.Println(rv.Interface().(S1))

	rv = reflect.ValueOf(&v1).Elem()
	rv.Set(reflect.ValueOf(S1{F1: 300, f2: 100}))
	fmt.Println(v1)

	fmt.Println(rv.Type().NumField())

	fmt.Println(rv.Field(0))
	fmt.Println(rv.FieldByName("F1"))

	rv2 := reflect.ValueOf(v2)
	fmt.Println(rv2.FieldByIndex([]int{1, 0}))

	rv = reflect.ValueOf(&v1).Elem()
	rv.Field(0).SetInt(500)
	fmt.Println(v1)

	rv = reflect.ValueOf(&v1)
	rt := rv.Type()
	fmt.Println(rt.NumMethod())

	m := rt.Method(0)
	fmt.Printf("%+v\n", m)

	fmt.Println(m.Type.NumIn())
	fmt.Println(m.Type.NumOut())

	fmt.Println(m.Type.In(0).String(), m.Type.In(1).String(), m.Type.In(2).String())

	fmt.Println(m.Type.Out(0).String())

	ret := m.Func.Call([]reflect.Value{rv, reflect.ValueOf(10), reflect.ValueOf(20)})
	fmt.Println(len(ret), ret[0])

}
func A_9() {
	v := map[int]string{
		1: "a",
		2: "b",
	}
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().String())
	fmt.Println(rv.Type().Key().Name())

	switch rv.Kind() {
	case reflect.Map:
		fmt.Println(rv, "is map")
	}
	fmt.Println(rv.Interface().(map[int]string))

	rv = reflect.ValueOf(&v).Elem()
	rv.Set(reflect.ValueOf(map[int]string{3: "c", 4: "d"}))
	fmt.Println(v)

	fmt.Println(rv.Len())

	for _, key := range rv.MapKeys() {
		val := rv.MapIndex(key)
		fmt.Println(key, val)
	}

	iter := rv.MapRange()
	for iter.Next() {
		key, val := iter.Key(), iter.Value()
		fmt.Println(key, val)
	}

	rv = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(0), reflect.TypeOf(true)))
	fmt.Println(rv.Type().String())

	rv.SetMapIndex(reflect.ValueOf(100), reflect.ValueOf(true))
	rv.SetMapIndex(reflect.ValueOf(200), reflect.ValueOf(false))
	fmt.Println(rv)
}
func A_8() {
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)

	fmt.Println(rv.Type().String())

	fmt.Println(rv.Type().Elem().Name())

	switch rv.Kind() {
	case reflect.Slice:
		fmt.Println(rv, "is Slice")
	}

	fmt.Println(rv.Interface().([]int))

	rv = reflect.ValueOf(&v).Elem()
	rv.Set(reflect.ValueOf([]int{4, 5, 6}))
	fmt.Println(v)

	fmt.Println(rv.Len())
	fmt.Println(rv.Cap())

	for i := 0; i < rv.Len(); i++ {
		fmt.Println(rv.Index(i))
	}

	rv.Index(0).SetInt(7)
	fmt.Println(v)

	rv = reflect.Append(rv, reflect.ValueOf(8))
	v = rv.Interface().([]int)
	fmt.Println(v)

	rv = reflect.AppendSlice(rv, reflect.ValueOf([]int{9, 10}))
	v = rv.Interface().([]int)
	fmt.Println(v)

	rv = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(0)), 3, 5)
	fmt.Println(rv)

	rv = reflect.ValueOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	rv2 := rv.Slice(0, 5)
	fmt.Println(rv2, rv2.Len(), rv2.Cap())

	rv = reflect.ValueOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	rv3 := rv.Slice3(0, 5, 5)
	fmt.Println(rv3, rv3.Len(), rv3.Cap())

	rv = reflect.ValueOf([]int{1, 2, 3, 4, 5})
	rv2 = reflect.ValueOf(make([]int, 5, 5))
	i := reflect.Copy(rv2, rv)
	fmt.Println(rv, rv2, i)
}
func A_7() {
	v := [3]int{1, 2, 3}
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.Array:
		fmt.Println(rv, "is Array")
	}

	fmt.Println(rv.Interface().([3]int))

	rv = reflect.ValueOf(&v).Elem()
	rv.Set(reflect.ValueOf([3]int{4, 5, 6}))
	fmt.Println(rv)

	fmt.Println(rv.Len())
	fmt.Println(rv.Cap())

	for i := 0; i < rv.Len(); i++ {
		fmt.Println(rv.Index(i))
	}
	rv.Index(0).SetInt(10)
	fmt.Println(rv)
	v1 := [3]int{1, 2, 3}
	v2 := [3]int{}
	rv1 := reflect.ValueOf(&v1).Elem()
	rv2 := reflect.ValueOf(&v2).Elem()
	i := reflect.Copy(rv2, rv1)
	fmt.Println(v1, v2, i)

	rv = reflect.New(reflect.ArrayOf(3, reflect.TypeOf(0)))
	fmt.Println(rv1, rv2)

	v3 := [5]int{1, 2, 3, 4, 5}
	rv1 = reflect.ValueOf(&v3).Elem()
	rv2 = rv1.Slice(1, 3)
	fmt.Println(rv1, rv2)

	v3 = [5]int{1, 2, 3, 4, 5}
	rv1 = reflect.ValueOf(&v3).Elem()
	rv2 = rv1.Slice3(0, 3, 5)
	fmt.Println(rv1, rv2)

}
func A_6() {
	v := "a"
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.String:
		fmt.Println(rv, "is string")
	}

	fmt.Println(rv.String())
	fmt.Println(rv.Interface().(string))

	rv = reflect.ValueOf(&v).Elem()
	rv.SetString("hello")
	fmt.Println(rv)

	rv.Set(reflect.ValueOf("hoge"))
	fmt.Println(rv)

}
func A_5() {
	v := true
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.Bool:
		fmt.Println(rv, "is bool")
	}

	fmt.Println(rv.Bool())
	fmt.Println(rv.Interface().(bool))

	rv = reflect.ValueOf(&v).Elem()
	rv.SetBool(false)
	fmt.Println(rv)

	rv.Set(reflect.ValueOf(true))
	fmt.Println(rv)

}
func A_4() {
	v := complex128(1)
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.Complex128:
		fmt.Println(rv, "is complex")
	}

	fmt.Println(rv.Complex())
	fmt.Println(rv.Interface().(complex128))

	rv = reflect.ValueOf(&v).Elem()
	rv.SetComplex(2)
	fmt.Println(rv)

	rv.Set(reflect.ValueOf(complex128(2)))
	fmt.Println(rv)

}
func A_3() {
	v := float64(1)
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.Float64:
		fmt.Println(rv, "is float")
	}

	fmt.Println(rv.Float())
	fmt.Println(rv.Interface().(float64))

	rv = reflect.ValueOf(&v).Elem()
	rv.SetFloat(2)
	fmt.Println(rv)

	rv.Set(reflect.ValueOf(float64(5.2)))
	fmt.Println(rv)
}
func A_2() {
	v := uint(1)
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.Uint:
		fmt.Println(rv, "is uint")
	}

	fmt.Println(rv.Uint())
	fmt.Println(rv.Interface().(uint))

	rv = reflect.ValueOf(&v).Elem()
	rv.SetUint(2)
	fmt.Println(rv)

	rv.Set(reflect.ValueOf(uint(5)))
	fmt.Println(rv)
}
func A_1() {
	v := 1
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Type().Name())

	switch rv.Kind() {
	case reflect.Int:
		fmt.Println(rv, "is int")
	}

	fmt.Println(rv.Int())
	fmt.Println(rv.Interface().(int))

	rv = reflect.ValueOf(&v).Elem()
	rv.SetInt(2)
	fmt.Println(rv)

	rv.Set(reflect.ValueOf(5))
	fmt.Println(rv)
}
