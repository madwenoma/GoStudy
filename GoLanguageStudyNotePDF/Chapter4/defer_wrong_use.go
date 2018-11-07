package main

import (
	"fmt"
	"os"
	"log"
)

//演示defer的错误用法
//defer用来关闭资源的时候，尽量缩小到一个小的方法体内，
// 而不要放在一大段长代码里或者循环里

func main() {
	for i := 0; i < 10000; i++ {
		path := fmt.Sprintf("./log/%d.txt", i)
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()
		//do something
	}
}


//正确用法
func rightUseDefer() {
	do := func(n int) {
		path := fmt.Sprintf("./log/%d.txt", n)
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		//do something
	}
	for i := 0; i < 10000; i++ {
		do(i)
	}
}
