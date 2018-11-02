package main

import (
	"reflect"
	"fmt"
	"unsafe"
)

func main() {
	c := make(chan int, 4)
	v := reflect.ValueOf(c)
	if v.TrySend(reflect.ValueOf(100)) { //尝试传入值 返回ok，进行接收
		fmt.Println(v.TryRecv()) //接收返回值和接受结果
	}

	////////////////
	fmt.Println("==================nil interface=================")
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	fmt.Println(a == nil)
	fmt.Println(b == nil, reflect.ValueOf(b).IsNil()) //false true

	iface := (*[2]uintptr)(unsafe.Pointer(&b))
	fmt.Println(iface, iface[1] == 0)

}
