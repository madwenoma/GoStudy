package main

import (
	"reflect"
	"fmt"
)

type XX int

func (XX) String() string {
	return ""
}
func main() {
	var a XX
	t := reflect.TypeOf(a)
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))
	it := reflect.TypeOf(0)
	fmt.Println(t.ConvertibleTo(it))
	fmt.Println(t.AssignableTo(st), t.AssignableTo(it))

}
