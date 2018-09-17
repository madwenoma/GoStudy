package main

import (
	"fmt"
	"math"
)

var aa = 3 //函数外面定义必须使用var  包内部变量
var ss = "kkk"
var bb = true
//简写
var (
	a = 3 //函数外面定义必须使用var  包内部变量
	s = "kkk"
	b = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Print(a, b, s, "\n")
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "str"
	fmt.Print(a, b, c, s, "\n")
}
func variableShort() {
	a, b, c, s := 3, 4, true, "tiger" //简写函数内变量定义并赋值
	b = 5
	fmt.Print(a, b, c, s, "\n")
}

func forceType() {
	var a, b = 3, 4
	var c = math.Sqrt(float64(a*a + b*b)) //必须强制转换a*b结果
	fmt.Println(int(c))
}

const filename = "abc.txt" //包内部的常量
const (
	file1    = "hehe.txt"
	classNum = 3
)

func consts() {

	const a, b = 3, 4             //没有指定类型的常量
	var c int                     //声明c的类型为int，就需要强制转为int
	c = int(math.Sqrt(a*a + b*b)) //没有指定类型的常量，运行时自动识别为float
	fmt.Println(filename, c)
}

func enumsDemo() {
	const (
		java   = iota
		_
		golang
		python
		csharp
	)
	const (
		b float32 = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(java, csharp, golang, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Print("hello go")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(a, s, b)
	forceType()
	consts()
	enumsDemo()
}
