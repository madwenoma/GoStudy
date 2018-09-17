package main

import (
	"goStudy/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(4)
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//q.Push("one")
	//fmt.Println(q.Pop())
}
