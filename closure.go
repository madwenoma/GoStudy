package main

import (
	"fmt"
	"io"
	"strings"
	"bufio"
)

/*
https://www.bilibili.com/video/av24365381/?p=25
闭包的用法
*/

//注意写法都不同，声明里有两个func关键字，方法体有一个

func addSum() func(i int) int {
	sum := 0 //函数访问函数外面的变量
	return func(v int) int {
		sum += v
		return sum
	}
}

//闭包实现斐波那契 声明的地方有2个func
//func fibonacci() func() int {
//或者 type声明一个 函数：func() int，如此一来，这里就直接返回函数对象就可以
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

//声明一个函数对象，函数也是一个基本单位，可以声明、可以返回、可以作为参数
type intGen func() int
//函数实现接口
func (p intGen) Read(b []byte) (n int, err error) {
	next := p() //调用一下就是下一个
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(b)
}

func main() {
	a := addSum() //闭包的使用，必须先用变量声明是一个函数
	for i := 0; i < 10; i++ {
		fmt.Println(a(i)) //调用函数
	}

	f := fibonacci()

	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	printWithRreader(f)

}

func printWithRreader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
