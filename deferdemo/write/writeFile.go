package main

import (
	"fmt"
	"goStudy/deferdemo"
	"os"
	"bufio"
)

/**
https://www.bilibili.com/video/av24365381/?p=28
7-1 defer的用法
后进先出
有点像final，return之后执行
 */

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		//错误处理
		if pathError,ok := err.(*os.PathError); ok{
			fmt.Println(pathError.Err)
			fmt.Println(pathError.Op)
			fmt.Println(pathError.Path)
		} else {
			fmt.Println(err)

		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := deferdemo.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}

	}
}

func main() {
	f := deferdemo.Fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	writeFile("fibonacci.txt")
	tryDefer()
}
