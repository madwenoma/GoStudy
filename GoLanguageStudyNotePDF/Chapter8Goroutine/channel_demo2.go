package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		//for {
		//	n, ok := <-c
		//	if !ok {
		//		return
		//	}
		//	fmt.Println(n)
		//}

		//用range读取，替代循环读取
		for x := range c {
			fmt.Println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done

	signalWithChan()
}

//通过chan接受阻塞，实现goroute的信号并发控制，
// 当close chan的时候，chan收到消息，多个goroutin执行下一步
func signalWithChan() {
	var wg sync.WaitGroup
	c := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Println(id, "ready..")
			<-c
			fmt.Println(id, "running..")
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("All Ready?GO!")
	close(c)
	wg.Wait()
}
