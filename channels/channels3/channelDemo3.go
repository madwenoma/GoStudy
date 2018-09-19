package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// for {
	// 	fmt.Printf("worker %d received %d\n", id, <-c) //go func函数内才是真正的worker可以坦度提炼为函数
	// }

	//用range接受chan知道发送方，发送完毕，防止当chan close之后仍然打印chan初始值
	//也可以用两个值接收
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n) //go func函数内才是真正的worker可以坦度提炼为函数
	}
}

//创建worker 返回一个chanel，
// func createWorker(num int) chan int {
func createWorker(num int) chan<- int { //返回一个只能发送数据的channel
	// func createWorker(num int) <-chan int { //定义一个只能收数据的channel 箭头表示方向，
	c := make(chan int)
	go worker(num, c)
	return c
}

func chanDemo() {
	// var channels [10]chan int
	var channels [10]chan<- int //声明的时候必须用<-表示和这个chan只能发送数据
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //“a” + i会报错必须使用单引号
	}
	// nn := <-channels[1] //如果这里想接受，就会编译错误
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //“a” + i会报错必须使用单引号
	}

	time.Sleep(time.Second)
}

/*缓冲区chan
1.正常的chan如果没有接受者，在发送时会报错，deadlock
2.缓冲区chan，表示建立一个缓冲，在缓冲区间内，不会进行发送，也不会报错，超过3之后，报错，对性能有一定提升
*/
func bufferedChannel() {
	c := make(chan int, 3)
	c <- 'a'
	c <- 'b'
	c <- 'b'
	// c <- 'd'
	time.Sleep(time.Second)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'b'
	c <- 'b'
	c <- 'b'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
