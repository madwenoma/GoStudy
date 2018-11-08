package main

import (
	"time"
	"fmt"
)

/**
知识点
1.go关键字，并非是立即执行并发操作，而是创建一个并发任务单元。
	加入一个调度队列，等待系统线程调度，因此并不保证执行顺序，保证不阻塞
2.go之后的代码，不会立即执行，但是会保存函数指针，调用参数。
3.go关键字并立即为该任务单元分配执行所需的栈内空间，默认大小只有2KB，所以相比
	其他语言动辄MB级别的线程空间而言，有巨大的优势，轻易支持上千万并发。
4.go的自定义栈是按需分配，需要时自动扩容，最大能到GB规模。



 */


var c = 2

func counter() int {
	c++
	return c
}

func main() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second)
		fmt.Println("go:", x, y)
	}(a, counter())//具有defer一样的延迟执行特点，会保存变量，并执行函数
	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("go:", a, counter())
	//}()
	a += 200
	fmt.Println("main:", a, counter())
	time.Sleep(time.Second * 3)
}
