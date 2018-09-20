package main

import (
	"fmt"
	"sync"
	"time"
)

/**
https://www.bilibili.com/video/av24365381/?p=46
go中锁的用法
1.关键字sync.Mutex
2.代码块是用匿名函数实现的
3.传统同步机制如（锁、waitGroup、Cond）很少用到，
    3.1因为传统同步机制基于共享内存，效率较低
    3.2go常用的还是通过共享通道
*/

type atomicInt struct {
	v    int
	lock sync.Mutex
}

func (atoInt *atomicInt) increment() {
	// atoInt.lock.Lock()
	// defer atoInt.lock.Unlock()
	// atoInt.v++
	fmt.Println("save incre..")
	func() {
		atoInt.lock.Lock()
		defer atoInt.lock.Unlock()
		atoInt.v++
	}()
}

func (atoInt *atomicInt) get() int {
	atoInt.lock.Lock()
	defer atoInt.lock.Unlock()
	return atoInt.v
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
