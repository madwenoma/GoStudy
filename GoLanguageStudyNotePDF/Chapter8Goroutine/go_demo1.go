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
	}(a, counter())
	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("go:", a, counter())
	//}()
	a += 200
	fmt.Println("main:", a, counter())
	time.Sleep(time.Second * 3)
}
