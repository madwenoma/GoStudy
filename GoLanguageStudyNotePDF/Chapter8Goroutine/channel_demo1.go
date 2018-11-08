package main

import "fmt"

func main() {
	done := make(chan struct{}) ////同步模式,成对出现，否则会阻塞
	c := make(chan string)
	go func() {
		s := <-c //数据通道，接受消息
		fmt.Println(s)
		close(done) //关闭通道，是向通道发送了消息
	}()
	c <- "hello channel" //数据通道，发送数据
	<-done               //通道接受消息

	syncChan()
}

func syncChan() {
	c := make(chan int, 3) //创建异步通道，3表示3个槽位的缓冲槽
	c <- 1
	c <- 2
	c <- 3
	//c <- 4 报错，只能接受3个
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
