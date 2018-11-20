package main

import (
	"time"
	"fmt"
	"os"
)

//P178  timeout  and tick demo

func main() {
	//timeout

	go func() {
		for {
			select {
			case <-time.After(time.Second * 3):
				fmt.Println("has timeout...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		t := time.Tick(time.Second)
		for {
			select {
			case <-t:
				fmt.Println("tick at ", time.Now())
			}
		}
	}()

	//不使用waitGroup阻塞进程调用goroutine的方法
	<-(chan struct{})(nil)

}
