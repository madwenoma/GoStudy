package main

import (
	"sync"
	"fmt"
)

//单向通道

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c //注意这个声明单向通道的写法
	var receive <-chan int = c

	go func() {
		defer wg.Done()
		for n := range receive {
			fmt.Println(n)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 5; i++ {
			send <- i
		}
	}()

	wg.Wait()
	wrongDemo()
}

func wrongDemo() {
	c := make(chan int)
	//var send chan<- int = c //注意这个声明单向通道的写法
	var receive <-chan int = c

	//<-send//编译无法通过
	receive <- 1//编译通过但运行报错：invalid operation: receive <- 1 (send to receive-only type <-chan int)
}
