package main

import "fmt"

func test(x int) func() {
	fmt.Println(&x)
	return func() {
		fmt.Println(&x, x)//&x地址是一样的，闭包完全引用了原变量地址
	}
}

func main() {
	f := test(123)
	f()
}
