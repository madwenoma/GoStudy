package main

import (
	"sync"
	"fmt"
)

//P176

type receiver struct {
	sync.WaitGroup
	data chan int
}

//工厂方法
func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			fmt.Println(x)
		}
	}()

	return r
}

func main() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	close(r.data)
	r.Wait()
}
