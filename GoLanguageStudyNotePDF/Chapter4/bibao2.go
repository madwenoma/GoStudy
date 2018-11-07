package main

import "fmt"

var s []func()

var t []func()

func testAddNimingToList() []func() {
	for i := 0; i < 2; i++ {
		s = append(s, func() {
			fmt.Println(&i, i)
		})
	}
	return s
}

func testCopyParam() []func() {
	for i := 0; i < 2; i++ {
		x := i
		t = append(t, func() {
			fmt.Println(&x, x)
		})
	}
	return t
}

func main() {
	for _, value := range testAddNimingToList() {
		value()
	}

	for _, value := range testCopyParam() {
		value()
	}
}
