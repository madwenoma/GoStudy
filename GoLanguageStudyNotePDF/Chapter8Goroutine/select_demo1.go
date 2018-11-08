package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()
		for {
			select {
			case n, ok := <-a:
				if !ok {
					a = nil
					break
				}
				fmt.Println(n)

			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				fmt.Println(x)
			}
			if a == nil && b == nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		for i := 1; i < 4; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)
		for i := 1; i < 4; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()
}
