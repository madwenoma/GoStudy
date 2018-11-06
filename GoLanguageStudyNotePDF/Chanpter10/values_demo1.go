package main

import (
	"reflect"
	"fmt"
	"unsafe"
)

//P208

type User struct {
	Name string
	code int
}

func main() {
	a := 100
	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()
	fmt.Println(va.CanAddr(), va.CanSet())
	fmt.Println(vp.CanAddr(), vp.CanSet())
	fmt.Println("==========================")

	u := new(User)
	v := reflect.ValueOf(u).Elem()
	name := v.FieldByName("Name")
	code := v.FieldByName("code")
	fmt.Println(name.CanAddr(), name.CanSet())
	fmt.Println(code.CanAddr(), code.CanSet()) //true false因为code字段是小写，非公开字段

	if name.CanSet() {
		name.SetString("Tom")
	}
	if code.CanAddr() {
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}

	fmt.Printf("%+v\n", *u)

	fmt.Println("=============通过反射改变struct属性值=============")
	type user struct {
		Name string
		Age  int
	}
	user1 := user{
		"q.yuhen",
		34,
	}

	tu := reflect.ValueOf(&user1)//取地址值

	if !tu.CanInterface() {
		fmt.Println("CanInterface:fail1.")
		return
	}

	if pu, ok := tu.Interface().(*user); ok {//tu.Interface().(*user);
		pu.Age++
		fmt.Printf("%+v\n", user1)
	} else {
		fmt.Println("CanInterface:fail.")
	}


}
