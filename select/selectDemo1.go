package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
select demos
1.语法
2.select代码块里用case接受，两个channel哪一个先传数据过来就显示哪一个
*/

func main() {
	fmt.Println("select...")
	var c1, c2 = createChan(), createChan()
	for {
		select {
		case n := <-c1:
			fmt.Printf("from c1: %d\n", n)
		case n := <-c2:
			fmt.Printf("from c2: %d\n", n)
		}
	}
	time.Sleep(time.Millisecond)
}

func createChan() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}
