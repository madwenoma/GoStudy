package main

import "fmt"

//本例闭包多函数引用同一外部变量时，修改变量导致的同步问题

func testMore(x int) (func(), func()) {
	return func() {
		fmt.Println(x)
		x += 10
	}, func() {
		fmt.Println(x) //x被改变过了，已经不是入参X了
	}
}

func main() {
	a, b := testMore(100)
	a()
	b()
}
