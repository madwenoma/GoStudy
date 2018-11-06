package main

import (
	"reflect"
	"fmt"
)

/**
差异巨大
BenchmarkSet-4          2000000000               0.67 ns/op            0 B/op          0 allocs/op
BenchmarkRSet-4          5000000               252 ns/op               8 B/op          2 allocs/op
 */

type Data struct {
	X int
}

var d = new(Data)

func set(v int) {
	d.X = v
}

func rset(x int) {
	v := reflect.ValueOf(d).Elem() //对实例反射取ele实例，可以取方法名
	rField := v.FieldByName("X")
	rField.Set(reflect.ValueOf(x))
}

///////////////////////////////
var ve = reflect.ValueOf(d).Elem() //对实例反射取ele实例，可以取方法名
var rField = ve.FieldByName("X")

func cachedrset(x int) {
	rField.Set(reflect.ValueOf(x))
}

///////////////////////
func (data *Data) Inc() {
	data.X++
}

var v = reflect.ValueOf(d)
var rMethod = v.MethodByName("Inc")

func directCall() {
	d.Inc()
}

func rCall() {
	rMethod.Call(nil) //nil必传
}

func main() {
	directCall()
	fmt.Println(d.X)
	rCall()
	fmt.Println(d.X)

}
