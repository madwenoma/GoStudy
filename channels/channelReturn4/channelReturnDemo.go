package main

import (
	"fmt"
)

func doWorker(id int, c chan int, signal chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n) //go func函数内才是真正的worker可以坦度提炼为函数
		// go func() {
		// 	signal <- true
		// }()
		signal <- true //完成任务信号
	}
}

func createWorker(num int) worker {
	w := worker{
		in:     make(chan int),
		signal: make(chan bool),
	}
	go doWorker(num, w.in, w.signal)
	return w
}

type worker struct {
	in     chan int
	signal chan bool
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for key, worker := range workers {
		worker.in <- 'a' + key //“a” + i会报错必须使用单引号
		<-worker.signal
	}
	for key, worker := range workers {
		worker.in <- 'A' + key //“a” + i会报错必须使用单引号
		<-worker.signal
	}
	//如果两个一起接受完毕信号，则在doWork函数里需要另外定义一个go func
	//但是这样稍微复杂一些，就改为简单的返，在各自的for里接收信号
	// for _, worker := range workers {
	// 	<-worker.signal
	// 	<-worker.signal
	// }

}

func main() {
	chanDemo()
}
