package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
goroutine基础
1.go关键字+匿名函数（或函数）
2.go的函数里要通过传值方式将外面的变量传入，为什么不用闭包？因为有并发安全性问题
    2.1在循环体外面打印a也会有问题，多线程操作同一个变量的问题。如何解决？通过channel解决
3.io和runtime.Gosched可以让协程交出控制权
4.
*/

func main() {
	var a [10]int
	// fmt.Println(a)
	for i := 0; i < 10; i++ {
		go func(j int) {
			// go func() {
			for {
				// fmt.Printf("hello goroutine from %d\n", j)
				a[j]++ //只有这个逻辑 ，没有io或其他逻辑，这里就会死机，因为不会交出控制权，main函数也拿不到控制权，无法进入下面
				// 加入io如fmt print操作或使用runtime.Gosched()
				runtime.Gosched()
			}
		}(i)
		// }()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
