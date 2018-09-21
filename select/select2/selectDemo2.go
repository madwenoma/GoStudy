package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
select demos
1.语法
2.通过select case进行调度控制
	2.1通过变量控制发送和接收者的写作，当发送了一个值之后，接收者初始化开始接收，接收完毕后重新设置值
		2.1.1通过bool值
		2.1.2通过int分片
	2.2time的两个用法
		2.2.1time.After多少时间之后返回一个chan，定时器用途
		2.2.2time.Tick每隔多少时间发送一个chan
3.select中可以使用nilChannel
*/

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

func createWorker(num int) chan<- int { //返回一个只能发送数据的channel
	// func createWorker(num int) <-chan int { //定义一个只能收数据的channel 箭头表示方向，
	c := make(chan int)
	go worker(num, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n) //go func函数内才是真正的worker可以坦度提炼为函数
	}
}

// func main() {
// 	fmt.Println("select...")
// 	var c1, c2 = createChan(), createChan()
// 	w := createWorker(0)
// 	n := 0
// 	hasValue := false
// 	for {
// 		var activeWorker chan<- int
// 		if hasValue {
// 			activeWorker = w
// 		}
// 		select {
// 		case n = <-c1:
// 			hasValue = true
// 		case n = <-c2:
// 			hasValue = true
// 		case activeWorker <- n:
// 			hasValue = false
// 		}
// 	}
// 	time.Sleep(time.Millisecond)
// }

//使用分片接收发送来的数据
func main() {
	fmt.Println("select...")
	var c1, c2 = createChan(), createChan()
	w := createWorker(0)
	var values []int                   //用来接收c1 c2发送的数据的缓冲分片
	tm := time.After(20 * time.Second) //总时间
	tick := time.Tick(time.Second)     //每隔多少时间返回一个channel，接收者根据tick进行相应操作

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 { //如果value有值了 ，表示c1 c2发送过数据
			activeWorker = w        //给worker通道赋值
			activeValue = values[0] //从缓冲里拿第一个数据进行发送
		}
		select {
		case n := <-c1:
			values = append(values, n) //将数据加入到分片里
		case n := <-c2:
			values = append(values, n) //将数据加入到分片里
		case activeWorker <- activeValue: //吧数据发送给worker
			values = values[1:] //删除已发送的数据
		case <-time.After(800 * time.Millisecond): //for每次循环，重置时间，超过800ms就会打印
			fmt.Println("timeout..")
		case <-tick: //当每次达到设定时间，通道都会收到数据
			fmt.Println(len(values))
		case <-tm: //当tm达到设定时间，会收到通道数据
			fmt.Println("bye")
			return
		}
	}
	time.Sleep(time.Millisecond)
}
