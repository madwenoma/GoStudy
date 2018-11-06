package main

import (
	"reflect"
	"strings"
	"fmt"
)

/**
作者说如果语言支持泛型，则不需要这么费力的实现
	基于类型的加法
	只能通过反射来实现
 */

func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value
	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, value := range args {
			n += int(value.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, str := range args {
			ss = append(ss, str.String())//reflect.Value类型强制转为string：str.String()
		}
		ret = reflect.ValueOf(strings.Join(ss, "")) //数组ss转为string
	}
	return append(results, ret);
}

func makeAdd(fptr interface{}) {
	fe := reflect.ValueOf(fptr).Elem()
	ff := reflect.MakeFunc(fe.Type(), add)
	fe.Set(ff)
}

func main() {
	var intAdd func(x, y int) int
	var strAdd func(a, b string) string
	makeAdd(&intAdd)
	makeAdd(&strAdd)
	fmt.Println(intAdd(100, 200))
	fmt.Println(strAdd("hello ", "make func"))
}
