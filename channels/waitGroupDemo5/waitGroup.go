package main

import (
	"fmt"
	"sync"
)

/**
展示waitGroup的用法
1.指针类型
2.声明 var wg  sync.WaitGroup
3.几个方法
	3.1Add(int)需要等待多少个
	3.2Wait()开始等待
	3.3Done()任务完成
*/

func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done()
	}
}

func createWorker(num int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(num, w.in, wg)
	return w
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
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
