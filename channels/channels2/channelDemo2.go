package main

import (
	"fmt"
	"time"
)

func myWorker(num int, c chan int) {
	for {
		fmt.Printf("worker %d received %c\n", num, <-c)
	}
}

func main() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go myWorker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //“a” + i会报错必须使用单引号
	}

	time.Sleep(time.Second)
}
