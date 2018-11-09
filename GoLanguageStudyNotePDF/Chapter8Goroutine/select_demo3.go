package main

import (
	"fmt"
	"time"
)

//default 的两个用法
// 1.default防止阻塞
// 2.default处理通道满了的时候扩容

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for {

			select {
			case n, ok := <-c:
				if !ok {
					return
				}
				fmt.Println(n)

			default: //用default防止select阻塞，所以下方代码能执行
			}

			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	c <- 100
	close(c)
	<-done
	fmt.Println("========================")
	defaultDemo2()
}

func defaultDemo2() {
	done := make(chan struct{})
	dataChannel := []chan int{make(chan int, 2)}//异步通道，缓冲，2表示只能发送接受2次
	fmt.Println(cap(dataChannel), len(dataChannel))
	go func() {
		defer close(done)
		for i := 0; i < 10; i++ {
			select {
			case dataChannel[len(dataChannel)-1] <- i:
			default:
				dataChannel = append(dataChannel, make(chan int, 2))
			}
		}
	}()

	<-done

	fmt.Println("dc len:", len(dataChannel))
	for i := 0; i < len(dataChannel); i++ {
		dc := dataChannel[i]
		close(dc)
		for x := range dc {
			fmt.Println(x)
		}
	}
}
