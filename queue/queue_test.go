package queue

import "fmt"

/**
go example的用法
1.文件名以_test结尾
2.类函数名以Example开头
3.逻辑写完后空一行 //OutPut:
4.下面按顺序写期望返回值
5.godoc -http :6060
 */

func ExampleQueue_IsEmpty() {
	q := Queue{1}
	fmt.Println(q.IsEmpty())

	//OutPut:
	//false
}