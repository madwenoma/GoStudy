package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
select demos
1.语法
2.select代码块里用case接受，两个channel哪一个先传数据过来就显示哪一个
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
	var values []int
	tm := time.After(20 * time.Second) //总时间
	tick := time.Tick(time.Second)     //每隔多少时间返回一个channel，接收者根据tick进行相应操作

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //for每次循环，重置时间，超过800ms就会打印
			fmt.Println("timeout..")
		case <-tick:
			fmt.Println(len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
	time.Sleep(time.Millisecond)
}
