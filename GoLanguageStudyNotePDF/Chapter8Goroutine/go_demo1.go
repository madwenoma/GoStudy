package main

import (
	"time"
	"fmt"
)

var c = 2

func counter() int {
	c++
	return c
}

func main() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second)
		fmt.Println("go:", x, y)
	}(a, counter())//具有defer一样的延迟执行特点，会保存变量，并执行函数
	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("go:", a, counter())
	//}()
	a += 200
	fmt.Println("main:", a, counter())
	time.Sleep(time.Second * 3)
}
