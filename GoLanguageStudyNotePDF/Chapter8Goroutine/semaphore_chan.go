package main

import (
	"runtime"
	"sync"
	"time"
	"fmt"
)

//P177 通过通道实现信号量

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	s := make(chan byte, 2)

	for i := 0; i < 9; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			s <- 1
			defer func() { <-s }()
			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
		fmt.Println("once")
	}
	wg.Wait()

}
