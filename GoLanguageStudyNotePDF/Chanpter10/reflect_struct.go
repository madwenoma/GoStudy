package main

import (
	"reflect"
	"fmt"
	"net/http"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

////////////////////
type A int
type B struct {
	A
}

func (A) av()  {}
func (*A) ap() {}

func (B) bv()  {}
func (*B) bp() {}

//////////////

func main() {
	var m manager
	tm := reflect.TypeOf(m)
	name, _ := tm.FieldByName("name")
	fmt.Println(name.Name, name.Type)

	age := tm.FieldByIndex([]int{0, 1})
	fmt.Println(age.Name, age.Type)

	fmt.Println("///////////")
	var b B
	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}

	for _, t := range s {
		fmt.Println(t, ":")
		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(" ", t.Method(i))
		}
	}

	fmt.Println("///////////")
	var ss http.Server
	tt := reflect.TypeOf(ss)
	for i := 0; i < tt.NumField(); i++ {
		fmt.Println(tt.Field(i).Name)
	}

	fmt.Println("/////struct tag demo//////")
	type user struct {
		name string `field:"name" type:"varchar(50)"`
		age  int    `field:"age" type:"int"`
	}

	var u user
	t = reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s:%s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}

}
