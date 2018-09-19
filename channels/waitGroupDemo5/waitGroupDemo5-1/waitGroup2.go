package main

import (
	"fmt"
	"sync"
)

//对waitGroup.go的重构，不在传递WaitGroup对象，抽象为一个函数
func doWork(id int, c chan int, d func()) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		d()
	}
}

func createWorker(num int, d func()) worker {
	w := worker{
		in:       make(chan int),
		doneFunc: d,
	}
	go doWork(num, w.in, w.doneFunc)
	return w
}

type worker struct {
	in       chan int
	doneFunc func()
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, func() {
			wg.Done()
		})
	}
	wg.Add(20)
	for key, worker := range workers {
		worker.in <- 'a' + key //“a” + i会报错必须使用单引号
	}
	for key, worker := range workers {
		worker.in <- 'A' + key //“a” + i会报错必须使用单引号
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
