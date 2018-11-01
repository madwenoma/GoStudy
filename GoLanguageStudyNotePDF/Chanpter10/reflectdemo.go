package main

import (
	"reflect"
	"fmt"
)

type X int

type Y int

func main() {
	var a X = 100
	t := reflect.TypeOf(a)
	fmt.Println(t.Name(), t.Kind())
	///////////////////
	a = 200;
	var b X = 300
	var c Y = 400
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	fmt.Println(ta == tb, ta == tc)
	fmt.Println(ta.Kind() == tc.Kind())
	//////////////////////////
	//通过反射创建类型
	arr := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Println(arr, m)

	///////////////////////
	x := 100
	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println(tx.Kind(), tp.Kind())
	fmt.Println(tx == tp.Elem())

	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())


}
