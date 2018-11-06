package main

import (
	"fmt"
	"reflect"
)

/**
无法反射调用非导出方法（小写方法）
 */
type M struct{}

func (M) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err:%d", x+y)
}

func (M) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

func main() {
	fmt.Println("=================通过反射调用对象方法（固定参数）=================")

	var m M
	vm := reflect.ValueOf(&m)
	tm := vm.MethodByName("Test")

	vs := []reflect.Value{
		reflect.ValueOf(12),
		reflect.ValueOf(8),
	}

	out := tm.Call(vs) //[]Value

	for _, outValue := range out {
		fmt.Println(outValue)
	}

	fmt.Println("=================通过反射调用对象方法（可变参数）=================")

	fm := vm.MethodByName("Format")

	out = fm.Call([]reflect.Value{
		reflect.ValueOf("%s == %d"),
		reflect.ValueOf("hello"),
		reflect.ValueOf(250),
	})

	fmt.Println(out)

	out = fm.CallSlice([]reflect.Value{
		reflect.ValueOf("%s == %d"),
		reflect.ValueOf([]interface{}{"hello slice", 500,}),
	})
	fmt.Println(out)

}
