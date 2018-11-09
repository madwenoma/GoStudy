package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer wg.Done()
		for {
			var n int
			var ok bool
			select {
			case n, ok = <-c:
				fmt.Println("a1:", n)
			case n, ok = <-c:
				fmt.Println("a2:", n)
			}
			if !ok {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 10; i++ {
			select {
			case c <- i:
			case c <- i * 10:
			}
		}
	}()

	wg.Wait()
}
